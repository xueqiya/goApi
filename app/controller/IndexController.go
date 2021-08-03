package controller

import (
	"github.com/gin-gonic/gin"
	"xueqiya/app/common"
)

type IndexController struct {
	display *common.Display
	data    map[string]interface{}
}

func (s *IndexController) Run(c *gin.Context) {
	s.display = &common.Display{Context: c}
	s.display.Show("来啦，老弟")
}
