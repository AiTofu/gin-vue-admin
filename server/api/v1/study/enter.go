package study

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct{ BookApi }

var bookService = service.ServiceGroupApp.StudyServiceGroup.BookService
