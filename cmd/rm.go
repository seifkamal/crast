package main

import (
	"strconv"

	"github.com/spf13/cobra"

	"github.com/safe-k/crast"
)

func rmCommand(locker *crast.Locker, dir string) *cobra.Command {
	cmd := &cobra.Command{
		Use:  "rm",
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			list := locker.Lists.Get(dir)
			if list == nil {
				list = &crast.List{}
			}

			id, err := strconv.Atoi(args[0])
			if err != nil {
				cmd.PrintErrln(err)
			}

			list.Remove(id)
			locker.Save(*list, dir)
		},
	}

	return cmd
}
