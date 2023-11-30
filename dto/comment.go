package dto

import "SociLinkApi/models"

type CommentResponseDto struct {
	Id      string
	User    UserResponseDto
	Content string
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
		Id:      comment.ID.String(),
		User:    UserToResponseDto(comment.User),
		Content: comment.Content,
	}
}
