package lib

type Courier[K comparable, T any] struct {
	Packet map[K]T
}

func (c *Courier[K, T]) Add(key K, value T) {
	c.Packet[key] = value
}

func (c *Courier[K, T]) Get(key K) T {
	return c.Packet[key]
}
