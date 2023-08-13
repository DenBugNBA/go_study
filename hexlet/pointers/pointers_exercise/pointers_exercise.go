package main

import "fmt"

type Parent struct {
	Name     string
	Children []Child
}

type Child struct {
	Name string
	Age  int
}

func CopyParent(p *Parent) Parent {
	if p == nil {
		return Parent{}
	}
	newParent := Parent{}

	newParent.Name = p.Name
	for _, child := range p.Children {
		newParent.Children = append(newParent.Children, Child{Name: child.Name, Age: child.Age})
	}

	return newParent
}

func CopyParentTeacher(p *Parent) Parent {
	if p == nil {
		return Parent{}
	}
	children := make([]Child, len(p.Children))
	copy(children, p.Children)
	return Parent{
		Name:     p.Name,
		Children: children,
	}
}

func main() {
	cpNil := CopyParent(nil) // Parent{}
	fmt.Println(cpNil)

	p := &Parent{
		Name: "Harry",
		Children: []Child{
			{
				Name: "Andy",
				Age:  18,
			},
		},
	}
	cp := CopyParent(p)

	// при мутациях в копии "cp"
	// изначальная структура "p" не изменяется
	cp.Children[0] = Child{}
	fmt.Println(p.Children) // [{Andy 18}]

}
