package engine

// RunAsync (Português): Roda todos os testes por goroutines
func (e *Engine) RunAsync() (err error) {
	return e.run(false)
}
