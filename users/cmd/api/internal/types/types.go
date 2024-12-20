// Code generated by goctl. DO NOT EDIT.
package types

type AuthorizeUserRequest struct {
	UserID    int64 `path:"user_id,range=[1:]"`
	RoleLevel int64 `form:"role_level,range=[1:2]"`
}

type AuthorizeUserResponse struct {
}

type CreateUserRequest struct {
	Username      string `form:"username"`
	Email         string `form:"email"`
	Password      string `form:"password"`
	CoverImageUrl string `form:"cover_image_url,default=http://124.223.74.196/oj-file/user_cover/0.jpg"` //提供默认头像
	EmailCode     string `form:"email_code"`
}

type CreateUserResponse struct {
}

type DeleteUserRequest struct {
	ID int64 `path:"user_id,range=[1:]"`
}

type DeleteUserResponse struct {
}

type GetProfileRequest struct {
	UserID int64 `path:"user_id,range=[1:]"`
}

type GetUserInfoRequest struct {
	ID int64 `path:"user_id,range=[1:]"`
}

type GetUserInfoResponse struct {
	User_info struct {
		ID            int64  `json:"id"`
		Password      string `json:"password"`
		RoleLevel     int64  `json:"role_level"`
		Username      string `json:"username"`
		Email         string `json:"email"`
		CoverImageUrl string `json:"cover_image_url"`
	} `json:"user_info"`
}

type ListProfileRequest struct {
	Page         int64  `form:"page,default=1"`
	Limit        int64  `form:"page_size,default=10"`
	Descrition   string `form:"description,optional"`
	School       string `form:"school,optional"`
	OrderByScore bool   `form:"score_order,optional"`
}

type ListProfileResponse struct {
	Profiles []Profile `json:"profiles"`
}

type ListUserRequest struct {
	Page      int    `form:"page,default=1"`
	PageSize  int    `form:"page_size,default=10"`
	RoleLevel int64  `form:"role_level,optional"`
	Username  string `form:"username,optional"`
	Order     bool   `form:"order,default=false"`
}

type ListUserResponse struct {
	User_list  []LoginRow `json:"user_list"`
	UsersCount int        `json:"users_count"`
}

type LoginRow struct {
	ID            int64  `json:"id"`
	Username      string `json:"username"`
	RoleLevel     int64  `json:"role_level"`
	CoverImageUrl string `json:"cover_image_url"`
}

type LoginUserRequest struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}

type LoginUserResponse struct {
	User_login struct {
		ID            int64  `json:"id"`
		Username      string `json:"username"`
		RoleLevel     int64  `json:"role_level"`
		CoverImageUrl string `json:"cover_image_url"`
	} `json:"user_info"`
	Token string `json:"token"`
}

type LoginWithCodeRequest struct {
	Email string `form:"email"`
	Code  string `form:"email_code"`
}

type Profile struct {
	ID          int64  `json:"id"`
	UserID      int64  `json:"user_id"`
	Phone       string `json:"phone"`
	Name        string `json:"name"`
	School      string `json:"school"`
	Description string `json:"description"`
	ACCount     int64  `json:"ac_count"`
	SubmitCount int64  `json:"submit_count"`
	Score       int64  `json:"score"`
	Rating      int64  `json:"rating"`
}

type ProfileResponse struct {
	Profile struct {
		ID          int64  `json:"id"`
		UserID      int64  `json:"user_id"`
		Phone       string `json:"phone"`
		Name        string `json:"name"`
		School      string `json:"school"`
		Description string `json:"description"`
		ACCount     int64  `json:"ac_count"`
		SubmitCount int64  `json:"submit_count"`
		Score       int64  `json:"score"`
		Rating      int64  `json:"rating"`
	} `json:"profile"`
	UserInfo struct {
		ID            int64  `json:"id"`
		Username      string `json:"username"`
		RoleLevel     int64  `json:"role_level"`
		CoverImageUrl string `json:"cover_image_url"`
	} `json:"user_info"`
}

type RefreshUserRatingRequest struct {
	UserID int64 `path:"user_id,range=[1:]"`
}

type RefreshUserRatingResponse struct {
	Profile struct {
		ID          int64  `json:"id"`
		UserID      int64  `json:"user_id"`
		Phone       string `json:"phone"`
		Name        string `json:"name"`
		School      string `json:"school"`
		Description string `json:"description"`
		ACCount     int64  `json:"ac_count"`
		SubmitCount int64  `json:"submit_count"`
		Score       int64  `json:"score"`
		Rating      int64  `json:"rating"`
	} `json:"profile"`
}

type SendEmailToUserRequest struct {
	Email string `form:"email"`
}

type SendEmailToUserResponse struct {
}

type UpdateProfileRequest struct {
	UserID      int64  `path:"user_id,range=[1:]"`
	Phone       string `form:"phone,optional"`
	Name        string `form:"name,optional"`
	School      string `form:"school,optional"`
	Description string `form:"description,optional"`
}

type UpdateProfileResponse struct {
}

type UpdateUserRequest struct {
	ID            int64  `path:"user_id,range=[1:]"`
	Password      string `form:"password,optional"`
	Username      string `form:"username,optional"`
	Email         string `form:"email,optional"`
	CoverImageUrl string `form:"cover_image_url,optional"`
}

type UpdateUserResponse struct {
}

type UserInfo struct {
	ID            int64  `json:"id"`
	Password      string `json:"password"`
	RoleLevel     int64  `json:"role_level"`
	Username      string `json:"username"`
	Email         string `json:"email"`
	CoverImageUrl string `json:"cover_image_url"`
}
