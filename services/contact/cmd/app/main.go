package main

import (
	"fmt"
	"github.com/ChristinaFomenko/slurm-clean-architecture/pkg/store/postgres"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"

	deliveryGrpc "github.com/ChristinaFomenko/slurm-clean-architecture/services/contact/internal/delivery/grpc"
	deliveryHttp "github.com/ChristinaFomenko/slurm-clean-architecture/services/contact/internal/delivery/http"
	repositoryStorage "github.com/ChristinaFomenko/slurm-clean-architecture/services/contact/internal/repository/storage/postgres"
	useCaseContact "github.com/ChristinaFomenko/slurm-clean-architecture/services/contact/internal/useCase/contact"
	useCaseGroup "github.com/ChristinaFomenko/slurm-clean-architecture/services/contact/internal/useCase/group"
)

func main() {
	conn, err := postgres.New(postgres.Settings{})
	if err != nil {
		panic(err)
	}

	defer conn.Pool.Close()

	fmt.Println(conn.Pool.Stat())

	var (
		repoStorage  = repositoryStorage.New(conn.Pool, repositoryStorage.Options{})
		ucContact    = useCaseContact.New(repoStorage, useCaseContact.Options{})
		ucGroup      = useCaseGroup.New(repoStorage, useCaseGroup.Options{})
		_            = deliveryGrpc.New(ucContact, ucGroup, deliveryGrpc.Options{})
		listenerHttp = deliveryHttp.New(ucContact, ucGroup, deliveryHttp.Options{})
	)

	go func() {
		fmt.Printf("service started successfully on http port: %d", viper.GetUint("HTTP_PORT"))
		if err = listenerHttp.Run(); err != nil {
			panic(err)
		}
	}()

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)
	<-signalCh

	fmt.Println("Hello World!")
}
