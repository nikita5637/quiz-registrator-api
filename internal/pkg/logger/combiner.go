package logger

import "io"

// Combiner ...
type Combiner struct {
	ws []io.Writer
}

// WithWriter ...
func (c *Combiner) WithWriter(w io.Writer) *Combiner {
	c.ws = append(c.ws, w)
	return c
}

// Write ...
func (c *Combiner) Write(b []byte) (int, error) {
	for _, w := range c.ws {
		_, _ = w.Write(b)
	}
	return len(b), nil
}
