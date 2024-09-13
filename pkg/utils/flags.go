package utils

import (
	"strconv"

	"github.com/spf13/cobra"
)

func ReadStringFlag(cmd *cobra.Command, name string) string {
	if cmd.Flags().Lookup(name) == nil {
		return ""
	} else {
		return cmd.Flags().Lookup(name).Value.String()
	}
}

func ReadIntFlag(cmd *cobra.Command, name string) int64 {
	str := ReadStringFlag(cmd, name)
	i, _ := strconv.ParseInt(str, 10, 64)
	return i
}
