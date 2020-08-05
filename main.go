package main

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/chenjiandongx/react/modules"
)

const version = "0.1.0"

var rootCmd = &cobra.Command{
	Use:     "oscar",
	Short:   "Next generation building tool for nothing",
	Version: version,
}

func NewListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List all available modules",
		Run: func(cmd *cobra.Command, args []string) {
			for _, m := range modules.Registry {
				fmt.Println(" -", m.Name())
			}
		},
	}
	return cmd
}

func NewRunCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "run",
		Short: "Run with the specify module",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				fmt.Println("> please choose a module to run")
				return
			}
			forever, _ := cmd.Flags().GetBool("forever")
			runCommand(args[0], forever)
		},
		Example: "  oscar run bootlog",
	}

	cmd.Flags().BoolP("forever", "f", false, "whether to run module forever")
	return cmd
}

func runCommand(m string, forever bool) {
	var moduler modules.Moduler
	for _, r := range modules.Registry {
		if r.Name() == m {
			moduler = r
			break
		}
	}

	if moduler == nil {
		fmt.Printf("> sorry, [%s] module does not exist\n", m)
		return
	}
	moduler.Display(forever)
}

func init() {
	rootCmd.AddCommand(
		NewListCommand(),
		NewRunCommand(),
	)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
