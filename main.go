package main

import (
     "fmt"
     "os" //  is for accesing command line arguments
     "net/http"// is for making HTTP requests
     "time"
)


func MakeRequest(url string, ch chan<-string) {
  start := time.Now()
  resp, _ := http.Get(url)

  secs := time.Since(start).Seconds()

  ch <- fmt.Sprintf("%.2f elapsed with domain: %s with response status %d", secs, url,resp.StatusCode)
}


func main()  {
 start := time.Now()
  ch := make(chan string)
  for _,url := range os.Args[1:]{
      go MakeRequest(url, ch)
  }

  for range os.Args[1:]{
    fmt.Println(<-ch)
  }
  fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}