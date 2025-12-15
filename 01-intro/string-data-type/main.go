package main;
import "fmt";
import "unicode/utf8";


func main() {
	// we can define string using string

	var myStr string = "My first string in go\n";

	fmt.Println(myStr);

	// getting length of the string
	// note
	fmt.Println(len(myStr)); // This does not give length, it gives the bytes in the binary form

	// for actual length, we can use rune
	// import the utf-8 package
	fmt.Printf(`Length of the string "%s" is "%d`, myStr, utf8.RuneCountInString(myStr));

	var thisIsRune rune = 'o';
	fmt.Println(thisIsRune);

}
