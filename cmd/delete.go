package main

import (
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

func init() {
	cmd := &cobra.Command{
		Use:     "delete",
		Short:   "Deletes the list",
		Long:    "Deletes the current directory list. Crast will then fallback to the parent list.",
		Example: "crast delete",
		Version: "1.0.0",
		Args:    cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			_, listDir := locker.Lists.Get(dir)
			prompt := promptui.Prompt{
				Label:     "Delete the list under " + listDir,
				IsConfirm: true,
			}

			_, err := prompt.Run()
			if err != nil {
				cmd.PrintErrln(err)
				return
			}

			locker.RemoveList(listDir)
		},
	}

	mainCmd.AddCommand(cmd)
}
