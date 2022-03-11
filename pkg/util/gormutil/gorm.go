package gormutil

const DefaultLimit = 1000

// LimitAndOffset contains offset and limit fields.
type LimitAndOffset struct {
	Offset int
	Limit  int
}

// Unpointer fill LimitAndOffset with default values if offset/limit is nil
// or it will be filled with the passed value.
func Unpointer(offset *int64, limit *int64) *LimitAndOffset {
	var o, l int = 0, DefaultLimit

	if offset != nil {
		o = int(*offset)
	}

	if limit != nil {
		l = int(*limit)
	}

	return &LimitAndOffset{
		Offset: o,
		Limit:  l,
	}
}
