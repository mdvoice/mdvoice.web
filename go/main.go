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
	out  = "public/"
)

func main() {
	err := os.RemoveAll(out)
	checknil(err)
	file, err := template.ParseFiles(tmpl+"index.html", md+"首頁.gtpl")
	checknil(err)
	put, err := os.Create(out + "index.html")
	checknil(err)
	err = file.Execute(put, "")
	checknil(err)
	fmt.Printf("index done!")
	mds, err := ioutil.ReadDir(md + "活動/")
	checknil(err)
	for _, el := range mds {
		file, err := template.ParseFiles(tmpl+"index.html", md+"活動/"+el.Name()+"/index.md")
		checknil(err)
		err = os.Mkdir(el.Name(), os.ModePerm)
		checknil(err)
		put, err := os.Create(out + el.Name() + "/index.html")
		checknil(err)
		err = file.Execute(put, "")
		checknil(err)
		fmt.Printf("activity done!")
	}

}
