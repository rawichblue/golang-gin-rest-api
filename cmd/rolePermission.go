package cmd

import (
	"app/models"
	"app/modules"
	"context"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var rolePermission = &cobra.Command{
	Use:     "migrate-rolePermission",
	Aliases: []string{"addition"},
	Short:   "Start server",
	Run: func(cmd *cobra.Command, args []string) {

		log.Println("Migrating rolePermission")
		db := modules.Get().DB
		if _, err := db.NewCreateTable().Model((*models.RolePermission)(nil)).Exec(context.Background()); err != nil {
			log.Panic(err)
			os.Exit(1)
			return
		}
		log.Println("Model rolePermission up success")
	},
}

func init() {
	rootCmd.AddCommand(rolePermission)
}
