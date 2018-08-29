package zapp

import "reflect"

func CallMethod(obj interface{}, methodName string, defaultValue interface{}) interface{} {
	method := GetMethod(obj, methodName)
	if method.IsValid() { // メソッドがある場合のみ実行
		response := method.Call(nil)[0]
		// nilじゃなければ
		if reflect.ValueOf(response.Interface()).IsValid() {
			return response.Interface()
		}
	}
	return defaultValue
}

func GetMethod(obj interface{}, methodName string) reflect.Value {
	r := reflect.ValueOf(obj)
	method := r.MethodByName(methodName) // これは個別に定義したいので回避
	return method
}

func HasMethod(obj interface{}, methodName string) bool {
	method := GetMethod(obj, methodName)
	if method.IsValid() { // メソッドがある場合のみ実行
		return true
	}
	return false
}
