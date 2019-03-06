package main

import (
     "fmt"
    // "os" //  is for accesing command line arguments
     "net/http"// is for making HTTP requests
     "time"
)


func MakeRequest(url string) {
  start := time.Now()
  resp, _ := http.Get(url)

  secs := time.Since(start).Seconds()

  fmt.Sprintf("%.2f elapsed with domain: %s with response status %d", secs, url,resp.StatusCode)
}


func main()  {
 start := time.Now()
 urls := []string{"http://localhost:8010/api/slow/test/22","http://localhost:8010/api/fast/test/1","http://localhost:8010/api/fast/test/2","http://localhost:8010/api/fast/test/3","http://localhost:8010/api/slow/test/66"}
  for _,url := range urls{
       MakeRequest(url)
  }

  fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}