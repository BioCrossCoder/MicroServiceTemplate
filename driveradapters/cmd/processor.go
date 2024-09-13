package cmd

import (
	"sync"

	"github.com/spf13/cobra"
)

type CommandExecutor interface {
	Register(rootCmd *cobra.Command)
}

type Processor interface {
	Run() error
}

var (
	p     Processor
	pOnce sync.Once
)

type processor struct {
	executors []CommandExecutor
	cmd       *cobra.Command
}

func NewProcessor() Processor {
	pOnce.Do(func() {
		p = &processor{
			cmd: &cobra.Command{Use: "exec"},
			executors: []CommandExecutor{
				newAppExecutor(),
				newServerExecutor(),
			},
		}
	})
	return p
}

func (p *processor) Run() error {
	for _, executor := range p.executors {
		executor.Register(p.cmd)
	}
	return p.cmd.Execute()
}
