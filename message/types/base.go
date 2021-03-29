package types

//Base ...
type Base struct {
	identity string
}

//Identity ...
func (b *Base) Identity() string {
	return b.identity
}
