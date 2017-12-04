package models

import (
 "log"
 db "myblog/database"
 "time"
)

type Blog struct {
 Id        int    `json:"id" form:"id"`
 Title  string `json:"title" form:"title"`
 Content   string `json:"content " form:"content "`
 Created   time.Time `json:"created " form:"created"`
}

type Blogs interface{
  AddBlog()(id int64, err error)
  GetBlogs() (blogs []Blog, err error) 
  GetBlog()(blog Blog, err error)
  UpdateBlog() (ra int64, err error)
  DelBlog()(ra int64, err error)
}




func (b *Blog) AddBlog() (id int64, err error) {
  if b!=nil && b.Title!="" && b.Content!="" && b.Title!=" " && b.Content!=" "{   
     rs, err := db.SqlDB.Exec("INSERT INTO blog(title,content,created) VALUES (?,?,?)",
      b.Title , b.Content,b.Created)
  if err != nil {
     log.Fatalln(err)
     return 0,err
  }
 //id=0 int64
 id, err = rs.LastInsertId()
 id=1
 return id,err
 }
 return 
}





func (b *Blog) GetBlogs() (blogs []Blog, err error) {
   blogs = make([]Blog, 0)
   rows, err := db.SqlDB.Query("SELECT id, title,content,created FROM blog")
  // defer rows.Close()
   if err != nil {
     return
   }

   for rows.Next() {
     var blog Blog
     rows.Scan(&blog.Id, &blog.Title, &blog.Content,&blog.Created)
     blogs = append(blogs, blog)
   }
   if err = rows.Err(); err != nil {
        return
   }
   return
}





func (b *Blog) GetBlog()(blog Blog, err error){
  err= db.SqlDB.QueryRow("SELECT id, title,content,created FROM blog WHERE id=?", b.Id).Scan(
   &blog.Id, &blog.Title, &blog.Content,&blog.Created,)
  if err != nil {
   return
 }

 return
}






func (b *Blog) UpdateBlog() (ra int64, err error) {
  if b!=nil && b.Title!="" && b.Content!="" && b.Title!=" " && b.Content!=" "{
     stmt, err:= db.SqlDB.Prepare("UPDATE blog SET title=?,content=?,created=? WHERE id=?")
     rs,err:= stmt.Exec(b.Title,b.Content,b.Created,b.Id)
     if err != nil {
        log.Fatalln(err)
        return 0,err
     }
     ra, err= rs.RowsAffected()
    if err != nil {
        log.Fatalln(err)
        return 0,err
     }
     return ra,err
 }
   return 0,err
}










func (b *Blog) DelBlog()(ra int64, err error){
  rs, err := db.SqlDB.Exec("DELETE FROM blog WHERE id=?", b.Id)
  if err != nil {
      log.Fatalln(err)
  }

  ra,err= rs.RowsAffected()
  if err != nil {
     log.Fatalln(err)
     return 
  }
  return 
  }
