package get_user

import (
	"gym-management/src/components/auth/core/domain"
	"gym-management/src/lib/primitives/application_specific"
)

type GetUserQueryHandler struct {
	UserRepository domain.UserRepository
}

func (h *GetUserQueryHandler) Handle(query *GetUserQuery) (*GetUserQueryResponse, *application_specific.ApplicationException) {
	user, err := h.UserRepository.FindByID(query.Id, query.Session.Session)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, application_specific.NewNotFoundException("AUTH.USER.NOT_FOUND", "user not found", map[string]string{
			"id": query.Id,
		})
	}

	return &GetUserQueryResponse{
		Id:        user.State().Id,
		Role:      user.State().Role,
		Email:     user.State().Profile.Email,
		Usernames: user.State().Usernames,
		FirstName: user.State().Profile.FirstName,
		LastName:  user.State().Profile.LastName,
		Phone:     user.State().Profile.Phone,
		LastLogin: user.State().LastLogin,
		DeletedAt: user.State().DeletedAt,
	}, nil

}
