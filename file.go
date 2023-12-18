package main

import (
	"bufio"
	"io"
	"os"
)

func createNewfile(name string, massage string) error {
	file, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	file.WriteString(massage)
	return nil
}

func readFile(name string) (string, error) {
	file, err := os.OpenFile(name, os.O_RDONLY, 0666)
	if err != nil {
		return "", err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	var massage string
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		massage += string(line) + "\n"
	}
	return massage, nil
}

func addToFile(name string, massage string) error {
	file, err := os.OpenFile(name, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	file.WriteString(massage)
	return nil
}

func main() {
	// createNewfile("sample.log", "this is sample log") // add newfile

	// result, _ := readFile("sample.log")
	// fmt.Println(result) // readfile

	addToFile("sample.log", "\nthis is add message")

}
