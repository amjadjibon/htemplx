package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"htemplx/migrations"
)

var migrateCmd = &cobra.Command{
	Use: "migrate",
	Run: func(cmd *cobra.Command, _ []string) {
		_ = cmd.Help()
		os.Exit(0)
	},
}

var migrateUpCmd = &cobra.Command{
	Use: "up",
	Run: func(cmd *cobra.Command, _ []string) {
		var dbURL string
		if dbURL = cmd.Parent().PersistentFlags().Lookup("db_url").Value.String(); dbURL == "" {
			if dbURL = os.Getenv("DB_URL"); dbURL == "" {
				fmt.Println("db_url is not provided")
				os.Exit(1)
			}
		}
		if err := migrations.Up(dbURL); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

var migrateDownCmd = &cobra.Command{
	Use: "down",
	Run: func(cmd *cobra.Command, _ []string) {
		var dbURL string
		if dbURL = cmd.Parent().PersistentFlags().Lookup("db_url").Value.String(); dbURL == "" {
			if dbURL = os.Getenv("DB_URL"); dbURL == "" {
				fmt.Println("db_url is not provided")
				os.Exit(1)
			}
		}

		if err := migrations.Down(dbURL); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	migrateCmd.
		PersistentFlags().
		StringP(
			"db_url",
			"d",
			"",
			"postgres://postgres:postgres@localhost:5432/postgres",
		)

	_ = migrateCmd.MarkPersistentFlagRequired("db_url")

	migrateCmd.AddCommand(migrateUpCmd)
	migrateCmd.AddCommand(migrateDownCmd)
}
