// go mod init my-protobuf
// my-protobuf> protoc --go_out=. ./proto/basic/*.proto
// my-protobuf> protoc --go_opt=module=my-protobuf --go_out=. ./proto/basic/*.proto
// go mod tidy
// go mod vendor
package main

import (
	"fmt"
	"log"
	"my-protobuf/basic"
	"time"
)

type logWriter struct {
}

// log wrapper
func (writer logWriter) Write(bytes []byte) (int, error) {
	return fmt.Print(time.Now().Format("15:04:05") + " " + string(bytes))
}

func main() {
	log.SetFlags(0)
	log.SetOutput(new(logWriter))
	basic.BasicHello()
}
