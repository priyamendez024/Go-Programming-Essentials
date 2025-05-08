// Chapter 12: Networking Basics with net and net/http
package main

import (
    "bufio"
    "crypto/tls"
    "fmt"
    "io"
    "log"
    "net"
    "net/http"
)

func tcpServer() {
    ln, _ := net.Listen("tcp", ":8080")
    for {
        conn, _ := ln.Accept()
        go func(c net.Conn) {
            io.Copy(c, c)
            c.Close()
        }(conn)
    }
}

func tcpClient() {
    conn, _ := net.Dial("tcp", "localhost:8080")
    defer conn.Close()
    fmt.Fprintln(conn, "ping")
    resp, _ := bufio.NewReader(conn).ReadString('\n')
    fmt.Println("Response:", resp)
}

func httpServer() {
    http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Hello, HTTP!")
    })
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func httpClient() {
    resp, _ := http.Get("http://localhost:8080/hello")
    defer resp.Body.Close()
    body, _ := io.ReadAll(resp.Body)
    fmt.Println(string(body))
}

func httpsServer() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Secure Hello")
    })
    cfg := &tls.Config{MinVersion: tls.VersionTLS12}
    server := &http.Server{Addr: ":8443", Handler: mux, TLSConfig: cfg}
    log.Fatal(server.ListenAndServeTLS("server.crt", "server.key"))
}

func main() {
    go tcpServer()
    go httpServer()
    // choose one
}
