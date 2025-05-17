package Server_Protection_System

import (
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var Amount int
var Unit int
var NameFile string
var Folder_name string
var Directory_name string
var Drive string

// ===================================

var FileName string
var ContentType string
var Path string

func saveToDatabase(fileName string, contentType string, path string) {
	FileName = fileName
	ContentType = contentType
	Path = path

}

func isVideoValid(filePath string) bool {
	cmd := exec.Command("ffmpeg", "-v", "error", "-i", filePath, "-f", "null", "-")
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("❌ ffmpeg : %s | %s\n", err, strings.TrimSpace(string(output)))
		return false
	}
	return true
}

// # This function is for capturing video from the user.
//
// # Parameters:
//
// - Amount int =  File size value, for example 1 MB
//
// - Unit int = The unit of value, for example 20, has three states:
//
// 1- 10 KB
//
// 2- 20 MB
//
// 3- 30 GB
//
// - NameFile string = The name under which the user uploads their file.
//
// - Folder_name string = The name of the folder where the video is saved, for example, Uploads.
//
// - Directory_name string = Directory path inside that folder, for example video
//
// - Drive string = The drive where the file is saved or the path itself, for example: D:/MyApp
//
// var FileName string
//
// var ContentType string
//
// var Path string
//
// var Emali string
//
// var User string
//
// var User_n string
//
// These are the return parameters of this function.
func Movie_get_Api(c *gin.Context) {

	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, int64(Amount)<<int64(Unit))
	file, err := c.FormFile(NameFile)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"response": "!!!Error downloading file!!!"})
		return
	}
	originalName := file.Filename
	if strings.HasPrefix(originalName, ".") || strings.Contains(originalName, "..") {
		log.Println("🚫 The file name is suspicious", originalName)
		c.String(http.StatusBadRequest, "⚠️ The file name is suspicious.")
		return
	}

	src, err := file.Open()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"response": "!!!Error processing file!!!"})
		return
	}
	defer src.Close()
	buffer := make([]byte, 512)
	_, err = src.Read(buffer)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"response": "!!!Error reading file!!!"})
		return
	}
	contentType := http.DetectContentType(buffer)
	allowedTypes := map[string]string{
		"video/mp4": ".mp4",
	}
	ext, ok := allowedTypes[contentType]
	if !ok {
		c.JSON(http.StatusOK, gin.H{"response": "!!!The submission format is not allowed!!!"})
		return
	}
	newFileName := uuid.New().String() + ext
	savePath := filepath.Join(Drive, Folder_name, Directory_name)
	os.MkdirAll(savePath, 0700)
	fullPath := filepath.Join(savePath, newFileName)

	_, err = src.Seek(0, io.SeekStart)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"response": "!!!Error in rewinding file to beginning of reading!!!"})

		return
	}

	dst, err := os.OpenFile(fullPath, os.O_CREATE|os.O_WRONLY, 0600) // dst, err := os.Create(fullPath)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"response": "!!!Error creating destination file!!!"})

		return
	}
	defer dst.Close()
	_, err = io.Copy(dst, src)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"response": "!!!Error copying file!!!"})

		return
	}
	if !isVideoValid(fullPath) {
		os.Remove(fullPath)
		c.JSON(http.StatusOK, gin.H{"response": "!!!Invalid or corrupted video!!!"})
		return
	}
	saveToDatabase(newFileName, contentType, fullPath)
	c.JSON(http.StatusOK, gin.H{"response": "!!!File uploaded successfully!!!"})

}
