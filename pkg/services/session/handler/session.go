package handler

import (
	"kisaanSathi/pkg/logger"
	"kisaanSathi/pkg/network"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (f *sessionObject) Validate(c *gin.Context) {
	logger.Log(c).Debug("SERVICE-START")
	defer logger.Log(c).Debug("SERVICE-END")

	// before validate the session check the flag to validate the session
	userId := c.GetHeader("userId")
	sessionId := c.GetHeader("sessionId")
	// Check if the slices are not nil and have at least one element
	if userId == "" || sessionId == "" {
		log.Println("userId or sessionId not found or empty")
		// unauthorized user request
		c.JSON(http.StatusUnauthorized, network.FailureResponse(network.ApiErrors.Unauthorized.WithErrorDescription("invalid auth token")))
		c.Abort()
		return
	}
	validSession, err := f.repo.Validate(c, userId, sessionId)
	if err != nil {
		// log the error and return
		log.Println(err)
	}

	if !validSession {
		// unauthorized user request
		log.Println("session validated successfully")
		c.JSON(http.StatusUnauthorized, network.FailureResponse(network.ApiErrors.Unauthorized.WithErrorDescription("invalid auth token")))
		c.Abort()
		return
	}
	log.Println("session validated successfully")
	logger.Log(c).Info("CALL STARTED")
	c.Next()
}
