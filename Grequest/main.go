package main

import (
	"fmt"
	"github.com/levigross/grequests"
)

func ResponsHandle(resp *grequests.Response, err error) {
	if err != nil {
		fmt.Println("Unable to make request: ", err)
	} else {
		fmt.Println(resp.String())
	}
}

func Basic() {
	resp, err := grequests.Get("http://httpbin.org/get?/", nil)
	// You can modify the request by passing an optional RequestOptions struct
	ResponsHandle(resp, err)
}

func RequestQuirks() {
	reqOption := &grequests.RequestOptions{
		Params: map[string]string{"Hello": "GoodBye"},
	}
	resp, err := grequests.Get("http:/httpbin.org/get?Hello=World", reqOption)
	ResponsHandle(resp, err)
}

func DownloadTextFile() {
	resp, err := grequests.Get("https://example-files.online-convert.com/document/txt/example.txt", nil)
	if err != nil {
		fmt.Println("Unable to get text file url", err)
	}
	if err = resp.DownloadToFile("randomFile"); err != nil {
		fmt.Println("Unable to download: ", err)
	}
}

func DownloadReuse() {
	resp, err := grequests.Get("https://example-files.online-convert.com/document/txt/example.txt", nil)
	if err != nil {
		fmt.Println("Unable to get text file url", err)
	}
	fmt.Println(resp.Bytes())
	fmt.Println(resp.String() == "file-string")
	if err = resp.DownloadToFile("randomFile"); err != nil {
		fmt.Println("Unable to download: ", err)
	}
	fmt.Println("Check if can be printed again\n", resp.Bytes())
	fmt.Println(resp.String())
}

func main() {
	Basic()
	RequestQuirks()
	DownloadTextFile()
	DownloadReuse()
}
