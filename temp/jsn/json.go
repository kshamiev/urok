package jsn

type Json struct {
	data   []byte
	cursor int
}

func NewJson(data []byte) *Json {
	return &Json{
		data:   data,
		cursor: 0,
	}
}

func (self Json) Get() (name string, value string) {
	var cur int
	var flag int8
	// name
	for _, e := range self.data[self.cursor:] {
		self.cursor++
		switch e {
		case 123: // {
		case 125: // }
		case 34: // "
			if flag == 1 {
				// fmt.Println(self.cursor)
				goto stepName
			}
			if flag == 0 {
				cur = self.cursor
			}
			flag = 1
		case 58: // :
		case 44: // ,
		case 92: // \
			flag = -1
		case 91: // [
		case 93: // ]
		}
	}
stepName:
	name = string(self.data[cur : self.cursor-1])

	// value
	flag = 0
	for _, e := range self.data[self.cursor:] {
		self.cursor++
		switch e {
		case 123: // {
		case 125: // }
		case 34: // "
			if flag == 1 {
				goto stepValue
			}
			if flag == 0 {
				cur = self.cursor
			}
			flag = 1
		case 58: // :
		case 44: // ,
		case 92: // \
			flag = -1
		case 91: // [
		case 93: // ]
		}
	}
stepValue:
	value = string(self.data[cur : self.cursor-1])

	return name, value
}
