# glog
super tiny go log, based on go official log, glog means geralt(witcher protagonist) log

# usage

```go
package main

import (
    "github.com/wcgwuxinwei/glog"
)

func main() {
    logger := glog.New(true)
    logger.SetLevel(glog.LogLevelDebug)
    
    logger.Debugf("Test %s", "test")
}

```
