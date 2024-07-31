package main

import (
	"fmt"
	"regexp"
	"strings"
)


func wordCount(s string) map[string]int {
	re := regexp.MustCompile(`\p{P}`)
	s=re.ReplaceAllString(s,"")
	dictionary := make(map[string]int)
	words := strings.Split(s, " ")
	for _, word := range words {
		if word ==""{
			continue
		}
		dictionary[word]++
	}
	return dictionary
}

func palindrome(s string) bool {
	s = strings.ToLower(s)
	s = strings.ReplaceAll(s, " ", "")
	re := regexp.MustCompile(``)
	s = re.ReplaceAllString(s, "")
	j:=len(s)-1
	i:=0
	for i<j{
		if s[i]!=s[j]{
			return false
		}
		i++
		j--
	}
	return true

}
func main() {
	
	fmt.Println(wordCount("Hello,??/ : ( ) * _ World! Hello World! Hello, World! Hello, Wo*/-rld! Hello, Wor)%@ld?"))
	fmt.Println(palindrome("123!321asda"))

}
