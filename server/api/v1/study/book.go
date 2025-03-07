package study

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/study"
    studyReq "github.com/flipped-aurora/gin-vue-admin/server/model/study/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type BookApi struct {}



// CreateBook 创建书籍
// @Tags Book
// @Summary 创建书籍
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body study.Book true "创建书籍"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /book/createBook [post]
func (bookApi *BookApi) CreateBook(c *gin.Context) {
	var book study.Book
	err := c.ShouldBindJSON(&book)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = bookService.CreateBook(&book)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteBook 删除书籍
// @Tags Book
// @Summary 删除书籍
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body study.Book true "删除书籍"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /book/deleteBook [delete]
func (bookApi *BookApi) DeleteBook(c *gin.Context) {
	ID := c.Query("ID")
	err := bookService.DeleteBook(ID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteBookByIds 批量删除书籍
// @Tags Book
// @Summary 批量删除书籍
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /book/deleteBookByIds [delete]
func (bookApi *BookApi) DeleteBookByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	err := bookService.DeleteBookByIds(IDs)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateBook 更新书籍
// @Tags Book
// @Summary 更新书籍
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body study.Book true "更新书籍"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /book/updateBook [put]
func (bookApi *BookApi) UpdateBook(c *gin.Context) {
	var book study.Book
	err := c.ShouldBindJSON(&book)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = bookService.UpdateBook(book)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindBook 用id查询书籍
// @Tags Book
// @Summary 用id查询书籍
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询书籍"
// @Success 200 {object} response.Response{data=study.Book,msg=string} "查询成功"
// @Router /book/findBook [get]
func (bookApi *BookApi) FindBook(c *gin.Context) {
	ID := c.Query("ID")
	rebook, err := bookService.GetBook(ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(rebook, c)
}
// GetBookList 分页获取书籍列表
// @Tags Book
// @Summary 分页获取书籍列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query studyReq.BookSearch true "分页获取书籍列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /book/getBookList [get]
func (bookApi *BookApi) GetBookList(c *gin.Context) {
	var pageInfo studyReq.BookSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := bookService.GetBookInfoList(pageInfo)
	if err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败:" + err.Error(), c)
        return
    }
    response.OkWithDetailed(response.PageResult{
        List:     list,
        Total:    total,
        Page:     pageInfo.Page,
        PageSize: pageInfo.PageSize,
    }, "获取成功", c)
}

// GetBookPublic 不需要鉴权的书籍接口
// @Tags Book
// @Summary 不需要鉴权的书籍接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /book/getBookPublic [get]
func (bookApi *BookApi) GetBookPublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    bookService.GetBookPublic()
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的书籍接口信息",
    }, "获取成功", c)
}
