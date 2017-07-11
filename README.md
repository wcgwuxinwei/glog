# glog [![Build Status](https://travis-ci.org/wcgwuxinwei/glog.svg?branch=master)](https://travis-ci.org/wcgwuxinwei/glog)
super tiny go log, based on go official log, glog means geralt(witcher protagonist) log

# Usage

```go
package main

import (
    "github.com/wcgwuxinwei/glog"
)

func main() {
    glog.SetLevel(glog.LogLevelDebug)
    glog.SetReload(true)
    glog.SetInterval(20)
    
    // start use the glog
    glog.Debugf("Test %s", "test")
}

```
