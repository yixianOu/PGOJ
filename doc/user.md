### 1. "发送邮件"

1. route definition

- Url: /api1/users/email
- Method: GET
- Request: `SendEmailToUserRequest`
- Response: `SendEmailToUserResponse`

2. request definition



```golang
type SendEmailToUserRequest struct {
	Email string `form:"email"`
}
```


3. response definition



```golang
type SendEmailToUserResponse struct {
}
```

### 2. "获取用户列表"

1. route definition

- Url: /api1/users
- Method: GET
- Request: `ListUserRequest`
- Response: `ListUserResponse`

2. request definition



```golang
type ListUserRequest struct {
	Page int `form:"page,default=1"`
	PageSize int `form:"page_size,default=10"`
	RoleLevel int64 `form:"role_level,optional"`
	Username string `form:"username,optional"`
	Order bool `form:"order,default=false"`
}
```


3. response definition



```golang
type ListUserResponse struct {
	User_list []LoginRow `json:"user_list"`
	UsersCount int `json:"users_count"`
}
```

### 3. "创建用户"

1. route definition

- Url: /api1/users
- Method: POST
- Request: `CreateUserRequest`
- Response: `CreateUserResponse`

2. request definition



```golang
type CreateUserRequest struct {
	Username string `form:"username"`
	Email string `form:"email"`
	Password string `form:"password"`
	CoverImageUrl string `form:"cover_image_url,default=https://oj-file.oss-cn-shenzhen.aliyuncs.com/1user_cover.jpg"` //提供默认头像
	EmailCode string `form:"email_code"`
}
```


3. response definition



```golang
type CreateUserResponse struct {
}
```

### 4. "用户验证码登录"

1. route definition

- Url: /api1/users/email/login
- Method: POST
- Request: `LoginWithCodeRequest`
- Response: `LoginUserResponse`

2. request definition



```golang
type LoginWithCodeRequest struct {
	Email string `form:"email"`
	Code string `form:"email_code"`
}
```


3. response definition



```golang
type LoginUserResponse struct {
	User_login LoginRow `json:"user_info"`
	Token string `json:"token"`
}

type LoginRow struct {
	ID int64 `json:"id"`
	Username string `json:"username"`
	RoleLevel int64 `json:"role_level"`
	CoverImageUrl string `json:"cover_image_url"`
}
```

### 5. "用户登录"

1. route definition

- Url: /api1/users/login
- Method: POST
- Request: `LoginUserRequest`
- Response: `LoginUserResponse`

2. request definition



```golang
type LoginUserRequest struct {
	Email string `form:"email"`
	Password string `form:"password"`
}
```


3. response definition



```golang
type LoginUserResponse struct {
	User_login LoginRow `json:"user_info"`
	Token string `json:"token"`
}

type LoginRow struct {
	ID int64 `json:"id"`
	Username string `json:"username"`
	RoleLevel int64 `json:"role_level"`
	CoverImageUrl string `json:"cover_image_url"`
}
```

### 6. "获取用户基础信息"

1. route definition

- Url: /api1/users/:user_id
- Method: GET
- Request: `GetUserInfoRequest`
- Response: `GetUserInfoResponse`

2. request definition



```golang
type GetUserInfoRequest struct {
	ID int64 `path:"user_id,range=[1:]"`
}
```


3. response definition



```golang
type GetUserInfoResponse struct {
	User_info UserInfo `json:"user_info"`
}

type UserInfo struct {
	ID int64 `json:"id"`
	Password string `json:"password"`
	RoleLevel int64 `json:"role_level"`
	Username string `json:"username"`
	Email string `json:"email"`
	CoverImageUrl string `json:"cover_url"`
}
```

### 7. "更新用户基础信息"

1. route definition

- Url: /api1/users/:user_id
- Method: POST
- Request: `UpdateUserRequest`
- Response: `UpdateUserResponse`

2. request definition



```golang
type UpdateUserRequest struct {
	ID int64 `path:"user_id,range=[1:]"`
	Password string `form:"password,optional"`
	Username string `form:"username,optional"`
	Email string `form:"email,optional"`
	CoverImageUrl string `form:"cover_url,optional"`
}
```


3. response definition



```golang
type UpdateUserResponse struct {
}
```

### 8. "删除用户"

1. route definition

- Url: /api1/users/:user_id
- Method: DELETE
- Request: `DeleteUserRequest`
- Response: `DeleteUserResponse`

2. request definition



```golang
type DeleteUserRequest struct {
	ID int64 `path:"user_id,range=[1:]"`
}
```


3. response definition



```golang
type DeleteUserResponse struct {
}
```

### 9. "授权用户"

1. route definition

- Url: /api1/users/:user_id/authenticate
- Method: POST
- Request: `AuthorizeUserRequest`
- Response: `AuthorizeUserResponse`

2. request definition



```golang
type AuthorizeUserRequest struct {
	UserID int64 `path:"user_id,range=[1:]"`
	RoleLevel int64 `form:"role_level,range=[1:2]"`
}
```


3. response definition



```golang
type AuthorizeUserResponse struct {
}
```

### 10. "更新用户详细信息"

1. route definition

- Url: /api1/users/:user_id/profile
- Method: POST
- Request: `UpdateProfileRequest`
- Response: `UpdateProfileResponse`

2. request definition



```golang
type UpdateProfileRequest struct {
	UserID int64 `path:"user_id,range=[1:]"`
	Phone string `form:"phone"`
	Name string `form:"name"`
	School string `form:"school"`
	Description string `form:"description"`
}
```


3. response definition



```golang
type UpdateProfileResponse struct {
}
```

### 11. "获取用户详细信息"

1. route definition

- Url: /api1/users/:user_id/profile
- Method: GET
- Request: `GetProfileRequest`
- Response: `ProfileResponse`

2. request definition



```golang
type GetProfileRequest struct {
	UserID int64 `path:"user_id,range=[1:]"`
}
```


3. response definition



```golang
type ProfileResponse struct {
	Profile Profile `json:"profile"`
	UserInfo LoginRow `json:"user_info"`
}

type Profile struct {
	ID int64 `json:"id"`
	UserID int64 `json:"user_id"`
	Phone string `json:"phone"`
	Name string `json:"name"`
	School string `json:"school"`
	Description string `json:"description"`
	ACCount int64 `json:"ac_count"`
	SubmitCount int64 `json:"submit_count"`
	Score int64 `json:"score"`
	Rating int64 `json:"rating"`
	ACProblem string `json:"ac_problem"`
}

type LoginRow struct {
	ID int64 `json:"id"`
	Username string `json:"username"`
	RoleLevel int64 `json:"role_level"`
	CoverImageUrl string `json:"cover_image_url"`
}
```

