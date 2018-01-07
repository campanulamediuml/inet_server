package common_unit

type BaseJsonBean struct {  
    Code    int         `json:"code"`  
    Data    interface{} `json:"data"`  
    Message string      `json:"message"`  
}  
  
func NewBaseJsonBean() *BaseJsonBean {  
    return &BaseJsonBean{}  
} 