package dto

import (
	"SociLinkApi/models"
	"time"
)

type AnswerFriendshipRequestDto struct {
	RequestId string
	Answer    bool
}

type FriendshipResponseDto struct {
	Id        string
	Friend    UserResponseDto
	Accepted  bool
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

func FriendshipToResponseDto(friendship models.Friendship) FriendshipResponseDto {
	return FriendshipResponseDto{
		Id:        friendship.ID.String(),
		Friend:    UserToResponseDto(friendship.Friend),
		Accepted:  friendship.Accepted,
		CreatedAt: friendship.CreatedAt,
	}
}
