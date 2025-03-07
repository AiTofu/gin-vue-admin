
package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type BookSearch struct{
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    Title  *string `json:"title" form:"title" `
    StartPrice  *float64  `json:"startPrice" form:"startPrice"`
    EndPrice  *float64  `json:"endPrice" form:"endPrice"`
    request.PageInfo
}
