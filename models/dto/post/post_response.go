package post

type PostResponse struct {
	Data []PostDto `json:"data"`
}

func NewPostResponse(dtos []PostDto) *PostResponse {
	return &PostResponse{
		Data: dtos,
	}
}

// VO 슬라이스를 DTO 슬라이스로 변환하는 함수
func ToDto(vos []PostVo) []PostDto {
	dtos := make([]PostDto, len(vos))
	for i, vo := range vos {
		dtos[i] = PostDto{
			UserId: vo.getUserId(),
			Id:     vo.getId(),
			Title:  vo.getTitle(),
			Body:   vo.getBody(),
		}
	}
	return dtos
}
