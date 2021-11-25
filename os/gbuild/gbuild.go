// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

// Package gbuild manages the build-in variables from "gf build".
package gbuild

import (
	"github.com/donetkit/gtool/container/gvar"
	"github.com/donetkit/gtool/encoding/gbase64"
	"github.com/donetkit/gtool/internal/json"
	"runtime"
)

var (
	builtInVarStr = ""                       // Raw variable base64 string, which is injected by go build flags.
	builtInVarMap = map[string]interface{}{} // Binary custom variable map decoded.
)

func init() {
	// The `builtInVarStr` is injected by go build flags.
	if builtInVarStr != "" {
		err := json.UnmarshalUseNumber(gbase64.MustDecodeString(builtInVarStr), &builtInVarMap)
		if err != nil {

		}
		builtInVarMap["gfVersion"] = runtime.Version()
		builtInVarMap["goVersion"] = runtime.Version()
	}
}

// Info returns the basic built information of the binary as map.
// Note that it should be used with gf-cli tool "gf build",
// which automatically injects necessary information into the binary.
func Info() map[string]string {
	return map[string]string{
		"gf":   Get("gfVersion").String(),
		"go":   Get("goVersion").String(),
		"git":  Get("builtGit").String(),
		"time": Get("builtTime").String(),
	}
}

// Get retrieves and returns the build-in binary variable with given name.
func Get(name string, def ...interface{}) *gvar.Var {
	if v, ok := builtInVarMap[name]; ok {
		return gvar.New(v)
	}
	if len(def) > 0 {
		return gvar.New(def[0])
	}
	return nil
}

// Map returns the custom build-in variable map.
func Map() map[string]interface{} {
	return builtInVarMap
}
