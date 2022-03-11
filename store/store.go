package store
type Factory interface {
	Users() UserStore
	Activity() ActivityStore
	Close() error
	MessageTemplate() MessageTemplate
}
