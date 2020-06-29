package builder

import (
	ctx "github.com/codefresh-io/helm3/src/app/context"
)

type KubernetesCommandsBuilder struct {
	context ctx.Context
}

func (commandBuilder KubernetesCommandsBuilder) Build() []string {
	context := commandBuilder.context
	lines := []string{}

	availableActions := map[string]bool{
		"install":   true,
		"promotion": true,
		"auth":      true,
	}
	if availableActions[context.Action] {
		if context.KubeContext != "" {
			cmd := "kubectl config use-context " + context.KubeContext
			if context.DryRun == true {
				cmd = "echo " + cmd
			}
			lines = append(lines, cmd)
		}
	}
	return lines
}
