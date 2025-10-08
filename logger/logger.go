package logger

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

func logFileName() string {
	return TimeNow() + ".json"
}

func createFile() (*os.File, error) {
	file, err := os.Create(logFileName())
	if err != nil {
		fmt.Printf("Ошибка при создании файла: %v", err)
		return nil, err
	}
	return file, nil
}

func checkFile(nameFile string) bool {
	_, err := os.Stat(nameFile)
	return err == nil
}

func CreateLogFile() (*os.File, error) {
	nameFile := logFileName()
	if !checkFile(nameFile) {
		file, err := createFile()
		if err != nil {
			fmt.Printf("Ошибка при создании файла: %v", err)
			return nil, err
		}
		return file, nil
	}
	return os.OpenFile(nameFile, os.O_APPEND|os.O_WRONLY, 0600)
}

func Logging(request []byte, mutex *sync.Mutex) error {
	mutex.Lock()
	file, err := CreateLogFile()
	if err != nil {
		return err
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	if _, err := writer.Write(request); err != nil {
		fmt.Printf("Ошибка при записи в файл: %v", err)
		return writer.Flush()
	}
	mutex.Unlock()
	return writer.Flush()
}
