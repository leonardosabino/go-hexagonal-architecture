package model

const MAX_LIMIT = 200

type Pageable struct {
	TotalRowCount int64       `json:"totalRowCount"`
	PageSize      int         `json:"pageSize"`
	PageNumber    int         `json:"pageNumber"`
	List          interface{} `json:"list"`
}

func ToPageable(list interface{}, pagesize int, count int64, pageNumber int) Pageable {
	if pagesize == 0 {
		list = make([]string, 0)
	}
	return Pageable{
		TotalRowCount: count,
		PageSize:      pagesize,
		PageNumber:    pageNumber,
		List:          list,
	}
}
