package service

import (
	"fmt"
	"mime/multipart"
	"path/filepath"
	"saikumo.org/simple-douyin/src/common"
	"saikumo.org/simple-douyin/src/dto"
	"saikumo.org/simple-douyin/src/entity"
	"saikumo.org/simple-douyin/src/repository"
	"saikumo.org/simple-douyin/src/util"
	"time"
)

type userPublishFlow struct {
	form   *multipart.Form
	title  string
	userId uint

	fileArray  []*multipart.FileHeader
	videoArray []*entity.Video

	response *dto.Response
}

func UserPublish(data *multipart.Form, title string, userId uint) (*dto.Response, error) {
	common.Logger.Infof("用户投稿，用户名为：%d", userId)
	return newUserPublishFlow(data, title, userId).do()
}

func newUserPublishFlow(form *multipart.Form, title string, userId uint) *userPublishFlow {
	return &userPublishFlow{
		form:   form,
		title:  title,
		userId: userId,
	}
}

func (flow *userPublishFlow) do() (*dto.Response, error) {
	//校验参数
	if err := flow.checkParam(); err != nil {
		return nil, err
	}
	//视频发布
	if err := flow.userPublish(); err != nil {
		return nil, err
	}
	//打包参数
	flow.packResponse()

	return flow.response, nil
}

func (flow *userPublishFlow) checkParam() error {
	//视频标题为空
	if flow.title == "" {
		return common.VideoTitleIsNullError
	}
	//视频标题超过最大长度
	if len(flow.title) > common.MAX_VIDEO_TITLE_LENGTH {
		return common.VideoTitleOverMaxLengthError
	}
	//不支持的视频文件类型
	flow.fileArray = flow.form.File["data"]
	for _, file := range flow.fileArray {
		suffix := filepath.Ext(file.Filename)
		if _, ok := common.SUPPORTED_VIDEO_TYPE_SET[suffix]; !ok {
			return common.UnsupportedVideoTypeError
		}
	}
	return nil
}

func (flow *userPublishFlow) userPublish() error {
	//多文件上传
	for _, file := range flow.fileArray {
		//保存文件
		filename := fmt.Sprintf("%d_%d_%s", flow.userId, time.Now().Unix(), file.Filename)
		filePath := filepath.Join("./static", filename)
		if err := util.SaveUploadedFile(file, filePath); err != nil {
			return err
		}

		//截取一帧画面作为封面
		snapshotName := fmt.Sprintf("%s_%s", filename, "snapshot")
		snapshotPath := filepath.Join("./static", snapshotName)
		if err := util.GenerateSnapshot(filePath, snapshotPath, 1); err != nil {
			return err
		}

		//添加到视频数组
		flow.videoArray = append(flow.videoArray, &entity.Video{
			UserId:   flow.userId,
			PlayUrl:  util.GetFileUrl(filename),
			CoverUrl: util.GetFileUrl(snapshotName),
			Title:    flow.title,
		})
	}
	//存到数据库
	userRepository := repository.GetUserRepository()
	user := userRepository.FindUserByUserId(flow.userId)
	//如果没这个用户
	if user.ID == 0 {
		return common.UserNotExistError
	}
	//保存
	user.VideoList = flow.videoArray
	userRepository.Save(user)
	return nil
}

func (flow *userPublishFlow) packResponse() {
	flow.response = &dto.Response{
		StatusCode: common.SUCCESS_STATUS_CODE,
	}
}
