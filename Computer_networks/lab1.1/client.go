package main

import (
	"bytes"
	"fmt"
	"github.com/jlaffaye/ftp"
	"github.com/skorobogatov/input"
	"io/ioutil"
	"log"
	"time"
)

func main() {
	c, err := ftp.Dial("localhost:2121", ftp.DialWithTimeout(5*time.Second))
	if err != nil {
		log.Fatal("Error ftp.Dial(): ", err)
	}

	err = c.Login("admin", "123456")
	if err != nil {
		log.Fatal("Error c.Login(): ", err)
	}

	fmt.Println("Enter a command")
	var command string = input.Gets()
	
	for command != "quit" {
		switch command {
		case "show":
			list, err := c.List("/")
			if err != nil {
				log.Fatal("Failed to show")
			}
			for _, v := range list {
				fmt.Println(v.Name)
			}

		case "upload":
			data := bytes.NewBufferString("Это файл Панова Севастьяна. Кто прочитал, тот прочитал")
			if err = c.Stor("test_Seva_file.txt", data); err != nil {
				log.Fatal("Failed to upload")
			} else {
				fmt.Println("Upload successful")
			}
		case "mkdir":
			fmt.Println("Enter dir name:")
			if err := c.MakeDir(input.Gets()); err != nil {
				log.Fatal("Failed to create a dir")
			}
		case "rmFile":
			fmt.Println("Enter file name:")
			if err := c.Delete(input.Gets()); err != nil {
				log.Fatal("Failed to dalete a file")
			}
		case "rmDir":
			fmt.Println("Enter dir name:")
			if err := c.RemoveDirRecur(input.Gets()); err != nil {
				log.Fatal("Failed to delete dir")
			}
		case "get":
			fmt.Println("Enter file name:")
			r, err := c.Retr(input.Gets())
			if err != nil {
				log.Fatal("Failed to do \"get\"")
			}
			defer r.Close()

			buf, err := ioutil.ReadAll(r)
			fmt.Println("File content is: \"", string(buf), "\"")
		default:
			fmt.Println("Unknown command")
		}
		command = input.Gets()
	}

	if err := c.Quit(); err != nil {
		log.Fatal("Error c.Quit(): ", err)
	} else {
		fmt.Println("Quit success")
	}
}
