package dto

type PostResponseDto struct {
	Id         string
	User       UserResponseDto
	Content    string
	Images     []string
	Visibility string
}

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
