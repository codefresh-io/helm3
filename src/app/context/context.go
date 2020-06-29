package context

type Action string

const (
	Install   Action = "Install"
	Promotion Action = "Promotion"
	Auth      Action = "Auth"
)

type Context struct {
	action      Action
	kubeContext string
}
