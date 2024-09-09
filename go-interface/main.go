package main

type Quacker interface {
	Quack()
}

type Duck struct{}

func (d Duck) Quack() {
	println("Quack!")
}

type Person struct{}

func (d Person) Quack() {
	println("I can quack like a duck!")
}

func main() {
	var duck Duck
	var pers Person

	duck.Quack()
	pers.Quack()

}
