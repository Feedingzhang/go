package test
import(
"testing"
m "myblog/models"
)
func TestDelete(t *testing.T){
	var tests = []struct {
		want int64
        id    int
    }{
    	{want:1,id:26},
    	{want:1,id:27},
    	{want:1,id:25},
    	{want:0,id:88},
    	{want:0,id:90},}
    	blog:=&m.Blog{}
	for _,test := range tests {
		blog.Id=test.id
        if ra,_:= blog.DelBlog(); ra != test.want{
            t.Errorf("DelBlog = %d",ra)
        }
    }
}