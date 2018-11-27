package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "abcdef"
	fmt.Println("Contains =", strings.Contains(text, "bc"))   // true
	fmt.Println("HasPrefix =", strings.HasPrefix(text, "ab")) // true
	fmt.Println("HasSuffix =", strings.HasSuffix(text, "ef")) // true
	fmt.Println("Index =", strings.Index(text, "cd"))         // 2

	names := []string{"Mark", "Tami", "Amanda", "Jeremy"}
	joined := strings.Join(names, " - ")
	fmt.Printf("joined = %+v\n", joined) // Mark - Tami - Amanda - Jeremy

	fmt.Println("Repeat =", strings.Repeat(text, 2))        // abcdefabcdef
	fmt.Println("Split =", strings.Split(text, "cd"))       // [ab ef]
	fmt.Println("ToUpper =", strings.ToUpper(text))         // ABCDEF
	fmt.Println("Trim =", strings.Trim("  foo bar  ", " ")) // "foo bar"
	// Also see TrimLeft and TrimRight.

	// The Builder type supports efficiently building dynamic strings.
	var b strings.Builder
	b.WriteString("Foo")
	b.WriteString("Bar")
	b.WriteString("Baz")
	// Other Write methods on the Builder include
	// Write to write a byte slice
	// WriteByte to write a single byte
	// WriteRune to write a single rune
	fmt.Println(b.String()) // FooBarBaz

	sentence := "Mark goes to the park."
	fmt.Println("Replace =", strings.Replace(sentence, "ar", "il", -1)) // Milk goes to the pilk.

	// The Replacer type provides a more powerful alternative to strings.Replace.
	// Create a Replace object with pairs of old/new strings.
	r := strings.NewReplacer("&", "&amp;", "'", "&apos;", "\"", "&quot;", "<", "&lt;", ">", "&gt;")
	// Call the Replace method once for each string to be processed.
	fmt.Println(r.Replace("It's \"true\" that 5 is '>' 4."))
	// It&apos;s &quot;true&quot; that 5 is &apos;&gt;&apos; 4.
}
