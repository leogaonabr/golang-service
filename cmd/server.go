/*
Package cmd ...
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/viper"

	"github.com/leogaonabr/golang-service/api"
	"github.com/spf13/cobra"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Starts the HTTP server with the REST API",
	Long:  "Starts the HTTP server with the REST API",
	Run: func(cmd *cobra.Command, args []string) {
		srv := api.StartServer()

		// creates the channel that will keep the daemon goroutine blocked
		// until a sigint, sigterm of sigkill is received by the running process
		shutdownChannel := make(chan os.Signal, 1)
		signal.Notify(shutdownChannel, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
		<-shutdownChannel

		// starts the gracefull shutdown process and connection draining
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		srv.Shutdown(ctx)
	},
}

func init() {
	// binds the flags used on the 'server' command
	serverCmd.PersistentFlags().StringP("ENV", "e", "development", "environment")
	serverCmd.PersistentFlags().IntP("PORT", "p", 3000, "server port")

	rootCmd.AddCommand(serverCmd)
	viper.AutomaticEnv()
}
