package main

import (
	"encoding/base64"
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

func decode(in []byte) ([]byte, error) {
	return applyData(in, func(data map[string]string) (map[string]string, error) {
		newData := make(map[string]string)
		for k, v := range data {
			decoded, err := base64.StdEncoding.DecodeString(v)
			if err != nil {
				return newData, err
			}
			newData[k] = string(decoded)
		}
		return newData, nil
	})
}

func encode(in []byte) ([]byte, error) {
	return applyData(in, func(data map[string]string) (map[string]string, error) {
		newData := make(map[string]string)
		for k, v := range data {
			encoded := base64.StdEncoding.EncodeToString([]byte(v))
			newData[k] = encoded
		}
		return newData, nil
	})
}

func applyData(in []byte, f func(map[string]string) (map[string]string, error)) ([]byte, error) {
	var obj map[string]interface{}
	if err := yaml.Unmarshal(in, &obj); err != nil {
		return in, err
	}

	data, ok := getData(obj)
	if !ok {
		return in, nil
	}

	newData, err := f(data)
	if err != nil {
		return in, nil
	}

	obj["data"] = newData
	return yaml.Marshal(obj)
}

func getData(obj map[string]interface{}) (map[string]string, bool) {
	rawData, ok := obj["data"]
	if !ok {
		fmt.Fprintln(os.Stderr, "data is not found")
		return nil, false
	}

	rawDataMap, ok := rawData.(map[interface{}]interface{})
	if !ok {
		fmt.Fprintln(os.Stderr, "data is not map")
		return nil, false
	}

	data := make(map[string]string)

	for ik, iv := range rawDataMap {
		if k, ok := ik.(string); ok {
			if v, ok := iv.(string); ok {
				data[k] = v
			} else {
				return nil, false
			}
		} else {
			return nil, false
		}
	}

	return data, true
}
