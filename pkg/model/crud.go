package model

type CrudReq struct {
	MessageId string
	Message   map[string]string
}

type CrudPostReq struct {
	Message map[string]string
	Dt      int64
}

type CrudRes struct {
	Action    string
	MessageId string
}

type CrudResWithBody struct {
	Action    string
	MessageId string
	Response  interface{}
}

type DBMessageRecordList struct {
	MessageList []DBMessageRecord
	Count       int64
}

type DBMessageRecord struct {
	MessageId string            `bson:"_id,omitempty"`
	Message   map[string]string `bson:"message,omitempty"`
	Dt        int64             `bson:"dt,omitempty"`
}

type DBMessageDeleted struct {
	DeletedCount int64
}

type Number interface {
	int | int32 | int64
}

type FindOptions struct {
	Limit int
}
