package common

type IBaseListParam interface {
	Adjust()
}

type IBaseListResp interface {
	Adjust()
}

type BaseListParam struct {
}

func (s *BaseListParam) Adjust() {

}

type BaseListResp struct {
}

func (s *BaseListResp) Adjust() {

}
