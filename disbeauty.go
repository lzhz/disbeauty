package disbeauty

import (
	"errors"
	//"fmt"
	"strings"
)

var gValueMap map[string]interface{} = map[string]interface{}{}
var gHandlerMap map[string]func(args []string) string = map[string]func(args []string) string{
	"Key": func(args []string) string {
		v, ok := gValueMap[args[0]]
		if !ok {
			return ""
		}

		if len(args) == 1 {
			return DumpSimple(v)
		} else {
			return DumpSimple(GetInValue(v, args[1:]))
		}
	},
	"Val": func(args []string) string {
		v, ok := gValueMap[args[0]]
		if !ok {
			return ""
		}

		if len(args) == 1 {
			return DumpString(v)
		} else {
			return DumpString(GetInValue(v, args[1:]))
		}
	},
}

func RegisterData(key string, val interface{}) error {
	if _, ok := gValueMap[key]; ok {
		return errors.New("duplicate key " + key)
	}

	gValueMap[key] = val
	return nil
}

func Exec(cmd string) string {
	/*
		defer func() {
			if r := recover(); r != nil {
				fmt.Println(r)
			}
		}()
	*/

	arr := strings.Fields(cmd)
	if len(arr) < 2 {
		return "invalid cmd: " + cmd
	}

	if f, ok := gHandlerMap[arr[0]]; ok {
		return f(arr[1:])
	} else {
		return "invalid cmd: " + cmd
	}
}
