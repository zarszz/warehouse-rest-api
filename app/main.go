package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/spf13/viper"
	_categoryHttpDelivery "github.com/zarszz/warehouse-rest-api/category/delivery/http"
	_categoryHttpDeliveryMiddleware "github.com/zarszz/warehouse-rest-api/category/delivery/http/middleware"
	_categoryRepo "github.com/zarszz/warehouse-rest-api/category/repository/mysql"
	_categoryUcase "github.com/zarszz/warehouse-rest-api/category/usecase"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func main() {
	dbHost := viper.GetString(`database.host`)
	dbPort := viper.GetString(`database.port`)
	dbUser := viper.GetString(`database.user`)
	dbPass := viper.GetString(`database.pass`)
	dbName := viper.GetString(`database.name`)
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Jakarta")
	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())
	dbConn, err := sql.Open(`mysql`, dsn)

	if err != nil {
		log.Fatal(err)
	}
	err = dbConn.Ping()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := dbConn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	e := echo.New()
	middL := _categoryHttpDeliveryMiddleware.InitMiddleware()
	e.Use(middL.CORS)

	categoryRepo := _categoryRepo.NewMysqlCategoryRepository(dbConn)

	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second
	au := _categoryUcase.NewCategoryUsecase(categoryRepo, timeoutContext)
	_categoryHttpDelivery.NewCategoryHandler(e, au)

	_ = e.Start(viper.GetString("server.address"))
}
