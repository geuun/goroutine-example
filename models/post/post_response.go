package post

type Response struct {
	Id     int    `json:"id"`
	UserId int    `json:"userId"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

type ResponseBuilder struct {
	response *Response
}

func NewResponseBuilder() *ResponseBuilder {
	return &ResponseBuilder{
		response: &Response{},
	}
}

func (b *ResponseBuilder) GetId() int {
	return b.response.Id
}

func (b *ResponseBuilder) Id(id int) *ResponseBuilder {
	b.response.Id = id
	return b
}

func (b *ResponseBuilder) GetUserId() int {
	return b.response.UserId
}

func (b *ResponseBuilder) UserId(userId int) *ResponseBuilder {
	b.response.UserId = userId
	return b
}

func (b *ResponseBuilder) GetTitle() string {
	return b.response.Title
}

func (b *ResponseBuilder) Title(title string) *ResponseBuilder {
	b.response.Title = title
	return b
}

func (b *ResponseBuilder) GetBody() string {
	return b.response.Body
}

func (b *ResponseBuilder) Body(body string) *ResponseBuilder {
	b.response.Body = body
	return b
}

func (b *ResponseBuilder) Build() *Response {
	return b.response
}
