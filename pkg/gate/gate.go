package gate

import (
	//"fmt"
	"context"
	"net/http"
	"strings"

	"io/ioutil"
)

type httpGate struct {
}

type HttpGate interface {
	DoesItContainStr(ctx context.Context, url string, str string) (bool, error)
}

func NewHttpGate() HttpGate {
	return &httpGate{
	}
}

func (g *httpGate)DoesItContainStr(ctx context.Context, url string, str string) (bool, error) {

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("accept-language","ja")

	client := new(http.Client)
	resp, _ := client.Do(req)
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	if strings.Contains(string(byteArray),str) {
		return true, nil
	} else {
		return false, nil
	}
}





