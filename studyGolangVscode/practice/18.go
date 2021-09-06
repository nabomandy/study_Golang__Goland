if filename, success := UploadFile(); success {
	fmt.Println("Upload success", filename)
} else {
	fmt.Println("Failed to upload")
}