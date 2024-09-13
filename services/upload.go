package services

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"time"

	"main/functions"

	"github.com/gin-gonic/gin"
)

func Upload(c *gin.Context) {
	fmt.Println("Upload")

	path := "/home/runner/go-gin-api/uploads/"

	// single file
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		c.Abort()
		return
	}
	log.Println(file.Filename)

	t := time.Now()
	prefix := fmt.Sprintf(t.Format("20060102150405"))
	fileName, ext := functions.FileNameAndExt(file.Filename)
	fmt.Printf("fileName: %s\v", fileName)
	fmt.Printf("ext: %s\v", ext)

	bv := []byte(prefix + fileName)
	hasher := sha1.New()
	hasher.Write(bv)
	shaText := base64.URLEncoding.EncodeToString(hasher.Sum(nil))

	dst := path + shaText + ext

	// Upload the file to specific dst.
	err = c.SaveUploadedFile(file, dst)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		c.Abort()
		return
	}

	//insert into DB

	c.JSON(http.StatusOK, gin.H{
		"message": dst,
	})
}
