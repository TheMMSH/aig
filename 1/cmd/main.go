package main

import (
	"aig/1/pkg"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/spf13/viper"
	"log"
	"os"
)

func main() {
	setUpViper()

	ctx := context.Background()
	conn, err := pgx.Connect(ctx, viper.GetString("postgresSource"))
	if err != nil {
		log.Fatalf("cannot connect to postgres due to %s\n", err.Error())
	}

	router := gin.New()
	service := pkg.NewService(conn)
	pkg.NewController(service).SetRoutes(router)
	router.Run(":" + viper.GetString("gin.port"))
}

func setUpViper() {
	viper.SetConfigName(getEnv("CONFIG_NAME", "config"))
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./1/conf")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
