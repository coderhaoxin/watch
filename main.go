package main

import "github.com/pkg4go/execx"
import "github.com/spf13/cobra"

func main() {
	var dir, cmd string

	root := &cobra.Command{
		Use:   "watch",
		Short: "Watch files and folders for changes and something more ...",
		Long:  "Watch files and folders for changes and something more ...",
		Run: func(c *cobra.Command, args []string) {
			if dir == "" || cmd == "" {
				exit("both dir and cmd are required")
			}

			arr := execx.Split(cmd)

			watchDir(dir, arr[0], arr[1:]...)
		},
	}

	root.Flags().StringVarP(&dir, "dir", "d", "", "the directory to watch")
	root.Flags().StringVarP(&cmd, "cmd", "c", "", "the command to execute")

	root.Execute()
}
