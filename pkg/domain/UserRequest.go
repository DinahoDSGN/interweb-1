package domain

import (
	"fmt"
	"time"
)

type UserRequest struct {
	ID          int64     `json:"id,omitempty"`
	ChatID      int64     `json:"chat_id,omitempty"`
	Request     string    `json:"request,omitempty"`
	RequestDate time.Time `json:"request_date"`
	Result      []byte    `json:"result,omitempty"`
}

func (ur UserRequest) String() string {
	if len(ur.Result) > 20 {
		return fmt.Sprintf("ID:%d UserID:%d Command:%s RequestDate:%v - Result [%s]...", ur.ID, ur.ChatID, ur.Request, ur.RequestDate, ur.Result[:20])
	}

	return fmt.Sprintf("ID:%d UserID:%d Command:%s RequestDate:%v - Result [%s]", ur.ID, ur.ChatID, ur.Request, ur.RequestDate, ur.Result)
}

type TotalUserRequests struct {
	Request string
	Count   uint64
}

func (t TotalUserRequests) String() string {
	return fmt.Sprintf("%s - %d", t.Request, t.Count)
}
