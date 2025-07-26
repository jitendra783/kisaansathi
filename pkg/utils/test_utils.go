package utils

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
	"net/http/httptest"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	oracle "github.com/godoes/gorm-oracle"
	"gorm.io/gorm"
)

func NewMockDB() (*sql.DB, *gorm.DB, sqlmock.Sqlmock) {
	sqlDB, sqlMock, err := sqlmock.New()
	if err != nil {
		fmt.Printf("Failed to initialize sql mock: %v", err)
	}

	expectedOutput := sqlmock.NewRows([]string{"version"}).AddRow("19c")
	sqlMock.
		ExpectQuery("^select version from product_component_version where rownum = 1").
		WillReturnRows(expectedOutput)

	gormDB, err := gorm.Open(oracle.New(oracle.Config{
		Conn: sqlDB,
	}), &gorm.Config{})

	if err != nil {
		fmt.Printf("Failed to initialize GORM DB: %v", err)
	}

	if err = sqlMock.ExpectationsWereMet(); err != nil {
		fmt.Printf("there were unfulfilled expectations: %v", err)
	}

	return sqlDB, gormDB, sqlMock
}

func NewSqlxMockDB() (*sqlx.DB, sqlmock.Sqlmock) {
	sqlDB, sqlMock, err := sqlmock.New()
	if err != nil {
		fmt.Printf("Failed to initialize sql mock: %v", err)
	}

	sqlxDB := sqlx.NewDb(sqlDB, "sqlmock")

	if err != nil {
		fmt.Printf("Failed to initialize SQLX DB: %v", err)
	}

	return sqlxDB, sqlMock
}

func CreateTestGinContext(httpMethod string, requestPayload interface{}, queries map[string]string, params map[string]string) (w *httptest.ResponseRecorder, c *gin.Context) {
	recorder := httptest.NewRecorder()
	engine, _ := gin.CreateTestContext(recorder)

	byteData, err := json.Marshal(requestPayload)
	if err != nil {
		log.Fatalln(err)
	}

	engine.Request, err = http.NewRequest(httpMethod, "/", bytes.NewBuffer(byteData))
	if err != nil {
		log.Fatalln(err)
	}
	engine.Request.Header = http.Header{}
	engine.Request.Header.Set("Content-Type", "application/json")

	if queries != nil {
		q := engine.Request.URL.Query()
		for k, v := range queries {
			q.Add(k, v)
		}
		engine.Request.URL.RawQuery = q.Encode()
	}

	for k, v := range params {
		engine.Params = append(engine.Params, gin.Param{
			Key:   k,
			Value: v,
		})
	}

	return recorder, engine
}
