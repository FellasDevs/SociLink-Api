package dto

import (
	"github.com/google/uuid"
	"time"
)

type AnswerFriendshipRequestDto struct {
	RequestId string
	Answer    bool
}

type FriendshipResponseDto struct {
	Id        uuid.UUID
	User      UserResponseDto
	CreatedAt time.Time
}

type GetFriendsResponseDto struct {
	Friends []FriendshipResponseDto
}

type GetFriendshipRequestsResponseDto struct {
	Requests []FriendshipResponseDto
}
