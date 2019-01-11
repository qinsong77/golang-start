package src

import "fmt"

type Human struct {
	Name string
	Age int
	Phone string
}

type Student struct {
	Human
	School string
	Loan float32
}

type Employee struct {
	Human
	Company string
	Money float32
}

func (h Human) SayHi()  {
	fmt.Printf("Hi, I am %s you can call me on %s\n",h.Name,h.Phone)
}
func (h Human) Sing(lyrics string)  {
	fmt.Println("La la la la la ...",lyrics)
}
func (e Employee) SayHi()  {
	fmt.Printf("Hi，I am %s, I work at %s.Call me on %s\n",e.Name,e.Company,e.Phone)
}

type Men interface {
	SayHi()
	Sing(lyrics string)
}

func Testinterface()  {
	mike := Student{Human{"Mike", 25, "222-222-XXX"}, "MIT", 0.00}
	paul := Student{Human{"Paul", 26, "111-222-XXX"}, "Harvard", 100}
	sam := Employee{Human{"Sam", 36, "444-222-XXX"}, "Golang Inc.", 1000}
	tom := Employee{Human{"Tom", 37, "222-444-XXX"}, "Things Ltd.", 5000}
	var i Men
	i = mike
	fmt.Println("This is Mike,a Student:")
	i.SayHi()
	i.Sing("November rain")

	i = tom
	fmt.Println("This is tom, an Employee:")
	i.SayHi()
	i.Sing("Born to be wild")
	//定义了slice Men
	fmt.Println("Let's use a slice of Men and see what happens")
	x := make([]Men,3)
	x[0],x[1],x[2] = paul,sam,mike
	for _,value := range x{
		value.SayHi()
	}
}