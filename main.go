// Package main
package main

import (
	"Price-Provider/internal/config"
	"Price-Provider/internal/repository"
	"Price-Provider/internal/service"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"time"
)

// maxPrice max price
const maxPrice float32 = 100

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		logrus.Fatal(err)
	}

	client := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", cfg.RedisHost, cfg.RedisPort),
	})

	redisRep := repository.NewRedis(client, cfg.StreamName)
	priceService := service.NewPrices(redisRep, maxPrice)

	for {
		priceService.RandPrices()
		err = priceService.PublishPrices(context.Background())
		if err != nil {
			logrus.Fatal(err)
		}
		time.Sleep(time.Second)
	}
}
