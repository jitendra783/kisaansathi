package feeds

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func(f *feedsHandler) GetFeeds(c *gin.Context) {
	c.JSON(http.StatusOK, []gin.H{
		{
			"title": "Prepare soil for Rabi crops",
			"summary": "Best practices for soil preparation.",
			"link": "https://krishi.gov.in/tips/rabi",
			"image": "https://cdn.krishi.gov.in/images/rabi.jpg",
			"type": "tip",
		},
		{
			"title": "Weather alert for next 3 days",
			"summary": "High rainfall expected in MP region.",
			"link": "https://weather.com/alert/mp",
			"image": "",
			"type": "news",
		},
	})
}
