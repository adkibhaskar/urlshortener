package shortkey

import "math/rand"



var Urls1=make(map[string]int)

func GenerateShortKey()string{

	const letterBytes="abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	const keyLength=6

	b:=make([]byte,keyLength)

	for i:=range b{

		b[i]=letterBytes[rand.Intn(len(letterBytes))]
	}

	return string(b)
}