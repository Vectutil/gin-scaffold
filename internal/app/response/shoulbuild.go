package response

import (
	"gin-scaffold/internal/app/types/common"
	"github.com/gin-gonic/gin"
)

func ShouldBindForList(c *gin.Context, req common.IBaseListParam) error {
	if err := c.ShouldBindQuery(req); err != nil {
		return err
	}
	req.Adjust()
	return nil
}
