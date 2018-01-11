package common_unit

import (
    "fmt"
    "strings"
    "io/ioutil"
    // "strconv"
    )

type BaseJsonBean struct {  
    Code    int         `json:"code"`  
    Data    interface{} `json:"data"`  
    Message string      `json:"message"`  
}  
  
func NewBaseJsonBean() *BaseJsonBean {  
    return &BaseJsonBean{}  
} 

func Get_user_list()(user_list map[string]string){
    user_list = make(map[string]string)
    // 初始化用户名密码数组
    data, err := ioutil.ReadFile("./common_unit/user_account.txt")
    if err != nil{
        fmt.Println("get user accounts fail")
    } 
    var user_account_string string = string(data)
    // 读取保存了用户名的txt文件  
    user_account_list := strings.Split(user_account_string,"\n")
    // 通过回车键分割每一行
    for _,line := range user_account_list{
        account_info := strings.Split(line,"\t") 
        user_list[account_info[0]] = account_info[1]
    }
    fmt.Println(user_list)
    // 制造一个带索引的数组结构，其中每一个key是用户名，value是密码
    return
}
