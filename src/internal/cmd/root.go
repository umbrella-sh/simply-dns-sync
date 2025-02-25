package cmd

import (
	"github.com/spf13/cobra"

	"github.com/umbrella-sh/simply-dns-cli/internal/api"
	"github.com/umbrella-sh/simply-dns-cli/internal/cmd/backup"
	"github.com/umbrella-sh/simply-dns-cli/internal/cmd/create"
	"github.com/umbrella-sh/simply-dns-cli/internal/cmd/list"
	"github.com/umbrella-sh/simply-dns-cli/internal/cmd/remove"
	"github.com/umbrella-sh/simply-dns-cli/internal/cmd/restore"
	"github.com/umbrella-sh/simply-dns-cli/internal/cmd/update"
	"github.com/umbrella-sh/simply-dns-cli/internal/configs"
	"github.com/umbrella-sh/simply-dns-cli/internal/objects"
)

var rootCmd = &cobra.Command{
	Use:   configs.AppName,
	Short: "",
}

func init() {
	rootCmd.AddCommand(list.ListCmd)
	rootCmd.AddCommand(create.CreateCmd)
	rootCmd.AddCommand(update.UpdateCmd)
	rootCmd.AddCommand(remove.RemoveCmd)
	rootCmd.AddCommand(backup.BackupCmd)
	rootCmd.AddCommand(restore.RestoreCmd)
}

func RootExecute() error {
	api.Init(objects.SimplyApiConfig{
		Url:           configs.Main.SimplyApi.Url,
		AccountNumber: configs.Main.SimplyApi.AccountNumber,
		AccountApiKey: configs.Main.SimplyApi.AccountApiKey,
	})

	return rootCmd.Execute()
}
