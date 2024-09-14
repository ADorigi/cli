package utils

import (
	"strconv"

	"github.com/spf13/cobra"
)

func ReadBoolFlag(cmd *cobra.Command, name string) bool {
	flag := cmd.Flags().Lookup(name)
	if flag == nil {
		return false
	}
	value, err := cmd.Flags().GetBool(name)
	if err != nil {
		return false
	}
	return value
}

func ReadStringArrayFlag(cmd *cobra.Command, name string) ([]string, error) {
	stringArray, err := cmd.Flags().GetStringArray(name)
	if err != nil {
		return nil, err
	}
	return stringArray, nil
}

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
