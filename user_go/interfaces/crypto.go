package interfaces

type CryptoLib interface {
	Encrypt(string) (string, error)
	Decrypt(string) (string, error)

	B64Decode(string) (string, error)
	B64Encode(string) string
}
