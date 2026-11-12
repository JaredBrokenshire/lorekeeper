package requests

import (
	"fmt"
	"io"
	"net/http"
)

type UploadedFile struct {
	Name        string
	ContentType string
	Size        int64
	Open        func() (io.ReadCloser, error)
}

func MultipartFileRequest(r *http.Request, field string) (UploadedFile, error) {
	if err := r.ParseMultipartForm(10 * 1024 * 1024); err != nil {
		return UploadedFile{}, fmt.Errorf("error parsing multipart form file: %v", err)
	}

	f, h, err := r.FormFile(field)
	if err != nil {
		return UploadedFile{}, fmt.Errorf("error getting multipart form file:: %w", err)
	}
	defer f.Close()

	// Read the first 512 bytes for MIME detection.
	buffer := make([]byte, 512)
	n, err := f.Read(buffer)
	if err != nil && err != io.EOF {
		return UploadedFile{}, fmt.Errorf("error reading file: %w", err)
	}

	contentType := http.DetectContentType(buffer[:n])

	return UploadedFile{
		Name:        h.Filename,
		ContentType: contentType,
		Size:        h.Size,
		Open: func() (io.ReadCloser, error) {
			file, _, err := r.FormFile(field)
			return file, err
		},
	}, nil
}
