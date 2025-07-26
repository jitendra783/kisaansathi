package controller

import (
	"context"
	"kisaanSathi/pkg/logger"
)

func (s *repoObject) Validate(ctx context.Context, sessionId string, userId string) (bool, error) {
	logger.Log(ctx).Debug("START")
	defer logger.Log(ctx).Debug("END")

	// write business logic here
	flag, err := s.db.ValidateSession(ctx, sessionId, userId)
	return flag, err
}
