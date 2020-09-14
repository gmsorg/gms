package common

/**
使用这个调用
{"service_name":"UserServiceImpl","method_name":"GetUser"}
{"service_name":"UserServiceImpl","method_name":"RegisterUser"}
*/
type ReqMessage struct {
	ServiceName string `json:"service_name"`
	MethodName  string `json:"method_name"`
	ReqData     []byte `json:"req_data"`
}

type ResMessage struct {
	Code    int    `json:"code"`
	Msg     string `json:"msg"`
	ResData []byte `json:"res_data"`
}
