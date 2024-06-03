package registered_replies

import (
	"gym-management/src/components/auth/core/domain"
	"gym-management/src/lib/messages_broker"
	"gym-management/src/lib/primitives/application_specific"
)

func BuildIsEmailUsedReply(userRepository domain.UserRepository) *messages_broker.Reply {
	var answer = &messages_broker.Reply{
		Message: "Emails.IsUsed",
		Handler: func(params map[string]string, session *application_specific.Session) (map[string]string, *application_specific.ApplicationException) {
			query, err := parseParams(params, session)
			if err != nil {
				return nil, err
			}

			response, err := isEmailUsed(userRepository, query)
			if err != nil {
				return nil, err
			}

			return parseResponse(response), nil
		},
	}

	return answer
}

func parseParams(params map[string]string, session *application_specific.Session) (*isEmailUsedQuery, *application_specific.ApplicationException) {
	if params["email"] == "" {
		return nil, application_specific.NewDeveloperException("ANSWERS.WRONG_PARAMS", "email is required in Emails.IsUsed question params")
	}

	return &isEmailUsedQuery{
		Email:   params["email"],
		Session: session,
	}, nil
}

func parseResponse(response *isEmailUsedQueryResponse) map[string]string {
	if response.IsUsed {
		return map[string]string{"used": "true"}
	}

	return map[string]string{"used": "false"}
}

type isEmailUsedQuery struct {
	Email   string
	Session *application_specific.Session
}

type isEmailUsedQueryResponse struct {
	IsUsed bool
}

func isEmailUsed(userRepository domain.UserRepository, query *isEmailUsedQuery) (*isEmailUsedQueryResponse, *application_specific.ApplicationException) {
	email, err := application_specific.NewEmail(query.Email)
	if err != nil {
		return nil, err
	}

	used, err := userRepository.UsernameUsed(domain.UsernameFromEmail(email), query.Session)
	if err != nil {
		return nil, err
	}

	return &isEmailUsedQueryResponse{
		IsUsed: used,
	}, nil
}
