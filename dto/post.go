package dto

import (
	"SociLinkApi/models"
	"time"
)

type PostResponseDto struct {
	Id         string
	User       UserResponseDto
	Content    string
	Images     []string
	Visibility string
	Likes      int
	Liked      bool
	CreatedAt  time.Time
}

type GetPostResponseDto struct {
	Post PostResponseDto
}

type CreatePostRequestDto struct {
	Content    string
	Images     []string
	Visibility string
}

type CreatePostResponseDto struct {
	Post PostResponseDto
}

type EditPostRequestDto struct {
	Id         string
	Content    string
	Images     []string
	Visibility string
}

type EditPostResponseDto struct {
	Post PostResponseDto
}

type SearchPostRequestDto struct {
	PaginationRequestDto
	Search string `form:"search"`
}

type SearchPostResponseDto struct {
	Posts []PostResponseDto
}

func PostToResponseDto(post models.Post, likes int, liked bool) PostResponseDto {
	if post.Images == nil {
		post.Images = []string{}
	}

	return PostResponseDto{
		Id:         post.ID.String(),
		User:       UserToResponseDto(post.User),
		Content:    post.Content,
		Images:     post.Images,
		Visibility: post.Visibility,
		Likes:      likes,
		Liked:      liked,
		CreatedAt:  post.CreatedAt,
	}
}
