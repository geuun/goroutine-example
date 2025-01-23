package post

/**
Go의 타입 가시성 규칙
- exported(대문자로 시작) 타입은 패키지 외부에서 접근 가능
- unexported(소문자로 시작) 타입은 패키지 내부에서만 접근 가능
- exported 인터페이스의 메서드는 unexported 타입을 반환할 수 없음

인터페이스의 경우
exported 된 인터페이스는 다른 패키지에서 구현해야 할 수도 있다.
만약 exported 인터페이스의 메서드가 unexported 타입을 반환한다면, 다른 패키지에서 이를 구현할 수 없게 된다.

구조체 메서드의 경우
구조체의 메서드는 같은 패키지 내에서 구현된다.
다른 패키지는 이 메서드를 `사용`만 할 뿐이고, `구현`할 필요가 없다.
따라서 exported 구조체의 메서드가 unexported 타입을 반환하는 것은 문제가 되지 않는다.
반환된 unexported 타입은 불투명 타입(opaque type)처럼 동작하며, 외부 패키지에서는 이 타입의 내부를 알 수 없다.
*/

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

type PostBuilder struct {
	post *post
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

func (b *PostBuilder) Id(id int) *PostBuilder {
	b.post.id = id
	return b
}

func (b *PostBuilder) UserId(userId int) *PostBuilder {
	b.post.userId = userId
	return b
}

func (b *PostBuilder) Title(title string) *PostBuilder {
	b.post.title = title
	return b
}

func (b *PostBuilder) Body(body string) *PostBuilder {
	b.post.body = body
	return b
}

func (b *PostBuilder) Build() *post {
	return b.post
}

func NewPostBuilder() *PostBuilder {
	return &PostBuilder{
		post: &post{},
	}
}

// FromResponse
// 단일 Response 를 Entity로 변환
func FromResponse(response *Response) *post {
	return NewPostBuilder().
		Id(response.Id).
		UserId(response.UserId).
		Title(response.Title).
		Body(response.Body).
		Build()
}

// FromResponses
// Response 배열을 Entity 배열로 변환
func FromResponses(responses []Response) []*post {
	var posts []*post
	for _, resp := range responses {
		posts = append(posts, FromResponse(&resp))
	}
	return posts
}
