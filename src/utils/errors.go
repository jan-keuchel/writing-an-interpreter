package utils

import "fmt"


func Error(line int, message string) {
	fmt.Printf("[Line %d] Error: %s\n", line, message)
}
