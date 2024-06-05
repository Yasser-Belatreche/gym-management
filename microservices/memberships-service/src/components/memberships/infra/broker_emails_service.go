package infra

import (
	"gym-management-memberships/src/lib/messages_broker"
	"gym-management-memberships/src/lib/primitives/application_specific"
)

type BrokerEmailsService struct {
	Broker messages_broker.MessagesBroker
}

func (p *BrokerEmailsService) IsUsed(email application_specific.Email, session *application_specific.Session) bool {
	res, err := p.Broker.GetReply("Emails.IsUsed", map[string]string{
		"email": email.Value,
	}, session)

	if err != nil {
		return true
	}

	if res["used"] == "true" {
		return true
	} else if res["used"] == "false" {
		return false
	}

	return true
}
