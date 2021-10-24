package jsn

type Json struct {
	data   []byte
	cursor int
	cnt    int
}

func NewJson(data []byte) *Json {
	return &Json{
		data:   data,
		cursor: 0,
		cnt:    len(data),
	}
}

func (self *Json) Get() (name string, value string) {
	var cur int
	var e byte

	// name
	flag := int8(0)
	for self.cursor < self.cnt {
		e = self.data[self.cursor]
		self.cursor++
		switch e {
		case 123: // {
		case 125: // }
		case 34: // "
			if flag == 34 {
				goto stepName
			}
			if flag == 0 {
				cur = self.cursor
			}
			flag = 34
		case 92: // \
			self.cursor++
		case 91: // [
		case 93: // ]
		}
	}
stepName:
	name = string(self.data[cur : self.cursor-1])

	// value
	flag = 0
	for self.cursor < self.cnt {
		e = self.data[self.cursor]
		self.cursor++
		switch e {
		case 123: // {
		case 125: // }
		case 34: // "
			if flag == 34 {
				self.cursor++ // если значение в кавычках перепрыгиваем следующую запятую
				goto stepValue
			}
			if flag == 0 {
				cur = self.cursor
			}
			flag = 34
		case 58: // :
			if flag != 34 {
				cur = self.cursor
			}
		case 44: // ,
			if flag != 34 {
				goto stepValue
			}
		case 92: // \
			self.cursor++
		case 91: // [
		case 93: // ]
		}
	}
stepValue:
	value = string(self.data[cur : self.cursor-1])

	return name, value
}
