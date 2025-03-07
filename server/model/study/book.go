
// 自动生成模板Book
package study
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"gorm.io/datatypes"
)

// 书籍 结构体  Book
type Book struct {
    global.GVA_MODEL
    Title  *string `json:"title" form:"title" gorm:"column:title;comment:标题;" binding:"required"`  //标题
    Desc  *string `json:"desc" form:"desc" gorm:"column:desc;comment:介绍;"`  //介绍
    Price  *float64 `json:"price" form:"price" gorm:"column:price;comment:价格;"`  //价格
    Info  datatypes.JSON `json:"info" form:"info" gorm:"column:info;comment:内容;" swaggertype:"array,object"`  //内容
    Pic  string `json:"pic" form:"pic" gorm:"column:pic;comment:封面;"`  //封面
    Type  *string `json:"type" form:"type" gorm:"column:type;comment:书籍类型;"`  //类型
}


// TableName 书籍 Book自定义表名 book
func (Book) TableName() string {
    return "book"
}





