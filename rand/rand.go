package rand

// #cgo CFLAGS: -I/usr/local/ssl/include
// #cgo LDFLAGS: -L /usr/local/ssl/lib -lcrypto
import "C"
import (
	"errors"
)

/*
 * Read populates the buffer argument with a pseudo-random sequence of bytes
 * read from /dev/random (this requires a system with /dev/random).
 * Behavior is as for golang's crypto/rand.Read()
 */
func Read(buf []byte) (int, error) {
	l := cap(buf)
	r := RAND_bytes(buf, l)
	if r != 1 {
		return 0, errors.New("Error generating random byte sequence")
	}

	if len(buf) != l {
		return 0, errors.New("Random byte sequence has wrong length")
	}

	return len(buf), nil
}
