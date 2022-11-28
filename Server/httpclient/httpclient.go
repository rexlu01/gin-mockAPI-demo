package httpclient

import "net/http"

type Requests interface {
	GetMethod() string                                                                   //需要获取的请求方法
	GetUrI() string                                                                      //需要获取的请求URI
	GetHeaders() string                                                                  //需要获取的请求Headers
	GetParams() []byte                                                                   //需要获取的请求参数
	HttpRequest(mothed string, url string, headers string, Params []byte) *http.Response //http请求的主体方法
}
