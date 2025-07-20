package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/ndkariyasena/gobolt-view/server"
)

var rootCmd = &cobra.Command{
	Use:   "boltdb-webui",
	Short: "A lightweight BoltDB UI built with Go and Svelte",
	Long: `BoltDB Web UI is a modern, single-binary application to browse
and inspect your BoltDB files through a secure and responsive web interface.`,
}

// startCmd launches the embedded web server (API + UI)
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the web UI server",
	Run: func(cmd *cobra.Command, args []string) {
		server.Start()
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}

// Execute runs the CLI
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
