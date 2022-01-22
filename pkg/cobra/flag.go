package cobra

import "github.com/spf13/cobra"

func GetInt(cmd *cobra.Command, flag string) (i int, err error) {
	if cmd.Flag(flag) != nil && cmd.Flag(flag).Changed {
		i, err = cmd.Flags().GetInt(flag)
	}
	return i, err
}

func GetBool(cmd *cobra.Command, flag string) (s bool, err error) {
	if cmd.Flag(flag) != nil && cmd.Flag(flag).Changed {
		s, err = cmd.Flags().GetBool(flag)
	}
	return s, err
}

func GetString(cmd *cobra.Command, flag string) (s string, err error) {
	if cmd.Flag(flag) != nil && cmd.Flag(flag).Changed {
		s, err = cmd.Flags().GetString(flag)
	}
	return s, err
}

func GetStringSlice(cmd *cobra.Command, flag string) (s []string, err error) {
	if cmd.Flag(flag) != nil && cmd.Flag(flag).Changed {
		s, err = cmd.Flags().GetStringSlice(flag)
	}
	return s, err
}
