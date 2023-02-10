package verify

import (
	p "DB/psql"
	"encoding/json"
	"fmt"
	"strings"
)
// Get PHM DB data
func phm() []string { 

	var deviceInfo []device_info
	device_id := make([]string, 0)
	p.DBC.Select("device_id").Find(&deviceInfo)
	
	for _, c := range deviceInfo {
		device_id = append(device_id, c.Device_id)
	}
	return device_id
}

func evo() []string{

	var assets []asset
	var res phm_info
	id := make([]string,0)
	p.DBB.Select("phm_config").Find(&assets)

	for _, b := range assets {
		json.Unmarshal([]byte(b.Phm_config), &res)
		resp, err := json.Marshal(res.Phm_info.Device_id)
			if err != nil {
			fmt.Println(err)
		}
		respo := strings.Trim(string(resp), `"`)
		id = append(id ,respo)
	}
	return id
}
func Check() bool {

	evo_data, phm_data := evo(), phm()

	if len(evo_data) != len(phm_data) {
		fmt.Println("The database have different number of devices")
	}

	evo_more := make([]string, 0)
	phm_more := make([]string, 0)
	
	for _, v := range evo_data {
		found := false
		for _, c := range phm_data {
			if v == c {
				found = true
				break
			}
		}
		if !found {
			evo_more = append(evo_more, v)
		}	
	}

	for _, v := range phm_data {
		found := false
		for _, c := range evo_data {
			if v == c {
				found = true
				break
			}
		}
		if !found {
			phm_more = append(phm_more, v)
		}	
	}
	fmt.Printf("phm lack %v\n", evo_more)
	fmt.Printf("evo lack %v\n", phm_more)
	
	return true
}