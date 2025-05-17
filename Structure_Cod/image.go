package Server_Protection_System

import (
	"image"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var Amount_i int
var Unit_i int
var NameFile_i string
var Folder_name_i string
var Directory_name_i string
var Drive_i string

// ===============================================
var FileName_i string
var ContentType_i string
var Path_i string

func saveToDatabase1(fileName string, contentType string, path string) {
	FileName_i = fileName
	ContentType_i = contentType
	Path_i = path

}

// # This function is for capturing Image from the user.
//
// # Parameters:
//
// - Amount_i int =  File size value, for example 1 MB
//
// - Unit_i int = The unit of value, for example 20, has three states:
//
// 1- 10 KB
//
// 2- 20 MB
//
// 3- 30 GB
//
// - NameFile_i string = The name under which the user uploads their file.
//
// - Folder_name_i string = The name of the folder where the Image is saved, for example, Uploads.
//
// - Directory_name_i string = Directory path inside that folder, for example Image
//
// - Drive_i string = The drive where the file is saved or the path itself, for example: D:/MyApp
//
// var FileName_i string
//
// var ContentType_i string
//
// var Path_i string
//
// These are the return parameters of this function.
func Image_Get_Api(c *gin.Context) {
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, int64(Amount_i)<<int64(Unit_i))
	file, err := c.FormFile(NameFile_i)
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
		c.JSON(http.StatusOK, gin.H{"response": "!!!Error opening file!!!"})

		return
	}
	defer src.Close()
	_, format, err := image.DecodeConfig(src)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"response": "!!!The file is not valid.!!!"})

		return
	}
	allowedFormats := map[string]string{
		"jpeg": ".jpg",
		"png":  ".png",
	}
	ext, ok := allowedFormats[format]
	if !ok {
		c.JSON(http.StatusOK, gin.H{"response": "!!!Format is not allowed.!!!"})

		return
	}
	_, err = src.Seek(0, io.SeekStart)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"response": "!!!Error restoring file!!!"})

		return
	}
	savePath := filepath.Join(Drive_i, Folder_name_i, Directory_name_i)
	os.MkdirAll(savePath, 0700)
	newFileName := uuid.New().String() + ext
	fullPath := filepath.Join(savePath, newFileName)

	dst, err := os.OpenFile(fullPath, os.O_CREATE|os.O_WRONLY, 0600) // dst, err := os.Create(fullPath)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"response": "!!!Error saving image!!!"})

		return
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"response": "!!!Error saving file!!!"})

		return
	}

	saveToDatabase1(newFileName, format, fullPath)
	c.JSON(http.StatusOK, gin.H{"response": "!!!File uploaded successfully!!!"})

}
