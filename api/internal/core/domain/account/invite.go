package account

import (
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
)

type InviteEmail string

func NewInviteEmail(email string) (InviteEmail, error) {
	if email == "" {
		return "", errors.New("empty email")
	}

	if !strings.Contains(email, "@") {
		return "", errors.New("invalid email")
	}

	return InviteEmail(email), nil
}

func convertStringToTime(timeString string) (time.Time, error) {
	t, err := time.Parse(time.RFC3339, timeString)
	if err != nil {
		return time.Time{}, err
	}

	return t, nil
}

type InviteCreatedAt time.Time

func NewInviteCreatedAt(createdAt string) (InviteCreatedAt, error) {
	t, err := convertStringToTime(createdAt)
	if err != nil {
		return InviteCreatedAt(time.Time{}), err
	}

	return InviteCreatedAt(t), nil
}

type InviteUpdatedAt time.Time

func NewInviteUpdatedAt(updatedAt string) (InviteUpdatedAt, error) {
	t, err := convertStringToTime(updatedAt)
	if err != nil {
		return InviteUpdatedAt(time.Time{}), err
	}

	return InviteUpdatedAt(t), nil
}

type InviteDeletedAt time.Time

func NewInviteDeletedAt(deletedAt string) (InviteDeletedAt, error) {
	t, err := convertStringToTime(deletedAt)
	if err != nil {
		return InviteDeletedAt(time.Time{}), err
	}

	return InviteDeletedAt(t), nil
}

type Invite struct {
	ID        uuid.UUID
	AccountID uuid.UUID
	Email     InviteEmail
	CreatedAt InviteCreatedAt
	UpdatedAt InviteUpdatedAt
	DeletedAt InviteDeletedAt
}

func NewInvite(accountId uuid.UUID, email InviteEmail, createdAt InviteCreatedAt, updatedAt InviteUpdatedAt, deletedAt InviteDeletedAt) Invite {
	return Invite{
		ID:        uuid.New(),
		AccountID: accountId,
		Email:     email,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		DeletedAt: deletedAt,
	}
}
