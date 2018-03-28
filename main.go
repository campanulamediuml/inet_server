package main
 
import ( 
    "inet_server/controllers"  
    "fmt"  
    "net/http"   
)  

func WebServerBase() {  
    fmt.Println("服务器开始运行")  
    //第一个参数为客户端发起http请求时的接口名，第二个参数是一个func，负责处理这个请求。  
    http.HandleFunc("/go/login", controllers.Login_interface)
    http.HandleFunc("/go/sign", controllers.Sign_interface)
    http.HandleFunc("/go/index",controllers.Index_interface)
    //服务器要监听的主机地址和端口号  
    err := http.ListenAndServe(":80", nil)  
    if err != nil {  
        fmt.Println("ListenAndServe error: ", err.Error())  
    }  
}  

func main(){
    WebServerBase()
} 
