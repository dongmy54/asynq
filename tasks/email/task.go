// tasks/task.go
package email

import (
	"encoding/json"

	"github.com/hibiken/asynq"
)

// 定义任务类型常量
const (
	TypeWelcomeEmail = "email:welcome"
)

// WelcomeEmailPayload 定义任务所需的数据结构
type WelcomeEmailPayload struct {
	UserID int
}

// NewWelcomeEmailTask 创建一个发送欢迎邮件的任务
func NewWelcomeEmailTask(userID int) (*asynq.Task, error) {
	payload, err := json.Marshal(WelcomeEmailPayload{UserID: userID})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeWelcomeEmail, payload), nil
}
