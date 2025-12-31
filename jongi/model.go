package jongi

import "github.com/golang-jwt/jwt/v5"

type AuthClaims struct {
	UserID    string   `json:"user_id"`
	Role      AuthRole `json:"roles"`
	CompanyID string   `json:"company_id"`
	jwt.RegisteredClaims
}

type AuthRole struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Level    int    `json:"level"`
	Policies any    `json:"policies"`
}

type AuthContext struct{}
