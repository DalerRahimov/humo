package IDfile

import (
	"HW5/ConstID"
	"fmt"
	"unicode/utf8"
)

func IDchec() (inputID string) {
	fmt.Println("Введите свой ID (A123456): ")
	fmt.Scan(&inputID)
	levelID := string(inputID[0])

	if utf8.RuneCountInString(inputID) > 7 {
		fmt.Println(ConstID.IDError)
		return
	} else if levelID == ConstID.A || levelID == ConstID.C || levelID == ConstID.B {
		inputID = string(inputID[0])
		return
	} else {
		fmt.Println(ConstID.IDError)
		return
	}

	return
}
