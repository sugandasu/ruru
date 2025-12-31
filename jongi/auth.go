package jongi

import (
	"context"
)

func GetAuthFromContext(ctx context.Context) *AuthClaims {
	auth := ctx.Value(AuthContext{})
	if auth == nil {
		return nil
	}
	claim, ok := auth.(*AuthClaims)
	if !ok {
		return nil
	}
	return claim
}
