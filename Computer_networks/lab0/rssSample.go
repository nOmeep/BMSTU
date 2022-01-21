package main

import (
	"fmt"
	"github.com/IzeBerg/rss-parser-go"
)

type MyError struct{}

func (m MyError) Error() string {
	panic("implement me")
}

func sayHello() (string, *MyError) {
	return "", &MyError{}
}

func main() {

	var (
		rssObject *rss.RSS
		err error
	)

	var num int
	fmt.Scan(&num)

	switch num {
	case 1:
		fallthrough
	case 2:
		fmt.Println("Second or First")
		rssObject, err = rss.ParseRSS("http://blagnews.ru/rss_vk.xml")
	case 3:
		fmt.Println("Three")
		rssObject, err = rss.ParseRSS("http://www.rssboard.org/files/sample-rss-2.xml")
	case 4:
		fmt.Println("Four")
		rssObject, err = rss.ParseRSS("https://lenta.ru/rss")
	case 5:
		fmt.Println("Five")
		rssObject, err = rss.ParseRSS("https://news.mail.ru/rss/90/")
	case 6:
		fmt.Println("Six")
		rssObject, err = rss.ParseRSS("http://technolog.edu.ru/index.php?option=com_k2&view=itemlist&layout=category&task=category&id=8&lang=ru&format=feed")
	case 7:
		fmt.Println("Seven")
		rssObject, err = rss.ParseRSS("https://vz.ru/rss.xml")
	case 8:
		fmt.Println("Eight")
		rssObject, err = rss.ParseRSS("http://news.ap-pa.ru/rss.xml")
	default:
		_, err = sayHello()
	}

	//rssObject, err = rss.ParseRSS("http://blagnews.ru/rss_vk.xml")
	if err == nil {
		fmt.Printf("Title           : %s\n", rssObject.Channel.Title)
		fmt.Printf("Generator       : %s\n", rssObject.Channel.Generator)
		fmt.Printf("PubDate         : %s\n", rssObject.Channel.PubDate)
		fmt.Printf("LastBuildDate   : %s\n", rssObject.Channel.LastBuildDate)
		fmt.Printf("Description     : %s\n", rssObject.Channel.Description)

		fmt.Printf("Number of Items : %d\n", len(rssObject.Channel.Items))

		for v := range rssObject.Channel.Items {
			item := rssObject.Channel.Items[v]
			fmt.Println()
			fmt.Printf("Item Number : %d\n", v)
			fmt.Printf("Title       : %s\n", item.Title)
			fmt.Printf("Link        : %s\n", item.Link)
			fmt.Printf("Description : %s\n", item.Description)
			fmt.Printf("Guid        : %s\n", item.Guid.Value)
		}
	} else {
		fmt.Println(err)
	}
}