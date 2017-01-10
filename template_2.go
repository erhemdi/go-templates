package main

import "html/template"
import "os"

type Friend struct {
	Fname string
}

type Person struct {
	Username string
	Emails   []string
	Friends  []*Friend
}

func main() {
	f1 := Friend{Fname: "Artour"}
	f2 := Friend{Fname: "Sumail"}
	t := template.New("fieldname example")
	t, _ = t.Parse(
		`hello {{.Username}} !
        {{range .Emails}}
            an email {{.}}
        {{end}}
        {{with .Friends}}
            {{range .}}
                my friend name is {{.Fname}}
            {{end}}
        {{end}}
        `)

	p := Person{
		Username: "Evil Geniuses",
		Emails:   []string{"arteezy@eg.gg", "sumail@eg.gg"},
		Friends:  []*Friend{&f1, &f2},
	}

	t.Execute(os.Stdout, p)
}
