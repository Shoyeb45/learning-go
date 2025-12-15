package main

import (
	"fmt"
	"strings"
)

func main()  {

	var str = []string{"Shoyeb", " ", "is", " ", "a", " ", "programmer."};
	var stringBuilder strings.Builder;

	for i := range str {
		stringBuilder.WriteString(str[i]);
	}
	var catStr = stringBuilder.String();

	fmt.Println(catStr);
}