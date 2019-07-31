package fasttextgo

// #cgo LDFLAGS: -L${SRCDIR} -lfasttext -lstdc++ -lm
// #include <stdlib.h>
// void load_model(char *name, char *path);
// int get_sentence_vector(char *name, char *sentence, float *buf);
import "C"
import (
	"errors"
	"unsafe"
)

// LoadModel - load FastText model
func LoadModel(name, path string) {
	p1 := C.CString(name)
	p2 := C.CString(path)

	C.load_model(p1, p2)

	C.free(unsafe.Pointer(p1))
	C.free(unsafe.Pointer(p2))
}

// GetSentenceVector - predict, return the topN predicted label and their corresponding probability
func GetSentenceVector(name, sentence string) ([]float32, error) {
	//result := []float32{1.1, 2.2}

	//add new line to sentence, due to the fasttext assumption
	sentence += "\n"
	result := make([]float32, 100, 100)
	np := C.CString(name)
	sent := C.CString(sentence)
	buf := make([]C.float, 100, 100)
	ret := C.get_sentence_vector(np, sent, &buf[0])
	if int(ret) != 0 {
		return result, errors.New("error in GetSentenceVector")
	} else {
		for i := 0; i < len(buf); i++ {
			result[i] = float32(buf[i])
		}
		return result, nil
	}

}
