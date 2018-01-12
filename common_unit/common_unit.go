package common_unit

import (
    "fmt"
    "strings"
    "io/ioutil"
    "crypto/md5"
    "encoding/hex"
    "os"
    )

var debug int = 1

type BaseJsonBean struct {  
    Code    int         `json:"code"`  
    Data    interface{} `json:"data"`  
    Message string      `json:"message"`  
}  
  
func NewBaseJsonBean() *BaseJsonBean {  
    return &BaseJsonBean{}  
} 

func Cal_md5(string_need_to_cal_md5 string)(md5_string string){
    md5Ctx := md5.New()
    md5Ctx.Write([]byte(string_need_to_cal_md5))
    cipherStr := md5Ctx.Sum(nil)
    md5_string = hex.EncodeToString(cipherStr)
    return
    // 计算md5
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
    // 制造一个带索引的数组结构，其中每一个key是用户名，value是密码
    // if debug == 1{
    //     for username,password := range user_list{
    //         fmt.Println(username,password)
    //     } 
    // }
    return
}

func Sign_to_file(content string) error {
   // 以只写的模式，打开文件
    f, err := os.OpenFile("./common_unit/user_account.txt", os.O_WRONLY, 0644)
    if err != nil {
        fmt.Println("cacheFileList.yml file create failed. err: " + err.Error())
    } else {
        // 查找文件末尾的偏移量
        n, _ := f.Seek(0, os.SEEK_END)
        // 从末尾的偏移量开始写入内容
        _, err = f.WriteAt([]byte(content), n)
    }   
    defer f.Close()   
    return err
}









