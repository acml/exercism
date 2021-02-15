package cipher

// Cipher is an encryption interface
type Cipher interface {
	Encode(string) string
	Decode(string) string
}
