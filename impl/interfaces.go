package impl

type Fifo interface {
	Set(key string, value string) error
	Get(key string) (string, bool)
	Invalidate(key string) bool
}
