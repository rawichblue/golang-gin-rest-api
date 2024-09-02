package cmd

import (
	"app/models"
	"app/modules"
	"context"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var employee = &cobra.Command{
	Use:     "migrate-employee",
	Aliases: []string{"addition"},
	Short:   "Start server",
	Run: func(cmd *cobra.Command, args []string) {

		log.Println("Migrating employee")
		db := modules.Get().DB
		if _, err := db.NewCreateTable().Model((*models.Employee)(nil)).Exec(context.Background()); err != nil {
			log.Panic(err)
			os.Exit(1)
			return
		}
		log.Println("Model employee up success")
	},
}

func init() {
	rootCmd.AddCommand(employee)
}
