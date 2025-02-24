package request

// User login structure
type Signin struct {
	Email     string `json:"email"`     // 用户名
	Password  string `json:"password"`  // 密码
	Captcha   string `json:"captcha"`   // 验证码
	CaptchaId string `json:"captchaId"` // 验证码ID
}

// Register User register structure
type Signup struct {
	Username  string `json:"userName" example:"用户名"`
	Email     string `json:"email"`     // 用户名
	Password  string `json:"password"`  // 密码
	Captcha   string `json:"captcha"`   // 验证码
	CaptchaId string `json:"captchaId"` // 验证码ID
}

// Modify password structure
type ChangePasswordReq struct {
	ID          uint   `json:"-"`           // 从 JWT 中提取 user id，避免越权
	Password    string `json:"password"`    // 密码
	NewPassword string `json:"newPassword"` // 新密码
}

type UserInfo struct {
	Username  string `json:"userName" example:"用户名"`
	Email     string `json:"email"`     // 用户名
	Phone     string `json:"phone"`     // 手机
	HeaderImg string `json:"headerImg"` // 头像
}
