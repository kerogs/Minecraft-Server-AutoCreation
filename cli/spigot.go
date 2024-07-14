package cli

import (
	"fmt"
    "net/http"
	"os"
	"path/filepath"
	"io"


    "github.com/PuerkitoBio/goquery"
)

// Spigot makes an HTTP GET request to the given spigotUrl and retrieves the list of versions
// from the response. It returns a slice of strings representing the versions and an error
// if any occurred.
//
// Parameters:
// - spigotUrl: a string representing the URL to make the GET request to.
//
// Returns:
// - []string: a slice of strings representing the versions retrieved from the response.
// - error: an error if any occurred during the HTTP request or parsing the response.
func Spigot(spigotUrl string) ([]string, error) {
    res, err := http.Get(spigotUrl)
    if err != nil {
        return nil, err
    }
    defer res.Body.Close()
    
    if res.StatusCode != 200 {
        return nil, fmt.Errorf("status error: %d %s", res.StatusCode, res.Status)
    }

    doc, err := goquery.NewDocumentFromReader(res.Body)
    if err != nil {
        return nil, err
    }

    var versionList []string

    doc.Find("h2").Each(func(i int, s *goquery.Selection) {
        version := s.Text()
        versionList = append(versionList, version)
    })

    return versionList, nil
}

// SpigotDownload downloads a file from the provided URL and saves it locally.
//
// It takes the spigotUrlJar string as input and returns an error if any.
func SpigotDownload(spigotUrlJar string) error {
    resp, err := http.Get(spigotUrlJar)
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return fmt.Errorf("statut HTTP incorrect: %s", resp.Status)
    }

    filename := filepath.Base(spigotUrlJar)

    out, err := os.Create(filename)
    if err != nil {
        return err
    }
    defer out.Close()

    _, err = io.Copy(out, resp.Body)
    if err != nil {
        return err
    }

    return nil
}