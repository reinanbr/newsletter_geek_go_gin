package controllers

import (
	"github.com/gin-gonic/gin"
)


func HomeController(c *gin.Context){
	c.HTML(200,"index.html",gin.H{"title":"Home"})
}

	
