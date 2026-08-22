curl -X GET "http://localhost:8080/v1/mandi/"
curl -X GET "http://localhost:8080/v1/mandi/crop/:crop"
curl -X GET "http://localhost:8080/v1/mandi/compare/:crop/:market"
curl -X GET "http://localhost:8080/v1/mandi/prices"
curl -X GET "http://localhost:8080/v1/mandi/state/:state"
curl -X GET "http://localhost:8080/v1/mandi/district/:district"
curl -X GET "http://localhost:8080/v1/mandi/trending"
curl -X GET "http://localhost:8080/v1/mandibhav"
curl -X GET "http://localhost:8080/v1/scheme/"
curl -X GET "http://localhost:8080/v1/scheme/eligible"
curl -X GET "http://localhost:8080/v1/scheme/state"
curl -X GET "http://localhost:8080/v1/scheme/:id"
curl -X GET "http://localhost:8080/v1/soil/report"
curl -X GET "http://localhost:8080/v1/soil/recommendation"
curl -X GET "http://localhost:8080/v1/soil/types"
curl -X GET "http://localhost:8080/v1/feeds"
curl -X GET "http://localhost:8080/v1/feed/"
curl -X GET "http://localhost:8080/v1/feedback/"
curl -X GET "http://localhost:8080/v1/faq/"
curl -X GET "http://localhost:8080/v1/faq/category/:category"
curl -X GET "http://localhost:8080/v1/faq/:id"
curl -X GET "http://localhost:8080/v1/forecast"
curl -X GET "http://localhost:8080/v1/weather/"
curl -X GET "http://localhost:8080/v1/weather/current"
curl -X GET "http://localhost:8080/v1/weather/hourly"
curl -X GET "http://localhost:8080/v1/weather/weekly"
curl -X GET "http://localhost:8080/v1/weather/alerts"
curl -X GET "http://localhost:8080/v1/weather/rainfall"
curl -X GET "http://localhost:8080/v1/crop/"
curl -X GET "http://localhost:8080/v1/crop/season"
curl -X GET "http://localhost:8080/v1/crop/recommended"
curl -X GET "http://localhost:8080/v1/crop/:id"
curl -X GET "http://localhost:8080/v1/community/posts"
curl -X GET "http://localhost:8080/v1/contact/info"
curl -X GET "http://localhost:8080/v1/video/"
curl -X GET "http://localhost:8080/v1/video/category/:category"
curl -X GET "http://localhost:8080/v1/video/:id"
curl -X GET "http://localhost:8080/v1/expert/"
curl -X GET "http://localhost:8080/v1/expert/appointments"
curl -X GET "http://localhost:8080/v1/expert/:id"
curl -X GET "http://localhost:8080/v1/home/"
curl -X GET "http://localhost:8080/v1/home/dashboard"
curl -X GET "http://localhost:8080/v1/user/details"
curl -X GET "http://localhost:8080/v1/banner/"
curl -X GET "http://localhost:8080/swagger/*any"
curl -X GET "http://localhost:8080/health"
curl -X POST "http://localhost:8080/v1/user/register" -H "Content-Type: application/json" -d '{}' 
curl -X POST "http://localhost:8080/v1/user/refresh-token" -H "Content-Type: application/json" -d '{}' 
curl -X POST "http://localhost:8080/v1/user/reset-password" -H "Content-Type: application/json" -d '{}' 
curl -X POST "http://localhost:8080/v1/user/login" -H "Content-Type: application/json" -d '{}' 
curl -X POST "http://localhost:8080/v1/user/logout" -H "Content-Type: application/json" -d '{}' 
curl -X POST "http://localhost:8080/v1/user/forgot-password" -H "Content-Type: application/json" -d '{}' 
curl -X POST "http://localhost:8080/v1/user/change-password" -H "Content-Type: application/json" -d '{}' 
curl -X POST "http://localhost:8080/v1/community/post" -H "Content-Type: application/json" -d '{}' 
curl -X POST "http://localhost:8080/v1/community/comment" -H "Content-Type: application/json" -d '{}' 
curl -X POST "http://localhost:8080/v1/community/like" -H "Content-Type: application/json" -d '{}' 
curl -X POST "http://localhost:8080/v1/community/share" -H "Content-Type: application/json" -d '{}' 
curl -X POST "http://localhost:8080/v1/contact/" -H "Content-Type: application/json" -d '{}' 
curl -X POST "http://localhost:8080/v1/crop/" -H "Content-Type: application/json" -d '{}' 
curl -X POST "http://localhost:8080/v1/soil/test" -H "Content-Type: application/json" -d '{}' 
curl -X POST "http://localhost:8080/v1/expert/book" -H "Content-Type: application/json" -d '{}' 
curl -X POST "http://localhost:8080/v1/feedback/" -H "Content-Type: application/json" -d '{}' 
