package download

type Download struct {
	DownloadLink string `json:"downloadLink"`
}

type DownloadUseCase interface {
	IsValidExtension(file_name string) bool
	CreateDownloadLink(local_file_path string) *Download
}
