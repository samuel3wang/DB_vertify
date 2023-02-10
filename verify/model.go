package verify

// 對應 DB Table Column
type device_info struct {
	Device_id		string	`json:"device_id"`
}

type asset struct {
	Phm_config		string	`json:"phm_config"`	
}
type phm_info struct {
	Phm_info struct{
		Device_id 	string `json:"device_id"`
	} `json:"phm_info"`
}