package controllers

import (
    "fmt"
    "net/http"
    // "time"
    "encoding/json"
    "inet_server/common_unit"
    )

func Sign_interface(w http.ResponseWriter, req *http.Request) {  
    var user_accounts = common_unit.Get_user_list()
    fmt.Println("a user sign in")  
    //模拟延时  
    // time.Sleep(time.Second * 2)  
    //获取客户端通过POST方式传递的参数  
    req.ParseForm()  
    result := common_unit.NewBaseJsonBean() 
    username_content, found1 := req.Form["username"]  
    password_content, found2 := req.Form["password"] 

    if !(found1 && found2) { 
        result.Code = -1
        result.Message = "注册数据不能为空"
        // 填写返回数据
        result_bytes, _ := json.Marshal(result) 
        // 格式化返回数据 
        fmt.Fprint(w, string(result_bytes))
        fmt.Println("注册数据为空，注册失败")  
        return  
    }  
    // 没有，瞎写的
 
    username := username_content[0]  
    password := password_content[0] 
    // 提取用户名和密码

    var user_exist int = 0
    for user_tmp,_ := range user_accounts{
        fmt.Println(user_tmp)
        if user_tmp == username{
            user_exist = 1
            break
        }
    }
    // 判断用户是否存在
    if user_exist == 1{
        result.Code = 101  
        result.Message = "用户名已存在，注册失败" 
        fmt.Println("注册失败，用户名",username,"已存在") 
        // 如果用户名已经存在则返回错误 
    }else{
        var user_token string = "\n"+username+"\t"+string(common_unit.Cal_md5(password))
        err := common_unit.Sign_to_file(user_token)
        if err != nil{
            fmt.Println("服务器内部错误")
        }
        result.Code = 100 
        result.Message = "注册成功" 
        fmt.Println("注册成功，用户名为",username,"密码MD5为",common_unit.Cal_md5(password)) 
    }
    // 用户名不存在则写入注册文件，ojbk
    bytes, _ := json.Marshal(result)  
    fmt.Fprint(w, string(bytes))  
    //向客户端返回JSON数据    
}