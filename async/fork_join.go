package async

import "context"

func ForkJoin(ctx context.Context, tasks []Task)  Task{
	return Invoke(ctx, func(context.Context) (i interface{}, e error) {
		for _,task := range tasks{
			task.Run(ctx)
		}
		WaitAll(tasks)
		return nil, nil
	})
}

func WaitAll(tasks []Task)  {
	for _,task := range tasks  {
		if task != nil{
			_,_ = task.Outcome()
		}
	}
}

func CancelAll(tasks []Task)  {
	for _, task := range tasks{
		task.Cancel()
	}
}