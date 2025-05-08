// Chapter 11: File I/O, JSON, XML & Data Encoding
package main

import (
    "bufio"
    "encoding/gob"
    "encoding/json"
    "encoding/xml"
    "fmt"
    "io"
    "os"
)

func readFile(path string) error {
    f, err := os.Open(path)
    if err != nil {
        return err
    }
    defer f.Close()
    reader := bufio.NewReader(f)
    for {
        line, err := reader.ReadString('\n')
        fmt.Print(line)
        if err == io.EOF {
            break
        }
        if err != nil {
            return err
        }
    }
    return nil
}

func writeFile(path, content string) error {
    f, err := os.Create(path)
    if err != nil {
        return err
    }
    defer f.Close()
    _, err = f.WriteString(content)
    return err
}

type User struct {
    ID    int    `json:"id" xml:"id"`
    Name  string `json:"name" xml:"name"`
    Email string `json:"email" xml:"email"`
}

func jsonExample() {
    u := User{ID: 1, Name: "Alice", Email: "alice@example.com"}
    data, _ := json.MarshalIndent(u, "", "  ")
    fmt.Println(string(data))
    var u2 User
    json.Unmarshal(data, &u2)
    fmt.Println(u2)
}

func xmlExample() {
    note := struct {
        To   string `xml:"to"`
        From string `xml:"from"`
        Body string `xml:"body"`
    }{"Bob", "Carol", "Hello XML!"}
    out, _ := xml.MarshalIndent(note, "", "  ")
    fmt.Println(string(out))
}

func gobExample() {
    type Record struct {
        Key   string
        Value int
    }
    file, _ := os.Create("data.gob")
    defer file.Close()
    enc := gob.NewEncoder(file)
    records := []Record{{"a", 1}, {"b", 2}}
    enc.Encode(records)
}

func main() {
    writeFile("example.txt", "Hello, Go I/O!\n")
    readFile("example.txt")
    jsonExample()
    xmlExample()
    gobExample()
}
