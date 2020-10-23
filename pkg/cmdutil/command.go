package cmdutil

import (
	"fmt"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"k8s.io/kubectl/pkg/cmd/util"
)

// AddPodRunningTimeoutFlag specifies the amount of time to wait for a pod to be in running state
func AddPodRunningTimeoutFlag(command *cobra.Command, duration time.Duration) {
	util.AddPodRunningTimeoutFlag(command, duration)
}

// NormalizeArgs performs various normalizations on args
// Converts Type Name => Type/Name
// This conversion is done because kubernetes completions only
// work with spaces and not slashes.
// Resource(s) => Resource
// Portforwarding works on singular resource rather than its plural
// version.
func NormalizeArgs(args []string) ([]string, error) {
	if len(args) < 2 {
		return nil, fmt.Errorf("Insufficient args")
	} else if strings.Index(args[0], "/") > -1 {
		return nil, fmt.Errorf("Please use space between resource type and name")
	}

	// Convert plural resource name to singular
	if strings.LastIndex(args[0], "s") == len(args[0])-1 {
		args[0] = args[0][0 : len(args[0])-1]
	}

	args[0] = fmt.Sprintf("%s/%s", args[0], args[1])
	return removeArg(args, 1), nil
}

func removeArg(args []string, idx int) []string {
	copy(args[idx:], args[idx+1:])
	args = args[:len(args)-1]
	return args
}
