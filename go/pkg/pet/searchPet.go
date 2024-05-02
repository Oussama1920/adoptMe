package pet

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/Oussama1920/adoptMe/go/pkg/db"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func SearchPet(c *gin.Context, dbHandler db.DbHandler, logger *logrus.Logger) {
	var searchRequest db.SearchPet
	// Call BindJSON to bind the received JSON to

	if err := c.BindJSON(&searchRequest); err != nil {
		logger.Error(err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Failed to Parse searchRequest"})
		return
	}
	queryRequest := createSearchQuery(searchRequest)
	logger.Info("request is ", queryRequest)
	pets, err := dbHandler.GetListPets(c, queryRequest)
	if err != nil {
		logger.Error("failed to get list of pets: ", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Failed to get list of pets", "error": err.Error()})
		return
	}
	if len(pets) == 0 {
		c.IndentedJSON(http.StatusOK, gin.H{"status": "failed", "message": "no results"})

	}
	// get list of photos:
	for _, pet := range pets {
		if pet.Photo != "" {

			images := strings.TrimSuffix(pet.Photo, ",")
			listImages := strings.Split(images, ",")
			imageBytes, err := os.ReadFile(listImages[0])
			if err != nil {
				logger.Error(err)
				c.JSON(http.StatusBadRequest, gin.H{"error": "failed to read image file"})
			}
			dataURL := "data:image/jpeg;base64," + base64.StdEncoding.EncodeToString(imageBytes)
			pet.Images = append(pet.Images, db.Image{DataURL: dataURL})
			/*
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
			*/
		}
	}
	c.IndentedJSON(http.StatusOK, gin.H{"status": "success", "pets": pets})
}

func createSearchQuery(searchRequest db.SearchPet) string {
	// Initialize the base query string
	query := "SELECT id,user_id,name,type,age,photo,created_at FROM pets WHERE 1=1"

	// Add conditions based on the non-empty fields of the searchRequest
	if searchRequest.Name != "" {
		query += fmt.Sprintf(" AND name='%s'", searchRequest.Name)
	}
	if searchRequest.Age != "" {
		query += fmt.Sprintf(" AND age='%s'", searchRequest.Age)
	}
	if searchRequest.Type != "" {
		query += fmt.Sprintf(" AND type='%s'", searchRequest.Type)
	}
	if searchRequest.CreatedBefore != "" {
		query += fmt.Sprintf(" AND created_at<'%s'", searchRequest.CreatedBefore)
	}
	if searchRequest.CreatedAfter != "" {
		query += fmt.Sprintf(" AND created_at>'%s'", searchRequest.CreatedAfter)
	}
	return query
}
