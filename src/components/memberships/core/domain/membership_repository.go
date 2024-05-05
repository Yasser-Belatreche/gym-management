package domain

import "gym-management/src/lib/primitives/application_specific"

type MembershipRepository interface {
	FindByID(id string, session *application_specific.Session) (*Membership, *application_specific.ApplicationException)

	FindByCode(code string, session *application_specific.Session) (*Membership, *application_specific.ApplicationException)

	Create(membership *Membership, session *application_specific.Session) *application_specific.ApplicationException

	Update(membership *Membership, session *application_specific.Session) *application_specific.ApplicationException

	FindEnabledMemberships(session *application_specific.Session) ([]*Membership, *application_specific.ApplicationException)

	FindEnabledMembershipsByGymID(gymId string, session *application_specific.Session) ([]*Membership, *application_specific.ApplicationException)

	CountTrainingSessionsThisWeek(membershipId string, session *application_specific.Session) (int, *application_specific.ApplicationException)
}
