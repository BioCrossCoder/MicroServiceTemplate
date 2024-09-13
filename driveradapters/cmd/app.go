package cmd

import (
	"fmt"
	"main/logics"
	"sync"

	"github.com/spf13/cobra"
)

type appExecutor struct {
	adSvc   logics.AppDeploymentService
	rootCmd *cobra.Command
}

var (
	ae     CommandExecutor
	aeOnce sync.Once
)

func newAppExecutor() CommandExecutor {
	aeOnce.Do(func() {
		ae = &appExecutor{
			adSvc: logics.NewAppDeploymentService(),
			rootCmd: &cobra.Command{
				Use:   "app",
				Short: "app deployment management",
			},
		}
	})
	return ae
}

func (e *appExecutor) Register(rootCmd *cobra.Command) {
	e.rootCmd.AddCommand(
		e.install(),
		e.uninstall(),
		e.list(),
	)
	rootCmd.AddCommand(e.rootCmd)
}

func (e *appExecutor) install() *cobra.Command {
	return &cobra.Command{
		Use:   "install",
		Short: "install app",
		RunE: func(cmd *cobra.Command, args []string) error {
			appName := args[0]
			return e.adSvc.InstallApp(appName)
		},
	}
}

func (e *appExecutor) uninstall() *cobra.Command {
	return &cobra.Command{
		Use:   "uninstall",
		Short: "uninstall app",
		RunE: func(cmd *cobra.Command, args []string) error {
			appName := args[0]
			return e.adSvc.UninstallApp(appName)
		},
	}
}

func (e *appExecutor) list() *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "list all apps",
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			appList, err := e.adSvc.ListApps()
			if err != nil {
				return
			}
			for _, app := range appList {
				fmt.Println(app)
			}
			return
		},
	}
}
