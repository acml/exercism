// Package variablelengthquantity implements variable length quantity encoding
// and decoding.
package variablelengthquantity

// DecodeVarint implements variable length quantity decoding.
func DecodeVarint(input []byte) ([]uint32, error) {
	res := []uint32{}
	var t uint32
	for _, b := range input {
		t |= uint32(b & 0x7F)
		if b&0x80 == 0 {
			res = append(res, t)
			t = 0
		}
		t <<= 7
	}
	return res, nil
}

// EncodeVarint implements variable length quantity encoding.
func EncodeVarint(input []uint32) []byte {
	res := []byte{}
	for _, n := range input {
		var cont uint32
		var varint []byte
		for {
			varint = append([]byte{byte(cont | (n & 0x7F))}, varint...)
			n >>= 7
			if n == 0 {
				res = append(res, varint...)
				break
			}
			cont = 0x80
		}
	}
	return res
}
