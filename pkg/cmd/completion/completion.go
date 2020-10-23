package completion

import (
	"github.com/arjit95/kubepf/pkg/cmdutil"
	"github.com/spf13/cobra"
)

// NewCmdCompletion returns a completion command, which can be used
// to generate shell completions for kubepf
func NewCmdCompletion(f *cmdutil.Factory) *cobra.Command {
	return &cobra.Command{
		Use:   "completion [bash|zsh|fish|powershell]",
		Short: "Generate completion script",
		Long: `To load completions:
	
	Bash:
	
	$ source <(kubepf completion bash)
	
	# To load completions for each session, execute once:
	Linux:
	  $ kubepf completion bash > /etc/bash_completion.d/kubepf
	MacOS:
	  $ kubepf completion bash > /usr/local/etc/bash_completion.d/kubepf
	
	Zsh:
	
	# If shell completion is not already enabled in your environment you will need
	# to enable it.  You can execute the following once:
	
	$ echo "autoload -U compinit; compinit" >> ~/.zshrc
	
	# To load completions for each session, execute once:
	$ kubepf completion zsh > "${fpath[1]}/_kubepf"
	
	# You will need to start a new shell for this setup to take effect.
	
	Fish:
	
	$ kubepf completion fish | source
	
	# To load completions for each session, execute once:
	$ kubepf completion fish > ~/.config/fish/completions/kubepf.fish
	`,
		DisableFlagsInUseLine: true,
		ValidArgs:             []string{"bash", "zsh", "fish", "powershell"},
		Args:                  cobra.ExactValidArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			switch args[0] {
			case "bash":
				cmd.Root().GenBashCompletion(f.IOStreams.Out)
			case "zsh":
				cmd.Root().GenZshCompletion(f.IOStreams.Out)
			case "fish":
				cmd.Root().GenFishCompletion(f.IOStreams.Out, true)
			case "powershell":
				cmd.Root().GenPowerShellCompletion(f.IOStreams.Out)
			}
		},
	}
}
