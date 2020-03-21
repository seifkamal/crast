package main

import (
	"github.com/spf13/cobra"

	"github.com/safe-k/crast"
)

func deleteCommand(locker *crast.Locker, dir string) *cobra.Command {
	cmd := &cobra.Command{
		Use: "delete",
		Run: func(cmd *cobra.Command, args []string) {
			_, listDir := locker.Lists.Get(dir)
			locker.RemoveList(listDir)
		},
	}

	return cmd
}
