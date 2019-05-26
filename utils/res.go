package utils




type Result struct{
	Code int `json:"code"`
	Msg string `json:"msg"`
	Data interface{} `json:"data"`
}


func newRes(code int, msg string, data...interface{}) (int,Result){
	if len(data)>0{
		return code,Result{
			Code: code,
			Msg: msg,
			Data: data[0],
		}
	}
	return code, Result{
		Code: code,
		Msg: msg,
	}
}

func Success(msg string, data...interface{}) (int,Result){
	return newRes(200,msg,data)
}

func Fail(msg string) (int,Result){
	return newRes(400,msg)
}