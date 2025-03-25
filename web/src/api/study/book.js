import service from '@/utils/request'
// @Tags Book
// @Summary 创建书籍
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Book true "创建书籍"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /book/createBook [post]
export const createBook = (data) => {
  return service({
    url: '/book/createBook',
    method: 'post',
    data
  })
}

// @Tags Book
// @Summary 删除书籍
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Book true "删除书籍"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /book/deleteBook [delete]
export const deleteBook = (params) => {
  return service({
    url: '/book/deleteBook',
    method: 'delete',
    params
  })
}

// @Tags Book
// @Summary 批量删除书籍
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除书籍"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /book/deleteBook [delete]
export const deleteBookByIds = (params) => {
  return service({
    url: '/book/deleteBookByIds',
    method: 'delete',
    params
  })
}

// @Tags Book
// @Summary 更新书籍
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Book true "更新书籍"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /book/updateBook [put]
export const updateBook = (data) => {
  return service({
    url: '/book/updateBook',
    method: 'put',
    data
  })
}

// @Tags Book
// @Summary 用id查询书籍
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.Book true "用id查询书籍"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /book/findBook [get]
export const findBook = (params) => {
  return service({
    url: '/book/findBook',
    method: 'get',
    params
  })
}

// @Tags Book
// @Summary 分页获取书籍列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取书籍列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /book/getBookList [get]
export const getBookList = (params) => {
  return service({
    url: '/book/getBookList',
    method: 'get',
    params
  })
}

// @Tags Book
// @Summary 不需要鉴权的书籍接口
// @Accept application/json
// @Produce application/json
// @Param data query studyReq.BookSearch true "分页获取书籍列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /book/getBookPublic [get]
export const getBookPublic = () => {
  return service({
    url: '/book/getBookPublic',
    method: 'get',
  })
}
