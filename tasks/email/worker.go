// tasks/worker.go
package email

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hibiken/asynq"
)

// HandleWelcomeEmailTask 处理发送欢迎邮件的任务
func HandleWelcomeEmailTask(ctx context.Context, t *asynq.Task) error {
	var p WelcomeEmailPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("failed to unmarshal payload: %v", err)
	}

	// 模拟发送邮件的逻辑
	fmt.Printf("Sending welcome email to user ID: %d\n", p.UserID)
	return nil
}
