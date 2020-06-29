package context

type Context struct {
	Action      string
	KubeContext string
	DryRun      bool
}
