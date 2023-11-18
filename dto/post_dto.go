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
