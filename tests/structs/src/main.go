package main

type Person struct {
	name string
	age  int
}

func newPerson(name string) Person {
	p := Person{name: name}
	p.age = 42
	return p
}

func main() {
	bob := Person{"Bob", 20}
	_ = bob

	alice := Person{name: "Alice", age: 30}
	_ = alice

	fred := Person{name: "Fred"}
	_ = fred

	ann := &Person{name: "Ann", age: 40}
	*ann = newPerson("Jon")
	_ = ann

	var sean Person
	sean.name = "Sean"
	sean.age = 50
	sp := &sean
	sp.age = 51
	_ = sean

	dog := struct {
		name   string
		isGood bool
	}{
		"Rex",
		true,
	}
	_ = dog
}
