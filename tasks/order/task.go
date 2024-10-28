// tasks/task.go
package order

import (
	"encoding/json"

	"github.com/hibiken/asynq"
)

// 定义任务类型常量
const (
	TypeMyTest = "critical:mytest"
)

// 消息内容结构体
type MyTest struct {
	Info string
}

// 生成消息方法
func NewMyTestMessage(info string) *asynq.Task {
	payload, _ := json.Marshal(MyTest{Info: info})
	return asynq.NewTask(TypeMyTest, payload)
}
