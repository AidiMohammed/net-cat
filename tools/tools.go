package tools

import (
	"fmt"
	"os"
)

const COLOR_RED = 31
const COLOR_GREEN = 32
const COLOR_YELLOW = 33
const COLOR_BLUE = 34

func ColorString (indexColor int,str string) string{
	return fmt.Sprintf("\033[%vm%v\033[0m",indexColor,str)
}

func WelcomMessage() (string,error) {
	file,err := os.ReadFile("welcome.txt")
	if err != nil {
		return "",err
	}

	return string(file),nil
	
}