package transform

import (
	"go-bbs/app/http/model"
	"go-bbs/app/http/model/response"
	"go-bbs/global"
	"strings"
	"time"
)

func TransformAttach(attach *model.Attach) (attachVo *response.AttachVo) {
	if attach != nil {
		return nil
	}
	attachVo = &response.AttachVo{
		Aid:      attach.Aid,
		Tid:      attach.Tid,
		Pid:      attach.Pid,
		Uid:      attach.Uid,
		Filesize: attach.Filesize,
		Width:    attach.Width,
		Height:   attach.Height,
		//Filename:    attach.Filename,
		Orgfilename: attach.Orgfilename,
		Filetype:    attach.Filetype,
		CreateDate:  time.Unix(attach.CreateDate, 0).Format(time.DateTime),
		Comment:     attach.Comment,
		Downloads:   attach.Downloads,
		Credits:     attach.Credits,
		Golds:       attach.Golds,
		Rmbs:        attach.Rmbs,
		Isimage:     attach.Isimage,
	}

	if strings.Contains(attach.Filename, "http") {
		attachVo.Filename = attach.Filename
	} else {
		attachVo.Filename = global.CONFIG.System.Host + "/" + attach.Filename
	}
	return nil
}
