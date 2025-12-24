package model

type ImageData struct {
	StorageID        uint64 `json:"storage_id"`
	Type             string `json:"type"`
	Path             string `json:"path"`
	Filename         string `json:"filename"`
	Mime             string `json:"mime"`
	OriginalFilename string `json:"original_filename"`
}
