package auth

import (
	"gym-management/src/components/auth/core/domain"
	"gym-management/src/components/auth/core/usecases/create_admin"
	"gym-management/src/components/auth/core/usecases/get_session"
	"gym-management/src/components/auth/core/usecases/get_user"
	"gym-management/src/components/auth/core/usecases/login"
	"gym-management/src/lib/primitives/application_specific"
)

type Facade struct {
	UserRepository domain.UserRepository
	JwtSecret      string
}

func (f *Facade) GetUser(query *get_user.GetUserQuery) (*get_user.GetUserQueryResponse, *application_specific.ApplicationException) {
	handler := &get_user.GetUserQueryHandler{
		UserRepository: f.UserRepository,
	}

	return handler.Handle(query)
}

func (f *Facade) Login(command *login.LoginCommand) (*login.LoginCommandResponse, *application_specific.ApplicationException) {
	handler := &login.LoginCommandHandler{
		UserRepository: f.UserRepository,
		JwtSecret:      f.JwtSecret,
	}

	return handler.Handle(command)
}

func (f *Facade) GetSession(query *get_session.GetSessionQuery) (*get_session.GetSessionQueryResponse, *application_specific.ApplicationException) {
	handler := &get_session.GetSessionQueryHandler{
		UserRepository: f.UserRepository,
		JwtSecret:      f.JwtSecret,
	}

	return handler.Handle(query)
}

func (f *Facade) CreateAdmin(command *create_admin.CreateAdminCommand) (*create_admin.CreateAdminCommandResponse, *application_specific.ApplicationException) {
	handler := &create_admin.CreateAdminCommandHandler{
		UserRepository: f.UserRepository,
	}

	return handler.Handle(command)
}
