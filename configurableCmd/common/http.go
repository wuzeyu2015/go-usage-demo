package common

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func HttpSend(method, path string, v interface{}) []byte {
	cfg := getConf()
	if cfg == nil {
		fmt.Printf("httpSend(): cannot read config file\n")
		os.Exit(1)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	retCh := make(chan *http.Response)
	buf, err := json.Marshal(v)
	if err != nil {
		fmt.Printf("httpSend(): json.Marshal error, %#v\n", err)
		os.Exit(1)
	}

	httpFunc := func(url string, ior io.Reader) {
		req, e := http.NewRequestWithContext(ctx, method, url, ior)
		if e != nil {
			fmt.Printf("httpSend(): new req err=%v\n", err)
			os.Exit(1)
		}
		cli := &http.Client{}
		rsp, err1 := cli.Do(req)
		if err1 != nil {
			//The error 'context canceled' is normal
			//fmt.Printf("httpSend(): cli.Do, err=%v\n", err1)
			return
		}
		retCh <- rsp
	}

	for _, ctrl := range cfg.Controllers {
		url := fmt.Sprintf("http://%s:23081%s", ctrl.Ip, path)
		ior := bytes.NewReader(buf)
		go httpFunc(url, ior)
	}
	var rsp *http.Response
	var data []byte
	select {
	case <-ctx.Done():
		fmt.Printf("httpSend(): error occurs, %v\n", ctx.Err())
		os.Exit(1)
	case rsp = <-retCh:
		data, err = ioutil.ReadAll(rsp.Body)
		if err != nil {
			fmt.Printf("httpSend(): ReadAll error=%v\n", err)
			os.Exit(1)
		}
	}
	cancel()
	return data
}
