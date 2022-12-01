package httpclient

import (
	"fmt"
	"ginvue/Server/util"
	"io/ioutil"
	"net/http"
	"net/url"
)

func MainHttpClient(Methood string, URL string, Headers map[string]string, Params map[string]string) (respStatus string, respBody string) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()

	req, err := http.NewRequest(Methood, URL, nil)
	if err != nil {
		//这里后续log模块
		util.Check(err)
	}

	params := make(url.Values)
	for k, v := range Params {
		params.Add(k, v)
	}
	req.URL.RawQuery = params.Encode()

	//设置headers
	for k, v := range Headers {
		req.Header.Set(k, v)
	}

	r, err := http.DefaultClient.Do(req)
	defer func() {
		_ = r.Body.Close()
	}()

	body, _ := ioutil.ReadAll(r.Body)
	return r.Status, string(body)

}
