package services

import (
	"kisaanSathi/pkg/network"
	"kisaanSathi/pkg/repo"
	banner "kisaanSathi/pkg/services/banner/handler"
	community "kisaanSathi/pkg/services/community/handler"
	contact "kisaanSathi/pkg/services/contact/handler"
	cropHandler "kisaanSathi/pkg/services/crop/handler"
	expert "kisaanSathi/pkg/services/expert/handler"
	faq "kisaanSathi/pkg/services/faq/handler"
	feedback "kisaanSathi/pkg/services/feedback/handler"
	"kisaanSathi/pkg/services/feeds"
	homeHandler "kisaanSathi/pkg/services/home/handler"
	mandi "kisaanSathi/pkg/services/mandi/handler"
	metadata "kisaanSathi/pkg/services/metadata/handler"
	schemes "kisaanSathi/pkg/services/schemes/handler"
	soil_handler "kisaanSathi/pkg/services/soil/handler"
	reg "kisaanSathi/pkg/services/user/handler"
	weather_handler "kisaanSathi/pkg/services/weather/handler"

	video "kisaanSathi/pkg/services/video/handler"

	"net/http"

	"github.com/gin-gonic/gin"
)

type serviceObject struct {
	metadata.MetadataHandler
	reg.UserHandler
	soil_handler.SoilHandler
	feeds.FeedsHandler
	mandi.MandiHandler
	homeHandler.HomeHandler
	banner.BannerHandler
	weather_handler.WeatherHandler
	cropHandler.CropHandler
	contact.ContactHandler
	feedback.FeedbackHandler
	faq.FAQHandler
	schemes.SchemeHandler
	video.VideoHandler
	community.CommunityHandler
	expert.ExpertHandler
}

type ServiceLayer interface {
	GetMFHealth(c *gin.Context)
	metadata.MetadataHandler
	reg.UserHandler
	soil_handler.SoilHandler
	feeds.FeedsHandler
	mandi.MandiHandler
	homeHandler.HomeHandler
	banner.BannerHandler
	weather_handler.WeatherHandler
	cropHandler.CropHandler
	contact.ContactHandler
	feedback.FeedbackHandler
	faq.FAQHandler
	schemes.SchemeHandler
	video.VideoHandler
	community.CommunityHandler
	expert.ExpertHandler
}

func NewServiceObject(repo repo.DataObject) ServiceLayer {
	return &serviceObject{
		metadata.NewMetadataHandler(metadata.NewMetaDataController(repo)),
		reg.NewUserHandler(reg.NewUserController(repo)),
		soil_handler.NewSoilHandler(soil_handler.SoilController(repo)),
		feeds.NewFeedsHandler(repo),
		mandi.NewMandiHandler(mandi.NewMandiController(repo)),
		homeHandler.NewHomeHandler(homeHandler.NewHomeController(repo)),
		banner.NewBannerHandler(banner.NewBannerController(repo)),
		weather_handler.NewWeatherHandler(weather_handler.NewWeatherController()),
		cropHandler.NewCropHandler(cropHandler.NewCropController(repo)),
		contact.NewContacthandler(contact.NewContactController(repo)),
		feedback.NewFeedbackHandler(feedback.NewFeedbackController(repo)),
		faq.NewFAQHandler(faq.NewFAQController(repo)),
		schemes.NewSchemeHandler(schemes.NewSchemeController(repo)),
		video.NewVideoHandler(video.NewVideoController(repo)),
		community.NewCommunityHandler(community.NewCommunityController(repo)),
		expert.NewExpertHandler(expert.NewExpertController(repo)),
	}
}

func (s *serviceObject) GetMFHealth(c *gin.Context) {
	c.JSON(http.StatusOK, network.SuccessResponse("I AM HEALTHY"))
}
