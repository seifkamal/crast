package main

import (
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"

	"github.com/safe-k/crast"
)

func init() {
	cmd := &cobra.Command{
		Use:     "move [taskId...]",
		Short:   "Moves task(s) to a different list",
		Long:    "Moves one or more tasks to a different directory list. The target list selection will be collected via an interactive prompt",
		Example: "crast move GeRW8krZgz G6RZUkrZgm",
		Version: "1.0.0",
		Args:    cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			dirs := []string{}
			for _, d := range locker.Dirs() {
				if d == dir {
					continue
				}
				dirs = append(dirs, d)
			}

			prompt := promptui.Select{
				Label: "Select directory list to move item(s) to",
				Items: dirs,
			}

			_, newDir, err := prompt.Run()
			if err != nil {
				cmd.PrintErrln(err)
				return
			}

			currentList, currentDir := locker.Lists.Get(dir)
			newList, _ := locker.Lists.Get(newDir)

			for _, strID := range args {
				id := crast.TaskID(strID)
				task := currentList.Get(id)
				currentList.Remove(id)
				newList.Add(task)
			}

			locker.SaveList(currentList, currentDir)
			locker.SaveList(newList, newDir)
		},
	}

	mainCmd.AddCommand(cmd)
}
