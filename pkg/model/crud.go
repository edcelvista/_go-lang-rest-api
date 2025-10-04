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

type DBRecordList struct {
	RecordList []DBRecord
	Count      int64
}

type DBRecord struct {
	RecordId string            `bson:"_id,omitempty"`
	Record   map[string]string `bson:"record,omitempty"`
	Dt       int64             `bson:"dt,omitempty"`
}

type DBRecordDeleted struct {
	DeletedCount int64
}

type Number interface {
	int | int32 | int64
}

type FindOptions struct {
	Limit int
}
