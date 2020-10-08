package phasmalib

const password = "CHANGE ME TO YOUR OWN RANDOM STRING"
var passwordBytes = []byte(password)

func XorData(input []byte) (output []byte){
	output = make([]byte,0)
	for idx, char := range input {
		xorByte := char ^ passwordBytes[idx % len(passwordBytes)]
		output = append(output, xorByte)
	}
	return
}