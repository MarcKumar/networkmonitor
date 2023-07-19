package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

var check_interval, _ = time.ParseDuration(os.Getenv("check_interval"))
var check_url = os.Getenv("check_url")
var tasmota_ip = os.Getenv("tasmota_ip")

func main() {
	fmt.Println("Starting...")
	fmt.Println("Checking connection every:	", check_interval)
	fmt.Println("Using following URL: 		", check_url)
	fmt.Println("Sending powercycle to:		", tasmota_ip)
	fmt.Println("")

	for {
		if connected() {
			fmt.Println(getTime(), " Connected")
		} else {
			fmt.Println(getTime(), " Not connected")
			cyclePower()
		}
		time.Sleep(check_interval)
	}
}

func getTime() (dt string) {
	t := time.Now()
	return t.Format(time.RFC1123)
}

func connected() (ok bool) {
	_, err := http.Get(check_url)
	if err != nil {
		return false
	}
	return true
}

func cyclePower() {
	fmt.Println("Power off")
	http.Get("http://" + tasmota_ip + "/cm?cmnd=Power%20Off")
	time.Sleep(5 * time.Second)
	fmt.Println("Power on")
	http.Get("http://" + tasmota_ip + "/cm?cmnd=Power%20On")
}
