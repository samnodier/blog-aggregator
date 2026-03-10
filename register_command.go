package main

func (c *commands) register(name string, f func(*state, command) error) {
	c.handlers[name] = f
}
