package main

import (
	"bytes"
	"html/template"
	"strings"
)

func checknil(e error) {
	if e != nil {
		panic(e)
	}
}
func checkstr(e string) {
	if e == "" {
		panic(e)
	}
}

func getvalue(v string) map[string]string {
	var (
		null bytes.Buffer
		temp = make(map[string]string)
	)
	file, err := template.New("config").
		Funcs(template.FuncMap{"kaiyo": func(k string, v string) string {
			temp[k] = v
			return ""
		},
		}).
		ParseFiles(v)
	checknil(err)
	err = file.Execute(&null, temp)
	checknil(err)
	for i, el := range temp {
		temp[i] = strings.Replace(el, "\n", "", -1)
	}
	return temp
}
