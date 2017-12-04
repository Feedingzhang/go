package test
import(
"testing"
m "myblog/models"
)
func TestGets (t *testing.T){
	 blog:=m.Blog{}
	 got := make([]m.Blog, 0)
    if got,_= blog.GetBlogs(); got == nil{
            t.Errorf("GetsBlog is nil",)
        }
    
}