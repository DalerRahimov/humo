package main

import (
	"HW5/ConstID"
	"HW5/Floors"
	"HW5/IDfile"
	"fmt"
)

/*
существует 3 Уровня доступа A,B,C
А- первый этаж B- вторй С третий
Id состои из указателя ур.доступа и номера сотрудника (А123456, B000012, C00001)
*/
func main() {

	flooring, _ := Floors.FloorNum()

	if flooring > 3 || flooring < 1 {

		fmt.Println(ConstID.FloorError)
		return

	} else {

		ID := IDfile.IDchec()
		switch flooring {

		case 1:
			fmt.Println(ConstID.AccessTrue)
		case 2:
			if ID == ConstID.B || ID == ConstID.C {
				fmt.Println(ConstID.AccessTrue)
			} else {
				fmt.Println(ConstID.AccessFalse)
			}
		case 3:
			if ID == ConstID.C {
				fmt.Println(ConstID.AccessTrue)
			} else {
				fmt.Println(ConstID.AccessFalse)
			}
		default:
			fmt.Println(ConstID.AccessFalse)

		}
	}
}
