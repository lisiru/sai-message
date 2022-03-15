package store
type Factory interface {
	Users() UserStore
	Close() error
	MessageTemplate() MessageTemplate
	SmsRecord() SmsRecord
}
