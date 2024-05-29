package get_session

import (
	"gym-management/src/components/auth/core/domain"
	"gym-management/src/lib/primitives/application_specific"
)

type GetSessionQueryHandler struct {
	UserRepository domain.UserRepository
	JwtSecret      string
}

func (h *GetSessionQueryHandler) Handle(query *GetSessionQuery) (*GetSessionQueryResponse, *application_specific.ApplicationException) {
	claims, err := domain.DecodeToken(query.Token, h.JwtSecret)
	if err != nil {
		return nil, err
	}

	user, err := h.UserRepository.FindByID(claims.UserId, query.Session)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, application_specific.NewAuthenticationException("AUTH.TOKEN.INVALID", "Invalid Token", map[string]string{
			"token": query.Token,
		})
	}

	session, err := user.GetSession(query.Session)
	if err != nil {
		return nil, err
	}

	return &GetSessionQueryResponse{
		Session: session,
	}, nil
}
