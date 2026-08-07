package service

import "context"

func (s *AuthService)Health(ctx context.Context) (string, error){
	return "up", nil
}