package cmd

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	verbosity   string
	configFile  string
	interactive = false
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gsync",
	Short: "gsync is a tool to handle easily project composed of multiple git repositories",
	Long:  `gsync let you easily clean up multiple git repositories with a simple command.`,
}

//NewGsyncCommand initiate the gsync root command
func NewGsyncCommand() *cobra.Command {

	rootCmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		if err := setUpLogs(os.Stdout, verbosity); err != nil {
			return err
		}
		return nil
	}

	rootCmd.AddCommand(NewCleanCommand())

	rootCmd.PersistentFlags().StringVarP(&verbosity, "verbosity", "v", logrus.InfoLevel.String(), "Log level (debug, info, warn, error, fatal, panic")
	rootCmd.PersistentFlags().StringVarP(&configFile, "configFile", "f", "gsync.yml", "Gsync file to use.")

	return rootCmd
}

func setUpLogs(out io.Writer, level string) error {

	logrus.SetOutput(out)
	lvl, err := logrus.ParseLevel(level)
	if err != nil {
		return err
	}
	logrus.SetLevel(lvl)
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
	return nil
}

func addInteractiveFlag(cmd *cobra.Command) {
	cmd.Flags().BoolVarP(&interactive, "interactive", "i",false, "active interactive mode")
}
