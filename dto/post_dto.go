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
	CreatedAt  time.Time
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

type SearchPostResponseDto struct {
	Posts []PostResponseDto
}

func PostToPostResponseDto(post models.Post) PostResponseDto {
	return PostResponseDto{
		Id:         post.ID.String(),
		User:       UserToUserResponseDto(post.User),
		Content:    post.Content,
		Images:     post.Images,
		Visibility: post.Visibility,
		CreatedAt:  post.CreatedAt,
	}
}
