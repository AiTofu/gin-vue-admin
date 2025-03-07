package main

import (
	"fmt"
	"sync"
)

type User struct {
	Name string
	Age  int
	Sex  bool
}

// 嵌入方式（当前方式）
type Student struct {
	User  // 嵌入 User 结构体
	Score int
}

// 字段方式（对比用）
type StudentWithField struct {
	User  User // 创建一个名为 User 的字段
	Score int
}

// 值接收器方法
func (u User) Say() {
	fmt.Printf("值接收器 Say() - 地址: %p, 名字: %s\n", &u, u.Name)
}

// 指针接收器方法
func (u *User) PrintAge() {
	fmt.Printf("指针接收器 PrintAge() - 地址: %p, 年龄: %d\n", u, u.Age)
}

// 值接收器方法 - 修改不会影响原对象
func (u User) ChangeName(newName string) {
	u.Name = newName
	fmt.Printf("值接收器 ChangeName() - 地址: %p, 新名字: %s\n", &u, u.Name)
}

// 指针接收器方法 - 修改会影响原对象
func (u *User) ChangeAge(newAge int) {
	u.Age = newAge
	fmt.Printf("指针接收器 ChangeAge() - 地址: %p, 新年龄: %d\n", u, u.Age)
}

// 1. 不可变对象示例
type ImmutablePoint struct {
	X, Y int
}

// 使用值接收器，确保点不会被修改
func (p ImmutablePoint) GetX() int {
	return p.X
}

func (p ImmutablePoint) GetY() int {
	return p.Y
}

// 2. 并发安全示例
type Counter struct {
	value int
	mu    sync.Mutex
}

// 使用指针接收器，因为包含互斥锁
func (c *Counter) GetValue() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

// 使用指针接收器，因为需要修改共享状态
func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

// 原子操作：读取并增加
func (c *Counter) GetAndIncrement() (int, int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	oldValue := c.value
	c.value++
	return oldValue, c.value
}

func (c *Counter) Decrement() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value--
}

// 3. 小对象示例
type RGB struct {
	R, G, B uint8
}

// 使用值接收器，因为RGB结构体很小，复制开销可以忽略
func (c RGB) IsBlack() bool {
	return c.R == 0 && c.G == 0 && c.B == 0
}

// 4. 实际业务场景：用户权限检查
type UserRole struct {
	Role     string
	Level    int
	IsActive bool
}

// 使用值接收器进行权限检查，不需要修改对象
func (r UserRole) CanAccessAdmin() bool {
	return r.Role == "admin" && r.Level >= 3 && r.IsActive
}

func (r UserRole) CanEditContent() bool {
	return r.Role == "editor" && r.Level >= 2 && r.IsActive
}

func testUser() {
	user := User{
		Name: "John",
		Age:  20,
	}

	fmt.Printf("原始对象地址: %p\n", &user)

	// 测试值接收器方法
	user.Say()
	user.ChangeName("Mike")
	fmt.Printf("值接收器修改后，原对象名字: %s\n", user.Name)

	// 测试指针接收器方法
	user.PrintAge()
	user.ChangeAge(25)
	fmt.Printf("指针接收器修改后，原对象年龄: %d\n", user.Age)
}

func demonstrateValueReceivers() {
	// 1. 不可变对象示例
	point := ImmutablePoint{X: 10, Y: 20}
	fmt.Printf("Point: (%d, %d)\n", point.GetX(), point.GetY())

	// 2. 并发安全示例
	counter := Counter{value: 100}

	// 多个goroutine同时读取和修改计数器
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			// 原子操作：读取并增加
			oldValue, newValue := counter.GetAndIncrement()
			fmt.Printf("%d::Counter value: %d -> %d\n", id, oldValue, newValue)
		}(i)
	}
	wg.Wait()

	// 3. 小对象示例
	color := RGB{R: 0, G: 0, B: 0}
	fmt.Printf("Is black: %v\n", color.IsBlack())

	// 4. 权限检查示例
	adminRole := UserRole{Role: "admin", Level: 3, IsActive: true}
	editorRole := UserRole{Role: "editor", Level: 2, IsActive: true}

	fmt.Printf("Admin can access admin panel: %v\n", adminRole.CanAccessAdmin())
	fmt.Printf("Editor can edit content: %v\n", editorRole.CanEditContent())
}

func testEmbedding() {
	// 测试嵌入方式
	student := Student{
		User:  User{Name: "John", Age: 20},
		Score: 100,
	}

	// 可以直接访问 User 的字段
	fmt.Printf("Student.Name: %s\n", student.Name) // 直接访问
	fmt.Printf("Student.Age: %d\n", student.Age)   // 直接访问
	fmt.Printf("Student.Score: %d\n", student.Score)

	// 也可以显式指定 User 字段
	fmt.Printf("Student.User.Name: %s\n", student.User.Name)
	fmt.Printf("Student.User.Age: %d\n", student.User.Age)

	// 测试字段方式
	studentWithField := StudentWithField{
		User:  User{Name: "Mike", Age: 21},
		Score: 90,
	}

	// 必须通过 User 字段访问
	fmt.Printf("StudentWithField.User.Name: %s\n", studentWithField.User.Name)
	fmt.Printf("StudentWithField.User.Age: %d\n", studentWithField.User.Age)
	fmt.Printf("StudentWithField.Score: %d\n", studentWithField.Score)
}

func testAssert() {
	checkType := func(v interface{}) {
		switch v.(type) {
		case int:
			fmt.Println("int")
		case User:
			fmt.Println("User")
		case Student:
			fmt.Println("Student")
		default:
			fmt.Println("unknown type")
		}
	}

	checkType(1)
	checkType(User{Name: "John", Age: 20})
	checkType(Student{User: User{Name: "John", Age: 20}, Score: 100})

	userFunc := func(v interface{}) {
		user, ok := v.(User)
		if ok {
			fmt.Println("User", user.Name)
		} else {
			fmt.Println("not User")
		}
	}

	userFunc(User{Name: "John", Age: 20})
	userFunc(1)
}

func main() {
	fmt.Println("====== testUser =======")
	testUser()
	fmt.Println("====== demonstrateValueReceivers =======")
	demonstrateValueReceivers()
	fmt.Println("====== testEmbedding =======")
	testEmbedding()
	fmt.Println("====== testAssert =======")
	testAssert()
}
