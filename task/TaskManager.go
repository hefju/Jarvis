package task
import (
    "fmt"
    "strings"
    "github.com/hefju/Jarvis/models"
)
type  TaskManager struct  {

}


func (x TaskManager)Listen(command string ){
    command=strings.ToLower(command)
    switch command {
        case "what to do today":

        fmt.Println("ok, wait a minute")
        tasks:=models.GetTodayTasks()
        if len(tasks)==0{
            fmt.Println("没有任务.")
        }
        default:
        fmt.Println("i don't understand this command:",command)
    }
}