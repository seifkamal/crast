package crast

import (
	"fmt"
	"io"

	"github.com/olekukonko/tablewriter"
)

type row []string

const (
	doneChar string = "✔"
	todoChar string = " "
)

// Table returns a tablewriter.Table representation of the attached List.
func (l List) Table(writer io.Writer, topic string, showDone bool) *tablewriter.Table {
	table := tablewriter.NewWriter(writer)
	table.SetHeader([]string{"ID", doneChar, "P", "Topic", "Summary"})
	table.SetCenterSeparator("-")

	for _, task := range l.ByPriority() {
		if (!showDone && task.Done) || (topic != "" && topic != task.Topic) {
			continue
		}

		table.Append(row{
			fmt.Sprintf("%v", task.ID),
			fmt.Sprintf("[%s]", stateChar(task.Done)),
			fmt.Sprintf("%v", task.Priority),
			task.Topic,
			task.Summary,
		})
	}

	table.SetHeaderColor(
		tablewriter.Colors{tablewriter.Bold, tablewriter.BgHiBlackColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.BgGreenColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiYellowColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.BgHiBlackColor, tablewriter.FgWhiteColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.BgCyanColor, tablewriter.FgWhiteColor},
	)
	table.SetColumnColor(
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiBlackColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiBlackColor},
		tablewriter.Colors{tablewriter.FgHiYellowColor},
		tablewriter.Colors{tablewriter.FgHiBlackColor},
		tablewriter.Colors{},
	)

	return table
}

func stateChar(done bool) string {
	switch done {
	case true:
		return doneChar
	default:
		return todoChar
	}
}
