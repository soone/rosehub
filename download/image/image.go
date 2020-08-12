package download

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/valyala/fasthttp"
)

func DownloadImage(imgURL, sourceDir, save, fileName string) error {
	vals, err := url.Parse(imgURL)
	if err != nil {
		return err
	}

	statusCode, body, err := fasthttp.Get(nil, imgURL)
	if err != nil {
		return err
	}

	if statusCode != http.StatusOK {
		return errors.New("statusCode is not 200")
	}

	var fSave, fName, suffix string
	lastSlashIndex := strings.LastIndex(vals.Path, "/")
	lastDotIndex := strings.LastIndex(vals.Path, ".")
	if lastDotIndex == -1 {
		lastDotIndex = len(vals.Path)
	}

	if lastSlashIndex == -1 {
		lastSlashIndex = 0
	}

	fSave = sourceDir + vals.Path[:lastSlashIndex]
	fName = vals.Path[lastSlashIndex:lastDotIndex]
	suffix = vals.Path[lastDotIndex:]

	if len(save) > 0 {
		fSave = save
	}

	if len(fileName) > 0 {
		fName = fileName
	}

	finalFileName := fSave + fName + suffix
	err = os.MkdirAll(fSave, 0700)
	if err != nil {
		return err
	}

	file, err := os.Create(finalFileName)
	if err != nil {
		return err
	}

	b := bytes.NewReader(body)
	_, err = io.Copy(file, b)
	if err != nil {
		return err
	}

	return nil
}
