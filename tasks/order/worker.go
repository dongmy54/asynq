package order

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/hibiken/asynq"
)

// 消息处理部分
type MyTestHandler struct {
}

func (mth *MyTestHandler) ProcessTask(ctx context.Context, t *asynq.Task) error {
	var mt MyTest
	if err := json.Unmarshal(t.Payload(), &mt); err != nil {
		return errors.New("发生错误")
	}

	fmt.Println("收到消息=====", mt)
	return nil
}
