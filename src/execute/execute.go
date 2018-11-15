package execute

import (
	"demo/src/model"
	"fmt"
)

//执行器

func Execute(task model.Task) {
	fmt.Println(task.Content)
}