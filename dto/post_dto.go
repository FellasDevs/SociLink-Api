package dto

type CreatePostRequestDto struct {
	Content    string
	Images     []string
	Visibility string
}
type EditPostRequestDto struct {
	Id         string
	Content    string
	Images     []string
	Visibility string
}
type GetPostResponseDto struct {
	Id         string
	User       PayloadUser
	Content    string
	Images     []string
	Visibility string
}
