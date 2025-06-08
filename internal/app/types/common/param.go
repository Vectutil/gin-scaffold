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

type IBaseParam interface {
	Adjust()
}

type IBaseResp interface {
	Adjust()
}

type BaseParam struct {
}

func (s *BaseParam) Adjust() {

}

type BaseResp struct {
}

func (s *BaseResp) Adjust() {

}
