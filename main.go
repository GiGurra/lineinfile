package main

import (
	"github.com/GiGurra/boa/pkg/boa"
	"github.com/spf13/cobra"
)

type Params struct {
	Mode          boa.Required[string] `name:"mode" short:"m" alts:"add,remove" default:"add"`
	Line          boa.Required[string] `name:"line" pos:"true" help:"The line to add or remove"`
	DetectPattern boa.Optional[string] `name:"detect-pattern" short:"d" help:"The pattern to detect the line"`
}

func main() {
	boa.CmdT[Params]{
		Use:         "lineinfile",
		Short:       "Add or remove lines in a file",
		ParamEnrich: boa.ParamEnricherNone,
		RunFunc: func(params *Params, cmd *cobra.Command, args []string) {
			panic("not implemented")
		},
	}.Run()
}
