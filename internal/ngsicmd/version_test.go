/*
MIT License

Copyright (c) 2020 Kazuhito Suda

This file is part of NGSI Go

https://github.com/lets-fiware/ngsi-go

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

*/

package ngsicmd

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v2"
)

func TestVersionV2(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/version"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})
	err := cbVersion(c)

	assert.NoError(t, err)
}

func TestVersionLD(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion-ld", "https://orion", "ld")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/version"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld"})
	err := cbVersion(c)

	assert.NoError(t, err)
}

// initCmd() Error: no host
func TestVersionErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	c := cli.NewContext(app, set, nil)
	err := cbVersion(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "Required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestVersionErrorNewClient(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion-ld", "https://orion", "ld")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/version"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host,link")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--link=abc"})
	err := cbVersion(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "abc not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestVersionErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/version"
	reqRes.Err = errors.New("error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})
	err := cbVersion(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
	} else {
		t.FailNow()
	}
}

func TestVersionErrorStatusCode(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/version"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})
	err := cbVersion(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
	} else {
		t.FailNow()
	}
}
