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

package ngsilib

import (
	"bytes"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNgsiIntiConfig(t *testing.T) {
	ngsi := testNgsiLibInit()
	ngsi.ConfigFile = &MockIoLib{}

	filename := ""
	err := ngsi.InitConfig(&filename)
	assert.NoError(t, err)
}

func TestIntiConfig(t *testing.T) {
	ngsi := testNgsiLibInit()
	ngsi.ConfigFile = &MockIoLib{}
	filename := ""
	ngsi.ConfigFile.SetFileName(&filename)

	err := initConfig(ngsi, ngsi.ConfigFile)
	assert.NoError(t, err)
}

func TestIntiConfigBrokerList(t *testing.T) {
	ngsi := testNgsiLibInit()
	ngsi.ConfigFile = &MockIoLib{}
	filename := ""
	ngsi.ConfigFile.SetFileName(&filename)
	ngsi.brokerList = make(BrokerList)
	broker := &Broker{BrokerHost: "http://orion"}
	ngsi.brokerList["orion"] = broker

	err := initConfig(ngsi, ngsi.ConfigFile)
	assert.NoError(t, err)
}

func TestIntiConfigNoFileName(t *testing.T) {
	ngsi := testNgsiLibInit()
	ngsi.ConfigFile = &MockIoLib{}

	err := initConfig(ngsi, ngsi.ConfigFile)
	assert.NoError(t, err)
}

func TestIntiConfigFileName(t *testing.T) {
	ngsi := testNgsiLibInit()
	ngsi.ConfigFile = &MockIoLib{}
	filename := "config"
	ngsi.ConfigFile.SetFileName(&filename)

	err := initConfig(ngsi, ngsi.ConfigFile)
	assert.NoError(t, err)
}

func TestIntiConfigExistsFile(t *testing.T) {
	ngsi := testNgsiLibInit()
	ngsi.ConfigFile = &MockIoLib{}
	filename := "config"
	ngsi.ConfigFile.SetFileName(&filename)

	err := initConfig(ngsi, ngsi.ConfigFile)
	assert.NoError(t, err)
}

func TestIntiConfigBatchFlag(t *testing.T) {
	ngsi := testNgsiLibInit()
	ngsi.ConfigFile = &MockIoLib{}
	filename := "config"
	ngsi.ConfigFile.SetFileName(&filename)
	f := true
	ngsi.BatchFlag = &f

	err := initConfig(ngsi, ngsi.ConfigFile)
	assert.NoError(t, err)
}

func TestIntiConfigErrorNoFileName(t *testing.T) {
	ngsi := testNgsiLibInit()
	ngsi.ConfigFile = &MockIoLib{HomeDirErr: errors.New("error homedir")}

	err := initConfig(ngsi, ngsi.ConfigFile)

	if assert.Error(t, err) {
		ngsiErr := err.(*NgsiLibError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "error homedir", ngsiErr.Message)
	}
}

func TestIntiConfigErrorFileName(t *testing.T) {
	ngsi := testNgsiLibInit()
	ngsi.ConfigFile = &MockIoLib{PathAbsErr: errors.New("error path abs")}
	filename := "config"
	ngsi.ConfigFile.SetFileName(&filename)

	err := initConfig(ngsi, ngsi.ConfigFile)

	if assert.Error(t, err) {
		ngsiErr := err.(*NgsiLibError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "error path abs config", ngsiErr.Message)
	}
}

func TestIntiConfigErrorOpen(t *testing.T) {
	ngsi := testNgsiLibInit()
	ngsi.ConfigFile = &MockIoLib{OpenErr: errors.New("open error")}
	filename := "config"
	ngsi.ConfigFile.SetFileName(&filename)

	err := initConfig(ngsi, ngsi.ConfigFile)

	if assert.Error(t, err) {
		ngsiErr := err.(*NgsiLibError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "open error", ngsiErr.Message)
	}
}

func TestIntiConfigErrorDecode(t *testing.T) {
	ngsi := testNgsiLibInit()
	ngsi.ConfigFile = &MockIoLib{DecodeErr: errors.New("decode error")}
	filename := "config"
	ngsi.ConfigFile.SetFileName(&filename)

	err := initConfig(ngsi, ngsi.ConfigFile)

	if assert.Error(t, err) {
		ngsiErr := err.(*NgsiLibError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "decode error", ngsiErr.Message)
	}
}

func TestIntiConfigErrorParam(t *testing.T) {
	ngsi := testNgsiLibInit()
	ngsi.ConfigFile = &MockIoLib{}
	filename := ""
	ngsi.ConfigFile.SetFileName(&filename)
	ngsi.LogWriter = &bytes.Buffer{}
	ngsi.brokerList = make(BrokerList)
	broker := &Broker{BrokerHost: "http://orion", NgsiType: "v1"}
	ngsi.brokerList["orion"] = broker

	err := initConfig(ngsi, ngsi.ConfigFile)

	if assert.Error(t, err) {
		ngsiErr := err.(*NgsiLibError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "error in config file", ngsiErr.Message)
	}
}

func TestIntiConfigErrorContext(t *testing.T) {
	ngsi := testNgsiLibInit()
	ngsi.ConfigFile = &MockIoLib{}
	filename := ""
	ngsi.ConfigFile.SetFileName(&filename)
	ngsi.LogWriter = &bytes.Buffer{}
	ngsi.contextList = make(ContextsInfo)
	ngsi.contextList["ld"] = "context"

	err := initConfig(ngsi, ngsi.ConfigFile)

	if assert.Error(t, err) {
		ngsiErr := err.(*NgsiLibError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "error in config file", ngsiErr.Message)
	}
}

func TestSaveConfigFile(t *testing.T) {
	ngsi := testNgsiLibInit()
	ngsi.ConfigFile = &MockIoLib{}
	filename := "config"
	ngsi.ConfigFile.SetFileName(&filename)

	err := ngsi.saveConfigFile()
	assert.NoError(t, err)
}

func TestSaveConfigFileErrorOpenFile(t *testing.T) {
	ngsi := testNgsiLibInit()
	ngsi.ConfigFile = &MockIoLib{OpenErr: errors.New("open error")}
	filename := "config"
	ngsi.ConfigFile.SetFileName(&filename)

	err := ngsi.saveConfigFile()
	if assert.Error(t, err) {
		ngsiErr := err.(*NgsiLibError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "open error", ngsiErr.Message)
	}
}

func TestSaveConfigFileErrorTrancate(t *testing.T) {
	ngsi := testNgsiLibInit()
	ngsi.ConfigFile = &MockIoLib{TruncateErr: errors.New("trancate error")}
	filename := "config"
	ngsi.ConfigFile.SetFileName(&filename)

	err := ngsi.saveConfigFile()
	if assert.Error(t, err) {
		ngsiErr := err.(*NgsiLibError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "trancate error", ngsiErr.Message)
	}
}

func TestSaveConfigFileErrorEncode(t *testing.T) {
	ngsi := testNgsiLibInit()
	ngsi.ConfigFile = &MockIoLib{EncodeErr: errors.New("encode error")}
	filename := "config"
	ngsi.ConfigFile.SetFileName(&filename)

	err := ngsi.saveConfigFile()
	if assert.Error(t, err) {
		ngsiErr := err.(*NgsiLibError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "encode error", ngsiErr.Message)
	}
}

func TestSaveConfigFileEmpty(t *testing.T) {
	ngsi := testNgsiLibInit()
	ngsi.ConfigFile = &MockIoLib{}
	filename := ""
	ngsi.ConfigFile.SetFileName(&filename)

	err := ngsi.saveConfigFile()
	assert.NoError(t, err)
}

func TestGetPreviousArgs(t *testing.T) {
	ngsi := testNgsiLibInit()

	actual := ngsi.GetPreviousArgs()
	expected := ngsi.PreviousArgs
	assert.Equal(t, expected, actual)
}

func TestSavePreviousArgs(t *testing.T) {
	ngsi := testNgsiLibInit()
	filename := ""
	ngsi.ConfigFile.SetFileName(&filename)

	err := ngsi.SavePreviousArgs()
	assert.NoError(t, err)
}

func TestSavePreviousArgsNoSave(t *testing.T) {
	ngsi := testNgsiLibInit()
	ngsi.PreviousArgs.UsePreviousArgs = false

	err := ngsi.SavePreviousArgs()
	assert.NoError(t, err)
}
