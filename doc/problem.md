### 1. "搜索题目数据"

1. route definition

- Url: /api2/problems_data
- Method: GET
- Request: `SearchProblemsDataRequest`
- Response: `SearchProblemsDataResponse`

2. request definition



```golang
type SearchProblemsDataRequest struct {
	Page int64 `form:"page,default=1"`
	PageSize int64 `form:"page_size,default=10"`
	ScoreFloor int64 `form:"score_floor,optional"`
	ScoreCeil int64 `form:"score_ceil,optional"`
	Auth int64 `form:"auth,optional"`
	Order bool `form:"order,optional"`
}
```


3. response definition



```golang
type SearchProblemsDataResponse struct {
	ProblemData []ProblemData `json:"problems_data"`
}
```

### 2. "获取题目数据"

1. route definition

- Url: /api2/problems_data/:problem_data_id
- Method: GET
- Request: `GetProblemDataRequest`
- Response: `GetProblemDataResponse`

2. request definition



```golang
type GetProblemDataRequest struct {
	ProblemDataId int64 `path:"problem_data_id"`
}
```


3. response definition



```golang
type GetProblemDataResponse struct {
	ProblemData ProblemData `json:"problem_data"`
}

type ProblemData struct {
	ProblemDataId int64 
	ProblemId int64 
	Submission int64 
	Accepted int64 
	MemoryLimitExceeded int64 
	TimeLimitExceeded int64 
	WrongAnswer int64 
	RuntimeError int64 
	CompileError int64 
	OutputLimitExceeded int64 
	UnknowError int64 
	SegmentFault int64 
	FloatError int64 
	Score int64 
	Auth int64 
}
```

### 3. "更新题目数据"

1. route definition

- Url: /api2/problems_data/update
- Method: POST
- Request: `UpdateProblemDataRequest`
- Response: `UpdateProblemDataResponse`

2. request definition



```golang
type UpdateProblemDataRequest struct {
	ProblemDataId int64 `form:"id"`
	Score int64 `form:"score,optional"`
	Auth int64 `form:"auth,optional"`
}
```


3. response definition



```golang
type UpdateProblemDataResponse struct {
}
```

### 4. "添加题目（为了保证problem和problem_data的id一致性，必须使用api添加题目，而不能直接在数据库中新增problem记录）"

1. route definition

- Url: /api2/problems
- Method: POST
- Request: `AddProblemRequest`
- Response: `AddProblemResponse`

2. request definition



```golang
type AddProblemRequest struct {
	Oj string `form:"oj,optional"`
	Title string `form:"title"`
	Description string `form:"description"`
	Input string `form:"input"`
	Output string `form:"output"`
	Hint string `form:"hint,optional"`
	Source string `form:"source,optional"`
	LimitTime int64 `form:"limit_time"`
	LimitMemory int64 `form:"limit_memory"`
	Auth int64 `form:"auth,default=1"`
	Level int64 `form:"level,options=[1,2,3,4,5]"`
	ProblemCode string `form:"problem_code,optional"`
	Score int64 `form:"score"`
	TagIds []int64 `form:"tag_ids"`
}
```


3. response definition



```golang
type AddProblemResponse struct {
}
```

### 5. "删除题目"

1. route definition

- Url: /api2/problems/delete
- Method: POST
- Request: `DeleteProblemRequest`
- Response: `DeleteProblemResponse`

2. request definition



```golang
type DeleteProblemRequest struct {
	ProblemId int64 `form:"problem_id"`
}
```


3. response definition



```golang
type DeleteProblemResponse struct {
}
```

### 6. "更新题目"

1. route definition

- Url: /api2/problems/update
- Method: POST
- Request: `UpdateProblemRequest`
- Response: `UpdateProblemResponse`

2. request definition



```golang
type UpdateProblemRequest struct {
	ProblemId int64 `form:"problem_id"`
	Author string `form:"author,optional"`
	Oj string `form:"oj,optional"`
	Title string `form:"title,optional"`
	Description string `form:"description,optional"`
	Input string `form:"input,optional"`
	Output string `form:"output,optional"`
	SampleInput string `form:"sample_input,optional"`
	SampleOutput string `form:"sample_output,optional"`
	Hint string `form:"hint,optional"`
	Source string `form:"source,optional"`
	LimitTime int64 `form:"limit_time,optional"`
	LimitMemory int64 `form:"limit_memory,optional"`
	Auth int64 `form:"auth,optional"`
	Level int64 `form:"level,options=[1,2,3,4,5],optional"`
	ProblemCode string `form:"problem_code,optional"`
	TestCount int64 `form:"test_count,optional"`
	TagIds []int64 `form:"tag_ids,optional"`
	TagOperation bool `form:"tag_operation,optional"`
}
```


3. response definition



```golang
type UpdateProblemResponse struct {
}
```

### 7. "搜索题目"

1. route definition

- Url: /api2/problems
- Method: GET
- Request: `SearchProblemsRequest`
- Response: `SearchProblemsResponse`

2. request definition



```golang
type SearchProblemsRequest struct {
	Page int64 `form:"page,default=1"`
	PageSize int64 `form:"page_size,default=10"`
	Author string `form:"author,optional"`
	Title string `form:"title,optional"`
	ProblemCode string `form:"problem_code,optional"`
	Oj string `form:"oj,optional"`
	Description string `form:"description,optional"`
	Source string `form:"source,optional"`
	Auth int64 `form:"auth,optional"`
	Level int64 `form:"level,options=[1,2,3,4,5],optional"`
	Order bool `form:"order,optional"`
}
```


3. response definition



```golang
type SearchProblemsResponse struct {
	Problems []Problem `json:"problems"`
}
```

### 8. "获取题目信息"

1. route definition

- Url: /api2/problems/:problem_id
- Method: GET
- Request: `GetProblemByIdRequest`
- Response: `GetProblemByIdResponse`

2. request definition



```golang
type GetProblemByIdRequest struct {
	ProblemId int64 `path:"problem_id"`
}
```


3. response definition



```golang
type GetProblemByIdResponse struct {
	Problem Problem `json:"problem"`
}

type Problem struct {
	ProblemId int64 
	Author string 
	Addtime int64 
	Oj string 
	Title string 
	Description string 
	Input string 
	Output string 
	SampleInput string 
	SampleOutput string 
	Hint string 
	Source string 
	LimitTime int64 
	LimitMemory int64 
	Auth int64 
	Level int64 
	Test_count int64 
	ProblemCode string 
}
```

### 9. "根据标签获取题目"

1. route definition

- Url: /api2/problems/tag/:tag_id
- Method: GET
- Request: `ListProblemsByTagIdRequest`
- Response: `ListProblemsByTagIdResponse`

2. request definition



```golang
type ListProblemsByTagIdRequest struct {
	TagId int64 `path:"tag_id"`
	Page int64 `form:"page,default=1"`
	PageSize int64 `form:"page_size,default=10"`
}
```


3. response definition



```golang
type ListProblemsByTagIdResponse struct {
	Problems []Problem `json:"problems"`
}
```

### 10. "添加标签"

1. route definition

- Url: /api2/tags
- Method: POST
- Request: `AddTagRequest`
- Response: `AddTagResponse`

2. request definition



```golang
type AddTagRequest struct {
	TagName string `form:"tag_name"`
}
```


3. response definition



```golang
type AddTagResponse struct {
}
```

### 11. "删除标签"

1. route definition

- Url: /api2/tags/delete
- Method: POST
- Request: `DeleteTagRequest`
- Response: `DeleteTagResponse`

2. request definition



```golang
type DeleteTagRequest struct {
	TagId int64 `form:"tag_id"`
}
```


3. response definition



```golang
type DeleteTagResponse struct {
}
```

### 12. "更新标签"

1. route definition

- Url: /api2/tags/update
- Method: POST
- Request: `UpdateTagRequest`
- Response: `UpdateTagResponse`

2. request definition



```golang
type UpdateTagRequest struct {
	TagId int64 `form:"tag_id"`
	TagName string `form:"tag_name"`
}
```


3. response definition



```golang
type UpdateTagResponse struct {
}
```

### 13. "搜索标签"

1. route definition

- Url: /api2/tags
- Method: GET
- Request: `SearchTagRequest`
- Response: `SearchTagResponse`

2. request definition



```golang
type SearchTagRequest struct {
	Page int64 `form:"page,default=1"`
	Limit int64 `form:"page_size,default=10"`
	TagName string `form:"tag_name,optional"`
	Order bool `form:"order,optional"`
}
```


3. response definition



```golang
type SearchTagResponse struct {
	Tag []Tag `json:"tag"`
}
```

### 14. "获取标签信息"

1. route definition

- Url: /api2/tags/:tag_id
- Method: GET
- Request: `GetTagByIdRequest`
- Response: `GetTagByIdResponse`

2. request definition



```golang
type GetTagByIdRequest struct {
	TagId int64 `path:"tag_id"`
}
```


3. response definition



```golang
type GetTagByIdResponse struct {
	Tag Tag `json:"tag"`
}

type Tag struct {
	TagId int64 
	TagName string 
}
```

### 15. "根据题目获取标签"

1. route definition

- Url: /api2/tags/problem/:problem_id
- Method: GET
- Request: `ListTagsByProblemIdRequest`
- Response: `ListTagsByProblemIdResponse`

2. request definition



```golang
type ListTagsByProblemIdRequest struct {
	ProblemId int64 `path:"problem_id"`
}
```


3. response definition



```golang
type ListTagsByProblemIdResponse struct {
	Tags []Tag `json:"tags"`
}
```

### 16. "搜索测试用例"

1. route definition

- Url: /api2/test_cases
- Method: GET
- Request: `SearchTestCaseRequest`
- Response: `SearchTestCaseResponse`

2. request definition



```golang
type SearchTestCaseRequest struct {
	ProblemId int64 `form:"problem_id"`
	TestGroup int64 `form:"test_group,default=0"`
}
```


3. response definition



```golang
type SearchTestCaseResponse struct {
	TestCases []TestCases `json:"test_cases"`
}
```

### 17. "获取测试用例信息"

1. route definition

- Url: /api2/test_cases/:test_id
- Method: GET
- Request: `GetTestCaseByIdRequest`
- Response: `GetTestCaseByIdResponse`

2. request definition



```golang
type GetTestCaseByIdRequest struct {
	TestId int64 `path:"test_id"`
}
```


3. response definition



```golang
type GetTestCaseByIdResponse struct {
	TestCase TestCases `json:"test_case"`
}

type TestCases struct {
	TestId int64 
	ProblemId int64 
	TestGroup int64 
	TestInputFileName string 
	TestOutputFileName string 
	UpdateAt int64 
}
```

