package main

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
)

func main() {
	{
		kv := map[string]string{}
		data := "{\"key1\":\"v1\",\"key2\":\"v2\"}"
		err := json.Unmarshal([]byte(data), &kv)
		if err != nil {
			fmt.Printf("json.unmarshal failed, data: %s, err: %v\n", data, err)
		} else {
			fmt.Printf("data: %s, kv: %v\n", data, kv)
		}

		kv["key3"] = "v3"
		data2, err := json.Marshal(kv)
		if err != nil {
			fmt.Printf("json.Marshal failed, data2: %s, err: %v\n", data2, err)
		} else {
			fmt.Println(data)
			fmt.Println(string(data2))
		}
	}

	{
		//extParam := "{\"enable-coredns\":\"false\"}"
		extParam := "{}"
		newExtParam, err := setDefaultValue(extParam)
		if err != nil {
			fmt.Printf("setDefaultValue failed, extParam: %s, err: %s\n", extParam, err.Error())
		} else {
			fmt.Println("newExtParam:", extParam)
			fmt.Println("newExtParam:", newExtParam)
		}
	}
}

func setDefaultValue(extParam string) (string, error) {
	if extParam == "" {
		extParam = "{}"
	}
	kv := map[string]string{}
	err := json.Unmarshal([]byte(extParam), &kv)
	if err != nil {
		return "", errors.Wrapf(err, "invalid extParam: %s", extParam)
	}
	k := "enable-coredns"
	if _, ok := kv[k]; !ok {
		kv[k] = "false"
	}

	newExtParam, err := json.Marshal(kv)
	if err != nil {
		return "", errors.Wrapf(err, "json marshal failed: %v", kv)
	}

	return string(newExtParam), nil
}
