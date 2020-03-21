package main

import (
	"github.com/spf13/cobra"

	"github.com/safe-k/crast"
)

func undoCommand(locker *crast.Locker, dir string) *cobra.Command {
	cmd := &cobra.Command{
		Use: "undo",
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
