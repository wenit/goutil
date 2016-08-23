package maputil

import (
	"strconv"
)


/**
根据key获取map对象中的数据v，如果v不存在，返回默认值def
@param	m	map对象
@param	key	获取map数据的key
@param	def	默认值
@return interface{}
 */
func GetValue(m map[string]interface{}, key string, def... interface{}) (interface{}) {
	v, ok := m[key]
	if ok {
		return v
	} else {
		if len(def) == 0 {
			return nil
		}
		return def[0]
	}
}

/**
根据key获取map对象中的string类型数据v，如果v不存在，返回默认值def
@param	m	map对象
@param	key	获取map数据的key
@param	def	默认值
@return string
 */
func GetStringValue(m map[string]interface{}, key string, def... interface{}) string {
	v := GetValue(m, key, def...)
	b := v.(string)
	return b
}

/**
根据key获取map对象中的int类型数据v，如果v不存在，返回默认值def
@param	m	map对象
@param	key	获取map数据的key
@param	def	默认值
@return int
 */
func GetIntValue(m map[string]interface{}, key string, def... interface{}) int {
	v := GetValue(m, key, def...)
	switch v.(type) {
	case string:
		i, _ := strconv.Atoi(v.(string))
		return i
	case int:
		return v.(int)
	default:
		if len(def) == 0 {
			return 0
		} else {
			return (def[0]).(int)
		}
	}
}

/**
根据key获取map对象中的Float类型数据v，如果v不存在，返回默认值def
@param	m	map对象
@param	key	获取map数据的key
@param	def	默认值
@return float64
 */
func GetFloatValue(m map[string]interface{}, key string, def... interface{}) float64 {
	v := GetValue(m, key, def...)
	switch v.(type) {
	case string:
		i, _ := strconv.ParseFloat(v.(string), 64)
		return i
	case float64:
		return v.(float64)
	case int:
		i, _ := strconv.ParseFloat(strconv.Itoa(v.(int)), 64)
		return i
	default:
		if len(def) == 0 {
			return 0
		} else {
			return (def[0]).(float64)
		}
	}
}

/**
根据key获取map对象中的Bool类型数据v，如果v不存在，返回默认值def
@param	m	map对象
@param	key	获取map数据的key
@param	def	默认值
@return bool
 */
func GetBoolValue(m map[string]interface{}, key string, def... interface{}) bool {
	v := GetValue(m, key, def...)
	switch v.(type) {
	case string:
		i, _ := strconv.ParseBool(v.(string))
		return i
	case bool:
		return v.(bool)
	case int:
		i, _ := strconv.ParseBool(strconv.Itoa(v.(int)))
		return i
	default:
		if len(def) == 0 {
			return false
		} else {
			return (def[0]).(bool)
		}
	}
	b, _ := strconv.ParseBool(v.(string))
	return b
}