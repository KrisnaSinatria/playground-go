package utils

import (
	"context"
	"mime/multipart"
	"os"
	"time"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

// UploadToCloudinary menerima file dari HTTP request dan mengunggahnya ke Cloudinary
func UploadToCloudinary(fileHeader *multipart.FileHeader) (string, error) {
	// Timeout 10 detik agar request tidak me-hang jika koneksi internet lambat
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 1. Inisialisasi Cloudinary menggunakan environment variable dari .env
	cld, err := cloudinary.NewFromParams(
		os.Getenv("CLOUDINARY_CLOUD_NAME"),
		os.Getenv("CLOUDINARY_API_KEY"),
		os.Getenv("CLOUDINARY_API_SECRET"),
	)
	if err != nil {
		return "", err
	}

	// 2. Buka stream file fisik
	file, err := fileHeader.Open()
	if err != nil {
		return "", err
	}
	defer file.Close()

	// 3. Unggah file ke folder "rooms" di akun Cloudinary Anda
	uploadResult, err := cld.Upload.Upload(ctx, file, uploader.UploadParams{
		Folder: "rooms",
	})
	if err != nil {
		return "", err
	}

	// 4. Kembalikan HTTPS Secure URL gambar
	return uploadResult.SecureURL, nil
}
