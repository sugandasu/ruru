package jongi

import (
	"context"
	"errors"
)

func AuthorizedTo(ctx context.Context, action string, object string, data any) error {
	if defaultPolicy == nil {
		return errors.New("policy does not exists")
	}

	act, ok := defaultPolicy.Actions[action]
	if !ok {
		return errors.New("unauthorized")
	}

	return act.Run(ctx, data)
}
