package app

import (
	"github.com/codefresh-io/helm3/src/app/builder"
	ctx "github.com/codefresh-io/helm3/src/app/context"
	"os"
	"strconv"
)

func main() {
	dryRun, _ := strconv.ParseBool(os.Getenv("DRY_RUN"))

	context := ctx.Context{
		Action:      os.Getenv("ACTION"),
		KubeContext: os.Getenv("KUBE_CONTEXT"),
		DryRun:      dryRun,
	}
	kubernetesCommandsBuilder := builder.KubernetesCommandsBuilder{context}
	kubernetesCommandsBuilder.Build()
}
