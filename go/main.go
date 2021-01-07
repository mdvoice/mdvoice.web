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
	act  = "activity/"
	des  [][]string
	err  error
)

// type index struct {
// 	首頁 map[string]string
// 	公告 []interface{}
// }

func main() {

	//init

	os.RemoveAll(out + act)
	os.Remove(out + "index.html")

	// checknil(err)
	err = os.Mkdir(out+act, os.ModePerm)
	checknil(err)

	// b, err := json.MarshalIndent(value, "", "  ")
	// checknil(err)
	// fmt.Print(string(b))

	mds, err := ioutil.ReadDir(md + "活動/")
	checknil(err)
	// news := make([]string, len(mds))
	for i, el := range mds {
		tmp := out + act + el.Name()
		mdfile := md + "活動/" + el.Name() + "/index.md"
		value := getvalue(mdfile)
		file, err := template.New("tmpl").
			Funcs(template.FuncMap{"kaiyo": func() string { return "" }}).
			ParseFiles(tmpl+"activity.html", mdfile)
		checknil(err)
		err = os.Mkdir(tmp, os.ModePerm)
		checknil(err)
		put, err := os.Create(tmp + "/index.html")
		checknil(err)
		err = file.Execute(put, value)
		checknil(err)
		des = append(des, []string{el.Name(), value["簡介"]})
		fmt.Printf("activity done! %d / %d \n", i+1, len(mds))
	}

	file, err := template.ParseFiles(tmpl + "index.html")
	checknil(err)
	value := getvalue(md + "首頁.gtpl")
	put, err := os.Create(out + "index.html")
	checknil(err)
	exe := struct {
		Home map[string]string
		News [][]string
	}{
		Home: value,
		News: des,
	}
	err = file.Execute(put, exe)
	checknil(err)
	fmt.Printf("index done! \n")

}
