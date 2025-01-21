package post

// https://jsonplaceholder.typicode.com/posts

type PostResponse []PostVo

func NewPostResponse(posts []PostVo) PostResponse {
	return PostResponse(posts)
}
