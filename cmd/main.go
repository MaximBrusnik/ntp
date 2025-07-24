package main

import (
	"github.com/MaximBrusnik/ntp"
)

const (
	ntpServer = "ntp1.stratum1.ru"
)

// хоть это и модуль (по идее тут не стоит использовать main, но для наглядности добавил)
func main() {
	ntp.PrintNTPTime(ntpServer)
}
