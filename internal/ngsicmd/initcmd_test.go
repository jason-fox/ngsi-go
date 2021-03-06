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
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v2"
)

func TestInitCmdFalse(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	c := cli.NewContext(app, set, nil)

	_, err := initCmd(c, "Testing", false)

	assert.NoError(t, err)
}

func TestInitCmdBatch(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	setupFlagBool(set, "batch")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--batch"})

	_, err := initCmd(c, "Testing", false)

	assert.NoError(t, err)
}

func TestInitCmdTrue(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	c := cli.NewContext(app, set, nil)
	setupFlagString(set, "host")
	_ = set.Parse([]string{"--host=orion"})

	_, err := initCmd(c, "Testing", true)

	assert.NoError(t, err)
}

func TestInitCmdDefaultValues(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	c := cli.NewContext(app, set, nil)

	ngsi.PreviousArgs.Stderr = "err"
	ngsi.PreviousArgs.Syslog = "err"
	ngsi.PreviousArgs.CacheFile = "cache"

	_, err := initCmd(c, "Testing", false)

	assert.NoError(t, err)
}

func TestInitCmdConfig(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	c := cli.NewContext(app, set, nil)
	setupFlagString(set, "host")
	_ = set.Parse([]string{"--config=config-file", "--host=orion"})

	_, err := initCmd(c, "Testing", true)

	assert.NoError(t, err)
}

func TestInitCmdStderr(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	setupFlagString(set, "stderr,host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--stderr=err", "--host=orion"})

	_, err := initCmd(c, "Testing", true)

	assert.NoError(t, err)
}

func TestInitCmdErrorStderr(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	setupFlagString(set, "stderr")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--stderr=abc"})

	_, err := initCmd(c, "Testing", true)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "stderr logLevel error", ngsiErr.Message)
	}
}

func TestInitCmdSyslog(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	setupFlagString(set, "syslog,host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--syslog=off", "--host=orion"})

	_, err := initCmd(c, "Testing", true)

	assert.NoError(t, err)
}

func TestInitCmdSyslogLevelDebug(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	setupFlagString(set, "syslog,host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--syslog=debug", "--host=orion"})
	ngsi.SyslogLib = &MockSyslogLib{}

	_, err := initCmd(c, "Testing", true)

	assert.NoError(t, err)
}

func TestInitCmdSyslogWindows(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	setupFlagString(set, "syslog,host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--syslog=err", "--host=orion"})
	ngsi.OsType = "windows"

	_, err := initCmd(c, "Testing", true)

	assert.NoError(t, err)
}

func TestInitCmdErrorSyslogLevel(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	setupFlagString(set, "syslog")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--syslog=on"})

	_, err := initCmd(c, "Testing", true)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "syslog logLevel error", ngsiErr.Message)
	}
}

func TestInitCmdErrorSyslog(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	setupFlagString(set, "syslog")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--syslog=debug"})
	ngsi.SyslogLib = &MockSyslogLib{Err: errors.New("syslog new error")}

	_, err := initCmd(c, "Testing", true)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "syslog new error", ngsiErr.Message)
	}
}

func TestInitCmdErrorHostNotFound(t *testing.T) {
	_, set, app, _ := setupTest()

	c := cli.NewContext(app, set, nil)

	_, err := initCmd(c, "Testing", true)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "Required host not found", ngsiErr.Message)
	}
}

func TestInitCmdErrorInitTokenMgr(t *testing.T) {
	ngsi, set, app, _ := setupTest2()

	// setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	setupFlagString(set, "cacheFile")
	_ = set.Parse([]string{"--cacheFile=abc"})
	c := cli.NewContext(app, set, nil)
	ngsi.CacheFile = &MockIoLib{PathAbs: errors.New("error")}

	_, err := initCmd(c, "Testing", false)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "error ", ngsiErr.Message)
	}
}
func TestInitCmdArgs(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	setupFlagString(set, "margin,timeout,maxCount")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--margin=1", "--timeout=1", "--maxCount=0"})

	_, err := initCmd(c, "Testing", false)

	assert.NoError(t, err)
}

func TestInitCmdArgs2(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	setupFlagString(set, "margin,timeout,maxCount")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--margin=10000", "--timeout=10000", "--maxCount=5000"})

	_, err := initCmd(c, "Testing", false)

	assert.NoError(t, err)
}

func TestInitCmdArgs3(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	setupFlagString(set, "margin,timeout,maxCount")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--margin=181", "--timeout=61", "--maxCount=101"})

	_, err := initCmd(c, "Testing", false)

	assert.NoError(t, err)
}
