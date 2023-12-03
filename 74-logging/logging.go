package main

import (
	"bytes"
	"fmt"
	"log"
	"log/slog"
	"os"
)

func main() {
	log.Println("msg from standard logger")

	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Llongfile | log.LUTC)
	log.Println("msg from standard logger with flags")

	mylog := log.New(os.Stdout, "mylog:", log.LstdFlags|log.Lshortfile|log.Lmsgprefix)
	mylog.Println("msg from mylog")

	mylog.SetPrefix("new-prefix-mylog:")
	mylog.Println("msg from mylog with a new prefix")
	fmt.Println()

	var buf bytes.Buffer
	buflog := log.New(&buf, "buflog:", log.LstdFlags)
	buflog.Println("hello from buflog")
	fmt.Println("buflog wrote to buf:")
	fmt.Print(buf.String())
	fmt.Println()

	jsonHandler := slog.NewJSONHandler(os.Stderr, nil)
	myslog := slog.New(jsonHandler)
	myslog.Info("hi from myslog, a json-formatting logger")
	myslog.Info("hi again from myslog, with key-value pairs", "name", "Chris", "age", 35)
}
