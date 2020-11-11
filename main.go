package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/labstack/echo"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"

	_categoryHttpDelivery "github.com/zarszz/warehouse-rest-api/category/delivery/http"
	_categoryRepo "github.com/zarszz/warehouse-rest-api/category/repository/mysql"
	_categoryUcase "github.com/zarszz/warehouse-rest-api/category/usecase"
	_middleware "github.com/zarszz/warehouse-rest-api/middleware"

	_userHttpDelivery "github.com/zarszz/warehouse-rest-api/user/delivery"
	_userRepo "github.com/zarszz/warehouse-rest-api/user/repository/mysql"
	_userUcase "github.com/zarszz/warehouse-rest-api/user/usecase"

	_warehouseHttpDelivery "github.com/zarszz/warehouse-rest-api/warehouse/delivery"
	_warehouseRepo "github.com/zarszz/warehouse-rest-api/warehouse/repository/mysql"
	_warehouseUcase "github.com/zarszz/warehouse-rest-api/warehouse/usecase"

	_roomHttpDelivery "github.com/zarszz/warehouse-rest-api/room/delivery"
	_roomRepo "github.com/zarszz/warehouse-rest-api/room/repository/mysql"
	_roomUcase "github.com/zarszz/warehouse-rest-api/room/usecase"

	_itemHttpDelivery "github.com/zarszz/warehouse-rest-api/item/delivery"
	_itemRepo "github.com/zarszz/warehouse-rest-api/item/repository/mysql"
	_itemUcase "github.com/zarszz/warehouse-rest-api/item/usecase"

	_rackHttpDelivery "github.com/zarszz/warehouse-rest-api/rack/delivery/http"
	_rackRepo "github.com/zarszz/warehouse-rest-api/rack/repository/mysql"
	_rackUcase "github.com/zarszz/warehouse-rest-api/rack/usecase"
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
	dbHost := viper.GetString(`database_prod.host`)
	dbUser := viper.GetString(`database_prod.user`)
	dbPass := viper.GetString(`database_prod.pass`)
	dbName := viper.GetString(`database_prod.name`)
	connection := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable&", dbUser, dbPass, dbHost, dbName)
	fmt.Println(connection)
	val := url.Values{}
	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())
	fmt.Println(connection)
	dbConn, err := sql.Open(`postgres`, dsn)

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
	middL := _middleware.InitMiddleware()
	e.Use(middL.CORS)

	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second

	categoryRepo := _categoryRepo.NewMysqlCategoryRepository(dbConn)
	categoryUsecase := _categoryUcase.NewCategoryUsecase(categoryRepo, timeoutContext)
	_categoryHttpDelivery.NewCategoryHandler(e, categoryUsecase)

	userRepo := _userRepo.NewMysqlUserRepository(dbConn)
	au := _userUcase.NewUserUsecase(userRepo, timeoutContext)
	_userHttpDelivery.NewUserHandler(e, au)

	userAddressRepo := _userRepo.NewMysqlUserAddressRepository(dbConn)
	userAddressUsecase := _userUcase.NewUserAddressUsecase(userAddressRepo, timeoutContext)
	_userHttpDelivery.NewUserAddressHandler(e, userAddressUsecase)

	warehouseAddressRepo := _warehouseRepo.NewMysqlWarehouseAddressRepository(dbConn)
	warehouseAddressUsecase := _warehouseUcase.NewWarehouseAddressUsecase(warehouseAddressRepo, timeoutContext)
	_warehouseHttpDelivery.NewUserAddressHandler(e, warehouseAddressUsecase)

	roomRepo := _roomRepo.NewMysqlWarehouseRepository(dbConn)
	roomUsecase := _roomUcase.NewRoomeUsecase(roomRepo, timeoutContext)
	_roomHttpDelivery.NewRoomHandler(e, roomUsecase)

	rackRepo := _rackRepo.NewMysqlRackRepository(dbConn)
	rackUsecase := _rackUcase.NewRackUsecase(rackRepo, timeoutContext)
	_rackHttpDelivery.NewRackHandler(e, rackUsecase)

	itemRepo := _itemRepo.NewMysqlItemRepository(dbConn)
	itemUsecase := _itemUcase.NewItemUsecase(itemRepo, timeoutContext)
	_itemHttpDelivery.NewItemHandler(e, itemUsecase)

	warehouseRepo := _warehouseRepo.NewMysqlWarehouseRepository(dbConn)
	warehouseUsecase := _warehouseUcase.NewWarehouseUsecase(warehouseRepo, timeoutContext)
	_warehouseHttpDelivery.NewWarehouseHandler(e, warehouseUsecase, roomUsecase, rackUsecase, itemUsecase)

	_ = e.Start(viper.GetString("server.address"))
}
