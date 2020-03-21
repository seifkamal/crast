package main

import (
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"

	"github.com/safe-k/crast"
)

func clearCommand(locker *crast.Locker, dir string) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "clear",
		Short:   "Removes all tasks from the list",
		Long:    "Removes all tasks from the current directory list",
		Example: "crast clear",
		Version: "1.0.0",
		Args:    cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			list, listDir := locker.Lists.Get(dir)
			prompt := promptui.Prompt{
				Label:     "Remove all tasks under " + listDir,
				IsConfirm: true,
			}

			_, err := prompt.Run()
			if err != nil {
				cmd.PrintErrln(err)
				return
			}

			list.Clear()
			locker.SaveList(list, listDir)
		},
	}

	return cmd
}
