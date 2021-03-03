package engine

// RunSync (Português): Roda todos os testes de forma linear.
func (e *Engine) RunSync() (err error) {
	return e.run(true)
}
