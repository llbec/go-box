package log

import "fmt"

//Recorder defines a log recorder
type Recorder struct{}

func (r *Recorder) write(title, data string) {
	fmt.Printf("%v:\t%v\n", title, data)
}

//Info info log
func (r *Recorder) Info(s string) {
	r.write("INFO", s)
}

//Info info log
func (r *Recorder) Error(s string, e error) error {
	r.write("ERROR", s)
	return e
}
