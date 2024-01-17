package main

func Test() {
	var FileExtMapping = map[string]*FileExtPayload{
		"jpg": NewFileExtPayload([]string{
			"image/jpeg",
			"image/pjpeg",
			"image/jpg",
			"image/x-jpeg",
			"image/x-pjpeg",
		}, MEDIA_IMAGE, DEFAULT_FILE_ICON_IMAGE),

		"jpeg": NewFileExtPayload([]string{
			"image/jpeg",
			"image/pjpeg",
			"image/jpg",
			"image/x-jpeg",
			"image/x-pjpeg",
		}, MEDIA_IMAGE, DEFAULT_FILE_ICON_IMAGE),

		"png": NewFileExtPayload([]string{
			"image/png",
			"image/x-png",
		}, MEDIA_IMAGE, DEFAULT_FILE_ICON_IMAGE),

		"gif": NewFileExtPayload([]string{
			"image/x-gif",
			"image/gif",
		}, MEDIA_IMAGE, DEFAULT_FILE_ICON_IMAGE),

		"mp4": NewFileExtPayload([]string{
			"video/mp4",
			"application/mp4",
			"audio/mp4",
		}, MEDIA_VIDEO, DEFAULT_FILE_ICON_VIDEO),

		"mov": NewFileExtPayload([]string{
			"video/quicktime",
			"video/x-quicktime",
			"video/x-mov",
			"video/mp4",
			"application/octet-stream",
		}, MEDIA_VIDEO, DEFAULT_FILE_ICON_VIDEO),

		"pdf": NewFileExtPayload([]string{
			"application/pdf",
		}, "others", DEFAULT_FILE_ICON_PDF),
	}
}
