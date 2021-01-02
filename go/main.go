package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
)

var (
	tmpl = "template/"
	md   = "markdown/"
	out  = "public/html/"
	news []string
)

func main() {

	//init
	err := os.RemoveAll(out)
	checknil(err)
	err = os.Mkdir(out, os.ModePerm)
	checknil(err)

	// b, err := json.MarshalIndent(value, "", "  ")
	// checknil(err)
	// fmt.Print(string(b))

	mds, err := ioutil.ReadDir(md + "活動/")
	checknil(err)
	// news := make([]string, len(mds))
	for i, el := range mds {
		value := getvalue(md + "活動/" + el.Name() + "/index.md")
		file, err := template.ParseFiles(tmpl + "activity.html")
		checknil(err)
		err = os.Mkdir(out+el.Name(), os.ModePerm)
		checknil(err)
		put, err := os.Create(out + el.Name() + "/index.html")
		checknil(err)
		err = file.Execute(put, value)
		checknil(err)
		news = append(news, value["簡介"])
		fmt.Printf("activity done! %d / %d", i+1, len(mds))
	}
	// b, err := json.MarshalIndent(temp, "", "  ")
	// checknil(err)
	// fmt.Print(string(b))

	file, err := template.ParseFiles(tmpl + "index.html")
	checknil(err)
	value := getvalue(md + "首頁.gtpl")
	put, err := os.Create(out + "index.html")
	checknil(err)
	value["最新消息"] = news
	err = file.Execute(put, value)
	checknil(err)
	fmt.Printf("index done! \n")

}
