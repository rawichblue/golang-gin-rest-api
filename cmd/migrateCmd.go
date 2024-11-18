package cmd

import (
	"app/config/migrations"
	"app/modules"
	"log"

	// "app/modules/log"
	"context"
	"os"

	"github.com/spf13/cobra"
)

// Migrate Command
func Migrate() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "db",
		Args: NotReqArgs,
	}
	cmd.AddCommand(migrateStructUp())
	cmd.AddCommand(migrateStructDown())
	cmd.AddCommand(migrateStructRefresh())
	return cmd
}

func migrateStructUp() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "up",
		Args: NotReqArgs,
		Run: func(cmd *cobra.Command, args []string) {
			log.Printf("Executing model up...")
			// log.Printf("Executing model up...")
			db := modules.Get().DB
			for _, ent := range migrations.Entities() {
				if _, err := db.NewCreateTable().Model(ent).Exec(context.Background()); err != nil {
					log.Printf("%s", err)
					os.Exit(1)
					return
				}
			}
			log.Printf("model up success.")
		},
	}
	return cmd
}

func migrateStructDown() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "down",
		Args: NotReqArgs,
		Run: func(cmd *cobra.Command, args []string) {
			log.Printf("Executing model down...")
			db := modules.Get().DB
			for _, ent := range migrations.Entities() {
				if _, err := db.NewDropTable().Model(ent).Exec(context.Background()); err != nil {
					log.Printf("%s", err)
					os.Exit(1)
					return
				}
			}
			log.Printf("model down success.")

		},
	}
	return cmd
}

func migrateStructRefresh() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "refresh",
		Args: NotReqArgs,
		Run: func(cmd *cobra.Command, args []string) {
			log.Printf("Executing model refresh...")

			db := modules.Get().DB
			for _, ent := range migrations.Entities() {
				if _, err := db.NewDropTable().Model(ent).Exec(context.Background()); err != nil {
					log.Printf("%s", err)
					os.Exit(1)
					return
				}
			}

			for _, ent := range migrations.Entities() {
				if _, err := db.NewCreateTable().Model(ent).Exec(context.Background()); err != nil {
					log.Printf("%s", err)
					os.Exit(1)
					return
				}
			}

			log.Printf("model refresh success.")

		},
	}
	return cmd
}

// func initCMD() *cobra.Command {
// 	cmd := &cobra.Command{
// 		Use:  "init",
// 		Long: "create migration tables",
// 		Run: func(cmd *cobra.Command, args []string) {
// 			migrator := migrate.NewMigrator(modules.Get().DB.Svc.DB(), migrations.Migrations)
// 			migrator.Init(cmd.Context())
// 		},
// 	}
// 	return cmd
// }

// func createSQL() *cobra.Command {
// 	cmd := &cobra.Command{
// 		Use:  "create_sql",
// 		Long: "create up and down SQL migrations",
// 		Run: func(cmd *cobra.Command, args []string) {
// 			migrator := migrate.NewMigrator(modules.Get().DB.Svc.DB(), migrations.Migrations)
// 			name := strings.Join(args, "_")
// 			files, err := migrator.CreateSQLMigrations(cmd.Context(), name)
// 			if err != nil {
// 				log.Printf(err.Printf())
// 				return
// 			}
// 			for _, mf := range files {
// 				fmt.Printf("created migration %s (%s)\n", mf.Name, mf.Path)
// 			}
// 		},
// 	}
// 	return cmd
// }
// func createGO() *cobra.Command {
// 	cmd := &cobra.Command{
// 		Use:  "create_go",
// 		Long: "create Go migration",
// 		Run: func(cmd *cobra.Command, args []string) {
// 			migrator := migrate.NewMigrator(modules.Get().DB.Svc.DB(), migrations.Migrations)
// 			name := strings.Join(args, "_")
// 			mf, err := migrator.CreateGoMigration(cmd.Context(), name)
// 			if err != nil {
// 				log.Printf(err.Printf())
// 				return
// 			}
// 			fmt.Printf("created migration %s (%s)\n", mf.Name, mf.Path)
// 		},
// 	}
// 	return cmd
// }
// func migrateCMD() *cobra.Command {
// 	cmd := &cobra.Command{
// 		Use: "migrate",
// 		Run: func(cmd *cobra.Command, args []string) {
// 			migrator := migrate.NewMigrator(modules.Get().DB.Svc.DB(), migrations.Migrations)
// 			if err := migrator.Lock(cmd.Context()); err != nil {
// 				log.Printf(err.Printf())
// 				return
// 			}
// 			defer migrator.Unlock(cmd.Context()) //nolint:errcheck

// 			group, err := migrator.Migrate(cmd.Context())
// 			if err != nil {
// 				log.Printf(err.Printf())
// 				return
// 			}
// 			if group.IsZero() {
// 				fmt.Printf("there are no new migrations to run (database is up to date)\n")
// 				return
// 			}
// 			fmt.Printf("migrated to %s\n", group)
// 			return
// 		},
// 	}
// 	return cmd
// }
// func rollbackCMD() *cobra.Command {
// 	cmd := &cobra.Command{
// 		Use: "rollback",
// 		Run: func(cmd *cobra.Command, args []string) {
// 			migrator := migrate.NewMigrator(modules.Get().DB.Svc.DB(), migrations.Migrations)
// 			group, err := migrator.Rollback(cmd.Context())
// 			if err != nil {
// 				log.Printf(err.Printf())
// 				return
// 			}

// 			if group.ID == 0 {
// 				fmt.Printf("there are no groups to roll back\n")
// 				return
// 			}

// 			fmt.Printf("rolled back %s\n", group)
// 			return
// 		},
// 	}
// 	return cmd
// }

// func statusCMD() *cobra.Command {
// 	cmd := &cobra.Command{
// 		Use:  "status",
// 		Long: "print migrations status",
// 		Run: func(cmd *cobra.Command, args []string) {
// 			migrator := migrate.NewMigrator(modules.Get().DB.Svc.DB(), migrations.Migrations)
// 			ms, err := migrator.MigrationsWithStatus(cmd.Context())
// 			if err != nil {
// 				log.Printf(err.Printf())
// 				return
// 			}
// 			fmt.Printf("migrations: %s\n", ms)
// 			fmt.Printf("unapplied migrations: %s\n", ms.Unapplied())
// 			fmt.Printf("last migration group: %s\n", ms.LastGroup())
// 			return
// 		},
// 	}
// 	return cmd
// }

// func markAppliedCMD() *cobra.Command {
// 	cmd := &cobra.Command{
// 		Use:  "mark_applied",
// 		Long: "mark migrations as applied without actually running them",
// 		Run: func(cmd *cobra.Command, args []string) {
// 			migrator := migrate.NewMigrator(modules.Get().DB.Svc.DB(), migrations.Migrations)
// 			group, err := migrator.Migrate(cmd.Context(), migrate.WithNopMigration())
// 			if err != nil {
// 				log.Printf(err.Printf())
// 				return
// 			}
// 			if group.IsZero() {
// 				fmt.Printf("there are no new migrations to mark as applied\n")
// 				return
// 			}
// 			fmt.Printf("marked as applied %s\n", group)
// 			return
// 		},
// 	}
// 	return cmd
// }

// func migrateSeed() *cobra.Command {
// 	cmd := &cobra.Command{
// 		Use:  "seed",
// 		Args: NotReqArgs,
// 		Run: func(cmd *cobra.Command, args []string) {
// 			log.Printf("Executing model seed...")

// 			mod := modules.Get()
// 			ctx := context.Background()

// 			startTime := time.Now()
// 			if err := seeds.Seeds(mod, ctx); err != nil {
// 				log.Printf("%s", err)
// 				os.Exit(1)
// 				return
// 			}
// 			endTime := time.Now()
// 			duration := endTime.Sub(startTime)

// 			log.Printf("model seed success.")
// 			log.Printf("Model seed success. Duration: %s", duration)

// 		},
// 	}
// 	return cmd
// }
