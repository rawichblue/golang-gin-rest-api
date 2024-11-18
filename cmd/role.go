package cmd

import (
	"app/models"
	"app/modules"
	"context"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var role = &cobra.Command{
	Use:     "migrate-role",
	Aliases: []string{"addition"},
	Short:   "Start server",
	Run: func(cmd *cobra.Command, args []string) {

		log.Println("Migrating role")
		db := modules.Get().DB
		if _, err := db.NewCreateTable().Model((*models.Role)(nil)).Exec(context.Background()); err != nil {
			log.Panic(err)
			os.Exit(1)
			return
		}
		log.Println("Model role up success")
	},
}

func init() {
	rootCmd.AddCommand(role)
}
