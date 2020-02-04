package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/openshift/clustercli/pkg/cmd/controller"
)

func main() {
	command := NewClusterCLICommand()
	if err := command.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func NewClusterCLICommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cli-manager",
		Short: "OpenShift cluster cli-manager",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
			os.Exit(1)
		},
	}

	cmd.AddCommand(controller.NewCLIController())
	return cmd
}
