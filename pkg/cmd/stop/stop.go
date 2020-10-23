package stop

import (
	"fmt"

	"github.com/MakeNowJust/heredoc"
	"github.com/gdamore/tcell"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"

	"github.com/arjit95/kubepf/pkg/cmdutil"
)

// NewCmdStop returns stop port forwarding command handler
func NewCmdStop(f *cmdutil.Factory) *cobra.Command {
	command := &cobra.Command{
		Use:   "stop",
		Short: "Stops port forwarding on resource",
		Long:  "Stops an existing port forwarding operation for a pod. Note: Only pods forwarded by kubepf can be stopped",
		Example: heredoc.Doc(`
			$ kubepf stop pod pod-name
			// interactive mode
			$ stop pod pod-name
		`),

		Args:                  cobra.NoArgs,
		DisableFlagsInUseLine: true,

		RunE: func(cmd *cobra.Command, args []string) error {
			resources := make([]string, 0, len(f.Resources))

			for key := range f.Resources {
				resources = append(resources, key)
			}

			prompt := promptui.Select{
				Label: "Select resource to stop",
				Items: resources,
			}

			_, result, promptErr := prompt.Run()

			if f.Interactive { // Swap the screen to redraw the view
				screen, err := tcell.NewScreen()
				if err != nil {
					return err
				}

				go func() {
					f.Root.App.QueueUpdateDraw(func() {
						f.Root.App.SetScreen(screen)
					})
				}()
			}

			// Prompt errors are ignored, but an error means we should not proceed further
			if promptErr != nil {
				return nil
			}

			resource := f.Resources[result]
			if resource == nil {
				return fmt.Errorf("Cannot find resource %s", result)
			}

			close(resource.StopChannel)
			fmt.Fprintf(f.IOStreams.Logger, "Successfully stopped %s\n", result)
			return nil
		},
	}

	// Hack: Some global flags are not being recognized without this line.
	command.Flags()
	return command
}
