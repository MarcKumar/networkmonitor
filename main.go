package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	for {
		if (connected()) {
			fmt.Println(getTime(), " Connected")
		} else {
			fmt.Println(getTime(), " Not connected")
			cyclePower()
		}
		time.Sleep(300 * time.Second)
	}
}

func getTime() (dt string) {
	t := time.Now()
	return t.Format("2006-01-02 15:04:05")
}

func connected() (ok bool) {
    _, err := http.Get("http://google.com")
    if err != nil {
        return false
    }
    return true
}

func cyclePower() {
	fmt.Println("Power off")
	http.Get("http://192.168.1.231/cm?cmnd=Power%20Off")
	time.Sleep(5 * time.Second)
	fmt.Println("Power on")
	http.Get("http://192.168.1.231/cm?cmnd=Power%20On")
}