package get_session

import "gym-management/src/lib/primitives/application_specific"

type GetSessionQuery struct {
	Token   string
	Session *application_specific.Session
}
