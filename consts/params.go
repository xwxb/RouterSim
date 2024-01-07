package consts

import "time"

var RouterRcvTickerChan = time.NewTicker(1 * time.Second).C
var HostRcvTickerChan = time.NewTicker(1 * time.Second).C
