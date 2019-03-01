package httprouter

import "testing"

func TestRequest_Get(t *testing.T) {
	data := make(map[string][]string)
	req := &Request{Data: data}
	strs := req.Data["name"]

	t.Log(len(strs))
}
