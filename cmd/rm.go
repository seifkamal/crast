package main

import (
	"github.com/spf13/cobra"

	"github.com/seifkamal/crast"
)

func init() {
	cmd := &cobra.Command{
		Use:     "rm [taskId...]",
		Short:   "Removes task(s)",
		Long:    "Removes one or more tasks from the current directory list",
		Example: "crast rm GeRW8krZgz G6RZUkrZgm",
		Version: "1.0.0",
		Args:    cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			list, listDir := locker.Lists.Get(dir)
			if list == nil {
				list = &crast.List{}
			}

			list.Remove(crast.TaskID(args[0]))
			locker.SaveList(list, listDir)
		},
	}

	mainCmd.AddCommand(cmd)
}
