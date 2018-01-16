package controllers

import (
    "fmt"
    "net/http"
    // "time"
    // "encoding/json"
    "inet_server/common_unit"
    "io/ioutil"
    )

func Login_interface(w http.ResponseWriter, req *http.Request) {
    var user_accounts = common_unit.Get_user_list()  
    fmt.Println("a user login")  
    //模拟延时  
    // time.Sleep(time.Second * 2)  
    //获取客户端通过POST方式传递的参数  
    req.ParseForm()  
    result := common_unit.NewBaseJsonBean() 
    username_content, _ := req.Form["username"]  
    password_content, _ := req.Form["password"] 
    fmt.Println(username_content,password_content)


    if username_content[0] == "" || password_content[0] == "" { 
        result.Code = -1
        result.Message = "登录数据不能为空"
        // 填写返回数据
        // result_bytes, _ := json.Marshal(result)
        // fmt.Println(result_bytes) 
        // 格式化返回数据 
        // fmt.Fprint(w, string(result_bytes))
        logging := "登录信息为空，登录失败"
        common_unit.Write_log(logging)
        // 之后自动跳转
        data, _ := ioutil.ReadFile("./views/login_fail.html")
        var content string = string(data)
        fmt.Fprint(w, content) 
        fmt.Println(logging)  
        return  
    }  
 
    username := username_content[0]  
    password := password_content[0]  
    
    if user_accounts[username] == string(common_unit.Cal_md5(password)){  
        result.Code = 100  
        result.Message = "登录成功" 
        logging := "登陆成功，用户名为"+username+"密码MD5为"+common_unit.Cal_md5(password)
        common_unit.Write_log(logging)
        fmt.Println(logging) 
        data, err := ioutil.ReadFile("./views/login_success.html")
        if err != nil{
            logging = "找不到主页"
            fmt.Println(logging)
            common_unit.Write_log(logging)
        } else {
            var content string = string(data)
            fmt.Fprint(w, content) 
        }
 
    } else {  
        result.Code = 101  
        result.Message = "用户名或密码不正确"  
        logging := "用户"+username+"尝试登陆，并登录失败"
        common_unit.Write_log(logging)
        // 之后自动跳转
        data, _ := ioutil.ReadFile("./views/login_fail.html")
        var content string = string(data)
        fmt.Fprint(w, content) 
        fmt.Println(logging)   
    }
    // bytes, _ := json.Marshal(result)  
    // fmt.Fprint(w, string(bytes))  
    //向客户端返回JSON数据    
}