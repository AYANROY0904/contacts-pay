package controllers

import (
	"contacts-pay/config"
	"contacts-pay/models"
	"fmt"
	"log"
	"net/http"

	//"github.com/aws/aws-sdk-go/aws"

	//"github.com/aws/aws-sdk-go/service/dynamodb"
	//"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/gin-gonic/gin"
)

// SyncContacts API
func SyncContacts(c *gin.Context) {
	var contacts []models.User
	if err := c.BindJSON(&contacts); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	for _, contact := range contacts {
		// Store in PostgreSQL
		config.DB.Create(&contact)

		// Store in Redis for fast lookup
		err := config.RDB.Set(config.Ctx, contact.Userid, fmt.Sprintf("%s,%s,%s", contact.PhoneNumber, contact.Username, contact.UciID), 0).Err()
		if err != nil {
			log.Fatalf("Failed to store data in Redis: %v", err)
		} else {
			fmt.Println("Data stored successfully!")
		}
		// Store in DynamoDB for record-keeping
		// 	item, err := dynamodbattribute.MarshalMap(contact)
		// 	if err != nil {
		// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to store in DynamoDB"})
		// 		return
		// 	}
		// 	_, err = config.DynamoDB.PutItem(&dynamodb.PutItemInput{
		// 		TableName: aws.String("contacts-pay"),
		// 		Item:      item,
		// 	})
		// 	if err != nil {
		// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to store in DynamoDB"})
		// 		return
		// 	}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Contacts synced"})
}

// Lookup Contact API
func LookupContact(c *gin.Context) {
	phone := c.Param("phone")

	// Check in Redis first
	if cached, err := config.RDB.Get(config.Ctx, phone).Result(); err == nil {
		c.JSON(http.StatusOK, gin.H{"data": cached})
		return
	}

	// If not found in Redis, check PostgreSQL
	var user models.User
	config.DB.First(&user, "phone_number = ?", phone)

	if user.PhoneNumber != "" {
		// Store in Redis for fast future lookups
		data := fmt.Sprintf("%s,%s", user.Username, user.UciID)
		config.RDB.Set(config.Ctx, phone, data, 0)

		c.JSON(http.StatusOK, gin.H{"data": data})
		return
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
}
