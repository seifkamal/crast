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
			list, listDir := locker.Lists.Get(dir)

			topic := cmd.Flag("topic").Value.String()
			task := &crast.Task{
				Topic:   topic,
				Summary: args[0],
			}

			list.Add(task)
			locker.SaveList(list, listDir)
		},
	}

	cmd.Flags().StringP("topic", "t", "general", "A topic to put this task under")

	return cmd
}
