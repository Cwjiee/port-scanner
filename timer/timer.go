package timer

import (
	"fmt"
	"time"
)

func TimeTrack(start time.Time) {
	elapsed := time.Since(start)
	fmt.Printf("\nscanning took %+v\n", elapsed)	
}
