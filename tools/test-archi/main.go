package main

import (
    "fmt"
    "os"
    "path/filepath"
    "regexp"
)

func main() {
    regex := regexp.MustCompile(`\.dist$`)

    err := filepath.Walk("./tests/data", func(path string, info os.FileInfo, err error) error {
        // Ignorer les erreurs
        if err != nil {
            return nil
        }

        if info.Mode().IsRegular() {
            filename := info.Name()
            if regex.MatchString(filename) {
                fmt.Println("good" + " " + path + " ")
            }else{
                fmt.Println("bad" + " " + path + " ")
            }
        }

        return nil
    })

    if err != nil {
        fmt.Println(err)
    }
}
