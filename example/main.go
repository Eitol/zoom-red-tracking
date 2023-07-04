package main

import (
	"fmt"
	"github.com/Eitol/zoom-red-tracking/zoom"
	"time"
)

func main() {
	var tracking *zoom.Shipment
	var err error
	c := zoom.NewDefaultClient()
	baseTracking := 1553486107
	for i := 0; i < 1000; i++ {
		startTime := time.Now()
		tracking, err = c.GetTrackingInfo(baseTracking + (1 * 10_000))
		fmt.Printf("elapsed time: %v\n", time.Since(startTime))
		if err == nil {
			fmt.Printf("tracking: %v\n", tracking)
		} else {
			fmt.Printf("error: %v\n", err)
		}
	}
}
