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
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v2"
)

func TestBatchCreateLd(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "ld")
	setupFlagString(set, "host,data")
	setupFlagBool(set, "append,keyValues")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})
	err := batch(c, "create")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "data is empty", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestBatchUpdateLd(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "ld")
	setupFlagString(set, "host,data")
	setupFlagBool(set, "append,keyValues")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})
	err := batch(c, "update")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "data is empty", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestBatchUsertLd(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "ld")
	setupFlagString(set, "host,data")
	setupFlagBool(set, "append,keyValues")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})
	err := batch(c, "upsert")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "data is empty", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestBatchDeleteLd(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "ld")
	setupFlagString(set, "host,data")
	setupFlagBool(set, "append,keyValues")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})
	err := batch(c, "delete")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "data is empty", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestBatchCreateV2(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")
	setupFlagString(set, "host,data")
	setupFlagBool(set, "append,keyValues")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})
	err := batch(c, "create")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "data is empty", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestBatchUpdateV2(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")
	setupFlagString(set, "host,data")
	setupFlagBool(set, "append,keyValues")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})
	err := batch(c, "update")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "data is empty", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestBatchAppendV2(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")
	setupFlagString(set, "host,data")
	setupFlagBool(set, "append,keyValues")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})
	err := batch(c, "upsert")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "data is empty", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}
func TestBatchReplaceV2(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")
	setupFlagString(set, "host,data")
	setupFlagBool(set, "append,keyValues")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})
	err := batch(c, "replace")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "data is empty", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestBatchDeleteV2(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")
	setupFlagString(set, "host,data")
	setupFlagBool(set, "append,keyValues")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})
	err := batch(c, "delete")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "data is empty", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestBatchErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host,data")
	setupFlagBool(set, "append,keyValues")

	c := cli.NewContext(app, set, nil)
	err := batch(c, "create")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "Required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestBatchErrorNewClient(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")
	setupFlagString(set, "host,data,link")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--link=abc", "--host=orion"})
	err := batch(c, "create")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "abc not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestBatchErrorModeV2(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")
	setupFlagString(set, "host,data,link")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})
	err := batch(c, "get")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "error: get", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestBatchErrorModeLD(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "LD")
	setupFlagString(set, "host,data,link")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})
	err := batch(c, "get")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "error: get", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestBatchCreate(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "LD")
	setupFlagString(set, "host,data,link")
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(testData)
	reqRes.Path = "/ngsi-ld/v1/entityOperations/create"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--data=" + testData})
	client, _ := newClient(ngsi, c, false)
	err := batchCreate(c, ngsi, client)

	assert.NoError(t, err)
}

func TestBatchCreateErrorReadAll(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "LD")
	setupFlagString(set, "host,data,link")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--data="})
	client, _ := newClient(ngsi, c, false)
	err := batchCreate(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "data is empty", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestBatchCreateErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "LD")
	setupFlagString(set, "host,data,link")
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/entityOperations/create"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--data={}"})
	client, _ := newClient(ngsi, c, false)
	err := batchCreate(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "url error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestBatchCreateErrorHTTPStatus(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "LD")
	setupFlagString(set, "host,data,link")
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/ngsi-ld/v1/entityOperations/create"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--data={}"})
	client, _ := newClient(ngsi, c, false)
	err := batchCreate(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
	} else {
		t.FailNow()
	}
}

func TestBatchUpdate(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "LD")
	setupFlagString(set, "host,data,link")
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.ResBody = []byte(testData)
	reqRes.Path = "/ngsi-ld/v1/entityOperations/update"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--data=" + testData})
	client, _ := newClient(ngsi, c, false)
	err := batchUpdate(c, ngsi, client)

	assert.NoError(t, err)
}

func TestBatchUpdateErrorReadAll(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "LD")
	setupFlagString(set, "host,data,link")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--data="})
	client, _ := newClient(ngsi, c, false)
	err := batchUpdate(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "data is empty", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestBatchUpdateErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "LD")
	setupFlagString(set, "host,data,link")
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/entityOperations/update"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--data={}"})
	client, _ := newClient(ngsi, c, false)
	err := batchUpdate(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "url error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestBatchUpdateErrorHTTPStatus(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "LD")
	setupFlagString(set, "host,data,link")
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/ngsi-ld/v1/entityOperations/update"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--data={}"})
	client, _ := newClient(ngsi, c, false)
	err := batchUpdate(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
	} else {
		t.FailNow()
	}
}

func TestBatchUpsert(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "LD")
	setupFlagString(set, "host,data,link")
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.ResBody = []byte(testData)
	reqRes.Path = "/ngsi-ld/v1/entityOperations/upsert"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--data=" + testData})
	client, _ := newClient(ngsi, c, false)
	err := batchUpsert(c, ngsi, client)

	assert.NoError(t, err)
}

func TestBatchUpsertErrorReadAll(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "LD")
	setupFlagString(set, "host,data,link")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--data="})
	client, _ := newClient(ngsi, c, false)
	err := batchUpsert(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "data is empty", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestBatchUpsertErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "LD")
	setupFlagString(set, "host,data,link")
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/entityOperations/upsert"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--data={}"})
	client, _ := newClient(ngsi, c, false)
	err := batchUpsert(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "url error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestBatchUpsertErrorHTTPStatus(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "LD")
	setupFlagString(set, "host,data,link")
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/ngsi-ld/v1/entityOperations/upsert"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--data={}"})
	client, _ := newClient(ngsi, c, false)
	err := batchUpsert(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
	} else {
		t.FailNow()
	}
}

func TestBatchDelete(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "LD")
	setupFlagString(set, "host,data,link")
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(testData)
	reqRes.Path = "/ngsi-ld/v1/entityOperations/delete"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--data=" + testData})
	client, _ := newClient(ngsi, c, false)
	err := batchDelete(c, ngsi, client)

	assert.NoError(t, err)
}

func TestBatchDeleteErrorReadAll(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "LD")
	setupFlagString(set, "host,data,link")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--data="})
	client, _ := newClient(ngsi, c, false)
	err := batchDelete(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "data is empty", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestBatchDeleteErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "LD")
	setupFlagString(set, "host,data,link")
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/entityOperations/delete"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--data={}"})
	client, _ := newClient(ngsi, c, false)
	err := batchDelete(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "url error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestBatchDeleteErrorHTTPStatus(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "LD")
	setupFlagString(set, "host,data,link")
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/ngsi-ld/v1/entityOperations/delete"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--data={}"})
	client, _ := newClient(ngsi, c, false)
	err := batchDelete(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
	} else {
		t.FailNow()
	}
}

var testData = `[
    {
      "id": "urn:ngsi-ld:TemperatureSensor:002",
      "type": "TemperatureSensor",
      "category": {
            "type": "Property",
            "value": "sensor"
      },
      "temperature": {
            "type": "Property",
            "value": 21,
            "unitCode": "CEL"
      }
    },
    {
      "id": "urn:ngsi-ld:TemperatureSensor:003",
      "type": "TemperatureSensor",
      "category": {
            "type": "Property",
            "value": "sensor"
      },
      "temperature": {
            "type": "Property",
            "value": 27,
            "unitCode": "CEL"
      }
    }
]`
