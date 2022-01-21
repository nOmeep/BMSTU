package main

import (
	"fmt"
	"github.com/IzeBerg/rss-parser-go"
	"github.com/mmcdole/gofeed"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func parseHandler(w http.ResponseWriter, r *http.Request) {
	var (
		rssObject *rss.RSS
		err error
	)

	var num int
	num, _ = strconv.Atoi(r.FormValue("number"))

	doParse := true

	switch num {
	case 1:
		fallthrough
	case 2:
		fmt.Fprintf(w, "One or two\n")
		rssObject, err = rss.ParseRSS("http://blagnews.ru/rss_vk.xml")
	case 3:
		fmt.Fprintf(w, "Three\n")
		rssObject, err = rss.ParseRSS("http://www.rssboard.org/files/sample-rss-2.xml")
	case 4:
		fmt.Fprintf(w, "Four\n")
		rssObject, err = rss.ParseRSS("https://lenta.ru/rss")
	case 5:
		fmt.Fprintf(w, "Five\n")
		rssObject, err = rss.ParseRSS("https://news.mail.ru/rss/90/")
	case 6:
		fmt.Fprintf(w, "Six\n")
		rssObject, err = rss.ParseRSS("http://technolog.edu.ru/index.php?option=com_k2&view=itemlist&layout=category&task=category&id=8&lang=ru&format=feed")
	case 7:
		fmt.Fprintf(w, "Seven\n")
		rssObject, err = rss.ParseRSS("https://vz.ru/rss.xml")
	case 8:
		fmt.Fprintf(w, "Eight\n")
		rssObject, err = rss.ParseRSS("http://news.ap-pa.ru/rss.xml")
	default:
		fmt.Fprintf(w, "Wrong input")
		doParse = false
	}

	if err == nil && doParse {
		fmt.Fprintf(w, "Title           : %s\n", rssObject.Channel.Title)
		fmt.Fprintf(w, "Generator       : %s\n", rssObject.Channel.Generator)
		fmt.Fprintf(w, "PubDate         : %s\n", rssObject.Channel.PubDate)
		fmt.Fprintf(w, "LastBuildDate   : %s\n", rssObject.Channel.LastBuildDate)
		fmt.Fprintf(w, "Description     : %s\n", rssObject.Channel.Description)

		fmt.Printf("Number of Items : %d\n", len(rssObject.Channel.Items))

		for v := range rssObject.Channel.Items {
			item := rssObject.Channel.Items[v]
			fmt.Println()
			fmt.Fprintf(w, "Item Number : %d\n", v)
			fmt.Fprintf(w, "Title       : %s\n", item.Title)
			fmt.Fprintf(w, "Link        : %s\n", item.Link)
			fmt.Fprintf(w, "Description : %s\n", item.Description)
			fmt.Fprintf(w, "Guid        : %s\n", item.Guid.Value)
		}
	} else {
		fmt.Println(err)
	}
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,"Get user:\n")

	r.ParseForm()
	fmt.Fprintln(w, "Вывод формы:", r.Form)
	fmt.Println("Вывод формы:", r.Form)

	name := r.FormValue("name")
	age := r.FormValue("age")

	fmt.Fprintf(w, "Имя: %s\nВозраст: %s", name, age)
}

func HomeRouterHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //анализ аргументов,
	fmt.Println(r.Form) // ввод информации о форме на стороне сервера
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	http.ServeFile(w, r, "user.html")
}

func linkHandler(w http.ResponseWriter, r *http.Request) {

	r.ParseForm() //анализ аргументов,

	fmt.Println(r.Form) // ввод информации о форме на стороне сервера

	fmt.Println("path", r.URL.Path)

	fmt.Println("scheme", r.URL.Scheme)

	fmt.Println(r.Form["url_long"])

	for k, v := range r.Form {

		fmt.Println("key:", k)

		fmt.Println("val:", strings.Join(v, ""))

	}
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL("https://news.mail.ru/rss/90/")

	content := ` <html lang="ru"> <head> <title>Тест</title> <meta http-equiv="Content-type" content="text/html;charset=UTF-8" /> </head> <body>`
	content += `<h1 style="font-size: 30px;">` + feed.Title + `</h1>`
	content += `<br>`
	for _, x := range feed.Items {
		content += `<div>`
		content += `<a href="` + x.Link + `" style="font-size: 23px;">` + x.Title + `</a>`
		content += `<p>` + x.Description + `</p>`
		content += `<p>` + x.Published + `</p>`
		content += `</div>`
		content += `<br>`
	}
	content += `</body> </html>`
	tmpl, _ := template.New("example").Parse(content)
	tmpl.Execute(w, content)

}

func main() {
	http.HandleFunc("/postUser", GetHandler)
	http.HandleFunc("/link", linkHandler)
	http.HandleFunc("/parse", parseHandler)
	http.HandleFunc("/", HomeRouterHandler)  // установим роутер
	err := http.ListenAndServe(":9000", nil) // задаем слушать порт
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}