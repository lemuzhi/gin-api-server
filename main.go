package main

import (
	"context"
	"flag"
	"gin-api-server/cmd"
	"github.com/spf13/cobra"
	"log"
)

const version = "v1.0.0"

func main() {
	flag.Parse()

	root := &cobra.Command{
		Use:     "gpt",
		Version: version,
		Short:   "This is a gin project template",
	}
	ctx := context.Background()

	root.AddCommand(cmd.NewServerStartCmd(ctx, version))

	if err := root.Execute(); err != nil {
		log.Println(err)
	}

}
