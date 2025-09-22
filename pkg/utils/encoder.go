package utils

func Encode(id int64) string {
	if id == 0 {
		return string(Charset[0])
	}

	result := make([]byte, 0)
	for id > 0 {
		remainder := id % Base
		result = append([]byte{Charset[remainder]}, result...)
		id = id / Base
	}
	return string(result)
}
