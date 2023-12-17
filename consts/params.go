package consts

import "time"

var RouterRcvTickerChan = time.NewTicker(3 * time.Second).C
var HostRcvTickerChan = time.NewTicker(3 * time.Second).C
