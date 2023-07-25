package api

import (
	"strconv"
)

var (
	OK = NewError(200, "请求成功")

	// 服务级错误码
	ErrServer    = NewError(1001, "服务异常，请联系管理员")
	ErrParam     = NewError(1002, "参数有误")
	ErrSignParam = NewError(1003, "签名参数有误")

	// 缓存服务器错误码
	ErrCache = NewError(10001, "缓存服务器异常")
    ErrCacheDate = NewError(10002, "缓存服务器获取失败，数据连接异常")

	// 业务异常
	UserAddFailErr        = NewError(20001, "添加用户失败")
	UserNameExistErr      = NewError(20002, "用户名已存在")
	UserNameOrPasswordErr = NewError(20003, "账号或密码错误")
	UserNameEmptyErr      = NewError(20004, "用户名不能为空")

	PhoneExistErr  = NewError(20101, "手机号码已存在")
	PhoneEmptyErr  = NewError(20102, "手机号码不能为空")
	PhoneNumberErr = NewError(20103, "手机号码错误")

	SmsCodeEmptyErr = NewError(20201, "短信验证码不能为空")
	SmsCodeSendErr  = NewError(20202, "短信验证码发送失败")
	SmsCodeErr      = NewError(20203, "短信验证码错误")

	PasswordEmptyErr = NewError(20302, "密码不能为空")
	PasswordErr      = NewError(20403, "两次密码不一致")
	LoginTypeErr     = NewError(20504, "登录方式错误")

	QueryUserFailErr = NewError(30001, "获取用户信息失败")
	QRCodeRetryErr   = NewError(30002, "重新获取二维码")
	QRCodeGetFailErr = NewError(30003, "获取二维码失败")

	AddArticleFailErr = NewError(50001, "添加文章失败")

	JwtValidationErr = NewError(90001, "令牌验证失败")
	JwtExpiresErr    = NewError(90002, "无效令牌")
	JwtGeneratorErr  = NewError(90003, "生成令牌失败")
)

func NewError(code int, text string) *CodeError {
	return &CodeError{code, text, nil}
}

func NewErrorData(code int, text string, data interface{}) *CodeError {
	return &CodeError{code, text, data}
}

type CodeError struct {
	Code    int
	Message string
	Data    interface{}
}

func (e *CodeError) Error() string {
	return strconv.Itoa(e.Code) + ": " + e.Message
}
