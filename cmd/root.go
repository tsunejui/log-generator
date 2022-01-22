package cmd

import (
	"fmt"
	pkgCobra "log-generator/pkg/cobra"
	pkgLog "log-generator/pkg/log"

	"github.com/spf13/cobra"
)

const (
	CMD_NAME        = "log-generator"
	limitErrorCount = 10
	limitWriteCount = 30
	intervalSec     = 3
	folderName      = "logs"
)

const (
	FLAG_MAX        = "max"
	FLAG_FOLDER     = "folder"
	FLAG_FORMAT     = "format"
	FLAG_INTERVAL   = "interval"
	FLAG_SEPARATION = "separation"
)

// usage: go run cmd/log-generator/main.go --folder="logs" --format=json --separation=true --max=30 --interval=3

func init() {
	rootCmd.Flags().Int(FLAG_MAX, limitWriteCount, "please specify a maximum limit for writing logs")
	rootCmd.Flags().String(FLAG_FOLDER, folderName, "please specify a folder path for storing logs")
	rootCmd.Flags().String(
		FLAG_FORMAT,
		pkgLog.FORMAT_JSON,
		fmt.Sprintf("please specify log format (%s or %s)", pkgLog.FORMAT_TEXT, pkgLog.FORMAT_JSON),
	)
	rootCmd.Flags().Int(FLAG_INTERVAL, intervalSec, "please specify interval time")
	rootCmd.Flags().Bool(FLAG_SEPARATION, false, "please input true or false")
}

var rootCmd = &cobra.Command{
	Use:   CMD_NAME,
	Short: fmt.Sprintf("%s is a simple cli tool to generate logs for testing", CMD_NAME),
	Long:  "A easy way to generate testing logs",
	RunE: func(cmd *cobra.Command, args []string) error {
		max, err := pkgCobra.GetInt(cmd, FLAG_MAX)
		if err != nil {
			return fmt.Errorf("failed to get maximum limit: %v", err)
		}

		folder, err := pkgCobra.GetString(cmd, FLAG_FOLDER)
		if err != nil {
			return fmt.Errorf("failed to get folder: %v", err)
		}

		format, err := pkgCobra.GetString(cmd, FLAG_FORMAT)
		if err != nil {
			return fmt.Errorf("failed to get format: %v", err)
		}
		if format == "" {
			return fmt.Errorf("format is required")
		}

		separation, err := pkgCobra.GetBool(cmd, FLAG_SEPARATION)
		if err != nil {
			return fmt.Errorf("failed to get separation: %v", err)
		}

		interval, err := pkgCobra.GetInt(cmd, FLAG_INTERVAL)
		if err != nil {
			return fmt.Errorf("failed to get separation: %v", err)
		}

		l := pkgLog.New(&pkgLog.Config{
			Interval:        interval,
			Name:            format,
			FilePath:        folder,
			Separation:      separation,
			Source:          format,
			ErrCountLimit:   limitErrorCount,
			WriteCountLimit: max,
		})
		return l.Mock()
	},
}

func Execute() error {
	return rootCmd.Execute()
}
