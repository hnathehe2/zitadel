package view

import (
	"github.com/zitadel/zitadel/internal/eventstore"
	"github.com/zitadel/zitadel/internal/repository/user"
	"github.com/zitadel/zitadel/internal/zerrors"
)

func UserByIDQuery(id, instanceID string, sequence uint64, eventTypes []eventstore.EventType) (*eventstore.SearchQueryBuilder, error) {
	if id == "" {
		return nil, zerrors.ThrowPreconditionFailed(nil, "EVENT-d8isw", "Errors.User.UserIDMissing")
	}
	return eventstore.NewSearchQueryBuilder(eventstore.ColumnsEvent).
		AwaitOpenTransactions().
		InstanceID(instanceID).
		SequenceGreater(sequence).
		AddQuery().
		AggregateTypes(user.AggregateType).
		AggregateIDs(id).
		EventTypes(eventTypes...).
		Builder(), nil
}
