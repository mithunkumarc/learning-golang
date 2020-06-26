package main

import (
	"fmt"
	"time"
)

func main() {
	// telling golang compiler that input will be in this format 
	// below is a sample format
	const layout = "Jan 2, 2006 at 3:04pm (MST)"
	
	// parsing(string format) to time.Time 
	// "Feb 4... " actual input
	tm, _ := time.Parse(layout, "Feb 4, 2014 at 6:05pm (PST)")
	
	fmt.Println(tm) // 2014-02-04 18:05:00 +0000 PST
	fmt.Printf("%T",tm)// time.Time
}
