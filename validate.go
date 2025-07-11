package creditcard

import (
	"fmt"
	"strconv"
)

func Validate(s string) {
	var validate_number int
	var card string
	card = "135413218942645a"

	last_char := card[len(card)-1]

	validate_number, err := strconv.Atoi(string(last_char))
	if err != nil {
		fmt.Println("Ошибка преоброзование:", err)
		return
	}

	fmt.Println(validate_number)
}
