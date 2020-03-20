package main

import (
	"strconv"

	"github.com/spf13/cobra"

	"github.com/safe-k/crast"
)

func doCommand(locker *crast.Locker, dir string) *cobra.Command {
	cmd := &cobra.Command{
		Use:  "do",
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			list, listDir := locker.Lists.Get(dir)
			if list == nil {
				list = &crast.List{}
			}

			ids := []int{}
			for _, strID := range args {
				id, err := strconv.Atoi(strID)
				if err != nil {
					cmd.PrintErrln(err)
					return
				}

				ids = append(ids, id)
			}

			list.Do(ids...)
			locker.SaveList(list, listDir)
		},
	}

	return cmd
}
