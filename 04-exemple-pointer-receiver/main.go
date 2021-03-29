package main

import (
	"fmt"
	"strings"
)

type Post struct {
	Title     string
	Text      string
	published bool
}

func (p Post) Headline() string {
	return fmt.Sprintf("%v - %v...", p.Title, p.Text[:50])
}

func (p Post) Published() bool { return p.published }

func (p *Post) Publish() {
	p.published = true
}

func (p *Post) Unpublish() {
	p.published = false
}

func UpperTitle(p *Post) {
	p.Title = strings.ToUpper(p.Title)
}

func main() {
	p := Post{
		Title: "Go release",
		Text: `Blablablabla Blablablabla Blablablabla BlablablablaBlablablabla Blablablabla 
		BlablablablaBlablablabla Blablablabla Blablablabla BlablablablaBlablablablaBlablablabla`,
	}

	fmt.Println(p.Headline())
	fmt.Printf("Post published? %v\n", p.Published())
	p.Publish()
	fmt.Printf("Post published? %v\n", p.Published())
	p.Unpublish()
	fmt.Printf("Post published? %v\n", p.Published())
	UpperTitle(&p)
	fmt.Println(p.Headline())

	pythonPost := &Post{
		Title: "Python Intro",
		Text: `Blablablabla Blablablabla Blablablabla BlablablablaBlablablabla Blablablabla 
		BlablablablaBlablablabla Blablablabla Blablablabla BlablablablaBlablablablaBlablablabla`,
	}

	UpperTitle(pythonPost)
	fmt.Println(pythonPost.Headline())

}
