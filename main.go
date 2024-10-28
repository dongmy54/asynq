// main.go
package main

import (
	"fmt"
	"log"
	"time"

	"example/user/hello/tasks/order"

	"github.com/hibiken/asynq"
)

func main() {
	// 设置 Redis 客户端连接信息
	redisOpt := asynq.RedisClientOpt{Addr: "127.0.0.1:6379"}

	// 1. 创建任务生产者并发送任务
	client := asynq.NewClient(redisOpt)
	defer client.Close()

	task := order.NewMyTestMessage("mymymy") // 假设 UserID 为 42
	// 发送任务到队列中
	info, err := client.Enqueue(task)
	if err != nil {
		log.Fatalf("could not enqueue task: %v", err)
	}
	fmt.Printf("Enqueued task: id=%s queue=%s\n", info.ID, info.Queue)

	// 发送一个延时任务
	client.Enqueue(task, asynq.ProcessIn(10*time.Second), asynq.Queue("critical"))

	// 2. 创建任务消费者并启动任务处理
	srv := asynq.NewServer(redisOpt, asynq.Config{
		// Specify how many concurrent workers to use
		Concurrency: 10,
		// Optionally specify multiple queues with different priority.
		Queues: map[string]int{
			"critical": 6,
			"default":  3,
			"low":      1,
		},
		// See the godoc for other configuration options
	})

	mux := asynq.NewServeMux()
	mux.Handle(order.TypeMyTest, &order.MyTestHandler{})

	// 再启动一个消费服务呢
	srv1 := asynq.NewServer(redisOpt, asynq.Config{
		// Specify how many concurrent workers to use
		Concurrency: 15,
		// Optionally specify multiple queues with different priority.
		Queues: map[string]int{
			"critical": 6,
			"default":  3,
			"low":      1,
		},
		// See the godoc for other configuration options
	})

	mux1 := asynq.NewServeMux()
	mux1.Handle(order.TypeMyTest, &order.MyTestHandler{})

	go func() {
		srv1.Run(mux1)
	}()

	// 定时任务
	scheduler := asynq.NewScheduler(
		redisOpt, &asynq.SchedulerOpts{
			EnqueueErrorHandler: func(task *asynq.Task, opts []asynq.Option, err error) {
				fmt.Printf("Scheduler EnqueueErrorHandler <<<<<<<===>>>>> err : %+v , task : %+v", err, task)
			},
		})

	// every one minute exec
	entryID, _ := scheduler.Register("*/1 * * * *", task)
	go func() {
		scheduler.Run()
	}()

	fmt.Println("定时任务id:", entryID)

	if err := srv.Run(mux); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}
