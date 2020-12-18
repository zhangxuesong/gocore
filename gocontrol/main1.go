package main

func main() {
	//numbers := [...]int8{0, 1, 2, 3, 4, 5, 6}
	//switch 1 + 3 {
	//case numbers[0], numbers[1]:
	//	fmt.Println("0 or 1")
	//case numbers[2], numbers[3]:
	//	fmt.Println("2 or 3")
	//case numbers[4], numbers[5], numbers[6]:
	//	fmt.Println("4 or 5 or 6")
	//}

	//	.\main1.go:8:2: invalid case numbers[0] in switch on 1 + 3 (mismatched types int8 and int)
	//	.\main1.go:8:2: invalid case numbers[1] in switch on 1 + 3 (mismatched types int8 and int)
	//	.\main1.go:10:2: invalid case numbers[2] in switch on 1 + 3 (mismatched types int8 and int)
	//	.\main1.go:10:2: invalid case numbers[3] in switch on 1 + 3 (mismatched types int8 and int)
	//	.\main1.go:12:2: invalid case numbers[4] in switch on 1 + 3 (mismatched types int8 and int)
	//	.\main1.go:12:2: invalid case numbers[5] in switch on 1 + 3 (mismatched types int8 and int)
	//	.\main1.go:12:2: invalid case numbers[6] in switch on 1 + 3 (mismatched types int8 and int)

	//numbers := [...]int8{0, 1, 2, 3, 4, 5, 6}
	//switch numbers[4] {
	//case 0, 1:
	//	fmt.Println("0 or 1")
	//case 2, 3:
	//	fmt.Println("2 or 3")
	//case 4, 5, 6:
	//	fmt.Println("4 or 5 or 6")
	//}

	//numbers := [...]int8{0, 1, 2, 3, 4, 5, 6}
	//switch numbers[4] {
	//case 0, 1, 2:
	//	fmt.Println("0 or 1 or 2")
	//case 2, 3, 4:
	//	fmt.Println("2 or 3 or 4")
	//case 4, 5, 6:
	//	fmt.Println("4 or 5 or 6")
	//}

//	.\main1.go:38:7: duplicate case 2 in switch
//	previous case at .\main1.go:36:13
//	.\main1.go:40:7: duplicate case 4 in switch
//	previous case at .\main1.go:38:13

	//numbers := [...]int8{0, 1, 2, 3, 4, 5, 6}
	//switch numbers[4] {
	//case numbers[0], numbers[1], numbers[2]:
	//	fmt.Println("0 or 1 or 2")
	//case numbers[2], numbers[3], numbers[4]:
	//	fmt.Println("2 or 3 or 4")
	//case numbers[4], numbers[5], numbers[6]:
	//	fmt.Println("4 or 5 or 6")
	//}

	//numbers := interface{}(byte(127))
	//switch t := numbers.(type) {
	//case uint8, uint16:
	//	fmt.Println("uint8 or uint16")
	//case byte:
	//	fmt.Println("byte")
	//default:
	//	fmt.Println("unsupported type: %T", t)
	//}

//	.\main1.go:63:2: duplicate case byte in type switch
//		previous case at .\main1.go:61:2

}
