package binance

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
)

// ProxyURLForWsHandshakeOverride is a function that can be used to override the proxy URL
var ProxyURLForWsHandshakeOverride = ""

// WsHandler handle raw websocket message
type WsHandler func(message []byte)

// ErrHandler handles errors
type ErrHandler func(err error)

// WsConfig webservice configuration
type WsConfig struct {
	Endpoint string
}

func newWsConfig(endpoint string) *WsConfig {
	return &WsConfig{
		Endpoint: endpoint,
	}
}

var wsServe = func(cfg *WsConfig, handler WsHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	Dialer := websocket.Dialer{
		Proxy: http.ProxyFromEnvironment,
	}
	if ProxyURLForWsHandshakeOverride != "" {
		proxyURL, e := url.Parse(ProxyURLForWsHandshakeOverride)
		if e != nil {
			return nil, nil, fmt.Errorf("unable to parse ProxyURLForWsHandshakeOverride (%s): %s", ProxyURLForWsHandshakeOverride, e)
		}
		Dialer.Proxy = func(req *http.Request) (*url.URL, error) {
			log.Printf("using proxy %s for websocket handshake", proxyURL)
			return proxyURL, nil
		}
	}

	c, resp, err := Dialer.Dial(cfg.Endpoint, nil)
	if err != nil {
		byts := []byte{}
		// check resp and resp.Body as they can be nil
		if resp != nil && resp.Body != nil {
			defer resp.Body.Close()

			var e error
			byts, e = ioutil.ReadAll(resp.Body)
			if e != nil {
				return nil, nil, fmt.Errorf("unable to read response from fialed dial call, error: %s, original dialer error: %s", e, err)
			}
		}
		return nil, nil, fmt.Errorf("response from failed Dial call: %v, orignial dialer error: %s", string(byts), err)
	}
	doneC = make(chan struct{})
	stopC = make(chan struct{})
	go func() {
		// This function will exit either on error from
		// websocket.Conn.ReadMessage or when the stopC channel is
		// closed by the client.
		defer close(doneC)
		if WebsocketKeepalive {
			keepAlive(c, WebsocketTimeout)
		}
		// Wait for the stopC channel to be closed.  We do that in a
		// separate goroutine because ReadMessage is a blocking
		// operation.
		silent := false
		go func() {
			select {
			case <-stopC:
				silent = true
			case <-doneC:
			}
			c.Close()
		}()
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				if !silent {
					errHandler(err)
				}
				return
			}
			handler(message)
		}
	}()
	return
}

func keepAlive(c *websocket.Conn, timeout time.Duration) {
	ticker := time.NewTicker(timeout)

	lastResponse := time.Now()
	c.SetPongHandler(func(msg string) error {
		lastResponse = time.Now()
		return nil
	})

	go func() {
		defer ticker.Stop()
		for {
			deadline := time.Now().Add(10 * time.Second)
			err := c.WriteControl(websocket.PingMessage, []byte{}, deadline)
			if err != nil {
				c.Close()
				return
			}
			<-ticker.C
			if time.Since(lastResponse) > timeout {
				c.Close()
				return
			}
		}
	}()
}
