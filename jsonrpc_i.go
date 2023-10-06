package coingecko

import (
	"os"
	"log"
	"bytes"
	"errors"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

type RPCRequest struct {
	JsonRPC string `json:"jsonrpc"`
	ID      string `json:"id"`
	Method  string `json:"method"`
	Params  interface{} `json:"params,omitempty"`
}

type RPCResponse struct {
	Error struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
	Result interface{} `json:"result"`
}

type RPC struct {
	URL  string
	User string
	Pass string
}

var VerboseRPC bool = os.Getenv("VERBOSE_RPC") != ""

func (c *RPC) RPCQuery(httpPath, httpMethod, rpcMethod string, in any, out any) (err error) {
	var iB, oB []byte
	var req   RPCRequest
	var res   RPCResponse
	var reqH *http.Request
	var resH *http.Response
	var cli  *http.Client = &http.Client{}
	
	// Serialize body.
	req.JsonRPC = "2.0"
	req.ID = "0"
	req.Method = rpcMethod
	req.Params = in
	iB, err = json.Marshal(req)
	if err != nil { return }
	if VerboseRPC { log.Print("Sending: " + string(iB)) }
	
	// Perform HTTP request.
	reqH, err = http.NewRequest(httpMethod, c.URL + httpPath, bytes.NewBuffer(iB))
	if err != nil { return }
	reqH.Header.Set("Accept", "application/json")
	reqH.Header.Set("Content-Type", "application/json")
	if c.User != "" && c.Pass != "" {
		reqH.SetBasicAuth(c.User, c.Pass)
	}
	
	resH, err = cli.Do(reqH)
	if err != nil { return }
	defer resH.Body.Close()
	
	// Retrieve the response.
	oB, err = ioutil.ReadAll(resH.Body)
	if err != nil { return }
	if VerboseRPC { log.Print("Received: " + string(oB)) }
	if resH.StatusCode != 200 { err = errors.New(string(oB)); return }
	
	// Parse error message if any.
	err = json.Unmarshal(oB, &res)
	if err != nil { return }
	if res.Error.Message != "" {
		err = errors.New(res.Error.Message)
		return
	}
	
	// Return result.
	if out != nil {
		oB, err = json.Marshal(res.Result)
		if err != nil { return }
		if VerboseRPC { log.Print("Result: " + string(oB)) }
		err = json.Unmarshal(oB, out)
		if err != nil { return }
	}
	
	return
}

func (c *RPC) SimQuery(httpPath, httpMethod string, out any) (err error) {
	var oB  []byte
	var reqH *http.Request
	var resH *http.Response
	var cli  *http.Client = &http.Client{}
	
	reqH, err = http.NewRequest(httpMethod, c.URL + httpPath, nil)
	if err != nil { return }
	reqH.Header.Set("Accept", "application/json")
	if c.User != "" && c.Pass != "" {
		reqH.SetBasicAuth(c.User, c.Pass)
	}
	
	resH, err = cli.Do(reqH)
	if err != nil { return }
	defer resH.Body.Close()
	
	oB, err = ioutil.ReadAll(resH.Body)
	if err != nil { return }
	if VerboseRPC { log.Print("Received: " + string(oB)) }
	if resH.StatusCode != 200 { err = errors.New(string(oB)); return }
	
	err = IsError(oB)
	if err != nil { return }
	
	if out != nil {
		if VerboseRPC { log.Print("Result: " + string(oB)) }
		err = json.Unmarshal(oB, out)
		if err != nil { return }
	}
	
	return
}

func IsError(oB []byte) (err error) {
	var err1 struct {
		Status struct {
			ErrorMessage string `json:"error_message"`
		} `json:"status"`
	}
	var err2 struct {
		Error string `json:"error"`
	}
	
	err = json.Unmarshal(oB, &err1)
	if err == nil && err1.Status.ErrorMessage != "" {
		return errors.New(err1.Status.ErrorMessage)
	}
	
	err = json.Unmarshal(oB, &err2)
	if err == nil && err2.Error != "" {
		return errors.New(err2.Error)
	}
	
	return nil
}
