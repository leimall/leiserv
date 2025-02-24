package response

import "leiserv/models/example"

type ExaFileResponse struct {
	File example.ExaFileUploadAndDownload `json:"file"`
}
