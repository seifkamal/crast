package main

import (
	"github.com/spf13/cobra"
	"github.com/teris-io/shortid"

	"github.com/safe-k/crast"
)

func addCommand(locker *crast.Locker, dir string) *cobra.Command {
	cmd := &cobra.Command{
		Use:  "add",
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			list, listDir := locker.Lists.Get(dir)
			topic := cmd.Flag("topic").Value.String()

			priorityLevel, err := cmd.Flags().GetInt("priority")
			if err != nil {
				cmd.PrintErrln(err)
				return
			}

			priority := crast.Priority(priorityLevel)
			if valid := priority.IsValid(); !valid {
				cmd.PrintErrln("Invalid priority level")
				return
			}

			for _, summary := range args {
				id, err := shortid.Generate()
				if err != nil {
					cmd.PrintErrln(err)
					return
				}

				list.Add(&crast.Task{
					ID:       crast.TaskID(id),
					Topic:    topic,
					Summary:  summary,
					Priority: priority,
				})
			}

			locker.SaveList(list, listDir)
		},
	}

	cmd.Flags().StringP("topic", "t", "general", "A topic to put this task under")
	cmd.Flags().IntP("priority", "p", 4, "The priority level of this task")

	return cmd
}
