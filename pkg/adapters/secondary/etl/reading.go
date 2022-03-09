package etl

import (
	"io/ioutil"
	"net/http"
)

func (a *etladapter) ExtractTodos() (rawTodos []byte, err error) {
	var resp *http.Response

	resp, err = a.client.Get("https://jsonplaceholder.typicode.com/todos")
	if err != nil {
		return
	}

	defer resp.Body.Close()

	rawTodos, err = ioutil.ReadAll(resp.Body)

	return
}
