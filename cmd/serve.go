package cmd

import (
	"drello-api/pkg/infrastracture/mysql"
	"drello-api/pkg/presentation/rest"
	"log"

	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the API server",
	Long:  `Start the API server`,
	Run: func(cmd *cobra.Command, args []string) {
		_, err := mysql.Open()
		if err != nil {
			log.Println(err)
		}

		rest.HandleRequests()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
