package engine

func (e *Engine) mountData() {
	for i := 0; i != e.sizeOfData; i += 1 {
		key, keyData := e.data.NewData()
		e.cache[key] = keyData
	}
}
