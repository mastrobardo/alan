package cmd

import (
	"fmt"

	"github.com/kelvindecosta/alan/machine"
	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run file",
	Short: "Run a Turing Machine",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		input, _ := cmd.Flags().GetString("input")
		maxSteps, _ := cmd.Flags().GetUint("max-steps")
		m := machine.NewMachine()
		m.Parse(args[0])
		h, a := m.Compute(input, maxSteps)
		fmt.Printf("Halted   : %v\nAccepted : %v\n", h, a)
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.Flags().StringP("input", "i", "", "input string")
	runCmd.Flags().UintP("max-steps", "m", 200, "maximum steps before forced halt")
}
