
package study

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/study"
    studyReq "github.com/flipped-aurora/gin-vue-admin/server/model/study/request"
)

type BookService struct {}
// CreateBook 创建书籍记录
// Author [yourname](https://github.com/yourname)
func (bookService *BookService) CreateBook(book *study.Book) (err error) {
	err = global.GVA_DB.Create(book).Error
	return err
}

// DeleteBook 删除书籍记录
// Author [yourname](https://github.com/yourname)
func (bookService *BookService)DeleteBook(ID string) (err error) {
	err = global.GVA_DB.Delete(&study.Book{},"id = ?",ID).Error
	return err
}

// DeleteBookByIds 批量删除书籍记录
// Author [yourname](https://github.com/yourname)
func (bookService *BookService)DeleteBookByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]study.Book{},"id in ?",IDs).Error
	return err
}

// UpdateBook 更新书籍记录
// Author [yourname](https://github.com/yourname)
func (bookService *BookService)UpdateBook(book study.Book) (err error) {
	err = global.GVA_DB.Model(&study.Book{}).Where("id = ?",book.ID).Updates(&book).Error
	return err
}

// GetBook 根据ID获取书籍记录
// Author [yourname](https://github.com/yourname)
func (bookService *BookService)GetBook(ID string) (book study.Book, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&book).Error
	return
}
// GetBookInfoList 分页获取书籍记录
// Author [yourname](https://github.com/yourname)
func (bookService *BookService)GetBookInfoList(info studyReq.BookSearch) (list []study.Book, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&study.Book{})
    var books []study.Book
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.Title != nil && *info.Title != "" {
        db = db.Where("title LIKE ?","%"+*info.Title+"%")
    }
        if info.StartPrice != nil && info.EndPrice != nil {
            db = db.Where("price BETWEEN ? AND ? ",info.StartPrice,info.EndPrice)
        }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&books).Error
	return  books, total, err
}
func (bookService *BookService)GetBookPublic() {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
