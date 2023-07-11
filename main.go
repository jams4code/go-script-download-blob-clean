package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/lf-edge/eve/libs/zedUpload/types"
)

// CDN Classic
// const (
// 	accountURL    = "https://accounturl.blob.iotlab.connect.bobst.com/" // replace with your account URL
// 	accountName   = "accounturl" // replace with your account name
// 	accountKey    = "XXXXXXXXXXXXXXXXXXXXXXXXX" // replace with your account master key
// 	containerName = "images" // replace with your container name
// 	remoteFile    = "develop-20230606.1-iotedge.packer" // replace with your remote file name
// 	objMaxSize    = 3608723968 // replace with your object max size
// 	cdnType	   	  = "classic"
// )

// CDN Verizon
const (
	accountURL    = "https://deviceimageverizon.azureedge.net" // replace with your account URL
	accountName   = "deviceimageverizon" // replace with your account name
	accountKey    = "XXXXXXXXXXXXXXXXXXXXXXXXX" // replace with your account master key
	containerName = "images" // replace with your container name
	remoteFile    = "develop-20230606.1-iotedge.packer" // replace with your remote file name
	objMaxSize    = 10000000000 // replace with your object max size
	cdnType	   	  = "verizon"
)


func main() {
	// Create an HTTP client
	httpClient := &http.Client{}

	// Initialize doneParts
	doneParts := types.DownloadedParts{}

	// Get the current working directory
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current working directory:", err)
		return
	}

	// Create the local file path
	localFile := filepath.Join(cwd, remoteFile)


	// counters
	var totalRequest int = 1000
	var successCounter int = 0
	var errorCounter int = 0
	var error403Counter int = 0

	// The for loop that will make the download attempt 10 times
	for i := 0; i < totalRequest; i++ {
		fmt.Println("Attempt number: ", i+1)
		fmt.Println("Downloading Blob...")

		_, err := DownloadAzureBlob(accountURL, accountName, accountKey, containerName, remoteFile, localFile, objMaxSize, httpClient, doneParts, nil)
		if err != nil {
			// Check if the error is a 403 error
			if strings.Contains(err.Error(), "RESPONSE Status: 403 Forbidden") {
				fmt.Println("Error downloading blob: 403")
				// Log to file to check when it works and when it fails
				// create a filename "error-403_"+ i +"_"+cdnType i being the iteration number
				filename := "error-403__" + cdnType

				LogToFile(filename, err.Error())

				error403Counter++
			} else {
				fmt.Println("---->Error downloading blob<----")
				LogToFile("error-"+cdnType, err.Error())

				errorCounter++
			}

			// wait 5 seconds before trying again
			time.Sleep(5 * time.Second)

			// If there is an error, we log it and go to the next iteration
			continue
		}

		// Create a file just to know that the run was successful "<timestamp>-download-success.txt"
		LogToFile("success"+cdnType, "Blob downloaded successfully")

		successCounter++

		fmt.Println("Blob downloaded successfully")
	}

	// print a separator
	fmt.Println("--------------------------------------------------")

	// Display the counters
	fmt.Println("Total requests: ", totalRequest)
	fmt.Println("Success: ", successCounter)
	fmt.Println("Error: ", errorCounter)
	fmt.Println("Error 403: ", error403Counter)

	// print a separator
	fmt.Println("--------------------------------------------------")
}