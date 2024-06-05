package create_admin

import (
	"gym-management-auth/src/components/auth/core/domain"
	"gym-management-auth/src/lib/primitives/application_specific"
)

type CreateAdminCommandHandler struct {
	UserRepository domain.UserRepository
}

func (h *CreateAdminCommandHandler) Handle(command *CreateAdminCommand) (*CreateAdminCommandResponse, *application_specific.ApplicationException) {
	email, err := application_specific.NewEmail(command.Email)
	if err != nil {
		return nil, err
	}

	used, err := h.UserRepository.UsernameUsed(domain.UsernameFromEmail(email), command.Session)
	if err != nil {
		return nil, err
	}
	if used {
		return nil, application_specific.NewValidationException("AUTH.EMAIL_USED", "Email is already used", map[string]string{
			"email": command.Email,
		})
	}

	password, err := domain.PasswordFromPlain(command.Password)
	if err != nil {
		return nil, err
	}

	admin := domain.CreateAdmin(
		command.FirstName,
		command.LastName,
		command.Phone,
		email,
		password,
	)

	err = h.UserRepository.Create(admin, command.Session)
	if err != nil {
		return nil, err
	}

	return &CreateAdminCommandResponse{
		Id: admin.State().Id,
	}, nil
}
