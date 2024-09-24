package controller

import (
	"strconv"
	"u_demo/logic"
	"u_demo/models"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreatePostHandler(c *gin.Context) {
	p := new(models.Post)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Debug("c.ShouldBindJSON(p) error", zap.Any("err", err))
		zap.L().Error("create post with invalid param")
		ResponseError(c, CodeInvalidParam)
	}

	// 获取数据
	userID, err := getCurrentUserID(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}
	p.AuthorID = userID
	if err := logic.CreatePost(p); err != nil {
		zap.L().Error("logic.CreatePost(p) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}

func GetPostDetailHandlerByID(c *gin.Context) {
	pidStr := c.Param("id")
	pid, err := strconv.ParseInt(pidStr, 10, 64)
	if err != nil {
		zap.L().Error("get post detail with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	data, err := logic.GetPostById(pid)
	if err != nil {
		zap.L().Error("logic.GetPostById(pid) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}

func GetPostListHandler(c *gin.Context) {
	page, size := getPageInfo(c)
	// 获取数据
	data, err := logic.GetPostList(page, size)
	if err != nil {
		zap.L().Error("logic.GetPostList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}

// GetPostListHandler2 升级版帖子列表接口
// 根据前端传来的参数动态的获取帖子列表
// 按创建时间排序 或者 按照 分数排序
// 1. 获取请求的query string参数
// 2. 去redis查询id列表
// 3. 根据id去数据库查询帖子详细信息
func GetPostListHandler2(c *gin.Context) {
	// GET请求参数(query string)：/api/v1/posts2?page=1&size=10&order=time
	// 初始化结构体时指定初始参数
	//p := &models.ParamPostList{
	//	Page:  1,
	//	Size:  10,
	//	Order: models.OrderTime, // magic string
	//}
	////c.ShouldBind()  根据请求的数据类型选择相应的方法去获取数据
	////c.ShouldBindJSON() 如果请求中携带的是json格式的数据，才能用这个方法获取到数据
	//if err := c.ShouldBindQuery(p); err != nil {
	//	zap.L().Error("GetPostListHandler2 with invalid params", zap.Error(err))
	//	ResponseError(c, CodeInvalidParam)
	//	return
	//}
	//data, err := logic.GetPostListNew(p) // 更新：合二为一
	//// 获取数据
	//if err != nil {
	//	zap.L().Error("logic.GetPostList() failed", zap.Error(err))
	//	ResponseError(c, CodeServerBusy)
	//	return
	//}
	//ResponseSuccess(c, data)
	//// 返回响应
}
