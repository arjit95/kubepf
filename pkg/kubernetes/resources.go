package kubernetes

import (
	"github.com/spf13/cobra"
	"k8s.io/client-go/discovery"

	"github.com/arjit95/kubepf/pkg/cmdutil"
)

// ListResources returns the names of all resources for a type
func ListResources(f *cmdutil.Factory, cmd *cobra.Command, args []string) ([]string, error) {
	namespace, err := cmd.InheritedFlags().GetString("namespace")
	if err != nil {
		return nil, err
	}

	r := f.Factory.NewBuilder().
		Unstructured().
		NamespaceParam(namespace).DefaultNamespace().AllNamespaces(namespace == "").
		ResourceTypeOrNameArgs(true, args...).
		ContinueOnError().
		Latest().
		Flatten().
		Do()

	resources, err := r.Infos()
	if err != nil {
		return nil, err
	}

	names := []string{}
	for _, resource := range resources {
		names = append(names, resource.Name)
	}

	return names, nil
}

// ListResourceTypes returns the list of available resource types for the kubernetes cluster
func ListResourceTypes(f *cmdutil.Factory) ([]string, error) {
	discoveryclient, err := f.Factory.ToDiscoveryClient()
	if err != nil {
		return nil, err
	}

	lists, err := discovery.ServerPreferredResources(discoveryclient)
	if err != nil {
		return nil, err
	}

	resources := []string{}

	for _, list := range lists {
		for _, resource := range list.APIResources {
			resources = append(resources, resource.Name)
			resources = append(resources, resource.ShortNames...)
		}
	}

	return resources, nil
}
