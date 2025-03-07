package study

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type BookRouter struct {}

// InitBookRouter 初始化 书籍 路由信息
func (s *BookRouter) InitBookRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	bookRouter := Router.Group("book").Use(middleware.OperationRecord())
	bookRouterWithoutRecord := Router.Group("book")
	bookRouterWithoutAuth := PublicRouter.Group("book")
	{
		bookRouter.POST("createBook", bookApi.CreateBook)   // 新建书籍
		bookRouter.DELETE("deleteBook", bookApi.DeleteBook) // 删除书籍
		bookRouter.DELETE("deleteBookByIds", bookApi.DeleteBookByIds) // 批量删除书籍
		bookRouter.PUT("updateBook", bookApi.UpdateBook)    // 更新书籍
	}
	{
		bookRouterWithoutRecord.GET("findBook", bookApi.FindBook)        // 根据ID获取书籍
		bookRouterWithoutRecord.GET("getBookList", bookApi.GetBookList)  // 获取书籍列表
	}
	{
	    bookRouterWithoutAuth.GET("getBookPublic", bookApi.GetBookPublic)  // 书籍开放接口
	}
}
