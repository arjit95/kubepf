package start

import (
	"fmt"
	"time"

	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	pf "k8s.io/kubectl/pkg/cmd/portforward"

	"github.com/arjit95/kubepf/pkg/cmd/completion"
	"github.com/arjit95/kubepf/pkg/cmdutil"
	"github.com/arjit95/kubepf/pkg/kubernetes"
)

const (
	// Amount of time to wait until at least one pod is running
	defaultPodPortForwardWaitTimeout = 60 * time.Second
)

// NewCmdStart returns start port forwarding command handler
func NewCmdStart(f *cmdutil.Factory) *cobra.Command {

	command := &cobra.Command{
		Use:   "start",
		Short: "Starts port forwarding on resource",
		Long:  "Port forwards a kubernetes resource, so it can be accessed by the host machine.",
		Example: heredoc.Doc(`
			$ kubepf start pod pod-name localPort:targetPort
			// interactive mode
			$ start pod pod-name localPort:targetPort
		`),

		Args:                  cobra.MinimumNArgs(2),
		ValidArgsFunction:     completion.ResourceCompletionFunc(f),
		DisableFlagsInUseLine: true,

		RunE: func(cmd *cobra.Command, args []string) error {
			address, err := cmd.Flags().GetStringSlice("address")
			if err != nil {
				return err
			}

			namespace, err := cmd.InheritedFlags().GetString("namespace")
			if err != nil {
				return err
			}

			opts := &pf.PortForwardOptions{
				PortForwarder: &kubernetes.DefaultPortForwarder{
					Factory: f,
					IOStreams: genericclioptions.IOStreams{
						In:     f.IOStreams.In,
						Out:    f.IOStreams.Logger,
						ErrOut: f.IOStreams.ErrOut,
					},
				},
				Address:   address,
				Namespace: namespace,
			}

			if args, err = cmdutil.NormalizeArgs(args); err != nil {
				return err
			}

			fmt.Fprintf(f.IOStreams.Logger, "Trying to port-forward %s\n", args[0])

			if err = opts.Complete(f, cmd, args); err != nil {
				return err
			}
			if err = opts.Validate(); err != nil {
				return err
			}

			if err = opts.RunPortForward(); err != nil {
				return err
			}

			return nil
		},
	}

	command.Flags().StringSlice("address", []string{"localhost"}, "Addresses to listen on (comma separated). Only accepts IP addresses or localhost as a value. When localhost is supplied, kubectl will try to bind on both 127.0.0.1 and ::1 and will fail if neither of these addresses are available to bind.")
	cmdutil.AddPodRunningTimeoutFlag(command, defaultPodPortForwardWaitTimeout)
	return command
}

func removeArg(args []string, idx int) []string {
	copy(args[idx:], args[idx+1:])
	args = args[:len(args)-1]
	return args
}
