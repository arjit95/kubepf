package completion

import (
	"strings"

	"github.com/spf13/cobra"

	"github.com/arjit95/kubepf/pkg/cmdutil"
	"github.com/arjit95/kubepf/pkg/kubernetes"
	"github.com/arjit95/kubepf/pkg/utils"
)

// CobraCompleter is a cobra completion function
type CobraCompleter func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective)

type cachedCompletions struct {
	namespace []string
	resources []string
}

var completions *cachedCompletions

func filterCompletions(completions []string, toComplete string) []string {
	return utils.FilterStr(completions, func(str string) bool {
		return strings.HasPrefix(str, toComplete)
	})
}

// ResourceCompletionFunc validates the args and returns a cobra completion function
// which can be used to provide dynamic completions for pod names
func ResourceCompletionFunc(f *cmdutil.Factory) CobraCompleter {
	if completions == nil {
		completions = &cachedCompletions{}
	}

	return func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		skipCache, err := cmd.InheritedFlags().GetBool("no-cache")
		if err != nil {
			return nil, cobra.ShellCompDirectiveDefault
		}

		resources := []string{}
		if len(args) == 0 {
			if !skipCache && len(completions.resources) > 0 {
				resources = completions.resources
			} else {
				resources, err = kubernetes.ListResourceTypes(f)
			}

			if err != nil {
				return nil, cobra.ShellCompDirectiveDefault
			}

			completions.namespace = resources
		} else if len(args) == 1 {
			resources, err = kubernetes.ListResources(f, cmd, args)
			if err != nil {
				return nil, cobra.ShellCompDirectiveDefault
			}
		}

		return filterCompletions(resources, toComplete), cobra.ShellCompDirectiveNoFileComp
	}
}

// NSCompletionFunc is a flag completion function used to provide
// namespace suggestions
func NSCompletionFunc(f *cmdutil.Factory) CobraCompleter {
	return func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		namespaces, err := kubernetes.ListNamespaces(f)
		if err != nil {
			return nil, cobra.ShellCompDirectiveDefault
		}

		return filterCompletions(namespaces, toComplete), cobra.ShellCompDirectiveNoFileComp
	}
}
