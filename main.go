package main
import (
	"geek_history/tools"
	"geek_history/src/controllers"
	"log"
	"github.com/gin-gonic/gin"
)



func main(){
	router := gin.Default() //criando o manipulador de rota

	router.HTMLRender = tools.LoadTemplates("./src/views/templates") //definindo a pasta que vai configurar os includes e bases de templates
	
	router.GET("/",controllers.HomeController) //definindo a primeira rota
	
	// fazendo o servidor rodar
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
	
}


