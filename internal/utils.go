package internal

import "fmt"

func RedText(text string) string {
	return fmt.Sprintf("%c[1;0;31m%s%c[0m", 0x1B, text, 0x1B)
}
