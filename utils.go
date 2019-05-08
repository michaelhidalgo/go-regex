package regex

import(
	"math/rand"
)

// RandomString util function to generate a fixed-lengh function
// taken from https://medium.com/@kpbird/golang-generate-fixed-size-random-string-dd6dbd5e63c0
func RandomString(len int) string {

   bytes := make([]byte, len)
   for i := 0; i < len; i++ {
		bytes[i] = byte(65 + rand.Intn(25))
	}
    return string(bytes)
}

