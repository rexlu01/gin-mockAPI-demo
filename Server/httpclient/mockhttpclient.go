package httpclient

type req struct {
	Method  string
	UrI     string
	Headers string
	Params  []byte
}
