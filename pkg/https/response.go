package https

type ResponseData struct {
	//返回编码
	Code int

	//返回信息
	Msg string

	//返回数据
	Data interface{}
}
