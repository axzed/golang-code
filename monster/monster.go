package monster

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Monster struct {
	Name  string
	Age   int
	Skill string
}

func (this *Monster) Store() bool {
	data, err := json.Marshal(this)
	if err != nil {
		fmt.Println("marshal err =", err)
		return false
	}

	// 保存到文件
	filePath := "C:\\Users\\薛文朝\\Desktop\\GolangCode\\src\\monster\\monster.ser"
	err = ioutil.WriteFile(filePath, data, 0666)
	if err != nil {
		fmt.Println("write file err =", err)
		return false
	}
	return true;
}

func (this *Monster) ReStore() bool {
	filePath := "C:\\Users\\薛文朝\\Desktop\\GolangCode\\src\\monster\\monster.ser"
	_, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Read file err = ", err)
		return false
	}
	return true
}
