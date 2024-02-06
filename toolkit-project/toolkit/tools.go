package toolkit

import (
	"crypto/rand"
	"net/http"
)


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



//uploadfile is a struct informactio to save information abou an upload file
type UploadedFile struct {
	newFileName string
	originalFileName string
	fileSize int64

}

func (t *Tools) UploadedFiles(r *http.Request, uploadDir string, rename ...bool) ([]*UploadedFile, error){

		rename := true
		if len(rename) > 0 {
			renameFile = rename[0]

		}


}