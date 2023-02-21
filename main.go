// Package main
package main

import (
	"context"
	"fmt"
	"time"

	"github.com/OVantsevich/PriceProvider/internal/config"
	"github.com/OVantsevich/PriceProvider/internal/repository"
	"github.com/OVantsevich/PriceProvider/internal/service"

	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
)

// maxPrice max price
const maxPrice float32 = 5

// sleepTime max price
const sleepTime = time.Second * 5

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
		time.Sleep(sleepTime)
	}
}
