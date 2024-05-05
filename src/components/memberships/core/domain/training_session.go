package domain

import (
	"gym-management/src/lib/primitives/application_specific"
	"gym-management/src/lib/primitives/generic"
	"time"
)

type TrainingSession struct {
	id        string
	startedAt time.Time
	endedAt   *time.Time
}

type TrainingSessionState struct {
	Id        string
	StartedAt time.Time
	EndedAt   *time.Time
}

func CreateTrainingSession() *TrainingSession {
	return &TrainingSession{
		id:        generic.GenerateUUID(),
		startedAt: time.Now(),
		endedAt:   nil,
	}
}

func TrainingSessionFromState(state *TrainingSessionState) *TrainingSession {
	return &TrainingSession{
		id:        state.Id,
		startedAt: state.StartedAt,
		endedAt:   state.EndedAt,
	}
}

func (t *TrainingSession) End() *application_specific.ApplicationException {
	if t.endedAt != nil {
		return application_specific.NewValidationException("MEMBERSHIPS.TRAINING_SESSIONS.ALREADY_ENDED", "Training session has already ended", map[string]string{
			"id": t.id,
		})
	}

	endedAt := time.Now()
	t.endedAt = &endedAt

	return nil
}

func (t *TrainingSession) State() *TrainingSessionState {
	return &TrainingSessionState{
		Id:        t.id,
		StartedAt: t.startedAt,
		EndedAt:   t.endedAt,
	}
}
