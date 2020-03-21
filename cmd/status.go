package main

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/safe-k/crast"
)

func statusCommand(locker *crast.Locker, dir string) *cobra.Command {
	cmd := &cobra.Command{
		Use:  "status",
		Args: cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			list, listDir := locker.Lists.Get(dir)
			if len(*list) == 0 {
				cmd.Println("List", dir, "is empty")
				return
			}

			topic := cmd.Flag("topic").Value.String()
			showDone, err := cmd.Flags().GetBool("all")
			if err != nil {
				cmd.PrintErrln(err)
				return
			}

			table := list.Table(os.Stdout, topic, showDone)
			table.SetCaption(true, "Dir: "+listDir)
			table.Render()
		},
	}

	cmd.Flags().StringP("topic", "t", "", "Filter tasks by topic")
	cmd.Flags().BoolP("all", "a", false, "Show all tasks (including ones marked as done)")

	return cmd
}
