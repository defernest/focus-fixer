package axis

import (
	"bytes"
	"fmt"
	"net"
	"net/http"
	"time"
)

type Camera struct {
	IP       net.IP `csv:"ip"`
	Login    string `csv:"login"`
	Password string `csv:"pass"`
}

func (cam Camera) Camera() net.IP {
	return cam.IP
}

func (cam Camera) Autofocus() error {
	uri := fmt.Sprintf("http://%s/axis-cgi/opticssetup.cgi?autofocus=perform", cam.IP)
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return fmt.Errorf("create new request error: %v", err)
	}
	if err != nil {
		return err
	}
	body, err := MakeRequest(req, cam)
	if err != nil {
		return fmt.Errorf("make request error: %v", err)
	}
	if !bytes.EqualFold(bytes.TrimSpace(body.Bytes()), []byte("ok")) {
		return fmt.Errorf("fail to set autofocus (response body !ok): %s", body)
	}
	return nil
}

func MakeRequest(request *http.Request, cred Camera) (*bytes.Buffer, error) {
	request.SetBasicAuth(cred.Login, cred.Password)
	client := &http.Client{Timeout: 60 * time.Second}
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode {
	case http.StatusOK:
		body := &bytes.Buffer{}
		_, err = body.ReadFrom(resp.Body)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		return body, err
	case http.StatusUnauthorized:
		return nil, fmt.Errorf("[%v] either you supplied the wrong credentials (e.g., bad password)", resp.Status)
	default:
		return nil, fmt.Errorf("http error [%v]", resp.Status)
	}
}
