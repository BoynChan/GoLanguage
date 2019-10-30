package main

import (
	"html/template"
	"log"
	"os"
)

/*
在这里,b的类型是html,所以其<b>不会被转义
而a的类型是string,所以其<b>会被认为是字符串一起输出
*/
func main() {
	const templ = `<p>A: {{.A}}</p><p>B: {{.B}}</p>`
	t := template.Must(template.New("escape").Parse(templ))
	var data struct {
		A string        // untrusted plain text
		B template.HTML // trusted HTML
	}
	data.A = "<b>Hello!</b>"
	data.B = "<b>Hello!</b>"
	if err := t.Execute(os.Stdout, data); err != nil {
		log.Fatal(err)
	}
}
