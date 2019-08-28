// jsonchan implements a streaming JSON encoder from an interface channel sink.
// By default it encodes the object stream as a JSON array.
package jsonchan

import (
	"encoding/json"
	"io"
)

// Stream encodes ac as a JSON array to w. Encoding/writing happens per-object.
// Call close() on ac to stop encoding.
func Stream(w io.Writer, ac chan interface{}) {
	w.Write([]byte("["))
	enc := json.NewEncoder(w)
	first := true
	for {
		select {
		case a, ok := <-ac:
			if ok {
				if !first {
					w.Write([]byte(","))
				}
				enc.Encode(a)
				first = false
			} else {
				w.Write([]byte("]"))
				return
			}
		}
	}
}
