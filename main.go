package main
import ("fmt"
  //  "github.com/hefju/Utils/datetime"
    "github.com/gin-gonic/gin"
  //  "github.com/hefju/Jarvis/task"
//    "github.com/donnie4w/go-logger/logger"
    "net/http"
)
func main(){

    router := gin.Default()
    router.GET("/", func(c *gin.Context) {//测试，获取数据表信息
//        logger.Debug("visit homepage")
        c.String(http.StatusOK, "Welcome to Jarvis...")
    })

    router.GET("/today",GetTodayList)
    router.Run(":8084")

  //manager:=  task.TaskManager{}
    //manager.Listen("gogogo")
  //  manager.Listen("What to do today")

    fmt.Println("end")
}

func GetTodayList(c *gin.Context){
    c.String(http.StatusOK,"获取今日任务")
}

