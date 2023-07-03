# Go-Script-Download-Blob

Go-Script-Download-Blob is a Go program that simulates the process of downloading a VM image from Azure storage behind a CDN. The aim is to reproduce a known issue experienced on edge devices during such downloads. T

It works by executing 10 downloads of a blob using specific hardcoded variables. Parameters can be adjusted to mimic different CDN configurations, specifically for Classic and Verizon CDN.

## Files

The project includes the following files:

- `azure.go`: Contains the functionality for creating a pipeline and downloading a blob from Azure storage. This files was extracted from EVE-OS, which is the vendor we use for the remote management of our edge devices. There is the EVE-Layer that runs and download the blobs corresponding to the VM which is the actual runtime for our edge devices. [Source code can be found here.(https://github.dev/lf-edge/eve/blob/master/libs/zedUpload/azureutil/azure.go)]
- `log.go`: Handles logging to files. Was just made to create a file for each download and make it easy to check the failure reason and how many success against, how many failures.
- `main.go`: The main entry point for the program.
- `go.mod` and `go.sum`: Go module and checksum files respectively, handling the project's dependencies.

## Building and Running

You will need to have Go installed on your machine. The version used at the time of the last update is Go 1.20.5

To initialize the module of the program, navigate to the directory containing the Go files, and use the `go mod init go-script-download-blob` command.

To install the dependencies use `go mod tidy`

To run the program, execute the command `go run .` from the directory

## Customizing the Download
You can modify the hardcoded variables in the main.go file to simulate different conditions for the download. By altering these variables, you can simulate downloads from different CDN configurations, including Classic and Verizon CDNs.

Please be aware that I was able to reproduce the known issue with the Classic CDN. It has been found to work correctly with the Verizon CDN.

## Logs and Error Handling
The program generates logs to record both successful downloads and errors. It differentiates between a general error and a 403 error, logging these separately.

## Conclusion
There is an issue on Microsoft Azure Classic CDN side. When downloading a blob directly from storage nothing wrong happens, same when it's a Verizon CDN. Otherwise, we keep getting 403 errors with azure classic cdn.