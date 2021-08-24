package command

import (
	"github.com/gin-gonic/gin"
	"github.com/salihkemaloglu/todo/pkg/handler"
	"github.com/salihkemaloglu/todo/pkg/util/config"
	"github.com/spf13/cobra"
)

// Backend API Service
func NewAPIRun(config *config.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "run-api",
		Short: "Run api service",
		Long:  `Run api backend service.`,
		Args:  cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			runAPI(config)
		},
	}

	return cmd
}

func runAPI(config *config.Config) {

	r := gin.Default()
	h := handler.NewHandler(config)
	health := r.Group("/")
	{
		health.POST("callback", h.Callback)
	}

	if err := r.Run(config.Service.Port); err != nil {
		panic(err)
	}
}
