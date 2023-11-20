package dto

import (
	"time"
)

type AnswerFriendshipRequestDto struct {
	RequestId string
	Answer    bool
}

type FriendshipResponseDto struct {
	Id        string
	Friend    UserResponseDto
	CreatedAt time.Time
}

type GetFriendsResponseDto struct {
	Friends []FriendshipResponseDto
}

type GetFriendshipResponseDto struct {
	Friendship FriendshipResponseDto
}

type GetFriendshipRequestsResponseDto struct {
	Requests []FriendshipResponseDto
}
