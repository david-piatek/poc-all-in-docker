package main

import (
    "fmt"
    "io/ioutil"
    "os"
    "path/filepath"
    "regexp"
)

func main() {
    // Définir la regex pour les fichiers à vérifier
    regex := regexp.MustCompile("^\w+\.\w+\.(ext|dist)$")

    // Parcourir le répertoire courant
    err := filepath.Walk("./tests/data", func(path string, info os.FileInfo, err error) error {
        // Ignorer les erreurs
        if err != nil {
            return nil
        }

        // Vérifier si le fichier est un fichier régulier
        if info.Mode().IsRegular() {
            // Extraire le nom du fichier
            filename := info.Name()

            // Vérifier si le nom du fichier ne comporte pas l'extension .dist
            if !regex.MatchString(filename) {
                // Afficher le chemin du fichier
                fmt.Println(path)
            }
        }

        return nil
    })

    if err != nil {
        fmt.Println(err)
    }
}
