/*
JSON操作工具类
 */
package json

import "encoding/json"


/*
将字符串转换成实体对象
示例代码：
	student := Student{}
	fmt.Println("初始：", student)
	json := `{"name":"章文兵","age":27,"Classes":["Math","English","Chinese"]}`
	ToJson(json, &student)
	fmt.Println("转换后：", student)
*/
func ToJson(s string, v interface{}) error {
	err := json.Unmarshal([]byte(s), v)
	return err
}

/*
将对象转换成字符串
示例代码：
	student := Student{Person:Person{"章文兵", 27}, Classes:[]string{"Math", "English", "Chinese"}, }
	json := JsonToString(student)
	fmt.Println(json)
 */
func JsonToString(v interface{}) (string) {
	bt, _ := json.Marshal(v)
	return string(bt)
}