### 1. "添加测试用例"

1. route definition

- Url: /api3/test_cases
- Method: POST
- Request: `AddTestCaseRequest`
- Response: `AddTestCaseResponse`

2. request definition



```golang
type AddTestCaseRequest struct {
	ProblemId int64 `form:"problem_id,range=[1:]"`
	TestGroup int64 `form:"test_group"`
}
```


3. response definition



```golang
type AddTestCaseResponse struct {
}
```

### 2. "删除测试用例"

1. route definition

- Url: /api3/test_cases/delete
- Method: POST
- Request: `DeleteTestCaseRequest`
- Response: `DeleteTestCaseResponse`

2. request definition



```golang
type DeleteTestCaseRequest struct {
	TestId int64 `form:"test_id,range=[1:]"`
}
```


3. response definition



```golang
type DeleteTestCaseResponse struct {
}
```

### 3. "更新用户头像"

1. route definition

- Url: /api3/users/:user_id/cover
- Method: POST
- Request: `UpdateUserCoverRequest`
- Response: `UpdateUserResponse`

2. request definition



```golang
type UpdateUserCoverRequest struct {
	ID int64 `path:"user_id,range=[1:]"`
}
```


3. response definition



```golang
type UpdateUserResponse struct {
}
```

