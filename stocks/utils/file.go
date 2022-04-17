package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"io"
	"os"
	"path"
	"ww/stocks/xerr"
)

func LoadExcelToTempFolder(ctx *gin.Context) (filePath string) {
	file, _, err := ctx.Request.FormFile("file")
	xerr.ThrowError(err)

	err = os.Chmod(os.TempDir(), 0777)
	xerr.ThrowError(err)

	filePath = path.Join(os.TempDir(), uuid.New().String()+".csv")
	out, err := os.Create(filePath)
	defer out.Close()
	xerr.ThrowError(err)

	_, err = io.Copy(out, file)
	xerr.ThrowError(err)

	return filePath
}
