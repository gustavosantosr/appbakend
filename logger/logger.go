package logger

import (
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/rs/xid"
)

/*WriteLogger create a logger */
func WriteLogger(writeErr string) string {
	f, err := os.OpenFile("Error.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	guid := xid.New()
	logger := log.New(f, "CodeError: "+guid.String()+" ", log.LstdFlags)
	_, fn, line, _ := runtime.Caller(1)
	logger.Println(fmt.Sprintf("Caller : %s, line: %d, Error: %s", fn, line, writeErr))
	return guid.String()
}
