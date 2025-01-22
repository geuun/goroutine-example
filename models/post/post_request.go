package post

type request struct {
	userId int    `json:"userId"`
	title  string `json:"title"`
	body   string `json:"body"`
}

type Request interface {
	UserId() int
	Title() string
	Body() string
}

func (r *request) UserId() int {
	return r.userId
}

func (r *request) Title() string {
	return r.title
}

func (r *request) Body() string {
	return r.body
}

type RequestBuilder interface {
	SetUserId(userId int) RequestBuilder
	SetTitle(title string) RequestBuilder
	SetBody(body string) RequestBuilder
	Build() Request
}

type requestBuilder struct {
	userId int
	title  string
	body   string
}

func (rb *requestBuilder) SetUserId(userId int) RequestBuilder {
	rb.userId = userId
	return rb
}

func (rb *requestBuilder) SetTitle(title string) RequestBuilder {
	rb.title = title
	return rb
}

func (rb *requestBuilder) SetBody(body string) RequestBuilder {
	rb.body = body
	return rb
}

func (rb *requestBuilder) Build() Request {
	return &request{
		userId: rb.userId,
		title:  rb.title,
		body:   rb.body,
	}
}

func NewRequest() RequestBuilder {
	return &requestBuilder{
		userId: 0,
		title:  "",
		body:   "",
	}
}
