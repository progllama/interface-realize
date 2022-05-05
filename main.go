package main

func main() {
	parser := NewParser()
	parser.Parse(text)
}

const text string = `package main
type Gopher interface {
	Hello()
	Say(string)
	Introduce() string
	Repeat(string) string
}

type RealGopher struct {
}

type Popher interface {
	GoodNight(Gopher)
	Write(string)
	Read() string
}
`
