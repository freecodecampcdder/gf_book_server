package model

type Address struct {
	Id             int64  `json:"id"`
	ReceiverName   string `json:"receiver_name"`
	ReceiverPhone  string `json:"receiver_phone"`
	AddressContent string `json:"address_content"`
	Status         uint   `json:"status"`
}
