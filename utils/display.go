package utils

import (
	"os"

	"github.com/olekukonko/tablewriter"
)

// GetTable return configured tablewriter.Table
func GetTable() *tablewriter.Table {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetBorder(false)
	table.SetRowSeparator("")
	table.SetColumnSeparator("  ")
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetHeaderLine(false)
	return table
}
