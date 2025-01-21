package post

type PostDto struct {
	UserId int    `json:"userId"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

// DTO를 VO로 변환하는 함수
func (dto PostDto) ToVo() PostVo {
	return &innerPostVo{
		UserId: dto.UserId,
		Id:     dto.Id,
		Title:  dto.Title,
		Body:   dto.Body,
	}
}
