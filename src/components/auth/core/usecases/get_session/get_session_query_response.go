package get_session

import "gym-management/src/lib/primitives/application_specific"

type GetSessionQueryResponse struct {
	Session *application_specific.UserSession
}
