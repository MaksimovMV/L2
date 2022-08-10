package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"
)

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type listOfUrls struct {
	u  map[string]string
	mu sync.Mutex
}

var (
	host   string
	scheme string
	wg     sync.WaitGroup
	urls   listOfUrls
)

// вызывается для каждого найденного элемента html
func processElement(index int, element *goquery.Selection) {
	// проверяем есть ли аттрибут href
	href, exists := element.Attr("href")
	if exists {
		if strings.HasPrefix(href, "/") {
			wg.Add(1)
			go download(scheme + "://" + host + href)
		} else {
			parsedUrl, err := url.Parse(href)
			if err != nil {
				return
			}
			if parsedUrl.Host == host {
				wg.Add(1)
				go download(href)
			}
		}
	}
}

// скачивание html - страниц со всеми ссылками на тот же домен
func download(complexUrl string) {
	defer wg.Done()
	urls.mu.Lock()
	if _, ok := urls.u[complexUrl]; ok {
		urls.mu.Unlock()
		return
	} else {
		urls.u[complexUrl] = complexUrl
		urls.mu.Unlock()
	}
	// Make request
	response, err := http.Get(complexUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	parsedUrl, err := url.Parse(complexUrl)
	if err != nil {
		log.Fatal(err)
	}

	if parsedUrl.Host != host {
		return
	}

	var fileName string

	path := parsedUrl.Path
	fmt.Println(path)
	if path == "" || path == "/" {
		fileName = host
	} else {
		strings.TrimSuffix(path, "/")
		arr := strings.Split(path, "/")
		fileName = arr[len(arr)-1]
	}
	if !strings.HasSuffix(fileName, ".html") {
		fileName += ".html"
	}
	fmt.Println(fileName)

	// Create output file
	outFile, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()

	// Copy data from HTTP response to file
	_, err = io.Copy(outFile, response.Body)
	if err != nil {
		log.Fatal(err)
	}

	response, err = http.Get(complexUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal("Error loading HTTP response body. ", err)
	}

	// Find all links and process them with the function
	// defined earlier
	document.Find("a").Each(processElement)
}

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("need just a url")
		return
	}
	complexUrl := args[1]
	parsedUrl, err := url.Parse(complexUrl)
	if err != nil {
		log.Fatal(err)
	}
	host = parsedUrl.Host
	scheme = parsedUrl.Scheme
	wg.Add(1)
	urls = listOfUrls{
		u:  make(map[string]string),
		mu: sync.Mutex{},
	}
	download(complexUrl)

	wg.Wait()
}
