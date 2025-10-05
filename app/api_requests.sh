curl -X GET "http://localhost:8080/v1/forecast"
curl -X GET "http://localhost:8080/v1/feeds"
curl -X GET "http://localhost:8080/v1/mandibhav"
curl -X GET "http://localhost:8080/health"
curl -X POST "http://localhost:8080/v1/user/login" -H "Content-Type: application/json" -d '{}' 
curl -X POST "http://localhost:8080/v1/user/logout" -H "Content-Type: application/json" -d '{}' 
curl -X POST "http://localhost:8080/v1/user/register" -H "Content-Type: application/json" -d '{}' 
