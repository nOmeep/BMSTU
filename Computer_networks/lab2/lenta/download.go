package main

import (
	"fmt"
	"github.com/mgutz/logxi/v1"
	"golang.org/x/net/html"
	"net/http"
)

func getAttr(node *html.Node, key string) string {
	for _, attr := range node.Attr {
		if attr.Key == key {
			return attr.Val
		}
	}
	return ""
}

func getChildren(node *html.Node) []*html.Node {
	var children []*html.Node
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		children = append(children, c)
	}
	return children
}

func isElem(node *html.Node, tag string) bool {
	return node != nil && node.Type == html.ElementNode && node.Data == tag
}

func isText(node *html.Node) bool {
	return node != nil && node.Type == html.TextNode
}

func isDiv(node *html.Node, class string) bool {
	return isElem(node, "div") && getAttr(node, "class") == class
}

type Item struct {
	Ref, Title, User string
}

func readItem(item *html.Node) *Item {
	log.Info("Starting reading")
	a := item.FirstChild.NextSibling.NextSibling.NextSibling
	name := a.FirstChild.NextSibling.NextSibling.NextSibling.FirstChild.NextSibling
	log.Info("Curr " + getAttr(name, "href"))
	a = a.FirstChild.NextSibling.FirstChild.NextSibling
	if len(getChildren(a)) == 3 {
		href := getAttr(a, "href")
		log.Info("href: " + href)
		data := a.FirstChild.NextSibling.FirstChild.Data
		fmt.Println("A IS: ", data)

		return &Item{
			Ref: href,
			Title: data,
			User: getAttr(name, "href"),
		}
	}
	return nil
}

func search(node *html.Node) []*Item {
	if isElem(node, "ul") && getAttr(node, "class") == "forum-section__list" {
		log.Info("Find section" + getAttr(node, "class"))
		var items []*Item
		for c := node.FirstChild.NextSibling; c != nil; c = c.NextSibling {
			if (getAttr(c, "class")) == "forum-section__item unreaded sticky" || (getAttr(c, "class")) == "forum-section__item readed sticky" || (getAttr(c, "class")) == "forum-section__item unreaded"{
				log.Info("Reading item: " + getAttr(c, "class"))
				if item := readItem(c); item != nil {
					items = append(items, item)
				}
			}
		}
		return items
	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		if items := search(c); items != nil {
			return items
		}
	}
	return nil
}

func downloadNews() []*Item {
	log.Info("sending request to https://dota2.ru/forum/forums/zhelezo-novosti-i-obsuzhdenija.166/")
	if response, err := http.Get("https://dota2.ru/forum/forums/zhelezo-novosti-i-obsuzhdenija.166/"); err != nil {
		log.Error("request to https://dota2.ru/forum/forums/zhelezo-novosti-i-obsuzhdenija.166/", "error", err)
	} else {
		defer response.Body.Close()
		status := response.StatusCode
		log.Info("got response from dota", "status", status)
		if status == http.StatusOK {
			if doc, err := html.Parse(response.Body); err != nil {
				log.Error("invalid HTML from dota", "error", err)
			} else {
				log.Info("HTML from dota parsed successfully")
				return search(doc)
			}
		}
	}
	return nil
}
