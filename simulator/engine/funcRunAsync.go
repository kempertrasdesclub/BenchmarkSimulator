package engine

func (e *Engine) RunAsync() {
	e.run(false)
}
