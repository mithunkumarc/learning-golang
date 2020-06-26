package main
import (
	"fmt"
	"time"
)

func main() {
	start_time := time.Now()
	time.Sleep(2000000000) // delay 2 seconds
	stop_time := time.Now()
	elapsed_time := stop_time.Sub(start_time)     // subtracting stop time and start time
	fmt.Println(elapsed_time) // 2s
}
