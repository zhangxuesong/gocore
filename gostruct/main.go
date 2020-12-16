package main

import "fmt"

// AnimalCategory 代表动物分类学中的基本分类法。
type AnimalCategory struct {
	kingdom string // 界。
	phylum  string // 门。
	class   string // 纲。
	order   string // 目。
	family  string // 科。
	genus   string // 属。
	species string // 种。
}

type Animal struct {
	scientificName string // 学名。
	AnimalCategory        // 动物基本分类。
}

type Cat struct {
	name string
	Animal
}

func (cat Cat) String() string {
	return fmt.Sprintf("%s (category: %s, name: %q)",
		cat.scientificName, cat.Animal.AnimalCategory, cat.name)
}

func (ac AnimalCategory) String() string {
	return fmt.Sprintf("%s%s%s%s%s%s%s",
		ac.kingdom, ac.phylum, ac.class, ac.order,
		ac.family, ac.genus, ac.species)
}

func (a Animal) Category() string {
	return a.AnimalCategory.String()
}

//func (a Animal) String() string {
//	return fmt.Sprintf("%s (category: %s)",
//		a.scientificName, a.AnimalCategory)
//}

type S1 struct {
	S1 string
}

func (s1 S1) String() string {
	return fmt.Sprint("this is : s1.\n")
}

type S2 struct {
	S2 string
}

func (s2 S2) String() string {
	return fmt.Sprint("this is : s2.\n")
}

type Base struct {
	Name string
	S1
	S2
}

//func (bs Base) String() string {
//	return fmt.Sprint("this is : base.\n")
//}

func main() {
	b := Base{}
	fmt.Printf("base %s.\n", b)

}
