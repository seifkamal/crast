package main

import (
	"github.com/spf13/cobra"

	"github.com/safe-k/crast"
)

func init() {
	cmd := &cobra.Command{
		Use:     "edit [taskId...]",
		Short:   "Edits existing task(s)",
		Long:    "Edits one or more existing tasks in the current directory list. Task summaries can only be edited for one task at a time",
		Example: "crast edit VvvggirWR G6RZUkrZgm -p 3 -t my-project",
		Version: "1.0.0",
		Args:    cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			list, listDir := locker.Lists.Get(dir)

			summaryFlag := cmd.Flag("summary")
			topicFlag := cmd.Flag("topic")
			priorityFlag := cmd.Flag("priority")

			if len(args) > 1 && summaryFlag.Changed {
				cmd.Println("Summaries can only be edited for one task at a time")
				return
			}

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

			for _, idStr := range args {
				id := crast.TaskID(idStr)
				task := list.Get(id)

				if summaryFlag.Changed {
					if summary := summaryFlag.Value.String(); summary != "" {
						task.Summary = summary
					}
				}
				if topicFlag.Changed {
					task.Topic = topicFlag.Value.String()
				}
				if priorityFlag.Changed {
					task.Priority = priority
				}

				list.Update(task)
			}

			locker.SaveList(list, listDir)
		},
	}

	cmd.Flags().StringP("summary", "s", "", "The task summary")
	cmd.Flags().StringP("topic", "t", "general", "A topic to put this task under")
	cmd.Flags().IntP("priority", "p", 4, "The priority level of this task")

	mainCmd.AddCommand(cmd)
}
