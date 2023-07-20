package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"
)

type Config struct {
	check_interval time.Duration
	check_url      string
	tasmota_ip     string
}

func main() {
	// Disable SSL verification for Docker
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	config := getConfig()

	fmt.Println("Starting...")
	fmt.Printf("Checking connection every:		%s\n", config.check_interval.String())
	fmt.Printf("Using following URL:			%s\n", config.check_url)
	fmt.Printf("Sending powercycle to:			%s\n\n", config.tasmota_ip)

	for {
		ok, duration := checkConnection(config.check_url)
		if ok {
			fmt.Printf("%s Got valid response in %s\n", getTime(), duration)
		} else {
			fmt.Printf("%s Got invalid response in %s execute power-cycle\n", getTime(), duration)
			doPowerCycle(config.tasmota_ip)
		}
		time.Sleep(config.check_interval)
	}
}

func getConfig() (config Config) {
	if os.Getenv("check_interval") != "" {
		config.check_interval, _ = time.ParseDuration(os.Getenv("check_interval"))
	} else {
		flag.DurationVar(&config.check_interval, "check_interval", 1*time.Hour, "Interval to check connection")
	}

	if os.Getenv("check_url") != "" {
		config.check_url = os.Getenv("check_url")
	} else {
		flag.StringVar(&config.check_url, "check_url", "http://google.com", "URL to check connection")
	}

	if os.Getenv("tasmota_ip") != "" {
		config.tasmota_ip = os.Getenv("tasmota_ip")
	} else {
		flag.StringVar(&config.tasmota_ip, "tasmota_ip", "127.0.0.1", "URL to check connection")
	}

	flag.Parse()

	return config
}

func getTime() (dt string) {
	t := time.Now()
	return t.Format(time.RFC1123)
}

func checkConnection(url string) (ok bool, duration time.Duration) {
	t := time.Now()
	_, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return false, time.Since(t)
	}
	return true, time.Since(t)
}

func doPowerCycle(tasmotaIp string) {
	fmt.Println("Power off")
	http.Get("http://" + tasmotaIp + "/cm?cmnd=Power%20Off")
	time.Sleep(5 * time.Second)
	fmt.Println("Power on")
	http.Get("http://" + tasmotaIp + "/cm?cmnd=Power%20On")
}
