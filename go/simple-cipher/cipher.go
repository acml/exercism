package cipher

// Cipher is an encyrption interface
type Cipher interface {
	Encode(string) string
	Decode(string) string
}
