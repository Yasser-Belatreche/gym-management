package auth

import (
	"gym-management/src/components/auth/core/usecases/get_session"
	"gym-management/src/components/auth/core/usecases/get_user"
	"gym-management/src/components/auth/core/usecases/login"
	"gym-management/src/lib/primitives/application_specific"
)

type Manager interface {
	Login(command *login.LoginCommand) (*login.LoginCommandResponse, *application_specific.ApplicationException)

	GetSession(query *get_session.GetSessionQuery) (*get_session.GetSessionQueryResponse, *application_specific.ApplicationException)

	GetUser(query *get_user.GetUserQuery) (*get_user.GetUserQueryResponse, *application_specific.ApplicationException)
}
