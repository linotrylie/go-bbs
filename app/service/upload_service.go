package service

import (
	"go-bbs/app/http/model"
	"go-bbs/global"
)

type FileUploadAndDownloadService struct{}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: Upload
//@description: 创建文件上传记录
//@param: file model.ExaFileUploadAndDownload
//@return: error

func (e *FileUploadAndDownloadService) Upload(file model.Attach) error {
	return global.DB.Create(&file).Error
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: FindFile
//@description: 查询文件记录
//@param: id uint
//@return: model.ExaFileUploadAndDownload, error

func (e *FileUploadAndDownloadService) FindFile(id uint) (model.Attach, error) {
	var file model.Attach
	err := global.DB.Where("aid = ?", id).First(&file).Error
	return file, err
}
