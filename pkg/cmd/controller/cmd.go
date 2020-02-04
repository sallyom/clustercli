package controller

import (
	"github.com/spf13/cobra"
	"github.com/openshift/clustercli/pkg/starter"
	"github.com/openshift/clustercli/pkg/version"
    "github.com/openshift/library-go/pkg/controller/controllercmd"
)

func NewCLIController() *cobra.Command {
	cmd := controllercmd.
	    NewControllerCommandConfig("cli-manager", version.Get(), starter.RunCLIController).
		NewCommand()
    cmd.Use = "cli-manager"
	cmd.Short = "Start the Cluster cli-manager"

	return cmd
}
