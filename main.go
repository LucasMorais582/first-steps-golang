// Conceitos básicos da linguagem
package main

import "fmt"

type UUID string
type User struct {
	id    UUID
	name  string
	email string
	age   int
}

// executando o arquivo: go run main.go

// Verbos de Formatação Comuns
// %v: Valor padrão
// %+v: Valor com nomes dos campos para structs
// %#v: Go-sintaxe de representação do valor
// %T: Tipo do valor
// %d: Números inteiros
// %f: Números de ponto flutuante
// %s: Strings
// %q: String entre aspas
// %t: Booleanos

func (userData User) displayName() string {
	return fmt.Sprintf("Dearly %s,", userData.name)
}

func main() {

	const firstString string = "Bora Lucas"
	secondString := "Aprendendo short cut" // variáveis não podem ser declaradas dessa forma globalmente

	println(firstString)
	println(secondString)

	// arrays
	fmt.Print("Arrays...\n\n")
	var firstArray [5]int = [5]int{10, 20, 30, 40, 50}
	// secondArray := [...]int{10, 20, 30}

	for index, value := range firstArray {
		fmt.Printf("Index value: %d, value set: %d \n\n", index, value)
	}

	// slices
	fmt.Print("Slices...\n\n")
	s := []int{10, 20, 30, 40, 50, 60, 70, 80, 90}
	s2 := []int{99, 89, 79, 69, 59, 49, 39, 29, 19}

	fmt.Printf("length= %d, capacity= %d %v\n\n", len(s), cap(s), s)
	fmt.Printf("length= %d, capacity= %d %v\n\n", len(s[:4]), cap(s[:4]), s[:4])

	// como retiramos os primeiros elementos, a capacidade do array muda
	fmt.Printf("length= %d, capacity= %d %v\n\n", len(s[4:]), cap(s[4:]), s[4:])

	//adicionando valores ao slice
	s = append(s, 110, 120, 130)
	fmt.Printf("length= %d, capacity= %d %v\n\n", len(s), cap(s), s)

	//adicionando slice à outro slice
	s = append(s, s2...)
	fmt.Printf("length= %d, capacity= %d %v\n\n", len(s), cap(s), s)

	// maps
	salaries := map[string]int{"Lucas": 1000, "Matheus": 1100, "Fernando": 2000}
	fmt.Println(salaries["Lucas"], "\n")
	delete(salaries, "Fernando")
	salaries["Sub"] = 3000

	for _, salary := range salaries {
		fmt.Printf("The salary's is: %d\n", salary)
	}

	// struct
	fmt.Print("Structs...\n\n")

	user := User{
		id:    "123e4567-e89b-12d3-a456-426614174000",
		name:  "Lucas",
		email: "lucasmorais582@example.com",
	}

	fmt.Println(user.displayName())

	// tipos de for
	for count := 0; count < 11; count++ {
		fmt.Println(count)
	}

	numbers := []string{"um", "dois", "três"}
	for key, value := range numbers {
		fmt.Println(key, value)
	}

	count2 := 0
	for count2 < 10 {
		fmt.Println(count2)
	}

	for {
		fmt.Println("Infinite for")
	}

}
