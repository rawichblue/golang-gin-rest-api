package cmd

import (
	"app/models"
	"app/modules"
	"context"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var permission = &cobra.Command{
	Use:     "migrate-permission",
	Aliases: []string{"addition"},
	Short:   "Start server",
	Run: func(cmd *cobra.Command, args []string) {

		log.Println("Migrating permission")
		db := modules.Get().DB
		if _, err := db.NewCreateTable().Model((*models.Permission)(nil)).Exec(context.Background()); err != nil {
			log.Panic(err)
			os.Exit(1)
			return
		}
		log.Println("Model permission up success")
	},
}

func init() {
	rootCmd.AddCommand(permission)
}
