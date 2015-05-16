package io

type Keyboard int

func (k *Keyboard) GetChar() (int, error) {
	return getChar()
}

func (k *Keyboard) Listen(handler func(byte), q byte) {
	for {
		ch, err := k.GetChar()
		if err != nil || ch < 0 {
			break
		}
		if byte(ch) == q {
			break
		}
		handler(byte(ch))
	}
}

func GetKeyboard() (*Keyboard, func()) {
	setRaw()
	return new(Keyboard), func() {
		setCooked()
	}
}
