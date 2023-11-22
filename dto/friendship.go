package dto

import (
	"SociLinkApi/models"
	types "SociLinkApi/types/pagination"
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
	types.PaginationResponse
	Friends []FriendshipResponseDto
}

type GetFriendshipResponseDto struct {
	Friendship FriendshipResponseDto
}

type GetFriendshipRequestsResponseDto struct {
	types.PaginationResponse
	Requests []FriendshipResponseDto
}

func FriendshipToFriendshipResponseDto(friendship models.Friendship) FriendshipResponseDto {
	return FriendshipResponseDto{
		Id:        friendship.ID.String(),
		Friend:    UserToUserResponseDto(friendship.Friend),
		Accepted:  friendship.Accepted,
		CreatedAt: friendship.CreatedAt,
	}
}
