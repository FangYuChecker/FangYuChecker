package Unit

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Monster struct {
	Name string `json:"name"`
	Age int	`json:"age"`
	Skill string `json:"skill"`
}

func (this *Monster) Store() bool{
	//先序列化
	data,err := json.Marshal(this)
	if err != nil{
		fmt.Println("Marshal Error !",err)
		return false
	}
	//保存文件
	//打开文件
	filepath := "/Users/fangyuchecker/Desktop/GoProject/src/UnitTest/Unit/Monster.json"
	err = ioutil.WriteFile(filepath,data,0666)
	if err != nil {
		fmt.Println("File Open Error",err)
		return false
	}
	//写入文件
	return true
}

func (this *Monster) Restore() bool{
	//读取文件
	filepath := "/Users/fangyuchecker/Desktop/GoProject/src/UnitTest/Unit/Monster.json"
	data,err := ioutil.ReadFile(filepath)
	if err != nil{
		fmt.Println("read file error !",err)
		return false
	}

	//反序列化
	err = json.Unmarshal([]byte(data),this)
	if err != nil {
		fmt.Println("Unmarshal fail !")
		return false
	}
	return true
}
