package main

import (
	"log"
	"os"

	"github.com/spf13/cobra"

	"github.com/safe-k/crast"
)

var (
	dir    string
	locker *crast.Locker

	mainCmd = &cobra.Command{
		Use: "crast [command]",
	}
)

func init() {
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}
	dir = currentDir

	locker, err = crast.NewLocker()
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	if err := mainCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
