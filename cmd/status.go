package main

import (
	"os"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

func init() {
	cmd := &cobra.Command{
		Use:     "status",
		Short:   "Prints the current list status",
		Long:    "Prints information about the current directory list and its contents",
		Example: "crast status -t my-project -a",
		Version: "1.0.0",
		Args:    cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			if len(locker.Lists) == 0 {
				cmd.Println("No lists initialised")
				return
			}

			interactive, err := cmd.Flags().GetBool("interactive")
			if err != nil {
				cmd.PrintErrln(err)
				return
			}

			selectedDir := dir
			if interactive {
				prompt := promptui.Select{
					Label: "Select directory list to show status for",
					Items: locker.Dirs(),
				}

				_, selectedDir, err = prompt.Run()
				if err != nil {
					cmd.PrintErrln(err)
					return
				}
			}

			list, listDir := locker.Lists.Get(selectedDir)
			if len(*list) == 0 {
				cmd.Println("List", selectedDir, "is empty")
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
	cmd.Flags().BoolP("interactive", "i", false, "Interactively select which list to show status for")

	mainCmd.AddCommand(cmd)
}
