package log

type Log struct {
	Messages map[string]string
}

func (l *Log) Info(key, value string) {
	l.Messages[key] = l.Messages[key] + "|" + value
}

func (l *Log) Error(value string) {
	l.Messages["error"] = l.Messages["error"] + "|" + value

}
