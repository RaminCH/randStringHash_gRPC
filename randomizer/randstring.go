package randomizer

import (
	"math/rand"
	// pb "github.com/RaminCH_self/Gransoft/v2/randomizer/proto/consignment"
)

const (
	address = "localhost:5700"
)

//RandomString function ...
func RandomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

//ListOfStrings func ...
func ListOfStrings(n string, i int) []string {
	c := make([]string, i)
	for v := range n {
		c = append(c, RandomString(i))
		_ = v
	}
	return c

}


// func main() {
// 	str := ListOfStrings(RandomString(4), 4)
// 	fmt.Println(str)
// }
