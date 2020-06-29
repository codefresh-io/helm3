package builder

import (
	main "github.com/codefresh-io/helm3/cmd"
)

type KubernetesCommandsBuilder struct {
}

func (commandBuilder KubernetesCommandsBuilder) build(context Context) []string {
	return []string{}
}

//
//def _build_kubectl_commands(self):
//lines = []
//if self.action in ['install', 'promotion', 'auth']:
//if self.kube_context is not None:
//kubectl_cmd = 'kubectl config use-context "%s"' % self.kube_context
//if self.dry_run:
//kubectl_cmd = 'echo ' + kubectl_cmd
//lines.append(kubectl_cmd)
//
//if self.kube_context is None and self.action != 'auth':
//raise Exception(
//'Must set KUBE_CONTEXT in environment (Name of Kubernetes cluster as named in Codefresh)')
//
//return lines
