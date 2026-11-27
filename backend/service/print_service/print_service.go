package print_service

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"printer/config"
	"printer/core"
	"printer/core/serial_code"
)

const (
	sumatraPdfAppName = "SumatraPDF.exe"
	// direct links and redirects to a pdf should not take forever, but some servers are slow
	downloadTimeout = 2 * time.Minute
	// SumatraPDF may hand the file to an already running instance and return before printing is finished,
	// so keep recently downloaded files and only sweep old leftovers
	tempFileMaxAge = 24 * time.Hour
)

// Print downloads the pdf file from pdfUrl and sends it to the default printer via SumatraPDF.
func Print(pdfUrl string) error {
	var appPath = core.GetAppPath()

	var tempDir = filepath.Join(appPath, config.PrintTempDir)
	if err := os.MkdirAll(tempDir, os.ModePerm); err != nil {
		return core.HandleError(fmt.Errorf("create print temp dir failed, %w", err))
	}

	sweepOldTempFiles(tempDir)

	localFile, err := downloadFile(pdfUrl, tempDir)
	if err != nil {
		return core.HandleError(err)
	}

	var sumatraPdfCommand = filepath.Join(appPath, "tools", sumatraPdfAppName)
	output, err := exec.Command(sumatraPdfCommand, "-print-to-default", "-silent", localFile).CombinedOutput()
	if err != nil {
		os.Remove(localFile)
		return core.HandleError(fmt.Errorf("run %s failed, %w, output: %s", sumatraPdfAppName, err, strings.TrimSpace(string(output))))
	}

	return nil
}

func downloadFile(url string, dir string) (string, error) {
	client := &http.Client{Timeout: downloadTimeout}
	resp, err := client.Get(url)
	if err != nil {
		return "", core.HandleError(fmt.Errorf("fetch pdf from %s failed, %w", url, err))
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", core.HandleError(fmt.Errorf("fetch pdf from %s failed, status: %s", url, resp.Status))
	}

	// file, err := os.CreateTemp(dir, "*.pdf")
	var tempFileName = fmt.Sprintf("%s.pdf", serial_code.CreateSerialCode())
	file, err := os.Create(filepath.Join(dir, tempFileName))
	if err != nil {
		return "", core.HandleError(fmt.Errorf("create pdf temp file failed, %w", err))
	}

	_, err = io.Copy(file, resp.Body)
	var closeErr = file.Close()
	if err == nil {
		err = closeErr
	}

	if err != nil {
		os.Remove(file.Name())
		return "", core.HandleError(fmt.Errorf("save pdf to %s failed, %w", file.Name(), err))
	}

	return file.Name(), nil
}

// sweepOldTempFiles removes pdf files left behind by previous prints. Failures are ignored: cleanup is best-effort.
func sweepOldTempFiles(dir string) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return
	}

	expiredBefore := time.Now().Add(-tempFileMaxAge)
	for _, entry := range entries {
		if entry.IsDir() || !strings.EqualFold(filepath.Ext(entry.Name()), ".pdf") {
			continue
		}

		info, err := entry.Info()
		if err != nil || info.ModTime().After(expiredBefore) {
			continue
		}

		os.Remove(filepath.Join(dir, entry.Name()))
	}
}
