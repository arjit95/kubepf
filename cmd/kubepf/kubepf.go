package main

import (
	"flag"
	"os"
	"path/filepath"

	"github.com/arjit95/cobi"
	cobEditor "github.com/arjit95/cobi/editor"
	"github.com/gdamore/tcell"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/kubectl/pkg/cmd/portforward"

	"github.com/arjit95/kubepf/pkg/cmd/root"
	"github.com/arjit95/kubepf/pkg/cmdutil"
	"github.com/arjit95/kubepf/pkg/iostreams"
)

func main() {
	var kubeconfig *string
	if home := homeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}

	clientset, err := kubernetes.NewForConfig(config)

	if err != nil {
		panic(err)
	}

	editor := cobEditor.NewEditor()
	editor.Input.SetFieldBackgroundColor(tcell.ColorBlack)
	editor.SetUpperPaneTitle("Kubernetes Port Forwarding")
	editor.SetLowerPaneTitle("Logs")

	factory := &cmdutil.Factory{
		RestConfig: config,
		Client:     clientset,
		IOStreams: iostreams.IOStreams{
			In:     os.Stdin,
			Out:    editor.Output,
			ErrOut: editor.Logger.Error,
			Logger: editor.Logger.Info,
		},
		Resources: make(map[string]*portforward.PortForwardOptions),
	}

	rootCmd := cobi.NewCommand(editor, root.NewCmdRoot(factory))

	err = rootCmd.ParseFlags(os.Args)
	if err != nil { // Execute root command to display additional info
		rootCmd.Execute()
		return
	}

	interactive, err := rootCmd.Flags().GetBool("interactive")
	if err != nil {
		rootCmd.Help()
		return
	}

	factory.Root = rootCmd
	if interactive {
		factory.Interactive = true

		rootCmd.ExecuteInteractive()
	} else {
		rootCmd.Execute()
	}
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}

	return os.Getenv("USERPROFILE") // windows
}
