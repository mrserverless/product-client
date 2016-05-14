package main

import (
        "github.com/parnurzeal/gorequest"
        "os"
        "fmt"
)

func main() {

        if (len(os.Args) <= 1) {
                fmt.Printf("invalid number of arguments: %d\n", len(os.Args))
                fmt.Println("USAGE: product-client [all|id]")
        } else {
                request := gorequest.New()
                cmd := os.Args[1]
                switch cmd {
                case "all":
                        request.Get("http://localhost:8080/products").EndBytes(printResponse)
                        return
                }
                request.Get("http://localhost:8080/products/" + cmd).EndBytes(printResponse)
        }

}

func printResponse(resp gorequest.Response, body string, errs []error){
        fmt.Printf("response: %d, error: %s\n", resp.StatusCode, errs)
        fmt.Println(body)
}