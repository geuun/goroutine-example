package post

type PostVo interface {
	getUserId() int
	getId() int
	getTitle() string
	getBody() string
}

/**
 * PostVo의 생성자
 */
func NewPostVo(userId int, id int, title string, body string) PostVo {
	return &innerPostVo{
		UserId: userId,
		Id:     id,
		Title:  title,
		Body:   body,
	}
}

/**
 * PostVo의 구현체
 */
type innerPostVo struct {
	UserId int    `json:"userId"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func (postVo innerPostVo) getUserId() int {
	return postVo.UserId
}

func (postVo innerPostVo) getId() int {
	return postVo.Id
}

func (postVo innerPostVo) getTitle() string {
	return postVo.Title
}

func (postVo innerPostVo) getBody() string {
	return postVo.Body
}
