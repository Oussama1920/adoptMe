package logging

import (
	"context"
	"net/http"
	"strings"
	"time"

	config "github.com/Oussama1920/adoptMe/go/pkg/config"
	db "github.com/Oussama1920/adoptMe/go/pkg/db"
	utilis "github.com/Oussama1920/adoptMe/go/pkg/utilis"
	"github.com/google/uuid"
	"github.com/thanhpk/randstr"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func SignUp(dbHandler db.DbHandler, ctx context.Context, logger *logrus.Logger) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var newUser db.User
		// Call BindJSON to bind the received JSON to
		if err := c.BindJSON(&newUser); err != nil {
			logger.Error(err)
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Failed to Parse User"})
			return
		}
		if newUser.Password != newUser.PasswordConfirm {
			logger.Error("Passwords do not match")
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Passwords do not match"})
			return
		}
		hashedPassword, err := utilis.HashPassword(newUser.Password)
		if err != nil {
			logger.Error(err)
			c.IndentedJSON(http.StatusBadGateway, gin.H{"status": "error", "message": err.Error()})
			return
		}
		// Generate Verification Code
		code := randstr.String(20)

		verification_code := utilis.Encode(code)

		// Update User in Database
		newUser.VerificationCode = verification_code

		now := time.Now()

		newUser.CreatedAt = now
		newUser.Email = strings.ToLower(newUser.Email)
		newUser.Role = "user"
		newUser.Provider = "local"
		newUser.UpdatedAt = now
		newUser.Password = hashedPassword
		newUser.ID = uuid.New().String()
		if err := dbHandler.AddUser(ctx, newUser); err != nil {
			logger.Error(err)
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Failed to insert user", "error": err.Error()})
			return
		}

		// ? Send Email
		emailData := utilis.EmailData{
			URL:       "http://localhost:8080" + "/auth/verifyemail/" + code,
			FirstName: newUser.FirstName,
			Subject:   "Your account verification code",
		}

		utilis.SendEmail(newUser, &emailData)

		message := "We sent an email with a verification code to " + newUser.Email
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

		c.IndentedJSON(http.StatusCreated, gin.H{"status": "success", "message": message})
	}

	return gin.HandlerFunc(fn)
}
func Login(dbHandler db.DbHandler, ctx context.Context, logger *logrus.Logger) gin.HandlerFunc {
	var tokenConfig utilis.TokenConfig

	if err := config.GetDataConfiguration("service.token", &tokenConfig); err != nil {
		logger.Errorf("can't read token configuration : %v", err.Error())
	}

	fn := func(c *gin.Context) {
		//to DO : check request
		var user db.SignInInput
		// Call BindJSON to bind the received JSON to
		if err := c.BindJSON(&user); err != nil {
			logger.Error("failed to parse input : ", c.Request.Body, "  error is : ", err)
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Failed to Parse User"})
			return
		}
		userResult, err := dbHandler.GetUserByEmail(ctx, strings.ToLower(user.Email))
		if err != nil {
			logger.Error(err)
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": "Failed to connect"})
			return
		}
		if !userResult.Verified {
			logger.Error("Please verify your email")

			c.JSON(http.StatusForbidden, gin.H{"status": "fail", "message": "Please verify your email"})
			return
		}
		if err := utilis.VerifyPassword(userResult.Password, user.Password); err != nil {
			logger.Error("Invalid email or Password ")
			c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Invalid email or Password " + err.Error()})
			return
		}
		// Generate Token
		token, err := utilis.GenerateToken(tokenConfig.TOKEN_EXPIRED_IN, userResult.ID, tokenConfig.TOKEN_SECRET)
		if err != nil {
			logger.Error(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
			return
		}
		c.SetCookie("token", token, tokenConfig.TOKEN_MAXAGE*60, "/", "localhost", false, true)
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.JSON(http.StatusOK, gin.H{"status": "success", "token": token})

	}

	return gin.HandlerFunc(fn)
}
func VerifyEmail(dbHandler db.DbHandler, ctx context.Context, logger *logrus.Logger) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		code := c.Params.ByName("verificationCode")
		verification_code := utilis.Encode(code)
		logger.Println("received with code : ", code)
		updatedUser, err := dbHandler.GetVerificationCode(ctx, verification_code)
		if err != nil {
			logger.Error("Invalid verification code or user doesn't exists")
			c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Invalid verification code or user doesn't exists"})
			return
		}

		if updatedUser.Verified {
			logger.Error("User already verified")

			c.JSON(http.StatusConflict, gin.H{"status": "fail", "message": "User already verified"})
			return
		}

		updatedUser.VerificationCode = ""
		updatedUser.Verified = true
		err = dbHandler.UpdateUser(ctx, *updatedUser)
		if err != nil {
			logger.Error("Failed to update user : ", err)

			c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Invalid verification code or user doesn't exists"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Email verified successfully"})

	}

	return gin.HandlerFunc(fn)
}
func LogOut(dbHandler db.DbHandler, ctx context.Context, logger *logrus.Logger) gin.HandlerFunc {

	fn := func(c *gin.Context) {
		c.SetCookie("token", "", -1, "/", "localhost", false, true)
		c.JSON(http.StatusOK, gin.H{"status": "success"})
	}

	return gin.HandlerFunc(fn)
}
func GetMe(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(*db.User)

	userResponse := &db.UserResponse{
		ID:        currentUser.ID,
		Name:      currentUser.Name,
		Email:     currentUser.Email,
		Photo:     currentUser.Photo,
		Role:      currentUser.Role,
		Provider:  currentUser.Provider,
		CreatedAt: currentUser.CreatedAt,
		UpdatedAt: currentUser.UpdatedAt,
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"user": userResponse}})
}
