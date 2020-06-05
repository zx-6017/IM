package Struct

type ImAction struct {
	State_StateChange string
	C2C_CallbackBeforeSendMsg string
}
func (this ImAction) CallbackBeforeSendMsg(){

}

type Data_struct struct {
	MsgBody []struct{
		MsgType string `json:"MsgType"`
		MsgContent struct{
			Desc string `json:"Desc"`
			Data string `json:"Data"`
			Ext string `json:"Ext"`
			Sound string `json:"Sound"`
		} `json:"MsgContent"`
	} `json:"MsgBody"`
	CallbackCommand string `json:"CallbackCommand"`
	From_Account string `json:"From_Account"`
	To_Account string `json:"To_Account"`
	MsgRandom int `json:"MsgRandom"`
	MsgSeq int `json:"MsgSeq"`
	MsgTime int `json:"MsgTime"`
}

type Msg_data struct {
	IsCorrection bool
	Revise string
	PrimaryText string
}