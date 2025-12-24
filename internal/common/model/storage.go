package model

type StorageData struct {
	StorageID        uint64 `json:"storage_id"` // used to unmarshal storage_id from elasticsearch
	FileName         string `json:"file_name,omitempty"`
	Type             string `json:"type"`
	Path             string `json:"path"`
	Filename         string `json:"filename"`
	Mime             string `json:"mime"`
	OriginalFilename string `json:"original_filename,omitempty"`
}
