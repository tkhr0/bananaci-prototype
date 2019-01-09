package cmd

import (
	"sync"

	"github.com/spf13/cobra"
	jobQueue "github.com/tkhr0/bananaci-prototype/lib/job_queue"
	"github.com/tkhr0/bananaci-prototype/lib/server"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run BananaCI server.",
	Long:  `Run worker and http server.`,
	Run:   serverCmdMain,
}

func init() {
	rootCmd.AddCommand(serverCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func serverCmdMain(cmd *cobra.Command, args []string) {
	maxWorkers := 3
	maxQueues := 10000
	d := jobQueue.NewDispatcher(maxWorkers, maxQueues)

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		server.Call(d)
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		server.CallHTTP(d)
		wg.Done()
	}()

	wg.Wait()
}
