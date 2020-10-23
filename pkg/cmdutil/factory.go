package cmdutil

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"

	"github.com/arjit95/cobi"
	"github.com/arjit95/kubepf/pkg/iostreams"
	"k8s.io/kubectl/pkg/cmd/portforward"
	cmdutil "k8s.io/kubectl/pkg/cmd/util"
)

// Factory contains the current state of app
type Factory struct {
	cmdutil.Factory
	IOStreams  iostreams.IOStreams
	RestConfig *rest.Config
	// Resources which are actively being portforwarded
	Resources   map[string]*portforward.PortForwardOptions
	Client      *kubernetes.Clientset
	Interactive bool
	Root        *cobi.Command
}
