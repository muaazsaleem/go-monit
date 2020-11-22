package main

import "github.com/muaazsaleem/go-monit/monit"

func main() {
	monit.RunUserTest(monit.UserHTTPTest{})
}
