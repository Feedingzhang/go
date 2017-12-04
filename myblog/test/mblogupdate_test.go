package test
import(
"testing"
m "myblog/models"
"time"
)

func TestUpdate(t *testing.T){
	var tests = []struct {
        blog  m.Blog
        want  int64
    }{
        {blog:m.Blog{20,"的是非得失","三方法的环境",time.Now()},want:1},
        {blog:m.Blog{21,"存放的贵航股份","而规划根据官方",time.Now()},want:1},
        {blog:m.Blog{22,"大锅饭的环境","吃个饭规划局建行卡",time.Now()},want:1},
        {blog:m.Blog{19,"东方化工话就会","的各个房间号接口",time.Now()},want:1},
        {blog:m.Blog{90,"认为太热与他人","返回集合快回家",time.Now()},want:0},
        {blog:m.Blog{20,"","三方法的环境",time.Now()},want:0},
        {blog:m.Blog{20,"","",time.Now()},want:0},
        {blog:m.Blog{20,"的是非得失","",time.Now()},want:0},}

        b:=m.Blog{}
	for _,test := range tests {
		b=test.blog
        if ra,_:= b.UpdateBlog(); ra != test.want{
            t.Errorf("UpdateBlog = %d",ra)
        }
    }
 }