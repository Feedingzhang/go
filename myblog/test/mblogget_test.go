package test
import(
"testing"
m "myblog/models"
"time"
)

func TestGet(t *testing.T){
	t1,_:=time.Parse("2006-01-02 15:04:05","2017-12-02 08:34:46")
	t2,_:=time.Parse("2006-01-02 15:04:05","2017-12-02 08:34:46")
	t3,_:=time.Parse("2006-01-02 15:04:05","2017-11-28 00:00:00")
	t4,_:=time.Parse("2006-01-02 15:04:05","2017-11-29 00:00:00")
	var tests = []struct {
		want  m.Blog
        id    int
    }{
      {want:m.Blog{8,"的是非得失","三方法的环境",t1},id:8},
      {want:m.Blog{14,"存放的贵航股份","而规划根据官方",t2},id:14},
      {want:m.Blog{3,"经济","似懂非懂刚发的的顺丰到付鼎折覆餗防守打法地方道德规范",t3},id:3},
      {want:m.Blog{2,"民生","群众脱贫",t4},id:2},
      {want:m.Blog{},id:50},
      {want:m.Blog{},id:51},}
    blog:=&m.Blog{}
    for _,test := range tests {
		blog.Id=test.id
        if got,_:= blog.GetBlog(); got != test.want{
            t.Errorf("GetBlog = %d",got.Id)
        }
    }
}