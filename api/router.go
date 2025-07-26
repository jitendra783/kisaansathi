package api

import (
	"bytes"
	"fmt"
	"io"
	"kisaanSathi/pkg/config"

	//"kisaanSathi/pkg/middlewares"
	serv "kisaanSathi/pkg/services"
	"kisaanSathi/pkg/utils"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

func getRouter(obj serv.ServiceLayer, logger *zap.Logger) *gin.Engine {
	router := gin.New()
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		utils.RegisterValidations(v)
	}
	router.Use(customLogger(logger))
	router.Use(gin.Recovery())
	router.GET("/health", obj.GetMFHealth)
	//router.Use(middlewares.AuthMiddleware())
	//router.Use(middlewares.AuthMiddlewareSession(obj))
	//NOTE : ADD ALLL ROUTES BELOW THIS POINT

	v1 := router.Group("/v1")
	v1.GET("/forecast", obj.GetForecast)
	v1.GET("/mandibhav", obj.GetMandiBhav)
	v1.GET("/feeds", obj.GetFeeds)
	user := v1.Group("/user")
	{
		user.POST("/login", obj.Login)
		user.POST("/logout", obj.Logout)
		user.POST("/register", obj.Register)
		//user.POST("/refreshtoken", obj.RefreshToken)
	}

	saveCurlCommands(router)
	return router
}

func saveCurlCommands(r *gin.Engine) {
	file, err := os.Create("api_requests.sh")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	for _, route := range r.Routes() {
		var curlCommand string
		url := fmt.Sprintf("http://localhost:8080%s", route.Path)

		switch route.Method {
		case "GET":
			curlCommand = fmt.Sprintf("curl -X GET \"%s\"\n", url)
		case "POST":
			data := "" //paylaod of api request body
			curlCommand = fmt.Sprintf("curl -X POST \"%s\" -H \"Content-Type: application/json\" -d '{%s}' \n", url, data)
		}

		_, err := file.WriteString(curlCommand)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}

	fmt.Println("cURL commands saved to api_requests.sh")
}

func customLogger(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			body string
		)
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery

		// Check if the content type is multipart/form-data (usually used for file uploads)
		contentType := c.Request.Header.Get("Content-Type")
		excludeBody := strings.HasPrefix(contentType, "multipart/form-data")

		// Read the request body unless it's excluded
		if !excludeBody {
			bodyBytes, err := io.ReadAll(c.Request.Body)
			if err == nil {
				bodyCopy := bytes.NewBuffer(bodyBytes)
				c.Request.Body = io.NopCloser(bodyCopy)
				body = string(bodyBytes)
			}
		}
		c.Next()

		if c.FullPath() != "/health" {
			latency := time.Since(start).Milliseconds()
			userID := c.GetString(config.USERID)
			uID := c.GetString(config.REQUESTID)
			ucc := c.GetString(config.UCC)
			isEncrypt := c.GetBool(config.ISENCRYPT)
			logger.Info("Call_Ended",
				zap.String("path", path),
				zap.String("requestID", uID),
				zap.String("ucc", ucc),
				zap.String("userId", userID),
				zap.Int("status", c.Writer.Status()),
				zap.String("method", c.Request.Method),
				zap.String("path", path),
				zap.String("query", query),
				zap.String("body", body),
				zap.String("user-agent", c.Request.UserAgent()),
				zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
				zap.Int64("latency", latency),
				zap.Bool("isEncrypt", isEncrypt),
			)
		}
	}
}
