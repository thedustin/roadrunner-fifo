package fifo

import "github.com/thedustin/roadrunner-fifo/impl"

// Wrapper for the plugin
type rpc struct {
	fifo impl.Fifo
}

func (r *rpc) Get(key string, out *string) error {
	value, ok := r.fifo.Get(key)

	if !ok {
		return nil
	}

	*out = value

	return nil
}

func (r *rpc) Set(key string, value string, out *bool) error {
	if err := r.fifo.Set(key, value); err != nil {
		*out = false
	}

	return nil
}

func (r *rpc) Invalidate(key string, out *bool) error {
	if ok := r.fifo.Invalidate(key); !ok {
		*out = false
		return nil
	}

	*out = true
	return nil
}
