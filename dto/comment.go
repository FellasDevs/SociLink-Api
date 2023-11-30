package dto

import (
	"SociLinkApi/models"
	"time"
)

type CommentResponseDto struct {
	Id        string
	User      UserResponseDto
	Content   string
	CreatedAt time.Time
}

type GetPostCommentsRequestDto struct {
	PaginationRequestDto
	PostId string `form:"postId"`
}

type GetPostCommentsResponseDto struct {
	Comments []CommentResponseDto
}

type CreateCommentRequestDto struct {
	PostId  string
	Content string
}

type EditCommentRequestDto struct {
	Content string `json:"content"`
}

func CommentToResponseDto(comment models.Comment) CommentResponseDto {
	return CommentResponseDto{
		Id:        comment.ID.String(),
		User:      UserToResponseDto(comment.User),
		Content:   comment.Content,
		CreatedAt: comment.CreatedAt,
	}
}
