## Factory Pattern

## Factory Pattern là gì

Factory Pattern là một pattern thuộc nhóm creational patterns và nó là một trong những pattern được sử dụng phổ biến nhất. 

Factory Pattern đưa toàn bộ logic của việc tạo mới object vào trong factory, che giấu logic của việc khởi tạo và giảm sự phụ thuộc nhằm tăng tính mở rộng

Client chỉ tương tác với factory struct và cho biết loại instance cần được khởi tạo. Factory struct sẽ tương tác với các struct tương ứng và sau đó trả về một instance cụ thể với struct đó

## Ví dụ

- Chúng ta có interface ```iPost``` xác định tất cả các hành vi mà một bài viết phải có
- Có struct ```post``` để implements interface ```iPost```
- Có 2 loại bài viết cụ thể là : ```newsPost``` và ```normalPost```. Cả 2 bài viết đều nhúng struct ```post``` vào làm thuộc tính, do đó ```newsPost``` và ```normalPost``` đều thực hiện gián tiếp các phương thức của interface ```iPost``` và implements interface này
- Chúng ta tạo ra struct postFactory để có thể tạo ra loại bài viết ```newsPost``` hoặc ```normalPost```
- Cuối cùng file ```main.go``` đóng vai trò như một client và thay vì tương tác trực tiếp với ```newsPost``` hoặc ```normalPost``` struct thì nó sẽ tương tác với postFactory để tạo ra các instance tương ứng với ```newsPost``` hoặc ```normalPost```


## Code demo

```ipost.go```

```go
package factory

type IPost interface {
	GetTitle() string
	GetAuthor() string
	SetTitle(title string)
	SetAuthor(author string)
}
```

Ở đây chúng ta chúng ta khởi tạo interface IPost xác định các hành vi có thể có của đối tượng bao gồm:

- GetTitle() string : Lấy tiêu đề bài viết
- GetAuthor() string : Lấy tên tác giả bài viết
- SetTitle(title string) : Đặt lại tiêu đề cho bài viết
- SetAuthor(author string) : Đặt tên tác giả cho bài viết

```post.go```

```go
package factory

type post struct {
	Title string
	Author string
}

func (p *post) GetTitle() string {
	return p.Title
}

func (p *post) GetAuthor() string {
	return p.Author
}

func (p *post) SetAuthor(author string) {
	p.Author = author
}

func (p *post) SetTitle(author string) {
	p.Title = author
}
```

Trong file post.go, chúng ta tạo 1 struct có tên là post struct này bao gồm 2 thuộc tính : Title và Author. Đồng thời chúng ta cũng khai báo thêm 4 method của post struct. Lúc này post struct sẽ implements interface IPost

```news.go```

```go
package factory

type newsPost struct {
	post
}

func newNewsPost() IPost {
	return &newsPost {
		post : post{
			Title: "Tin tức buổi sáng",
			Author : "Nguyễn Thu Hằng",
		},
	}
}
```

Trong file ```news.go``` chúng ta tạo ra newsPost struct có 1 thuộc tính là post (ở đây là 1 dạng kế thừa), lúc này newsPost cũng gián tiếp inplements interface ```IPost```

Tiếp theo chúng ta tạo func newNewsPost() và trả về 1 đối tượng cụ thể


```normal.go```

```go
package factory

type normalPost struct {
	post
}

func newNormalPost() IPost {
	return &normalPost {
		post : post{
			Title: "Lập trình và cuộc sống",
			Author : "Nguyễn Minh Duy",
		},
	}
}
```

Chúng ta sẽ làm tương tự như với news.go

```postFactory.go```

```go
package factory

import "errors"

func GetPost(typePost string) (IPost, error) {
	if typePost == "news" {
		return newNewsPost(), nil
	}
	if typePost == "normal" {
		return newNormalPost(), nil
	}
	return nil, errors.New("typePost incorrect")
}
```

Chúng ta định nghĩa func ```GetPost``` ở đây có thể hiểu như là một factory để tạo ra các instance tương ứng với các typePost chúng ta truyền vào. Ví dụ nếu truyền vào là type "news" thì sẽ tạo ra instance của newsPost, truyền vào là type "news" thì sẽ tạo ra instance của newsPost, còn các type khác sẽ trả về 1 message error

```main.go```

```go
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
```

File ```main.go``` đóng vai trò như một client, ở đây để tạo ra đối tượng chúng ta chỉ cần tương tác với factory bằng cách truyền vào loại đối tượng mà chúng ta muốn khởi tạo vào trong func GetPost() chứ không tương tác trực tiếp với struct, việc tương tác với struct đã có factory làm cho chúng ta rồi

Chạy ```go run main.go``` và kết quả chúng ta đạt được như sau:

```
Title: Tin tức buổi sáng
Author: Nguyễn Thu Hằng

Title: Lập trình và cuộc sống
Author: Nguyễn Minh Duy

Title: Tin nóng tuần qua
Author: Nguyễn Bích Ngọc

typePost incorrect
```