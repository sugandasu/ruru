package jongi

import (
	"context"
	"errors"
	"fmt"
)

var defaultPolicy *policy

type ValidateFunc func(ctx context.Context, data any) error

type Action struct {
	Action string
	Object string
	Run    ValidateFunc
}

type policy struct {
	Actions map[string]Action
}

type Policy interface {
	RegisterPolicy(action string, object string, run ValidateFunc)
	AuthorizedTo(ctx context.Context, action string, object string, data any) error
}

func (p *policy) AuthorizedTo(ctx context.Context, action string, object string, data any) error {
	act, ok := p.Actions[action]
	if !ok {
		return errors.New("unauthorized")
	}

	return act.Run(ctx, data)
}

func (p *policy) RegisterPolicy(action string, object string, run ValidateFunc) {
	key := fmt.Sprintf("%s.%s", action, object)
	p.Actions[key] = Action{
		Action: action,
		Object: object,
		Run:    run,
	}
}

func NewPolicy() Policy {
	pol := &policy{
		Actions: make(map[string]Action),
	}

	defaultPolicy = pol

	return pol
}
