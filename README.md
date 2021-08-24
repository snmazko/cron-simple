# Go/Golang a package for running scheduled tasks.

## Installation <a id="installation"></a>
```
go get github.com/snmazko/cron-simple
```

## Example code <a id="example"></a>

```go
package main

import (
	"fmt"
	"github.com/snmazko/cron-simple"
	"time"
)

type logger struct{}

func (l logger) Errorf(format string, a ...interface{}) {
	fmt.Printf(format, a...)
}

func (l logger) Info(data interface{}) {
	fmt.Printf("%v\n", data)
}

func main() {
	worker := cron.NewWorker().
		SetLogger(logger{}).
		Start()

	worker.MustCreateTask("* * * * *", func1)           //every minute
	worker.MustCreateTask("*/2 * * * *", func() error { //every two minute
		return func2("Run func2")
	})

	time.Sleep(5 * time.Minute)

    worker.Stop()
}

func func1() error {
	fmt.Printf("Run func1, time: %v\n", time.Now())
	return nil
}

func func2(msg string) error {
	fmt.Printf("%s, time: %v\n", msg, time.Now())
	return nil
}

```

### Example syntax <a id="example_syntax"></a>
```
*     *     *     *     *        

^     ^     ^     ^     ^
|     |     |     |     |
|     |     |     |     +----- day of week (0-6) (Sunday=0)
|     |     |     +------- month (1-12)
|     |     +--------- day of month (0-31) (0 - last day of month)
|     +----------- hour (0-23)
+------------- min (0-59)
```

+ `*	*	*	*	*`	- every minute
+ `*/5	*	*	*	*`	- every 5 minutes
+ `0	*	*	*	*`	- every hour
+ `0	*/3	*	*	*`	- every 3 hours
+ `0	13	*	*	*`	- every day at 13:00
+ `30	2	*	*	*`	- every day at 2:30
+ `0	0	*	*	*`	- every day at midnight
+ `0	0	1	*	*`	- on the first day of every month
+ `0	0	0	*	*`	- on the last day of every month
+ `0	0	6	1	*`	- on the sixth day of the first month
+ `0	0	*	*	0`	- at midnight every Sunday
