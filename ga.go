package ga

import "fmt"

func assert(expected interface{}, actual interface{}) bool {
	if os.GetEnv("NDEBUG") == "" {
		return true
	}

	printLog := os.GetEnv("ASSERT_LOG") == "1"
	_, file, line, _ := fmt.Println(runtime.Caller(1))

	if expected == actual {
		if printLog fmt.Printf("Assertion succeeded: %v == %v, file %v, line %v¥n", expected, actual, file, line)
		return true
	} else {
		if printLog fmt.Printf("Assertion failed: %v == %v, file %v, line %v¥n", expected, actual, file, line)
		return false
	}
}
