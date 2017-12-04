package router

import (
 "gopkg.in/gin-gonic/gin.v1"
 a "myblog/apis"
)

func InitRouter() *gin.Engine {
	 router := gin.Default()
	 router.GET("/",a.IndexApi)
	 router.GET("/blog/insert/:title/:content",a.AddBlogApi)

	 router.GET("/blog/selectall", a.GetBlogsApi)

	 router.GET("/blog/selectone/:id", a.GetBlogApi)

	 router.GET("/blog/update/:id/:title/:content", a.UpdateBlogApi)

	 router.GET("/blog/delete/:id", a.DelBlogApi)

 	 return router
}

// router := gin.Default()

//  router.GET("/", IndexApi)

//  router.POST("/person", AddPersonApi)

//  router.GET("/persons", GetPersonsApi)

//  router.GET("/person/:id", GetPersonApi)

//  router.PUT("/person/:id", ModPersonApi)

//  router.DELETE("/person/:id", DelPersonApi)

//  return router

