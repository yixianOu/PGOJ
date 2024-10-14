### 1. "列出判题记录"

1. route definition

- Url: /api4/judge
- Method: GET
- Request: `ListJudgeStatusRequest`
- Response: `ListJudgeStatusResponse`

2. request definition



```golang
type ListJudgeStatusRequest struct {
	Page int64 `form:"page,default=1"`
	Limit int64 `form:"limit,default=10"`
	UserId int64 `form:"user_id"`
	ProblemId int64 `form:"problem_id,optinal"`
	Result string `form:"result,optional"`
	Language string `form:"language,optional"`
	SubmitTime int64 `form:"submit_time,optional"`
	Contest int64 `form:"contest,optional"`
	ProblemTitle string `form:"problem_title,optional"`
	Order bool `form:"order,optinal"`
}
```


3. response definition



```golang
type ListJudgeStatusResponse struct {
	JudgeStatus []JudgeStatus `json:"judge_status"`
}
```

### 2. "发送判题请求"

1. route definition

- Url: /api4/judge
- Method: POST
- Request: `SendJudgeStatusRequest`
- Response: `SendJudgeStatusResponse`

2. request definition



```golang
type SendJudgeStatusRequest struct {
	ProblemId int64 `form:"problem_id"`
	ProblemTitle string `form:"problem_title,optional"`
	Oj string `form:"oj,optional"`
	Language string `form:"language"`
	Code string `form:"code"`
	Length int64 `form:"length,optinal"`
	Contest int64 `form:"contest,default=0"`
	ContestProblem int64 `form:"contest_problem,default=0"`
	Rating int64 `form:"rating,optinal"`
	Ip string `form:"ip,optional"`
}
```


3. response definition



```golang
type SendJudgeStatusResponse struct {
	JudgeStatus JudgeStatus `json:"judge_status"`
}

type JudgeStatus struct {
	JudgeId int64 
	UserId int64 
	ProblemId int64 
	ProblemTitle string 
	Oj string 
	Result string 
	TimeCost int64 
	MemoryCost int64 
	Length int64 
	Language string 
	SubmitTime int64 
	Judger string 
	Contest int64 
	ContestProblem int64 
	Code string 
	TestCases string 
	Message string 
	Rating int64 
	Ip string 
	InputData string 
	SampleOutput string 
	UserOutput string 
}
```

### 3. "删除判题记录"

1. route definition

- Url: /api4/judge/:judge_id
- Method: DELETE
- Request: `DeleteJudgeStatusRequest`
- Response: `DeleteJudgeStatusResponse`

2. request definition



```golang
type DeleteJudgeStatusRequest struct {
	JudgeId int64 `path:"judge_id"`
}
```


3. response definition



```golang
type DeleteJudgeStatusResponse struct {
}
```

### 4. "获取判题记录"

1. route definition

- Url: /api4/judge/:judge_id
- Method: GET
- Request: `GetJudgeStatusRequest`
- Response: `GetJudgeStatusResponse`

2. request definition



```golang
type GetJudgeStatusRequest struct {
	JudgeId int64 `path:"judge_id"`
}
```


3. response definition



```golang
type GetJudgeStatusResponse struct {
	JudgeStatus JudgeStatus `json:"judge_status"`
}

type JudgeStatus struct {
	JudgeId int64 
	UserId int64 
	ProblemId int64 
	ProblemTitle string 
	Oj string 
	Result string 
	TimeCost int64 
	MemoryCost int64 
	Length int64 
	Language string 
	SubmitTime int64 
	Judger string 
	Contest int64 
	ContestProblem int64 
	Code string 
	TestCases string 
	Message string 
	Rating int64 
	Ip string 
	InputData string 
	SampleOutput string 
	UserOutput string 
}
```

