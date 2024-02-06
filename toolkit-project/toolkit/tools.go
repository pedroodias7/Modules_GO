package toolkit

import "crypto/rand"


const randomStringSource = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTU_+"
// Tools is the type used to instantiate this module. Any variable of this type will have acees
//to all methods with the reciever
type Tools struct {}


// RandomString return a string a random string charescters of lengh n, using randomStringSource
// as main source
func (t *Tools) RandomString(n int) string {
	s, r := make([]rune, n), []rune(randomStringSource)

	for i := range s {
		p, _ := rand.Prime(rand.Reader, len(r))


		x, y := p.Uint64(), uint64(len(r))
		s[i]= r[x%y]
	}

	return string(s)
}