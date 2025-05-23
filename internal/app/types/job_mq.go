package types

type GetJobMqOneReq struct {
	Id         int64 `json:"id"`
	Type       int   `json:"type"`
	MainUserId int64 `json:"mainUserId"`
}
