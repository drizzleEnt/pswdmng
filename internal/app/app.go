package app

import (
	"pswdmng/internal/commands"
	"pswdmng/internal/repository"
	"pswdmng/internal/repository/dbrepo"
	"pswdmng/internal/service"
	"pswdmng/internal/service/password"
)

type Option func(*App)

type App struct {
	passwordService service.PasswordService
	repo repository.Repository
	root *commands.Root
}

func WithStorePath(path string) Option {
	return func(a *App) {

	}
}

func New(opts ...Option) *App {
	a := &App{}

	a.Root().InitCommands()

	for _, opt := range opts {
		opt(a)
	}

	return a
}

func (a *App) Run() {
	a.Root().Execute()
}

func (a *App) Repository() repository.Repository {
	if a.repo == nil {
		a.repo = dbrepo.New()
	}

	return a.repo
}

func (a *App) Root() *commands.Root {
	if a.root == nil {
		a.root = commands.New(a.Repository(), a.PasswordService())
	}

	return a.root
}

func (a *App)PasswordService() service.PasswordService {
	if a.passwordService == nil{
		a.passwordService = password.New()
	}

	return a.passwordService
}