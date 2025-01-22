package post

type Request interface {
	UserId() Request
	Title() Request
	Body() Request
}

type RequestBuilder interface {
	SetUserId(userId int) RequestBuilder
	SetTitle(title string) RequestBuilder
	SetBody(body string) RequestBuilder
	Build() (Request, error)
}

type postRequest struct {
	UserId int    `json:"userId"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

type postRequestBuilder struct {
	request *postRequest
}

func (r *postRequest) GetUserId() int {
	return r.UserId
}

func (r postRequest) GetTitle() string {
	return r.Title
}

func (r postRequest) GetBody() string {
	return r.Body
}
