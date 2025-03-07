package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string
	Age  int
	Sex  bool
}

type Student struct {
	User
	Score int
}

func (u User) SayHello() {
	fmt.Println("Hello, my name is", u.Name)
}

func (u *User) SetName(name string) {
	u.Name = name
	fmt.Println("SetName", name)
}

func testReflect() {
	checkType := func(p any) {
		t := reflect.TypeOf(p)
		v := reflect.ValueOf(p)
		fmt.Println("Type:", t.Name())
		fmt.Println("Value:", v)

		// 检查是否为结构体
		if t.Kind() == reflect.Struct {
			// 获取字段
			for i := 0; i < t.NumField(); i++ {
				field := t.Field(i)
				fieldValue := v.Field(i)
				fmt.Printf("Field: %s, Type: %v, Value: %v\n",
					field.Name, field.Type, fieldValue)
			}

			// 获取值接收器的方法
			fmt.Println("\nValue Receiver Methods:")
			for i := 0; i < t.NumMethod(); i++ {
				method := t.Method(i)
				fmt.Printf("Method: %s\n", method.Name)
			}

			// 获取指针接收器的方法
			fmt.Println("\nPointer Receiver Methods:")
			// 创建指针类型
			ptrType := reflect.TypeOf((*User)(nil))
			for i := 0; i < ptrType.NumMethod(); i++ {
				method := ptrType.Method(i)
				fmt.Printf("Method: %s\n", method.Name)
			}
		} else {
			fmt.Printf("Type %s is not a struct\n", t.Name())
		}
		fmt.Println("--------------------------------")
	}

	// checkType(1)
	// checkType("hello")
	// checkType(User{Name: "John", Age: 20, Sex: true})
	checkType(Student{User: User{Name: "John", Age: 20, Sex: true}, Score: 100})
}

func testReflectModifyStruct() {
	stu := Student{User: User{Name: "John", Age: 20, Sex: true}, Score: 100}
	// SayHello 是值接收器定义的方法，可以用 stu 或 &stu 来调用
	reflect.ValueOf(stu).MethodByName("SayHello").Call(nil)
	// SetName 是指针接收器定义的方法，必须用 &stu 来调用
	reflect.ValueOf(&stu).MethodByName("SetName").Call([]reflect.Value{reflect.ValueOf("Tom1")})
	reflect.ValueOf(&stu).MethodByName("SayHello").Call(nil)
	reflect.ValueOf(&stu).MethodByName("SetName").Call([]reflect.Value{reflect.ValueOf("Tom2")})
	reflect.ValueOf(&stu).MethodByName("SayHello").Call(nil)
	// 修改 Score 字段
	reflect.ValueOf(&stu).Elem().FieldByName("Score").SetInt(88)
	// 修改 User.Age 字段
	reflect.ValueOf(&stu).Elem().FieldByName("Age").SetInt(21)
	fmt.Println(stu)
}

func main() {
	fmt.Println("====== testReflect =======")
	testReflect()
	fmt.Println("====== testReflectModifyStruct =======")
	testReflectModifyStruct()
}
