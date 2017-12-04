package apis
import (
 "net/http"
 "log"
 "fmt"
 "strconv"
 "gopkg.in/gin-gonic/gin.v1"
 m "myblog/models"
 "time"
)

func IndexApi(c *gin.Context) {
   c.String(http.StatusOK, "It works")
}



func AddBlogApi(c *gin.Context) {
    title :=c.Param("title")
    content :=c.Param("content")
    b := m.Blog{Title:title,Content:content,Created:time.Now()}
    ra,err := b.AddBlog()
    if err != nil {
      log.Println(err)
      c.JSON(http.StatusOK, gin.H{
        "msg": "An error occurred while the insertion was performed",
      })
      return
   }
   var msg string
   if ra==0{
      msg = fmt.Sprintf("insert fail %d,The format you entered is incorrect", ra)
   }else{
      msg = fmt.Sprintf("insert successful %d", ra)
    }
   c.JSON(http.StatusOK, gin.H{
      "msg": msg,
   })
}



func GetBlogsApi (c *gin.Context){
    blogs := make([]m.Blog, 0)
    var b m.Blog
    blogs,err:=b.GetBlogs()
    if err != nil {
       log.Println(err)
       c.JSON(http.StatusOK, gin.H{
          "msg": "An error occurred while executing the query",
       })
       return
    }
    c.JSON(http.StatusOK, gin.H{
         "blogs": blogs,
    })
}


func GetBlogApi (c *gin.Context){
  cid := c.Param("id")
  //var p m.Person
  id, err := strconv.Atoi(cid)
  if err != nil {
      log.Println(err)
      c.JSON(http.StatusOK, gin.H{
        "msg": "The format you entered is incorrect",
      })
      return
  }
  b := m.Blog{Id:id}
  b,err=b.GetBlog()
  if err != nil {
     log.Println(err)
     c.JSON(http.StatusOK, gin.H{
      "blog": nil,
     })
     return
  }
  c.JSON(http.StatusOK, gin.H{
     "blog": b,
  })

}



func UpdateBlogApi(c *gin.Context) {
    title :=c.Param("title")
    content :=c.Param("content")
    cid := c.Param("id")
    id, err := strconv.Atoi(cid)
    if err != nil {
       log.Println(err)
       c.JSON(http.StatusOK, gin.H{
       "msg": "The format you entered is incorrect",
       })
       return
    }
   b := m.Blog{Id:id}
   b,err=b.GetBlog()
   if err != nil {
      log.Println(err)
      c.JSON(http.StatusOK, gin.H{
      "msg": "An error occurs when the update is performed",
     })
      return
   }
   bb := m.Blog{Id:id,Title:title,Content:content,Created:time.Now()}
  ra,err := bb.UpdateBlog()
  if err != nil {
    log.Fatalln(err)
   }
    var msg string
   if ra==0{
      msg = fmt.Sprintf("update fail %d,The format you entered is incorrect", ra)
   }else{
      msg = fmt.Sprintf("update successful %d", ra)
    }
   c.JSON(http.StatusOK, gin.H{
       "msg": msg,
   })
}






 func DelBlogApi (c *gin.Context){
    cid := c.Param("id")
    id, err := strconv.Atoi(cid)
    if err != nil {
      log.Println(err)
      c.JSON(http.StatusOK, gin.H{
        "msg": "The format you entered is incorrect",
      })
      return
    }
   
    b:= m.Blog{Id:id}
    ra,err:= b.DelBlog()
    if err != nil {
      log.Println(err)
      c.JSON(http.StatusOK, gin.H{
      "msg": "An error occurs when the delete is performed",
      })
      return
    }
    msg := fmt.Sprintf("Delete blog%d successful %d", id, ra)
    c.JSON(http.StatusOK, gin.H{
      "msg": msg,
    })
 
 }
