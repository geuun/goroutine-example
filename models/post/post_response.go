package post

type Response struct {
	ID     int    `json:"id"`
	UserID int    `json:"userId"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

type IResponse interface {
	GetId() int
	GetUserId() int
	GetTitle() string
	GetBody() string
}

func (r *Response) GetId() int {
	return r.ID
}

func (r *Response) GetUserId() int {
	return r.UserID
}

func (r *Response) GetTitle() string {
	return r.Title
}

func (r *Response) GetBody() string {
	return r.Body
}

type responseBuilder struct {
	ID     int
	UserID int
	Title  string
	Body   string
}

type ResponseBuilder interface {
	SetId(id int) ResponseBuilder
	SetUserId(userId int) ResponseBuilder
	SetTitle(title string) ResponseBuilder
	SetBody(body string) ResponseBuilder
	Build() IResponse
}

func (rb *responseBuilder) SetId(id int) ResponseBuilder {
	rb.ID = id
	return rb
}

func (rb *responseBuilder) SetUserId(userId int) ResponseBuilder {
	rb.UserID = userId
	return rb
}

func (rb *responseBuilder) SetTitle(title string) ResponseBuilder {
	rb.Title = title
	return rb
}

func (rb *responseBuilder) SetBody(body string) ResponseBuilder {
	rb.Body = body
	return rb
}

func (rb *responseBuilder) Build() IResponse {
	return &Response{
		ID:     rb.ID,
		UserID: rb.UserID,
		Title:  rb.Title,
		Body:   rb.Body,
	}
}

func NewResponse() ResponseBuilder {
	return &responseBuilder{}
}
