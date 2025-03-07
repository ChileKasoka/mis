package util

import (
	"fmt"
	"net/http"

	"github.com/ChileKasoka/mis/middleware"
	"github.com/google/uuid"
)

func ExtractUserId(r *http.Request) (uuid.UUID, error) {
	userIDStr, ok := r.Context().Value(middleware.UserIDKey).(string)
	if !ok || userIDStr == "" {
		return uuid.Nil, fmt.Errorf("unauthorized: missing or invalid user ID")

	}

	// Convert userID to uuid.UUID
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return uuid.Nil, fmt.Errorf("format invalid")

	}
	return userID, nil
}
