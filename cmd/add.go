package main

import (
	"github.com/spf13/cobra"

	"github.com/safe-k/crast"
)

func addCommand(locker *crast.Locker, dir string) *cobra.Command {
	cmd := &cobra.Command{
		Use:  "add",
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			list, listDir := locker.Lists.Get(dir)
			topic := cmd.Flag("topic").Value.String()

			for _, summary := range args {
				list.Add(&crast.Task{
					Topic:   topic,
					Summary: summary,
				})
			}

			locker.SaveList(list, listDir)
		},
	}

	cmd.Flags().StringP("topic", "t", "general", "A topic to put this task under")

	return cmd
}
