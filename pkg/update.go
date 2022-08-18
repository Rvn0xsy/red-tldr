package pkg

import (
	"archive/zip"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"red-tldr/utils"
	"strings"
	"time"
)

const (
	userName = "Rvn0xsy"
	repoName = "red-tldr-db"
)

type githubTagUrlData struct {
	Name string `json:"name"`
	ZipBallUrl  string`json:"zipball_url"`
}


func GetLatestReleaseFromGithub()  {
	ctx := context.Background()
	latest , err := githubFetchLatestTagRepo(userName + "/" + repoName)
	if err != nil{
		panic(err)
	}
	fmt.Println("[Updating Database version: ", strings.Trim(latest.Name," ") ,"]")
	err = downloadReleaseAndUnzip(ctx, latest.ZipBallUrl)
	if err != nil {
		panic(err)
	}
	fmt.Println("[Update Database Success.]")
}

func githubFetchLatestTagRepo(repo string) (latest githubTagUrlData, err error) {
	var latests []githubTagUrlData
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	url := fmt.Sprintf("https://api.github.com/repos/%s/tags", repo)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return latest, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return latest, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return latest, err
	}
	err = json.Unmarshal(body, &latests)
	if err != nil {
		return latest, err
	}
	if len(latests) == 0 {
		return latest, fmt.Errorf("no tags found for %s", repo)
	}
	return latests[0], nil
}

func unzipFile(r * zip.Reader, dest string)(err error) {
	for _, f := range r.File {
		rc, err := f.Open()
		if err != nil {
			return err
		}
		directory, name := filepath.Split(f.Name)
		paths := strings.Split(directory, "/")
		finalPath := strings.Join(paths[1:], "/")

		if strings.HasPrefix(name, ".") || strings.HasPrefix(finalPath, ".") {
			continue
		}
		newPath := filepath.Join(dest, finalPath + name)
		if f.FileInfo().IsDir() {
			os.MkdirAll(newPath, f.Mode())
		} else {
			f, err := os.OpenFile(
				newPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			_, err = io.Copy(f, rc)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func downloadReleaseAndUnzip(ctx context.Context, downloadURL string) (err error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, downloadURL, nil)
	if err != nil {
		return fmt.Errorf("failed to create HTTP request to %s: %s", downloadURL, err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to download a release file from %s: %s", downloadURL, err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to download a release file from %s: Not successful status %d", downloadURL, res.StatusCode)
	}

	buf, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("failed to create buffer for zip file: %s", err)
	}

	err = os.MkdirAll(utils.GetDatabasePath(), os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to create template base folder: %s", err)
	}

	reader := bytes.NewReader(buf)
	z, err := zip.NewReader(reader, reader.Size())
	err = unzipFile(z, utils.GetDatabasePath())
	if err != nil {
		return fmt.Errorf("failed to uncompress zip file: %s", err)
	}
	return err
}
