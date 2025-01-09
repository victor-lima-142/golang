package files

import (
	"bufio"
	"io"
	"os"
	"path/filepath"
	"second-steps/requests"
	"strings"
)

func ReadingFileAndFetchUrl() {
	read()
}

func read() {
	workingDir, err := os.Getwd()
	if err != nil {
		println("Error:", err.Error())
		os.Exit(1)
	}

	filePath := filepath.Join(workingDir, "rfiles", "sites.txt")
	file, err2 := os.Open(filePath)
	if err2 != nil {
		println("Error:", err.Error())
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
	for i, site := range sites {
		println("\n\nSite", i, "| Url", site, "Result:\n")
		requests.GetRequestToUrl(site)
	}
}
