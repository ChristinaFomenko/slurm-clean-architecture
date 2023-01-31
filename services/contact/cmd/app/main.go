package main

import (
	"fmt"
	"github.com/ChristinaFomenko/slurm-clean-architecture/pkg/store/postgres"
)

func main() {
	conn, err := postgres.New(postgres.Settings{})
	if err != nil {
		panic(err)
	}

	defer conn.Pool.Close()

	fmt.Println(conn.Pool.Stat())

	fmt.Println("Hello World!")
}
