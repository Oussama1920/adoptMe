package logging

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
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
			URL:       "http://localhost:3000" + "/verify-email/" + code,
			FirstName: newUser.Name,
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
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		var user db.SignInInput
		// Call BindJSON to bind the received JSON to
		if err := c.BindJSON(&user); err != nil {
			logger.Error("failed to parse input : ", c.Request.Body, "  error is : ", err)
			c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Failed to Parse User"})
			return
		}
		userResult, err := dbHandler.GetUserByEmail(ctx, strings.ToLower(user.Email))
		if err != nil {
			logger.Error(err)
			c.JSON(http.StatusInternalServerError, gin.H{"status": "failed", "message": "Failed to connect"})
			return
		}
		logger.Info("user email: ", user.Email)
		logger.Info("user password: ", user.Password)

		// check if user exist
		if userResult.Email == "" {
			logger.Error("no user found with for this email")
			c.JSON(http.StatusNotFound, gin.H{"status": "failed", "message": "Failed to connect"})
			return

		}
		if err := utilis.VerifyPassword(userResult.Password, user.Password); err != nil {
			logger.Error("Invalid email or Password ")
			c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Invalid email or Password " + err.Error()})
			return
		}
		if !userResult.Verified {
			logger.Error("Please verify your email")

			c.JSON(http.StatusForbidden, gin.H{"status": "fail", "message": "Please verify your email"})
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
		c.JSON(http.StatusOK, gin.H{"status": "success", "token": token})

	}

	return gin.HandlerFunc(fn)
}
func VerifyEmail(dbHandler db.DbHandler, ctx context.Context, logger *logrus.Logger) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

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
		err = dbHandler.VerifyUser(ctx, *updatedUser)
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
		ID:          currentUser.ID,
		Name:        currentUser.Name,
		Email:       currentUser.Email,
		Photo:       currentUser.Photo,
		Role:        currentUser.Role,
		Provider:    currentUser.Provider,
		CreatedAt:   currentUser.CreatedAt,
		UpdatedAt:   currentUser.UpdatedAt,
		FirstName:   currentUser.FirstName,
		DateOfBirth: currentUser.DateOfBirth,
		PhoneNumber: currentUser.PhoneNumber,
		Address:     currentUser.Address,
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "user": userResponse})
}
func GetUser(dbHandler db.DbHandler, ctx context.Context, logger *logrus.Logger) gin.HandlerFunc {

	fn := func(c *gin.Context) {
		var token string

		cookie, err := c.Cookie("token")

		authorizationHeader := c.Request.Header.Get("Authorization")
		fields := strings.Fields(authorizationHeader)

		if len(fields) != 0 && fields[0] == "Bearer" {
			token = fields[1]
		} else if err == nil {
			token = cookie
		}

		if token == "" {
			logger.Error("Invalid email or Password ")
			c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "You are not logged in"})
		}
		var tokenConfig utilis.TokenConfig

		if err := config.GetDataConfiguration("service.token", &tokenConfig); err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail to parse config", "message": err.Error()})
		}
		sub, err := utilis.ValidateToken(token, tokenConfig.TOKEN_SECRET)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": err.Error()})
			return
		}

		user, err := dbHandler.GetUserById(ctx, sub.(string))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": "the user belonging to this token no logger exists"})
			return
		}
		logger.Infof("founded user is : %v", user)
		c.JSON(http.StatusOK, gin.H{"status": "success", "user": user})

	}

	return gin.HandlerFunc(fn)
}

func UpdateUser(c *gin.Context, dbHandler db.DbHandler, logger *logrus.Logger) {
	currentUser := c.MustGet("currentUser").(*db.User)
	// Call BindJSON to bind the received JSON to
	var receivedUser db.User
	if err := c.BindJSON(&receivedUser); err != nil {
		logger.Error("failed to parse input : ", c.Request.Body, "  error is : ", err)
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Failed to Parse User"})
		return
	}
	logger.Infof("current user %v: ", currentUser)
	logger.Infof("new user %v: ", receivedUser)
	err := dbHandler.UpdateUser(c, *currentUser, receivedUser)
	if err != nil {
		logger.Error("failed to update user input :  error is : ", err)
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Failed to updated User"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"user": receivedUser}})
}

func GetPet(c *gin.Context, dbHandler db.DbHandler, logger *logrus.Logger) {

	id := c.Param("id")
	if id == "" {
		logger.Error("id should not be empty")
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "id is empty"})
		return
	}
	// Convert the id parameter to an integer
	petID, err := strconv.Atoi(id)
	if err != nil {
		// If conversion fails, return an error response
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID, id should be an int"})
		return
	}

	pet, err := dbHandler.GetPet(c, petID)
	if err != nil {
		logger.Error(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Failed to insert user", "error": err.Error()})
		return
	}
	if pet.Photo != "" {
		images := strings.TrimSuffix(pet.Photo, ",")
		listImages := strings.Split(images, ",")
		for _, val := range listImages {
			logger.Info("found image : ", val)
			// now let's get the list of images:
			imageBytes, err := os.ReadFile(val)
			if err != nil {
				logger.Error(err)
				c.JSON(http.StatusBadRequest, gin.H{"error": "failed to read image file"})
			}
			dataURL := "data:image/jpeg;base64," + base64.StdEncoding.EncodeToString(imageBytes)
			pet.Images = append(pet.Images, db.Image{DataURL: dataURL})
		}
	}
	c.IndentedJSON(http.StatusCreated, gin.H{"status": "success", "pet": pet})
}

func generateUniqueFilename(userId, extension string, photoId int) string {

	return userId + strconv.Itoa(photoId) + extension
}

func getFileExtension(dataURL string) string {
	// Split the data URL by comma
	parts := strings.Split(dataURL, ",")
	// Check if there are parts after splitting
	if len(parts) != 2 {
		return "" // Invalid data URL format
	}

	// Extract the base64 encoded data
	firstSplit := parts[0]

	parts2 := strings.Split(firstSplit, ";")
	// Check if there are parts after splitting
	if len(parts2) != 2 {
		return "" // Invalid data URL format
	}

	// Extract the MIME type from the data URL
	mimeType := strings.Split(parts2[0], ";")[0]
	fmt.Println("mimeType : ", mimeType)

	// Extract the file extension from the MIME type
	switch mimeType {
	case "data:image/png":
		return ".png"
	case "data:image/jpeg":
		return ".jpg"
	case "data:image/gif":
		return ".gif"
	default:
		return "" // Unsupported MIME type or no extension found
	}
}

func AddPet(c *gin.Context, dbHandler db.DbHandler, logger *logrus.Logger) {

	currentUser := c.MustGet("currentUser").(*db.User)
	var newPet db.Pet
	// Call BindJSON to bind the received JSON to

	if err := c.BindJSON(&newPet); err != nil {
		logger.Error(err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Failed to Parse Pet"})
		return
	}
	newPet.CreatedAt = time.Now()
	listOfphotoPaths := ""
	// now we saved the annonce to database so let's save the images :
	// Process each image
	for _, img := range newPet.Images {
		// Decode base64 data URL
		data, err := base64.StdEncoding.DecodeString(img.DataURL[strings.IndexByte(img.DataURL, ',')+1:])
		if err != nil {
			logger.Errorf("failed to decode image: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode image"})
			return
		}

		// Save image to file system
		fileExtension := getFileExtension(img.DataURL)
		if fileExtension == "" {
			logger.Error("unsupported extension")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "unsupported extension"})

		}
		filename := generateUniqueFilename(currentUser.ID, fileExtension, newPet.ID) // Implement this function to generate a unique filename
		filePath := filepath.Join("/home/odjebbi/uploads", filename)
		if err := os.WriteFile(filePath, data, 0644); err != nil {
			logger.Errorf("failed to save file: %v", err)

			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save image"})
			return
		}
		listOfphotoPaths = listOfphotoPaths + filePath + ","

		// Now you have the image saved in your file system, you can save the filename or URL to your database
		// Example: Save filename to database
		// YourDatabaseModel.Create(filename)
	}
	newPet.Photo = listOfphotoPaths
	id, err := dbHandler.AddPet(c, newPet, currentUser.ID)
	if err != nil {
		logger.Error(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Failed to insert user", "error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusCreated, gin.H{"status": "success", "id": id})
}
