package fs

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
)

// ExtractTarGz extracts a tar.gz file.
func ExtractTarGz(stream io.Reader, destDir string, logHandler func(filename string)) error {
	uncompressedStream, err := gzip.NewReader(stream)
	if err != nil {
		return err
	}

	tarReader := tar.NewReader(uncompressedStream)
	for {
		header, err := tarReader.Next()

		if err != nil {
			if err == io.EOF {
				break
			}

			return err
		}

		logHandler(header.Name)

		switch header.Typeflag {
		case tar.TypeDir:
			target := JoinPath(destDir, header.Name)

			if err := Mkdirp(target); err != nil {
				return fmt.Errorf("ExtractTarGz: Mkdir() failed: %w", err)
			}

			setTarFileAttrs(target, header)
		case tar.TypeReg:
			target := JoinPath(destDir, header.Name)

			outFile, err := os.Create(target)
			if err != nil {
				return fmt.Errorf("ExtractTarGz: Create() failed: %w", err)
			}

			if _, err := io.Copy(outFile, tarReader); err != nil {
				outFile.Close()
				return fmt.Errorf("ExtractTarGz: Copy() failed: %w", err)
			}

			if err := outFile.Close(); err != nil {
				return fmt.Errorf("ExtractTarGz: Close() failed: %w", err)
			}

			setTarFileAttrs(target, header)
		default:
			return fmt.Errorf("ExtractTarGz: Unknown type: %b in %s", header.Typeflag, header.Name)
		}
	}

	return nil
}

func setTarFileAttrs(target string, header *tar.Header) {
	os.Chmod(target, os.FileMode(header.Mode))
	os.Chtimes(target, header.AccessTime, header.ModTime)
}
