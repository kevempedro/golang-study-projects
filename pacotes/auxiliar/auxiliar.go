package auxiliar

import (
	"fmt"

	"github.com/badoux/checkmail"
)

func ValidadorEmail() {
	err := checkmail.ValidateFormat("kevem@gmail.com")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("E-mail válido")
}
