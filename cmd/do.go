package main

import (
	"github.com/spf13/cobra"

	"github.com/safe-k/crast"
)

func doCommand(locker *crast.Locker, dir string) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "do [taskId]",
		Short:   "Marks task(s) as done",
		Long:    "Marks one or more tasks in the current directory list as done",
		Example: "crast do GeRW8krZgz G6RZUkrZgm",
		Version: "1.0.0",
		Args:    cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			list, listDir := locker.Lists.Get(dir)
			if list == nil {
				list = &crast.List{}
			}

			ids := []crast.TaskID{}
			for _, strID := range args {
				ids = append(ids, crast.TaskID(strID))
			}

			list.Do(ids...)
			locker.SaveList(list, listDir)
		},
	}

	return cmd
}
