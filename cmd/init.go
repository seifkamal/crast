package main

import (
	"github.com/spf13/cobra"

	"github.com/safe-k/crast"
)

func initCommand(locker *crast.Locker, dir string) *cobra.Command {
	cmd := &cobra.Command{
		Use:  "init",
		Args: cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			if exists := locker.Lists.Has(dir); exists {
				cmd.Println("Directory list already exists")
				return
			}

			list := &crast.List{}
			locker.SaveList(list, dir)
		},
	}

	return cmd
}
