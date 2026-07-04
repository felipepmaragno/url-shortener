package codegen

const (
	alphabet                   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	base                       = uint64(len(alphabet))
	MinimumGeneratedCodeLength = 5
)

func EncodeBase62(id uint64) string {
	if id == 0 {
		return alphabet[:1]
	}

	var buf [11]byte
	i := len(buf)

	for id > 0 {
		i--
		buf[i] = alphabet[id%base]
		id /= base
	}

	return string(buf[i:])
}

func MinimumGeneratedID() uint64 {
	var id uint64 = 1
	for i := 1; i < MinimumGeneratedCodeLength; i++ {
		id *= base
	}
	return id
}
