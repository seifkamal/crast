package main

import (
	"github.com/spf13/cobra"

	"github.com/safe-k/crast"
)

func undoCommand(locker *crast.Locker, dir string) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "undo [taskId...]",
		Short:   "Marks task(s) as todo",
		Long:    "Marks one or more tasks in the current directory list as todo",
		Example: "crast undo GeRW8krZgz G6RZUkrZgm",
		Version: "1.0.0",
		Run: func(cmd *cobra.Command, args []string) {
			list, listDir := locker.Lists.Get(dir)
			if list == nil {
				list = &crast.List{}
			}

			ids := []crast.TaskID{}
			for _, strID := range args {
				ids = append(ids, crast.TaskID(strID))
			}

			list.Undo(ids...)
			locker.SaveList(list, listDir)
		},
	}

	return cmd
}
