package server

type FileLogger struct {
	filename string
}

func (fl FileLogger) Print(...interface{}) {
	return
}

func (fl FileLogger) Info(...interface{}) {
	return
}

func (fl FileLogger) Warn(...interface{}) {
	return
}
