package controllers

import (
    "fmt"
    "net/http"
    // "time"
    // "encoding/json"
    "inet_server/common_unit"
    // "os"
    "io/ioutil"
    )

func Index_interface(w http.ResponseWriter, req *http.Request) {
    fmt.Println("somebody visits index")  
    //模拟延时  
    // time.Sleep(time.Second * 2)  
    //获取客户端通过POST方式传递的参数  
    req.ParseForm() 
    data, err := ioutil.ReadFile("./views/index.html")
    if err != nil{
        logging := "get index fail"
        common_unit.Write_log(logging)
        fmt.Println(logging)
        fmt.Fprint(w,"open index error")
    }
    var index_content string = string(data)
    // bytes, _ := json.Marshal(result)  
    fmt.Fprint(w, index_content)  
    //向客户端返回JSON数据    
}