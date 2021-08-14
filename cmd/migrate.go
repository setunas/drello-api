package cmd

import (
	"context"
	"drello-api/pkg/infrastracture/mysql"
	"log"

	"github.com/spf13/cobra"
)

// migrateCmd represents the migrate command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Run database migration",
	Long: `migrate command makes the database schema aligned with
the schema objects defined in 'ent/migrate/schema.go'.`,
	Run: func(cmd *cobra.Command, args []string) {
		client, err := mysql.Open()
		if err != nil {
			log.Println(err)
		}

		ctx := context.TODO()
		if err := client.Schema.Create(ctx); err != nil {
			log.Fatalf("failed creating schema resources: %v", err)
		}

		log.Println("Successfully created schema resources.")
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}
