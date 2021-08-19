package main

import (
	"factory-pattern/factory"
	"fmt"
)

func main() {
    // Tạo ra các instance của đối tượng
	news, _ := factory.GetPost("news")
	normal, _ := factory.GetPost("normal")

    // In ra thông tin của các instance
	printDetailsPost(news)
	printDetailsPost(normal)

    // Cập nhật lại 1 số thông tin cho instance đã được khởi tạo trước đó
	news.SetAuthor("Nguyễn Bích Ngọc")
	news.SetTitle("Tin nóng tuần qua")
	printDetailsPost(news)

    // Trường hợp truyền vào loại đối tượng không chính xác
    _, err := factory.GetPost("other")
    if err != nil {
        fmt.Println(err)
    }
}

func printDetailsPost(p factory.IPost) {
    fmt.Printf("Title: %s \n", p.GetTitle())
    fmt.Printf("Author: %s \n\n", p.GetAuthor())
}