package main

import (
	"context"
	"fmt"
	"github.com/mdev5000/runnr"
	"github.com/mdev5000/vectest2/internal/pkg/runnrlocal"
	"github.com/spf13/cobra"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	runner := runnr.NewRunner()

	runner.AddCommand(&cobra.Command{
		Use: "hello",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("hello world")
		},
	})

	runner.AddCommand(runnrlocal.WatchCommand())

	ctx := context.Background()
	return runner.Run(ctx)
}
