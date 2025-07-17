package common

type ListReq struct {
	BaseListParam
	Page     int `json:"page" form:"page"`
	PageSize int `json:"pageSize" form:"pageSize"`
}

func (l *ListReq) Adjust() {
	if l.Page == 0 {
		l.Page = 1
	}
	if l.PageSize == 0 {
		l.PageSize = 10
	}
}

func (l *ListReq) GetOffset() int {
	return (l.Page - 1) * l.PageSize
}

//------------- resp -------------//

type ListResp struct {
	BaseListResp
	Total     int64 `json:"total"`
	Page      int   `json:"page"`
	PageSize  int   `json:"pageSize"`
	TotalPage int   `json:"totalPage"`
}

func (l *ListResp) Adjust() {
	if l.PageSize == 0 {
		l.TotalPage = 0
		return
	}
	l.TotalPage = (int(l.Total) + l.PageSize - 1) / l.PageSize
}

func (l *ListResp) GetTotalPage() int {
	if l.PageSize == 0 {
		return 0
	}
	return (int(l.Total) + l.PageSize - 1) / l.PageSize
}

type IdReq struct {
	Id int64 `json:"id" binding:"required"`
}
