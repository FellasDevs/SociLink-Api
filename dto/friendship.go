package dto

import (
	"github.com/google/uuid"
	"time"
)

type AnswerFriendshipRequestDto struct {
	RequestId string
	Answer    bool
}

type GetFriendshipRequestsResponseDto struct {
	Id        uuid.UUID
	User      UserResponseDto
	CreatedAt time.Time
}
