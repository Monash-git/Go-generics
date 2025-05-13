package errs

import (
	"fmt"
	"time"
)

//NewErr方法用于创建对应的错误信息

func NewErrIndexOutOfRange(length int, index int) error {
	return fmt.Errorf("下标超出范围，长度%d, 下标%d",length,index)
}

func NewErrInvalidType(want string, got any) error {
	return fmt.Errorf("类型转换失败，与其类型:%s，实际值:%#v",want,got)
}

func NewErrInvalidIntervalValue(interval time.Duration) error {
	return fmt.Errorf("无效的间隔时间 %d，预期值应大于 0",interval)
}

func NewErrInvalidMaxIntervalValue(maxInterval, initialInterval time.Duration) error {
	return fmt.Errorf("最大重试的间隔时间 [%d] 应大于等于初始重试的间隔时间 [%d] ",maxInterval,initialInterval)
}

func NewErrRetryExhausted(lastErr error) error {
	return fmt.Errorf("超过最大重试次数，业务返回的最后一个 error %w",lastErr)
}