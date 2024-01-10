/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	pkg "github.com/kanzifucius/go-postgress-to-s3/pkg"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// backupS3Cmd represents the backupS3 command
var backupS3Cmd = &cobra.Command{
	Use:   "backupS3",
	Short: "back up psotgres to s3",
	Long:  `back up psotgres to s3`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("backupS3 called")

		// Create a new backupPostgres instance
		backupPostgres := pkg.NewBackupPostgres(
			cmd.Flag("postgres-host").Value.String(),
			cmd.Flag("postgres-user").Value.String(),
			cmd.Flag("postgres-password").Value.String(),
			cmd.Flag("postgres-database").Value.String(),
		)

		err := backupPostgres.Backup(cmd.Flag("backup-file").Value.String())
		if err != nil {
			panic(err)

		}

		// Create a new backupS3 instance
		backupS3 := pkg.NewBackupS3(
			cmd.Flag("s3-bucket").Value.String(),
			cmd.Flag("s3-prefix").Value.String(),
		)

		err = backupS3.BackupToS3(cmd.Flag("backup-file").Value.String())
		if err != nil {
			panic(err)
		}

	},
}

func init() {
	rootCmd.AddCommand(backupS3Cmd)

	backupS3Cmd.Flags().String("s3-bucket", viper.GetString("S3-BUCKET"), "S3 bucket name")
	backupS3Cmd.Flags().String("s3-prefix", viper.GetString("S3-PREFIX"), "S3 prefix")
	backupS3Cmd.Flags().String("postgres-host", viper.GetString("POSTGRES-HOST"), "Postgres host")
	backupS3Cmd.Flags().String("postgres-user", viper.GetString("POSTGRES-USER"), "Postgres user")
	backupS3Cmd.Flags().String("postgres-password", viper.GetString("POSTGRES-PASSWORD"), "Postgres password")
	backupS3Cmd.Flags().String("postgres-database", viper.GetString("POSTGRES-DATABASE"), "Postgres database")
	backupS3Cmd.Flags().String("backup-file", viper.GetString("BACKUP-FILE"), "Backup file")

}
