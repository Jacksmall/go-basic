package errors

// Error 接口, error 接口类型的扩展
type Error interface {
	error
	Caller() []CallerInfo // 调用栈信息
	Wrapped() []error     // 错误的多级嵌套
	Code() int            // 错误码
	private()
}

type CallerInfo struct {
	FuncName string
	FileName string
	FileLine int
}

// 构建新的错误类型
func New(msg string) error {
	return nil
}

// 构造一个带错误码的错误
func NewWithCode(code int, msg string) error {
	return nil
}

// 错误二次包装函数
func Wrap(err error, msg string) error {
	return nil
}

func WrapWithCode(code int, err, error, msg string) error {
	return nil
}

// 从json字符串编码的错误中恢复错误对象
func FromJson(json string) (Error, error) {
	return nil, nil
}

func ToJson(err error) string {
	return ""
}
