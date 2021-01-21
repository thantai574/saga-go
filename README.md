# Go Saga Pattern 
#### Introduction 
Implement golang for distributed transaction 

##### Installation 
```shell
go get github.com/thantai574/saga-go
```
##### Import 
```go
import github.com/thantai574/saga-go/saga_go
```
##### Example 
```go
import github.com/thantai574/saga-go/saga_go

sg := saga_go.NewSaga(Options{
TypeSaga: TypeMemory,
})

c := context.TODO()
sg.AddStep(entities.Step{
    Func: func(ctx context.Context) error {
    c = context.WithValue(c, "step1ctx", "step1ctx")
    return tt.errorStep[0]
},
Compensate: func(ctx context.Context) error {
    fmt.Println("Cancel Step 1 ")
    return nil
    },
})

sg.AddStep(entities.Step{
    Func: func(ctx context.Context) error {
    fmt.Println(c.Value("step1ctx"))
    return tt.errorStep[1]
},
Compensate: func(ctx context.Context) error {
    fmt.Println("Cancel Step 2 ")
    return nil
},
})

coor := NewCoordinator(c, sg)

coor.Start()
```
##