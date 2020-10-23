package root

import (
	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/kubectl/pkg/cmd/util"

	"github.com/arjit95/kubepf/pkg/cmd/completion"
	"github.com/arjit95/kubepf/pkg/cmd/start"
	"github.com/arjit95/kubepf/pkg/cmd/stop"
	"github.com/arjit95/kubepf/pkg/cmdutil"
)

// NewCmdRoot returns kubepf root command handler
func NewCmdRoot(factory *cmdutil.Factory) *cobra.Command {
	root := &cobra.Command{
		Use:   "kubepf",
		Short: "Manage port-forwarding for your kubernetes nodes",
		Long: heredoc.Doc(`
			kubepf handles all the different resources in a single
			session, so you don't have to switch terminals or write
			bash scripts to port-forward multiple resources.
		`),
	}

	flags := root.PersistentFlags()
	flags.BoolP("interactive", "i", false, "Start kubepf in interactive mode")
	flags.Bool("no-cache", false, "Do not use cached state")

	kubeConfigFlags := genericclioptions.NewConfigFlags(true).WithDeprecatedPasswordFlag()
	kubeConfigFlags.AddFlags(flags)
	matchVersionKubeConfigFlags := util.NewMatchVersionFlags(kubeConfigFlags)
	matchVersionKubeConfigFlags.AddFlags(flags)

	f := util.NewFactory(matchVersionKubeConfigFlags)
	factory.Factory = f

	// Redirect streams and add child command.
	// Streams will be redirected to cobi if running in
	// interactive mode
	root.SetOut(factory.IOStreams.Out)
	root.SetErr(factory.IOStreams.ErrOut)
	root.AddCommand(start.NewCmdStart(factory))
	root.AddCommand(stop.NewCmdStop(factory))
	root.AddCommand(completion.NewCmdCompletion(factory))

	root.RegisterFlagCompletionFunc("namespace", completion.NSCompletionFunc(factory))

	return root
}
