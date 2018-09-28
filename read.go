package main


import (
	"fmt"
	"io/ioutil"
	"strings"
	"unicode"
	"github.com/ahamidi/go-mapreduce"
)


func main(){
	contents, err := ioutil.ReadFile("/media/smita/Smita/readfile")
	fmt.Print(string(contents))
	if err != nil {
		fmt.Println("Failed to read file ")

}

	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}

	//func FieldsFunc(s string, f func(rune) bool) []string

	words := strings.FieldsFunc(string(contents), f)
	fmt.Printf("Fields are: %q", words);

	kvs := make([]mapreduce.KeyValue, 0)
	for _, w := range words {
		kvs = append(kvs, mapreduce.KeyValue{Key: w, Value: "1"})
		}
		fmt.Println("Output is %s", kvs)
}


// echo "hello" > /media/smita/Smita/read
//go run read.go


//o/p without echo 
//smita
//is = 2
//a 
//good 5
//girl




/*smita
is = 2
a 
good 5
girl
"  foo1;bar2,baz3...",
Fields are: ["smita" "is" "2" "a" "good" "5" "girl" "foo1" "bar2" "baz3"]*/

