package tools

// This part of tools package will download rabbit-backup with curl
// bash: curl --silent --user "$rabbitUser:$rabbitPass" --header "Content-type: application/json" http://$rabbitAddr/api/definitions | tee $backupFile
//

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func RabbitBackup(rabbitUrl string, rabbitUser string, rabbitPass string) error {
	// Create the file
	out, err := os.Create("backup.json")
	if err != nil {
		return err
	}
	defer out.Close()

	rabbit_download_url := rabbitUrl + "/api/definitions"
	req, err := http.NewRequest("GET", rabbit_download_url, nil)
	if err != nil {
		return err
	}
	req.SetBasicAuth(rabbitUser, rabbitPass)
	req.Header.Set("Content-type", "application/json")

	// Get the data
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check server response
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}
	return nil
}
