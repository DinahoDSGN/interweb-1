package logger

type Logger interface {
	Info(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
}

var impl Logger

func SetLogger(logger Logger) {
	impl = logger
}

func Info(args ...interface{}) {
	impl.Info(args)
}

func Error(args ...interface{}) {
	impl.Error(args)
}

func Fatal(args ...interface{}) {
	impl.Fatal(args)
}
