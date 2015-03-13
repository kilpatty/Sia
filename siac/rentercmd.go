package main

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/NebulousLabs/Sia/modules"
)

var (
	renterCmd = &cobra.Command{
		Use:   "renter",
		Short: "Perform renter actions",
		Long:  "Upload and download files, or view a list of previously uploaded files.",
		Run:   wrap(renterstatuscmd),
	}

	renterUploadCmd = &cobra.Command{
		Use:   "upload [filename] [nickname] [pieces]",
		Short: "Upload a file",
		Long:  "Upload a file using a given nickname and split across 'pieces' hosts.",
		Run:   wrap(renteruploadcmd),
	}

	renterDownloadCmd = &cobra.Command{
		Use:   "download [nickname] [filename]",
		Short: "Download a file",
		Long:  "Download a previously-uploaded file to a specified destination.",
		Run:   wrap(renterdownloadcmd),
	}

	renterStatusCmd = &cobra.Command{
		Use:   "status",
		Short: "View a list of uploaded files",
		Long:  "View a list of files that have been uploaded to the network.",
		Run:   wrap(renterstatuscmd),
	}
)

// siac does not support /renter/upload, only /renter/uploadpath
func renteruploadcmd(filename, nickname, pieces string) {
	err := callAPI(fmt.Sprintf("/renter/uploadpath?filename=%s&nickname=%s&pieces=%s", filename, nickname, pieces))
	if err != nil {
		fmt.Println("Could not upload file:", err)
		return
	}
	fmt.Println("Uploaded", filename, "as", nickname)
}

func renterdownloadcmd(nickname, filename string) {
	err := callAPI(fmt.Sprintf("/renter/download?nickname=%s&filename=%s", nickname, filename))
	if err != nil {
		fmt.Println("Could not download file:", err)
		return
	}
	fmt.Println("Downloaded", nickname, "to", filename)
}

func renterstatuscmd() {
	status := new(modules.RentInfo)
	err := getAPI("/renter/status", status)
	if err != nil {
		fmt.Println("Could not get file status:", err)
		return
	}
	if len(status.Files) == 0 {
		fmt.Println("Not files have been uploaded.")
		return
	}
	fmt.Println("Uploaded", len(status.Files), "files:")
	for _, file := range status.Files {
		fmt.Println("\t", file)
	}
}