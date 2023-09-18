package service

import (
	"go-bbs/app/http/model"
	"go-bbs/app/repository"
	"go-bbs/global"
	"go-bbs/utils/upload"
	"mime/multipart"
	"strings"
	"time"
)

type uploadService struct{}

var UploadService = newUploadService()

func newUploadService() *uploadService {
	return new(uploadService)
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: Upload
//@description: 创建文件上传记录
//@param: file model.ExaFileUploadAndDownload
//@return: error

func (e *uploadService) Upload(file model.Attach) (model.Attach, error) {
	return file, global.DB.Create(&file).Error
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: FindFile
//@description: 查询文件记录
//@param: id uint
//@return: model.ExaFileUploadAndDownload, error

func (e *uploadService) FindFile(id int) (model.Attach, error) {
	var file = model.Attach{Aid: id}
	err := repository.AttachRepository.First(&file, nil)
	if err != nil {
		return model.Attach{}, err
	}
	return file, err
}

func (e *uploadService) UploadFile(header *multipart.FileHeader) (file model.Attach, err error) {
	oss := upload.NewOss()
	filePath, _, uploadErr := oss.UploadFile(header)
	if uploadErr != nil {
		panic(err)
	}
	s := strings.Split(header.Filename, ".")
	f := model.Attach{
		Filename:    filePath,
		Orgfilename: header.Filename,
		Filesize:    header.Size,
		Uid:         global.User.Uid,
		CreateDate:  time.Now().Unix(),
		Filetype:    s[len(s)-1],
	}
	return e.Upload(f)
}
