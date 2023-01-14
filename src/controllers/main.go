package controllers

import (
	"encoding/json"
	"fmt"
	"gotest/src/models"
	"gotest/src/services"
	"io/ioutil"
	"net/http"
)

func Init() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		arr := models.ArrayObj{}
		var m models.ResponseRandom
		c := NewClient(http.DefaultClient, 25)
		for i := 0; i < 25; i++ {
			go c.AsyncGet("https://api.chucknorris.io/jokes/random")
		}
		for i := 0; i < 25; i++ {
			select {
			case resp := <-c.Resp:
				body, _ := ioutil.ReadAll(resp.Body) // response body is []byte
				json.Unmarshal([]byte(string(body)), &m)
				arr.Obj = append(arr.Obj, m.ID)
				fmt.Printf("Status received for %s: %d\n", resp.Request.URL, resp.StatusCode)
			case err := <-c.Err:
				fmt.Printf("Error received: %s\n", err)
			}
		}
		services.SendJSONResponse("Succes Request", http.StatusOK, arr, w, r)
	}
}

// NewClient crea un nuevo cliente y establece sus canales apropiados
func NewClient(client *http.Client, bufferSize int) *Client {
	respch := make(chan *http.Response, bufferSize)
	errch := make(chan error, bufferSize)
	return &Client{
		Client: client,
		Resp:   respch,
		Err:    errch,
	}
}

// Client contiene un cliente y tiene dos canales para agregar
// respuestas y errores
type Client struct {
	*http.Client
	Resp chan *http.Response
	Err  chan error
}

// realiza una peticion GET y luego devuelve
// la respuesta/error al canal apropiado
func (c *Client) AsyncGet(url string) {
	resp, err := c.Get(url)
	if err != nil {
		c.Err <- err
		return
	}
	c.Resp <- resp
}
