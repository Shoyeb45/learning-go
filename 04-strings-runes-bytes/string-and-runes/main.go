package main
import "fmt";

// strings in go is nothing but the array of uint8, it uses utf-8 charset but dynamically expands to 
// other characters than ascii, like gamma in greek or some hindi alphabet



func main() {

	myString := "γ i";

	fmt.Println("the length is: ", len(myString)); // len gives bytes not actual character length
	
	// if we try to iterate, pay attention in index
	for idx, val := range myString {
		fmt.Printf("myString[%v] = %d\n", idx, val);
	}

	var indexed = myString[0];
	// type
	fmt.Printf("%v %T\n", indexed, indexed);

	// strings are immutable
	// myString[0] = '1'; // not allowed

	// runes:

	myRune := 'a';

	fmt.Printf("%v, %T", myRune, myRune); // runes are stored as int32

	// convert string to runes
	var myStr = []rune("rësũme");

	fmt.Printf("The len is %v", len(myStr));
}