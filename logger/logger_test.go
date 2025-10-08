package logger

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"testing"
	"time"
)

func TestCreateFile(t *testing.T) {
	file, err := createFile()
	if err != nil {
		t.Errorf("Ошибка при создании файла: %v", err)
	}
	if _, err := file.Stat(); err != nil {
		t.Errorf("Ошибка при получении информации о файле: %v", err)
	}
	if file.Close() != nil {
		t.Errorf("Ошибка при закрытии файла: %v", err)
	}
	t.Cleanup(func() {
		file.Close()
		os.Remove(file.Name())
	})
}

func TestCheckFile(t *testing.T) {
	file, _ := createFile()
	if !checkFile(file.Name()) {
		t.Errorf("Файл не создан")
	}

	if checkFile("incorrect_file") {
		t.Errorf("Файл не существует")
	}
	t.Cleanup(func() {
		file.Close()
		os.Remove(file.Name())
	})
}

func TestCreateLogFile(t *testing.T) {
	file1, err := CreateLogFile()
	if err != nil {
		t.Errorf("Ошибка при создании файла: %v", err)
	}
	if _, err := file1.Stat(); err != nil {
		t.Errorf("Ошибка при получении информации о файле: %v", err)
	}
	time.Sleep(1 * time.Second)
	// Проверяем, что файл создан и не будет создан новый файл (с учетом времени)
	file2, err := CreateLogFile()
	if err != nil {
		t.Errorf("Ошибка при создании файла: %v", err)
	}
	if _, err := file2.Stat(); err != nil {
		t.Errorf("Ошибка при получении информации о файле: %v", err)
	}
	if file1.Name() != file2.Name() {
		t.Errorf("Файлы не совпадают")
	}

	t.Cleanup(func() {
		file1.Close()
		file2.Close()
		os.Remove(file1.Name())
		os.Remove(file2.Name())
	})
}

func TestLogging(t *testing.T) {
	type fakeRequest struct {
		RequestURL    string `json:"requestURL"`
		RequestMethod string `json:"requestMethod"`
		Message       string `json:"message"`
		Timestamp     string `json:"timestamp"`
	}

	var fakeRequests = []fakeRequest{
		{
			RequestURL:    "https://example.com",
			RequestMethod: "GET",
			Message:       "test message",
			Timestamp:     "2021-01-01 12:00:00",
		},
		{
			RequestURL:    "https://example1.com",
			RequestMethod: "POST",
			Message:       "test message1",
			Timestamp:     "2021-01-01 12:00:01",
		},
		{
			RequestURL:    "https://example2.com",
			RequestMethod: "PUT",
			Message:       "test message2",
			Timestamp:     "2021-01-01 12:00:02",
		},
	}

	for index, request := range fakeRequests {
		t.Run(fmt.Sprintf("Запись запроса %d", index), func(t *testing.T) {
			var mutex sync.Mutex
			jsonRequest, err := json.Marshal(request)
			if err != nil {
				t.Errorf("Ошибка при маршалинге запроса: %v", err)
			}
			err = Logging(jsonRequest, &mutex)
			if err != nil {
				t.Errorf("Ошибка при записи запроса: %v", err)
			}
			time.Sleep(1 * time.Second)
		})
	}

}
