package slice

import "fmt"

func SumOfLetters(inputStr string) int{

	alphabets := make([]string,0)

	alphabets = append(alphabets, "")
	for i := 'a'; i <= 'z'; i++ {
		char := fmt.Sprintf("%c", i)
		alphabets = append(alphabets, char)
		//println(fmt.Sprintf("%v", i))
	}

	for index, value := range inputStr{
		println(index)
		println(value)
		_ = fmt.Sprintf("%c", value)
	}
	return 0;
}