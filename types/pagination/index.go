package types

import (
	"SociLinkApi/models"
)

type PaginationResponse struct {
	Page       int
	PageSize   int
	TotalCount int
}

type FriendshipListing struct {
	PaginationResponse
	Friendships []models.Friendship
}

type PostListing struct {
	PaginationResponse
	Posts []models.Post
}

type UserListing struct {
	PaginationResponse
	Users []models.User
}
