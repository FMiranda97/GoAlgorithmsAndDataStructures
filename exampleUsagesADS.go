package main

// This file serves as an example on how to use the ads package

import (
	ads "./ADS"
	"fmt"
	"math/rand"
)

type person struct {
	id   int
	name string
}

func (p person) CompareTo(t interface{}) int8 { // necessary for collections and sorts usage
	target := t.(*person) // (*person) to use collections with pass by referece, (person) to use collections with pass by value
	if p.id < target.id {
		return -1
	} else if p.id > target.id {
		return 1
	} else {
		return 0
	}
}

func createPerson(id int, nome string) *person { // constructor
	return &person{
		id:   id,
		name: nome,
	}
}

func retrievePerson(id int, tree ads.AVLTree) *person { // utility to retrieve person from collection searching by id
	a, err := tree.Get(person{id: id})
	if err != nil {
		fmt.Println("ERROR:", err)
	}
	return a.(*person)
}

func exampleCollectionUsage() {
	const nPersons = 20
	tree := ads.NewAVLTree() // initiate an AVLTree

	//populate the tree with 20 random individuals
	for i := 0; i < nPersons; i++ {
		p := createPerson(rand.Int()%(nPersons*100), "john doe") // returns pointer to a created person
		err := tree.Insert(p)                                    // insert pointer generated on previous line
		if err != nil {
			fmt.Println("Error:", err) // no errors
		}
	}

	err := tree.Insert(person{id: -1, name: "oops"})
	if err != nil {
		// error here because the first element inserted on a tree was a pointer
		// while here we decided to insert the actual structure
		fmt.Println("Error:", err)
	}

	err = tree.Insert(&person{id: -1, name: "nice one"})
	if err != nil {
		// this works because we pass the pointer to the struct
		fmt.Println("Error:", err)
	}

	tree.Print() // print contents of tree in order

	p := retrievePerson(1410, tree) // p is the element in the tree with id == 1410
	// since we populated the tree with pointers, changing p will affect the structure present in the tree
	p.name = "this a name we changed after retrieving the element from the tree"

	// let's print the tree again and verify that the name did change
	fmt.Println("------------------")
	tree.Print()
}

func exampleSortUsage() {
	const nPersons = 20
	var arr [nPersons]*person

	//populate the array with 20 random individuals
	for i := 0; i < nPersons; i++ {
		arr[i] = createPerson(rand.Int()%(nPersons*100), "john doe")
	}
	//print initial array
	for _, v := range arr {
		fmt.Println(v)
	}

	// directly sorting an array fails because Sort takes a slice as parameter
	err := ads.Sort(arr)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// sorting a slice of the whole array works and sorts the array. Here we use our package's default sorting algorithm
	err = ads.Sort(arr[:])
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("------------------")
	for _, v := range arr {
		fmt.Println(v)
	}

	fmt.Println("------------------")
	fmt.Println("------------------")
	//populate a new slice with 20 random individuals
	var slc []*person
	for i := 0; i < nPersons; i++ {
		slc = append(slc, createPerson(rand.Int()%(nPersons*100), "john doe"))
	}
	//print initial array
	for _, v := range slc {
		fmt.Println(v)
	}

	// sorting the slice works. Here we use QuickSortC which is a version of quick sort taking advantage of concurrency
	err = ads.QuickSortC(slc)
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("------------------")
	for _, v := range slc {
		fmt.Println(v)
	}
}

func main() {
	exampleCollectionUsage()
	fmt.Println("------------------")
	fmt.Println("------------------")
	fmt.Println("------------------")
	exampleSortUsage()
}
