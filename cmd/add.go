package main

import (
	"github.com/spf13/cobra"

	"github.com/safe-k/crast"
)

func addCommand(locker *crast.Locker, dir string) *cobra.Command {
	cmd := &cobra.Command{
		Use:  "add",
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			list := locker.Lists.Get(dir)
			if list == nil {
				list = &crast.List{}
			}

			topic := cmd.Flag("topic").Value.String()
			task := &crast.Task{
				Topic:   topic,
				Summary: args[0],
			}

			list.Add(task)
			locker.Save(*list, dir)
		},
	}

	cmd.Flags().StringP("topic", "t", "general", "A topic to put this task under")

	return cmd
}
