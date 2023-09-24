package wait

import "fmt"

// Log is used to log with prefix wait-for-it
func Log(message string) {
	fmt.Println("wait-for-it: " + message)
}
