// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/donetkit/gtool.

// Package gyaml provides accessing and converting for YAML content.
package gyaml

import (
	"github.com/donetkit/gtool/internal/json"
	"github.com/donetkit/gtool/util/gconv"
	"gopkg.in/yaml.v3"
)

func Encode(v interface{}) ([]byte, error) {
	return yaml.Marshal(v)
}

func Decode(v []byte) (interface{}, error) {
	var result map[string]interface{}
	if err := yaml.Unmarshal(v, &result); err != nil {
		return nil, err
	}
	return gconv.MapDeep(result), nil
}

func DecodeTo(v []byte, result interface{}) error {
	return yaml.Unmarshal(v, result)
}

func ToJson(v []byte) ([]byte, error) {
	if r, err := Decode(v); err != nil {
		return nil, err
	} else {
		return json.Marshal(r)
	}
}
