package packerservice

import (
	"archive/zip"
	"codegen-service/internal/s3"
	"io"
	"log/slog"
	"os"
	"path/filepath"
)

type PackerService struct {
	log      *slog.Logger
	s3Client *s3.S3Client
}

func NewPackerService(
	log *slog.Logger,
	s3Client *s3.S3Client,
) *PackerService {
	return &PackerService{
		log:      log,
		s3Client: s3Client,
	}
}

func (s *PackerService) ProcessProject(projectId string) error {
	const op = "packer.ProcessProject"

	log := s.log.With(
		slog.String("op", op),
	)

	projectDirPath := filepath.Join("output", projectId)
	zipPath, err := s.packProject(projectDirPath)
	if err != nil {
		log.Error("Error packing project: ", err)
		return err
	}

	err = s.uploadProject(zipPath, projectId)
	if err != nil {
		log.Error("Error uploading project: ", err)
		return err
	}

	s.safeDeleteProject(projectDirPath)

	return nil
}

func (s *PackerService) packProject(projectDirPath string) (string, error) {
	zipPath := filepath.Join(projectDirPath, "saf.zip")

	zipFile, err := os.Create(zipPath)
	if err != nil {
		return zipPath, err
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	err = filepath.Walk(projectDirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(projectDirPath, path)
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		zipEntry, err := zipWriter.Create(relPath)
		if err != nil {
			return err
		}

		_, err = io.Copy(zipEntry, file)
		return err
	})

	return zipPath, err
}

func (s *PackerService) uploadProject(zipPath string, projectId string) error {
	finalZipFile, err := os.Open(zipPath)
	if err != nil {
		return err
	}

	err = s.s3Client.UploadFile(finalZipFile, projectId)
	return err
}

func (s *PackerService) safeDeleteProject(projectDirPath string) {
	const op = "packer.safeDeleteProject"

	log := s.log.With(
		slog.String("op", op),
	)

	if err := os.RemoveAll(projectDirPath); err != nil {
		log.Error("Project dir was not removed - delete manually", projectDirPath)
	}
}
