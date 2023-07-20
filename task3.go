package main

import "fmt"

func maskLinks(s string) string {
	b := []byte(s)

	for i := 0; i < len(b)-6; i++ {
		if string(b[i:i+7]) == "http://" {
			for j := i + 7; j < len(b); j++ {
				if b[j] == ' ' {
					break
				}
				b[j] = '*'
			}
		}
	}

	return string(b)
}

func main() {
	s := "Here's my spammy page: http://hehefouls.netHAHAHA see you."
	fmt.Println(maskLinks(s))
}
