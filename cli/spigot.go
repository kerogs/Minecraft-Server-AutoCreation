package cli

import (
	"fmt"
    "net/http"
	"os"
	"path/filepath"
	"io"


    "github.com/PuerkitoBio/goquery"
)

func Spigot(spigotUrl string) ([]string, error) {
    // Faites la requête HTTP pour obtenir le contenu de la page
    res, err := http.Get(spigotUrl)
    if err != nil {
        return nil, err
    }
    defer res.Body.Close()
    
    if res.StatusCode != 200 {
        return nil, fmt.Errorf("status error: %d %s", res.StatusCode, res.Status)
    }

    // Chargez le document HTML
    doc, err := goquery.NewDocumentFromReader(res.Body)
    if err != nil {
        return nil, err
    }

    // Créez un tableau pour stocker les valeurs des balises h2
    var versionList []string

    // Utilisez une sélection CSS pour trouver toutes les balises h2
    doc.Find("h2").Each(func(i int, s *goquery.Selection) {
        version := s.Text()
        versionList = append(versionList, version)
    })

    // Retourne le tableau des versions
    return versionList, nil
}

func SpigotDownload(spigotUrlJar string) error {
    // Effectuer la requête HTTP pour télécharger le fichier
    resp, err := http.Get(spigotUrlJar)
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    // Vérifier le code de statut HTTP
    if resp.StatusCode != http.StatusOK {
        return fmt.Errorf("statut HTTP incorrect: %s", resp.Status)
    }

    // Extraire le nom du fichier du chemin d'accès URL
    filename := filepath.Base(spigotUrlJar)

    // Créer un fichier local pour enregistrer le contenu téléchargé
    out, err := os.Create(filename)
    if err != nil {
        return err
    }
    defer out.Close()

    // Copier le contenu du corps de la réponse HTTP vers le fichier local
    _, err = io.Copy(out, resp.Body)
    if err != nil {
        return err
    }

    return nil
}