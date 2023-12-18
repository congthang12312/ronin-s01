package main

import (
	"context"
	"fmt"
	"log"
	"ronin/cmd/serverd/router"
	"ronin/internal/controller/airport"
	"ronin/internal/repository"
	"ronin/pkg/env"
	"ronin/pkg/redis"
	"ronin/pkg/runner"
	"strconv"

	pkgerrors "github.com/pkg/errors"
)

type controllers struct {
	airportCrl airport.Controller
}

func main() {
	ctx := context.Background()

	if err := run(ctx); err != nil {
		log.Fatal("error when running application ")
	}
	fmt.Print("Exiting...")
}

func run(ctx context.Context) error {
	// TODO: init database in here

	// Init redis
	redisDB, err := strconv.Atoi(env.GetAndValidateF("REDIS_DB"))
	if err != nil {
		return pkgerrors.WithStack(fmt.Errorf("invalid redis db: %w", err))
	}
	redisClient, err := redis.NewClient(ctx, env.GetAndValidateF("REDIS_URL"),
		redisDB,
	)
	if err != nil {
		return err
	}
	defer redisClient.Close()

	// Init Controllers
	ctrls, err := initControllers(redisClient)
	if err != nil {
		return err
	}

	rtr := initRouter(ctrls)

	srv := runner.NewServer(rtr.Handler(),
		env.GetAndValidateF("PORT"),
	)

	runner.ExecParallel(ctx, srv.Start)
	fmt.Println("App initialization completed")
	return nil
}

func initControllers(redisClient *redis.Client) (controllers, error) {
	repo := repository.New(redisClient)
	return controllers{
		airportCrl: airport.New(repo),
	}, nil
}

func initRouter(ctrls controllers) router.Router {
	return router.New(ctrls.airportCrl)
}
