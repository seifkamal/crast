package main

import (
	"log"
	"os"

	"github.com/spf13/cobra"

	"github.com/safe-k/crast"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	locker, err := crast.NewLocker()
	if err != nil {
		log.Println(err)
		return
	}

	cmd := &cobra.Command{}
	cmd.AddCommand(initCommand(locker, dir))
	cmd.AddCommand(statusCommand(locker, dir))
	cmd.AddCommand(addCommand(locker, dir))
	cmd.AddCommand(rmCommand(locker, dir))
	cmd.AddCommand(doCommand(locker, dir))

	if err := cmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
