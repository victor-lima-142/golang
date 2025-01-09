package files

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"second-steps/requests"
	"strings"
)

func ReadingFileAndFetchUrl() {
	read()
}

func getWorkingDir() string {
	workingDir, err := os.Getwd()
	if err != nil {
		println("Error:", err.Error())
		os.Exit(1)
	}
	return filepath.Join(workingDir, "wfiles")
}

func read() {
	workingDir := getWorkingDir()
	filePath := filepath.Join(workingDir, "sites.txt")
	file, err2 := os.Open(filePath)
	if err2 != nil {
		println("Error:", err2.Error())
		os.Exit(1)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	fetching(reader)
}

func fetching(reader *bufio.Reader) {
	sites := []string{}
	for {
		line, err := reader.ReadString('\n')
		sites = append(sites, strings.TrimSpace(line))
		if err == io.EOF {
			break
		}
	}
	file := getFileToEdit()
	for i, site := range sites {
		httpRes, err2 := requests.RawGetRequestToUrl(site)
		writeResult(httpRes, i, site, file)
		if err2 != nil {
			println("Error:", err2.Error())
			os.Exit(1)
		}
	}
	defer file.Close()
}

func getFileToEdit() *os.File {
	workingDir := getWorkingDir()
	filePath := filepath.Join(workingDir, "result.txt")
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		println("Error:", err.Error())
		os.Exit(1)
	}
	return file
}

func writeResult(response *http.Response, index int, url string, file *os.File) {
	var text string = "Site "
	text += fmt.Sprint(index)
	text += " | Url "
	text += url
	text += " Result:"
	text += "\nStatus: "
	text += response.Status
	text += "\nStatus Code: "
	text += fmt.Sprint(response.StatusCode)
	text += "\n\n"
	println(text)
	if response.StatusCode == 200 {
		text += "\nRequest was successful\n\n"
		_, err := file.WriteString(text)
		if err != nil {
			println("Error:", err.Error())
			os.Exit(1)
		}
		return
	}
	text += "\nRequest was not successful\n\n"
	_, err := file.WriteString(text)
	if err != nil {
		println("Error:", err.Error())
		os.Exit(1)
	}
}
