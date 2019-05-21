package utils




type Result struct{
	Code int `json:"code"`
	Msg string `json:"msg"`
	Data interface{} `json:"data"`
}


func NewRes(code int, msg string, data...interface{})(int, Result){
	if len(data)>0{
		return 200,Result{
			Code: code,
			Msg: msg,
			Data: data[0],
		}
	}
	return 200, Result{
		Code: code,
		Msg: msg,

	}
}