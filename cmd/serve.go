package cmd

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/go-to/bcrd_backend/model"
	"github.com/go-to/bcrd_backend/repository"
	"github.com/go-to/bcrd_backend/router"
	"github.com/go-to/bcrd_backend/usecase"
	"github.com/go-to/bcrd_backend/util"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the gRPC server",
	Long:  `Start the BCRD gRPC server with the configuration from environment variables.`,
	Run: func(cmd *cobra.Command, args []string) {
		startServer()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}

func startServer() {
	// 環境変数読み込み
	loadEnv()

	// タイムゾーン
	locationName := os.Getenv("LOCATION_NAME")
	util.Init(locationName)

	// DB接続
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbSchema := os.Getenv("DB_SCHEMA")
	dbName := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s search_path=%s port=%s sslmode=disable TimeZone=%s", dbHost, dbUser, dbPass, dbName, dbSchema, dbPort, locationName)
	sqlDebug := os.Getenv("SQL_DEBUG")
	dbConn, err := model.Init(dsn, sqlDebug)
	if err != nil {
		log.Fatal(err)
	}

	// model設定
	configModel := model.NewConfigModel(dbConn)
	eventModel := model.NewEventModel(dbConn)
	shopModel := model.NewShopModel(dbConn)
	stampModel := model.NewStampModel(dbConn)

	// repository設定
	configRepository := repository.NewConfigRepository(*configModel)
	eventRepository := repository.NewEventRepository(*eventModel)
	shopRepository := repository.NewShopRepository(*shopModel)
	stampRepository := repository.NewStampRepository(*stampModel)

	// usecase設定
	shopUsecase := usecase.NewShopUseCase(
		*configRepository,
		*eventRepository,
		*shopRepository,
	)
	stampUsecase := usecase.NewStampUseCase(
		*configRepository,
		*stampRepository,
	)

	apiPortStr := os.Getenv("API_PORT")
	apiPort, err := strconv.Atoi(apiPortStr)
	if err != nil {
		log.Fatal(err)
	}

	u := router.Usecase{
		Shop:  usecase.IShopUsecase(shopUsecase),
		Stamp: usecase.IStampUsecase(stampUsecase),
	}
	router.New(apiPort, u)
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
