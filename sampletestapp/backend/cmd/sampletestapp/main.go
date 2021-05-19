package main

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"os"
	"time"
)

type Myresponse struct {
	Date     string   `json:"timestamp"`
	Origen   string   `json:"request_origin"`
	Xff      string   `json:"request_x-forwarded-for"`
	Hostname string   `json:"response_hostname"`
	Ips      []string `json:"response_ips"`
}

func main() {

	port := 8080

	http.HandleFunc("/", handlerPing)

	var address = fmt.Sprintf(":%d", port)
	fmt.Printf("backend atendiendo en %d\n", port)
	err := http.ListenAndServe(address, nil)
	if err != nil {
		panic(err)
	}

}

func logPrint(message string, status bool) {

	now := time.Now().String()
	var slog string
	if status {
		slog = "OK"
	} else {
		slog = "ERROR"
	}
	log := fmt.Sprintf("[%s] %s => %s\n", slog, now, message)
	fmt.Println(log)

}

func handlerPing(w http.ResponseWriter, r *http.Request) {

	hostname, err := os.Hostname()
	if err != nil {
		logPrint(err.Error(), false)
		return
	}

	var ips []string
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		logPrint(err.Error(), false)
		return
	}

	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ips = append(ips, ipnet.IP.String())
			}
		}
	}

	t := time.Now().String()

	data := Myresponse{
		Date:     t,
		Origen:   r.RemoteAddr,
		Xff:      r.Header.Get("X-Forwarded-For"),
		Hostname: hostname,
		Ips:      ips,
	}

	PayLoad, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		logPrint(err.Error(), false)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(PayLoad)
	logPrint(" => respondiendo a "+r.RemoteAddr+"(xff: "+r.Header.Get("X-Forwarded-For"), true)

}
