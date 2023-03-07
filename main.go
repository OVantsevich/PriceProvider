// Package main
package main

import (
	"context"
	"fmt"
	"net"

	"time"

	"github.com/OVantsevich/PriceProvider/internal/config"
	"github.com/OVantsevich/PriceProvider/internal/handler"
	"github.com/OVantsevich/PriceProvider/internal/repository"
	"github.com/OVantsevich/PriceProvider/internal/service"
	pr "github.com/OVantsevich/PriceProvider/proto"

	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

// maxPrice max price
const maxPrice float64 = 5

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

	go startRand(context.Background(), priceService)
	priceHandler := handler.NewPrice(priceService)

	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.Host, cfg.Port))
	if err != nil {
		defer logrus.Fatalf("error while listening port: %e", err)
	}
	ns := grpc.NewServer()
	pr.RegisterPriceProviderServer(ns, priceHandler)

	if err = ns.Serve(listen); err != nil {
		defer logrus.Fatalf("error while listening server: %e", err)
	}

}

func startRand(ctx context.Context, ps *service.Prices) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			ps.RandPrices()
			err := ps.PublishPrices(context.Background())
			if err != nil {
				logrus.Fatal(err)
			}
			time.Sleep(sleepTime)
		}
	}
}
