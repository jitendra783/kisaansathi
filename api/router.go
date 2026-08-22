package api

import (
	"bytes"
	"fmt"
	"io"
	"kisaanSathi/pkg/config"

	serv "kisaanSathi/pkg/services"
	"kisaanSathi/pkg/utils"
	"os"
	"strings"
	"time"

	_ "kisaanSathi/docs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
)
// @title KisaanSathi API
// @version 1.0
// @description This is the API documentation for KisaanSathi.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email
func getRouter(obj serv.ServiceLayer, logger *zap.Logger) *gin.Engine {
	router := gin.New()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	// Register custom validations
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		utils.RegisterValidations(v)
	}
	
	router.Use(customLogger(logger))
	
	router.Use(gin.Recovery())
	router.GET("/swagger/*any",
		ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/health", obj.GetMFHealth)
	//router.Use(middlewares.AuthMiddleware())
	//router.Use(middlewares.AuthMiddlewareSession(obj))
	//NOTE : ADD ALLL ROUTES BELOW THIS POINT

	v1 := router.Group("/v1")
	v1.GET("/forecast", obj.GetForecast)
	v1.GET("/mandibhav", obj.GetMandiBhav)
	v1.GET("/feeds", obj.GetFeeds)
	// ==========================
	// Authentication & User
	// ==========================
	user := v1.Group("/user")
	{
		user.POST("/login", obj.Login)
		user.POST("/logout", obj.Logout)
		user.POST("/register", obj.Register)
		user.POST("/refresh-token", obj.RefreshToken)
		user.POST("/forgot-password", obj.ForgotPassword)
		user.POST("/reset-password", obj.ResetPassword)
		user.POST("/change-password", obj.ChangePassword)
		// user.POST("/verify-otp", obj.VerifyOTP)
		// user.POST("/resend-otp", obj.ResendOTP)

		user.GET("/details", obj.GetUserDetails)
		// user.GET("/dashboard", obj.GetDashboard)
		// user.GET("/activity", obj.GetUserActivity)

		user.PUT("/update", obj.UpdateUserDetails)
		user.PUT("/avatar", obj.UpdateAvatar)

		// user.DELETE("/delete", obj.DeleteUser)
	}

	// ==========================
	// Home
	// ==========================
	home := v1.Group("/home")
	{
		home.GET("/", obj.GetHome)
		home.GET("/dashboard", obj.GetDashboardData)
	}

	// ==========================
	// Weather
	// ==========================
	weather := v1.Group("/weather")
	{
		weather.GET("/current", obj.GetCurrentWeather)
		weather.GET("/hourly", obj.GetHourlyForecast)
		weather.GET("/", obj.GetForecast)
		weather.GET("/weekly", obj.GetForecastWeekly)
		weather.GET("/alerts", obj.GetWeatherAlerts)
		weather.GET("/rainfall", obj.GetRainfall)
	}

	// ==========================
	// Mandi Bhav
	// ==========================
	mandi := v1.Group("/mandi")
	{
		mandi.GET("/", obj.GetMandiBhav)
		mandi.GET("/prices", obj.GetMandiPrices)
		mandi.GET("/crop/:crop", obj.GetCropPrices)
		mandi.GET("/state/:state", obj.GetStatePrices)
		mandi.GET("/district/:district", obj.GetDistrictPrices)
		mandi.GET("/trending", obj.GetTrendingPrices)
		mandi.GET("/compare/:crop/:market", obj.GetPriceComparison)
	}

	// ==========================
	// Crops
	// ==========================
	crop := v1.Group("/crop")
	{
		crop.GET("/", obj.GetCrops)
		crop.GET("/:id", obj.GetCrop)
		crop.POST("/", obj.CreateCrop)
		crop.PUT("/:id", obj.UpdateCrop)
		crop.DELETE("/:id", obj.DeleteCrop)

		crop.GET("/season", obj.GetCropSeason)
		crop.GET("/recommended", obj.GetRecommendedCrops)
	}

	// ==========================
	// Farm
	// ==========================
	// farm := v1.Group("/farm")
	// {
	// 	farm.GET("/", obj.GetFarm)
	// 	farm.POST("/", obj.CreateFarm)
	// 	farm.PUT("/", obj.UpdateFarm)
	// 	farm.DELETE("/:id", obj.DeleteFarm)
	// 	farm.GET("/history", obj.GetFarmHistory)
	// }

	// ==========================
	// Soil
	// ==========================
	soil := v1.Group("/soil")
	{
		soil.GET("/types", obj.GetSoilTypes)
		soil.GET("/report", obj.GetSoilReport)
		soil.POST("/test", obj.CreateSoilTest)
		soil.GET("/recommendation", obj.GetSoilRecommendation)
	}

	// ==========================
	// Fertilizer
	// ==========================
	// fertilizer := v1.Group("/fertilizer")
	// {
	// 	fertilizer.GET("/", obj.GetFertilizers)
	// 	fertilizer.GET("/recommendation", obj.GetFertilizerRecommendation)
	// 	fertilizer.GET("/dose", obj.GetFertilizerDose)
	// }

	// ==========================
	// Disease Detection
	// ==========================
	// disease := v1.Group("/disease")
	// {
	// 	disease.POST("/detect", obj.DetectDisease)
	// 	disease.POST("/image", obj.UploadDiseaseImage)
	// 	disease.GET("/history", obj.GetDiseaseHistory)
	// 	disease.GET("/solution", obj.GetDiseaseSolutions)
	// }

	// ==========================
	// Pest
	// ==========================
	// pest := v1.Group("/pest")
	// {
	// 	pest.GET("/", obj.GetPests)
	// 	pest.GET("/:crop", obj.GetCropPests)
	// 	pest.GET("/solution", obj.GetPestSolutions)
	// }

	// ==========================
	// Government Schemes
	// ==========================
	scheme := v1.Group("/scheme")
	{
		scheme.GET("/", obj.GetSchemes)
		scheme.GET("/:id", obj.GetScheme)
		scheme.GET("/eligible", obj.GetEligibleSchemes)
		scheme.GET("/state", obj.GetStateSchemes)
	}

	// ==========================
	// News Feed
	// ==========================
	feed := v1.Group("/feed")
	{
		feed.GET("/", obj.GetFeeds)
		// feed.GET("/latest", obj.GetLatestFeeds)
		// feed.GET("/category/:category", obj.GetFeedsByCategory)
		// feed.GET("/:id", obj.GetFeed)
	}

	// ==========================
	// Videos
	// ==========================
	video := v1.Group("/video")
	{
		video.GET("/", obj.GetVideos)
		video.GET("/:id", obj.GetVideo)
		video.GET("/category/:category", obj.GetVideosByCategory)
	}

	// ==========================
	// Community
	// ==========================
	community := v1.Group("/community")
	{
		community.GET("/posts", obj.GetPosts)
		community.POST("/post", obj.CreatePost)
		community.PUT("/post/:id", obj.UpdatePost)
		community.DELETE("/post/:id", obj.DeletePost)

		community.POST("/comment", obj.CreateComment)
		community.DELETE("/comment/:id", obj.DeleteComment)

		community.POST("/like", obj.LikePost)
		community.POST("/share", obj.SharePost)
	}

	// ==========================
	// Expert
	// ==========================
	expert := v1.Group("/expert")
	{
		expert.GET("/", obj.GetExperts)
		expert.GET("/:id", obj.GetExpert)
		expert.POST("/book", obj.BookConsultation)
		expert.GET("/appointments", obj.GetConsultationHistory)
	}

	// ==========================
	// Chat Assistant
	// ==========================
	// chat := v1.Group("/chat")
	// {
	// 	chat.POST("/message", obj.SendMessage)
	// 	chat.GET("/history", obj.GetChatHistory)
	// 	chat.DELETE("/history", obj.ClearChatHistory)
	// }

	// // ==========================
	// // Notifications
	// // ==========================
	// notification := v1.Group("/notification")
	// {
	// 	notification.GET("/", obj.GetNotifications)
	// 	notification.PUT("/read", obj.MarkNotificationRead)
	// 	notification.DELETE("/:id", obj.DeleteNotification)
	// }

	// ==========================
	// Crop Calendar
	// ==========================
	// calendar := v1.Group("/calendar")
	// {
	// 	calendar.GET("/", obj.GetCalendar)
	// 	calendar.GET("/tasks", obj.GetTasks)
	// 	calendar.POST("/task", obj.CreateTask)
	// 	calendar.PUT("/task/:id", obj.UpdateTask)
	// 	calendar.DELETE("/task/:id", obj.DeleteTask)
	// }

	// ==========================
	// Marketplace
	// ==========================
	// market := v1.Group("/market")
	// {
	// 	market.GET("/products", obj.GetProducts)
	// 	market.GET("/products/:id", obj.GetProduct)
	// 	market.POST("/products", obj.CreateProduct)
	// 	market.PUT("/products/:id", obj.UpdateProduct)
	// 	market.DELETE("/products/:id", obj.DeleteProduct)
	// }

	// ==========================
	// Orders
	// ==========================
	// order := v1.Group("/order")
	// {
	// 	order.POST("/", obj.CreateOrder)
	// 	order.GET("/", obj.GetOrders)
	// 	order.GET("/:id", obj.GetOrder)
	// 	order.PUT("/cancel/:id", obj.CancelOrder)
	// }

	// ==========================
	// Payments
	// ==========================
	// payment := v1.Group("/payment")
	// {
	// 	payment.POST("/create", obj.CreatePayment)
	// 	payment.POST("/verify", obj.VerifyPayment)
	// 	payment.POST("/webhook", obj.PaymentWebhook)
	// 	payment.GET("/history", obj.GetPaymentHistory)
	// }

	// ==========================
	// Feedback
	// ==========================
	feedback := v1.Group("/feedback")
	{
		feedback.POST("/", obj.CreateFeedback)
		feedback.GET("/", obj.GetFeedbacks)
	}

	// ==========================
	// FAQ
	// ==========================
	faq := v1.Group("/faq")
	{
		faq.GET("/", obj.GetFAQs)
		faq.GET("/:id", obj.GetFAQByID)
		faq.GET("/category/:category", obj.GetFAQsByCategory)
	}

	// ==========================
	// Contact
	// ==========================
	contact := v1.Group("/contact")
	{
		contact.POST("/", obj.ContactUs)
		contact.GET("/info", obj.GetContactInfo)
	}

	// ==========================
	// Banner
	// ==========================
	banner := v1.Group("/banner")
	{
		banner.GET("/", obj.GetBanners)
	}

	// ==========================
	// Search
	// ==========================
	// search := v1.Group("/search")
	// {
	// 	search.GET("/", obj.Search)
	// 	search.GET("/crop", obj.SearchCrop)
	// 	search.GET("/news", obj.SearchNews)
	// 	search.GET("/video", obj.SearchVideos)
	// 	search.GET("/expert", obj.SearchExperts)
	// }

	// ==========================
	// Location
	// ==========================
	// location := v1.Group("/location")
	// {
	// 	location.GET("/states", obj.GetStates)
	// 	location.GET("/districts", obj.GetDistricts)
	// 	location.GET("/villages", obj.GetVillages)
	// }

	// ==========================
	// Admin
	// ==========================
	// admin := v1.Group("/admin")
	// {
	// 	admin.GET("/dashboard", obj.AdminDashboard)

	// 	admin.GET("/users", obj.GetUsers)
	// 	admin.GET("/farmers", obj.GetFarmers)
	// 	admin.GET("/experts", obj.GetExpertsAdmin)

	// 	admin.POST("/banner", obj.CreateBanner)
	// 	admin.PUT("/banner/:id", obj.UpdateBanner)
	// 	admin.DELETE("/banner/:id", obj.DeleteBanner)

	// 	admin.POST("/news", obj.CreateNews)
	// 	admin.PUT("/news/:id", obj.UpdateNews)
	// 	admin.DELETE("/news/:id", obj.DeleteNews)

	// 	admin.POST("/video", obj.CreateVideo)
	// 	admin.PUT("/video/:id", obj.UpdateVideo)
	// 	admin.DELETE("/video/:id", obj.DeleteVideo)

	// 	admin.GET("/analytics", obj.GetAnalytics)
	// 	admin.GET("/reports", obj.GetReports)
	// }

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
			logger.Info("Call_Ended",
				zap.String("path", path),
				zap.String("requestID", uID),
				zap.String("userId", userID),
				zap.Int("status", c.Writer.Status()),
				zap.String("method", c.Request.Method),
				zap.String("path", path),
				zap.String("query", query),
				zap.String("body", body),
				zap.String("user-agent", c.Request.UserAgent()),
				zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
				zap.Int64("latency", latency),
			)
		}
	}
}
