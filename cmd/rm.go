package main

import (
	"github.com/spf13/cobra"

	"github.com/safe-k/crast"
)

func rmCommand(locker *crast.Locker, dir string) *cobra.Command {
	cmd := &cobra.Command{
		Use:  "rm",
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			list, listDir := locker.Lists.Get(dir)
			if list == nil {
				list = &crast.List{}
			}

			list.Remove(crast.TaskID(args[0]))
			locker.SaveList(list, listDir)
		},
	}

	return cmd
}
