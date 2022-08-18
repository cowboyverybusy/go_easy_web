package data

import "time"

type Movie struct {
	ID          int       `json:"id,omitempty"`        //omitempty代表当为空值时忽略字段，即隐藏字段
	Version     int32     `json:"-"`                   //-代表隐藏不展示字段
	TicketPrice float32   `json:"ticket_price,string"` //string代表输出字符串格式(只有int/uint/float/bool才能转换为字符串)
	Title       string    `json:"title"`
	CreateAt    time.Time `json:"create_time"`
	Runtime     Runtime   `json:"runtime,omitempty"`
}
