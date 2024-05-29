package login

import (
	"gym-management/src/components/auth/core/domain"
	"gym-management/src/lib/primitives/application_specific"
)

type LoginCommandHandler struct {
	JwtSecret      string
	UserRepository domain.UserRepository
}

func (h *LoginCommandHandler) Handle(command *LoginCommand) (*LoginCommandResponse, *application_specific.ApplicationException) {
	username, err := domain.UsernameFrom(command.Username)
	if err != nil {
		return nil, err
	}

	user, err := h.UserRepository.FindByUsername(username, command.Session)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, application_specific.NewValidationException("AUTH.USERNAME.INVALID", "Invalid Credentials", map[string]string{
			"username": command.Username,
			"password": command.Password,
		})
	}

	token, err := user.Login(command.Password, h.JwtSecret)
	if err != nil {
		return nil, err
	}

	err = h.UserRepository.Update(user, command.Session)
	if err != nil {
		return nil, err
	}

	return &LoginCommandResponse{
		Token: token.Value(),
	}, nil
}
