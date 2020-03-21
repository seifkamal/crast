package main

import (
	"github.com/spf13/cobra"

	"github.com/safe-k/crast"
)

func initCommand(locker *crast.Locker, dir string) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "init",
		Short:   "Initialises a list",
		Long:    "Initialises a list under the current directory",
		Example: "crast init",
		Version: "1.0.0",
		Args:    cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			if exists := locker.Lists.Has(dir); exists {
				cmd.Println("Directory list already exists")
				return
			}

			list := &crast.List{}
			locker.SaveList(list, dir)

			cmd.Println("New list initialised under", dir)
		},
	}

	return cmd
}
