package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"hbbapi/util/dingding"
	"io/ioutil"
	"runtime"
	"strings"
)

type PanicLog struct {
	TraceId  string   `json:"trace_id"`
	ErrStack []string `json:"err_stack"`
}

// Recovery 重写Recovery代码，支持报错的提醒
func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				if strings.Contains(strings.ToLower(err.(error).Error()), "broken pipe") ||
					strings.Contains(strings.ToLower(err.(error).Error()), "connection reset by peer") {
					return
				}
				traceId := c.GetString("traceId")
				errStack := stack(3)
				var panicLog = PanicLog{
					TraceId:  traceId,
					ErrStack: errStack,
				}
				jsonStr, _ := json.Marshal(panicLog)
				fmt.Println(string(jsonStr))
				_ = dingding.Ding{}.AtAll().
					SetTitle("gin出现panic").
					SetContent(err.(error).Error() + "\n traceId:" + traceId).
					Send()
			}
		}()
		c.Next()
	}
}

func stack(skip int) []string {
	var errString []string
	// As we loop, we open files and read them. These variables record the currently
	// loaded file.
	var lines [][]byte
	var lastFile string
	for i := skip; ; i++ { // Skip the expected number of frames
		pc, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		// Print this much at least.  If we can't find the source, it won't show.
		errString = append(errString, fmt.Sprintf("%s:%d (0x%x)\n", file, line, pc))
		if file != lastFile {
			data, err := ioutil.ReadFile(file)
			if err != nil {
				continue
			}
			lines = bytes.Split(data, []byte{'\n'})
			lastFile = file
		}
		errString = append(errString, fmt.Sprintf("\t%s: %s\n", function(pc), source(lines, line)))
	}
	return errString
}

var (
	dunno     = []byte("???")
	centerDot = []byte("·")
	dot       = []byte(".")
	slash     = []byte("/")
)

// source returns a space-trimmed slice of the n'th line.
func source(lines [][]byte, n int) []byte {
	n-- // in stack trace, lines are 1-indexed but our array is 0-indexed
	if n < 0 || n >= len(lines) {
		return dunno
	}
	return bytes.TrimSpace(lines[n])
}

// function returns, if possible, the name of the function containing the PC.
func function(pc uintptr) []byte {
	fn := runtime.FuncForPC(pc)
	if fn == nil {
		return dunno
	}
	name := []byte(fn.Name())
	// The name includes the path name to the package, which is unnecessary
	// since the file name is already included.  Plus, it has center dots.
	// That is, we see
	//	runtime/debug.*T·ptrmethod
	// and want
	//	*T.ptrmethod
	// Also the package path might contains dot (e.g. code.google.com/...),
	// so first eliminate the path prefix
	if lastSlash := bytes.LastIndex(name, slash); lastSlash >= 0 {
		name = name[lastSlash+1:]
	}
	if period := bytes.Index(name, dot); period >= 0 {
		name = name[period+1:]
	}
	name = bytes.Replace(name, centerDot, dot, -1)
	return name
}
