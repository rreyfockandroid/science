package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"multi_projects/http_putter/data"
	"net/http"
	"time"
)

const (
	timeout = time.Millisecond * 500

	port = "8787"
	paramSignal = "signal"
	signalKill = "kill"
)

var client = initClient()
var close = make(chan bool)

func main() {
	logrus.Info("--- start ---")
	go serverStart()

	//resp := get("http://www.wp.pl")
	resp := post("http://localhost:8787", getBikeJson())
	body, _ := ioutil.ReadAll(resp.Body)
	logrus.Infof("Response: %v, body: %s", resp, string(body))

	<-close
	logrus.Info("--- end ---")
}

func get(url string) (resp *http.Response) {
	if request, err := http.NewRequest("GET", url, nil); err == nil {
		resp, err = client.Do(request)
		if err != nil {
			logrus.Errorf("Error when http client GET req=%v, err=%v", request, err)
		}
	} else {
		logrus.Errorf("Error when create request %v", err)
	}
	return resp
}

func post(url string, body []byte) (resp *http.Response) {
	if request, err := http.NewRequest("POST", url, bytes.NewBuffer(body)); err == nil {
		resp, err = client.Do(request)
		if err != nil {
			logrus.Errorf("Error when http client POST req=%v, err=%v", request, err)
		}
	} else {
		logrus.Errorf("Error when create request %v", err)
	}
	return resp
}

func initClient() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			MaxConnsPerHost: 1,
			DisableKeepAlives: false,
			ResponseHeaderTimeout: timeout,
			ExpectContinueTimeout: timeout,
		},
	}
}


func handler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logrus.Errorf("Error when read body %v", err)
	}
	logrus.Infof("handler: %v", string(body))
	if r.URL.Query().Get(paramSignal) == signalKill {
		close <- true
		return
	}
	mess := fmt.Sprintf("Przeslales %s\n", body)
	w.Write([]byte(mess))
}

func serverStart() {
	logrus.Infof("start server listen on %s", port)
	http.HandleFunc("/", handler)
	if err := http.ListenAndServe(":" + port, nil); err != nil {
		logrus.Fatal("Start server fail %v", err)
	}
}

func getBikeJson() []byte {
	res, err := json.Marshal(getBikeIdx())
	if err == nil {
		logrus.Infof("JsonBike: %s", string(res))
		return res
	}
	logrus.Errorf("Error when marshal bike idx %v", err)
	return nil
}

func getBikeIdx() *data.BikeIdx {
	return &data.BikeIdx {
		Name: "Kona",
		AddDate: time.Now(),
		Destination: &data.Destination{
			Xc: 0,
			Dh: 1,
			Fr: 1,
		},
	}
}










