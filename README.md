# Go Saga Pattern 
#### Introduction 
Implement golang for distributed transaction 

##### Installation 
```shell
go get github.com/thantai574/saga-go
```
##### Import 
```go
import saga_go "github.com/thantai574/saga-go"
```
##### Example 
```go
package saga

import (
	"context"
	"fmt"
	saga_go "github.com/thantai574/saga-go"
	"github.com/thantai574/saga-go/entities"
)

func main() {
	sg := saga_go.NewSaga(saga_go.Options{
		TypeSaga: saga_go.TypeMemory,
	})

	c := context.TODO()
	sg.AddStep(entities.Step{
		Func: func(ctx context.Context) error {
			c = context.WithValue(c, "step1ctx", "step1ctx")
			return nil
		},
		Compensate: func(ctx context.Context) error {
			fmt.Println("Cancel Step 1 ")
			return nil
		},
	})

	sg.AddStep(entities.Step{
		Func: func(ctx context.Context) error {
			fmt.Println(c.Value("step1ctx"))
			return nil
		},
		Compensate: func(ctx context.Context) error {
			fmt.Println("Cancel Step 2 ")
			return nil
		},
	})

	coor := saga_go.NewCoordinator(c, sg)

	coor.Start()
}

```
##