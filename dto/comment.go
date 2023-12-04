// dto.go
package dto

import (
	"SociLinkApi/models"
	"time"
)

type CreateCommentReplyRequestDto struct {
	CommentId string `json:"commentId"`
	Content   string `json:"content"`
}
type CommentResponseDto struct {
	Id        string
	User      UserResponseDto
	Content   string
	CreatedAt time.Time
	Replies   []CommentReplyResponseDto
}

type CommentReplyResponseDto struct {
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
	var replyDtos []CommentReplyResponseDto
	for _, reply := range comment.Replies {
		replyDto := CommentReplyToResponseDto(reply)
		replyDtos = append(replyDtos, replyDto)
	}
	return CommentResponseDto{
		Id:        comment.ID.String(),
		User:      UserToResponseDto(comment.User),
		Content:   comment.Content,
		CreatedAt: comment.CreatedAt,
		Replies:   replyDtos,
	}
}

func CommentReplyToResponseDto(reply models.CommentReply) CommentReplyResponseDto {
	return CommentReplyResponseDto{
		Id:        reply.ID.String(),
		User:      UserToResponseDto(reply.User),
		Content:   reply.Content,
		CreatedAt: reply.CreatedAt,
	}
}
