package study

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct{ BookRouter }

var bookApi = api.ApiGroupApp.StudyApiGroup.BookApi
