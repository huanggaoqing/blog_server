package response

type PagingResponse[T interface{}] struct {
	CurrentPage int `json:"currentPage"`
	List        T   `json:"list"`
	Size        int `json:"size"`
	Total       int `json:"total"`
}

func NewPaging[T interface{}](page int, size int, list T, total int) *PagingResponse[T] {
	return &PagingResponse[T]{
		CurrentPage: page,
		Size:        size,
		List:        list,
		Total:       total,
	}
}
