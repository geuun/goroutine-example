package post

type post struct {
	id     int
	userId int
	title  string
	body   string
}

type Post interface {
	GetId() int
	GetUserId() int
	GetTitle() string
	GetBody() string
}

func (p *post) GetId() int {
	return p.id
}

func (p *post) GetUserId() int {
	return p.userId
}

func (p *post) GetTitle() string {
	return p.title
}

func (p *post) GetBody() string {
	return p.body
}
