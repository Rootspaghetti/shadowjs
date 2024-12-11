package main

import (
        "bufio"
        "flag"
        "fmt"
        "io/ioutil"
        "net/http"
        "net/url"
        "os"
        "regexp"
        "strings"
        "sync"
)

var visited = make(map[string]bool)
var jsFileSet = make(map[string]bool) // URL'lere göre kontrol

var mutex sync.Mutex

func showBanner() {
        banner := `
____  _               _                  _ ____
/ ___|| |__   __ _  __| | _____      __  | / ___|
\___ \| '_ \ / _` + "`" + ` |/ _` + "`" + ` |/ _ \ \ /\ / /  | \___ \
 ___) | | | | (_| | (_| | (_) \ V  V / |_| |___) |
|____/|_| |_|\__,_|\__,_|\___/ \_/\_/ \___/|____/
            JavaScript Discovery Tool

by:Root@Spaghetti
        `
        fmt.Println(banner)
}

func fetchHTML(targetURL string) (string, int) {
        resp, err := http.Get(targetURL)
        if err != nil {
                fmt.Printf("HTTP request failed (%s): %s\n", targetURL, err)
                return "", 0
        }
        defer resp.Body.Close()

        body, err := ioutil.ReadAll(resp.Body)
        if err != nil {
                fmt.Printf("HTML read error (%s): %s\n", targetURL, err)
                return "", resp.StatusCode
        }

        return string(body), resp.StatusCode
}

func findJSFiles(baseURL, html string) []string {
        var jsFiles []string
        re := regexp.MustCompile(`<script[^>]+src="([^"]+)"`)
        matches := re.FindAllStringSubmatch(html, -1)

        for _, match := range matches {
                if len(match) > 1 {
                        jsURL := match[1]
                        fullURL := toAbsoluteURL(baseURL, jsURL)
                        jsFiles = append(jsFiles, fullURL)
                }
        }
        return jsFiles
}

func findLinks(baseURL, html string) []string {
        var links []string
        re := regexp.MustCompile(`<a[^>]+href="([^"]+)"`)
        matches := re.FindAllStringSubmatch(html, -1)

        for _, match := range matches {
                if len(match) > 1 {
                        link := match[1]
                        fullURL := toAbsoluteURL(baseURL, link)
                        links = append(links, fullURL)
                }
        }
        return links
}

func toAbsoluteURL(baseURL, relativeURL string) string {
        parsedBaseURL, err := url.Parse(baseURL)
        if err != nil {
                fmt.Println("Base URL parse error:", err)
                return relativeURL
        }

        parsedRelativeURL, err := url.Parse(relativeURL)
        if err != nil {
                fmt.Println("Relative URL parse error:", err)
                return relativeURL
        }

        return parsedBaseURL.ResolveReference(parsedRelativeURL).String()
}

func crawl(baseURL, currentURL string, wg *sync.WaitGroup, results *[]string) {
        defer wg.Done()

        mutex.Lock()
        if visited[currentURL] {
                mutex.Unlock()
                return
        }
        visited[currentURL] = true
        mutex.Unlock()

        fmt.Printf("Scanning in progress: %s\n", currentURL)

        html, statusCode := fetchHTML(currentURL)
        if statusCode != 0 {
                fmt.Printf("HTTP Status Code: %d (%s)\n", statusCode, currentURL)
        }
        if html == "" {
                return
        }

        // JavaScript dosyalarını bul
        jsFiles := findJSFiles(baseURL, html)
        mutex.Lock()
        for _, file := range jsFiles {
                if !jsFileSet[file] { // Eğer URL daha önce eklenmediyse
                        jsFileSet[file] = true
                        *results = append(*results, file)
                        fmt.Println("Found JS:", file)
                }
        }
        mutex.Unlock()

        // Sayfa içindeki linkleri bul
        links := findLinks(baseURL, html)
        var innerWG sync.WaitGroup
        for _, link := range links {
                if strings.HasPrefix(link, baseURL) {
                        innerWG.Add(1)
                        go crawl(baseURL, link, &innerWG, results)
                }
        }
        innerWG.Wait()
}

func readSubdomains(filename string) ([]string, error) {
        file, err := os.Open(filename)
        if err != nil {
                return nil, err
        }
        defer file.Close()

        var subdomains []string
        scanner := bufio.NewScanner(file)
        for scanner.Scan() {
                subdomains = append(subdomains, scanner.Text())
        }

        if err := scanner.Err(); err != nil {
                return nil, err
        }

        return subdomains, nil
}

func saveResults(filename string, results []string) error {
        file, err := os.Create(filename)
        if err != nil {
                return err
        }
        defer file.Close()

        for _, result := range results {
                if strings.HasSuffix(result, ".js") {
                        _, err := file.WriteString(result + "\n")
                        if err != nil {
                                return err
                        }
                }
        }

        return nil
}

func main() {
        showBanner()
        domain := flag.String("d", "", "Target domain  (example: https://example.com)")
        subdomainsFile := flag.String("subs", "", "Subdomain list")
        outputFile := flag.String("o", "results.txt", "file to write output results")
        flag.Parse()

        if *domain == "" && *subdomainsFile == "" {
                fmt.Println("Usage: -d <domain> or -subs <subdomain list>")
                flag.Usage()
                return
        }

        var results []string
        var wg sync.WaitGroup

        if *domain != "" {
                fmt.Printf("Scanning target domain: %s\n", *domain)
                wg.Add(1)
                go crawl(*domain, *domain, &wg, &results)
        }

        if *subdomainsFile != "" {
                subdomains, err := readSubdomains(*subdomainsFile)
                if err != nil {
                        fmt.Println("Could not read subdomain list:", err)
                        return
                }

                for _, subdomain := range subdomains {
                        wg.Add(1)
                        go crawl(subdomain, subdomain, &wg, &results)
                }
        }

        wg.Wait()

        err := saveResults(*outputFile, results)
        if err != nil {
                fmt.Println("Results could not be recorded:", err)
        } else {
                fmt.Printf("\nResults recorded: %s\n", *outputFile)
        }
}
