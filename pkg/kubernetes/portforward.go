package kubernetes

import (
	"fmt"
	"net/http"
	"net/url"

	"k8s.io/cli-runtime/pkg/genericclioptions"
	pf "k8s.io/client-go/tools/portforward"
	"k8s.io/client-go/transport/spdy"
	"k8s.io/kubectl/pkg/cmd/portforward"

	"github.com/arjit95/kubepf/pkg/cmdutil"
)

// DefaultPortForwarder creates port forwarder instance for the resource
type DefaultPortForwarder struct {
	Factory *cmdutil.Factory
	genericclioptions.IOStreams
}

// ForwardPorts starts port forwarding for a specific resource
func (f *DefaultPortForwarder) ForwardPorts(method string, url *url.URL, opts portforward.PortForwardOptions) error {
	transport, upgrader, err := spdy.RoundTripperFor(opts.Config)
	if err != nil {
		return err
	}
	dialer := spdy.NewDialer(upgrader, &http.Client{Transport: transport}, method, url)
	fw, err := pf.NewOnAddresses(dialer, opts.Address, opts.Ports, opts.StopChannel, opts.ReadyChannel, f.Out, f.ErrOut)
	if err != nil {
		return err
	}

	if f.Factory.Interactive {
		go func() {
			podKey := fmt.Sprintf("%s/%s", opts.Namespace, opts.PodName)
			f.Factory.Resources[podKey] = &opts
			fw.ForwardPorts()
			delete(f.Factory.Resources, podKey)
		}()

		return nil
	}

	return fw.ForwardPorts()
}
