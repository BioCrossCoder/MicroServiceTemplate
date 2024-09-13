package cmd

import (
	"fmt"
	"main/common"
	"main/driveradapters/api"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

type serverExecutor struct {
	rootCmd *cobra.Command
	router  api.Router
}

var (
	se     CommandExecutor
	seOnce sync.Once
)

func newServerExecutor() CommandExecutor {
	seOnce.Do(func() {
		executor := &serverExecutor{
			rootCmd: &cobra.Command{
				Use:   "server",
				Short: "Run the server",
			},
			router: api.NewRouter(),
		}
		executor.rootCmd.Run = func(cmd *cobra.Command, args []string) {
			executor.runServer()
		}
		se = executor
	})
	return se
}

func (e *serverExecutor) Register(rootCmd *cobra.Command) {
	rootCmd.AddCommand(e.rootCmd)
}

func (e *serverExecutor) createServer(pathPrefix string) *gin.Engine {
	server := gin.New()
	server.Use(func(ctx *gin.Context) {
		apiPath := strings.TrimPrefix(ctx.FullPath(), pathPrefix)
		if apiPath != "" && !strings.HasPrefix(apiPath, "/health") {
			gin.Logger()(ctx)
		}
	})
	server.Use(gin.Recovery())
	return server
}

func (e *serverExecutor) runServer() {
	gin.SetMode(gin.ReleaseMode)
	go func() {
		apiPrefix := "/api/demo/v1"
		server := e.createServer(apiPrefix)
		group := server.Group(apiPrefix)
		e.router.RegisterPublicAPI(group)
		addr := fmt.Sprintf("%v:%v", common.SelfConfig.Host, common.SelfConfig.PublicPort)
		err := server.Run(addr)
		if err != nil {
			panic(err)
		}
	}()
	go func() {
		apiPrefix := "/api/demo/v1"
		server := e.createServer(apiPrefix)
		group := server.Group(apiPrefix)
		e.router.RegisterPrivateAPI(group)
		addr := fmt.Sprintf("%v:%v", common.SelfConfig.Host, common.SelfConfig.PrivatePort)
		err := server.Run(addr)
		if err != nil {
			panic(err)
		}
	}()
	for {
	}
}