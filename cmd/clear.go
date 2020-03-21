package main

import (
	"github.com/spf13/cobra"

	"github.com/safe-k/crast"
)

func clearCommand(locker *crast.Locker, dir string) *cobra.Command {
	cmd := &cobra.Command{
		Use:  "clear",
		Args: cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			list, listDir := locker.Lists.Get(dir)
			list.Clear()
			locker.SaveList(list, listDir)
		},
	}

	return cmd
}
