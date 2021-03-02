package engine

func (e *Engine) RunSync() {
	e.run(true)
}
