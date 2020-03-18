package main

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/safe-k/crast"
)

func viewCommand(locker *crast.Locker, dir string) *cobra.Command {
	cmd := &cobra.Command{
		Use: "view",
		Run: func(cmd *cobra.Command, args []string) {
			list := locker.Lists.Get(dir)
			if list == nil {
				cmd.Printf("List for %s is empty", dir)
				return
			}

			topic := cmd.Flag("topic").Value.String()
			showDone, err := cmd.Flags().GetBool("all")
			if err != nil {
				cmd.PrintErrln(err)
				return
			}

			output := "Dir: " + dir + "\n"
			for id, task := range *list {
				if topic == "" || task.Topic == topic {
					state := " "
					if task.Done {
						if !showDone {
							continue
						}
						state = "X"
					}

					output += fmt.Sprintf("%v - [%s] (%s) %s\n", strconv.Itoa(id), state, task.Topic, task.Summary)
				}
			}

			cmd.Println(output)
		},
	}

	cmd.Flags().StringP("topic", "t", "", "Filter tasks by topic")
	cmd.Flags().BoolP("all", "a", false, "Show all tasks (including ones marked as done)")

	return cmd
}
