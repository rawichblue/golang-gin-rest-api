package cmd

import (
	"app/models"
	"app/modules"
	"context"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var customer = &cobra.Command{
	Use:     "migrate-customer",
	Aliases: []string{"addition"},
	Short:   "Start server",
	Run: func(cmd *cobra.Command, args []string) {

		log.Println("Migrating customers")
		db := modules.Get().DB
		if _, err := db.NewCreateTable().Model((*models.Customer)(nil)).Exec(context.Background()); err != nil {
			log.Panic(err)
			os.Exit(1)
			return
		}
		log.Println("Model customers up success")
	},
}

func init() {
	rootCmd.AddCommand(customer)
}
