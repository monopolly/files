package file

import (
	"io/ioutil"
	"net"
	"net/http"
	"time"

	"github.com/monopolly/errors"
	"github.com/monopolly/useragent"
)

func Download(link string) (file []byte, err errors.E) {
	transport := &http.Transport{Dial: (&net.Dialer{Timeout: 20 * time.Second}).Dial, TLSHandshakeTimeout: 20 * time.Second}
	c := &http.Client{Timeout: time.Second * 20, Transport: transport}
	resp, er := c.Get(link)
	if er != nil {
		err = errors.Server()
		return
	}
	defer resp.Body.Close()

	file, er = ioutil.ReadAll(resp.Body)
	if er != nil {
		err = errors.File()
		return
	}
	return
}

//добавляет хедеры и генерит юзер агента как реальный юзер
func Downloads(link string) (b []byte, err errors.E) {

	req, _ := http.NewRequest("GET", link, nil)
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Accept-Languag", "en-us")
	req.Header.Add("DNT", "1")
	/* генерим агента */
	req.Header.Add("User-agent", useragent.Generate())
	resp, er := http.DefaultClient.Do(req)
	if er != nil {
		err = errors.Server(er)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		err = errors.Server("Status code is", resp.StatusCode)
		return
	}

	b, er = ioutil.ReadAll(resp.Body)
	if er != nil {
		err = errors.Unmarshal(er)
		return
	}

	if b == nil {
		err = errors.Server(er)
		return
	}

	return
}
