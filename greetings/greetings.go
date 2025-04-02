package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Hello 为指定的人返回问候语.
func Hello(name string) (string, error) {
	// 如果没有给出名字，返回一个带有消息的错误.
	if name == "" {
		return "", errors.New("empty name")
	}

	// 如果接收到名称，则返回一个嵌入名称的值
	// 在问候消息中.
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	// 一个映射来返回结果.
	messages := make(map[string]string)
	// 遍历切片中的名称，并请求来自每个名称的问候语.
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		// 在映射中将名称与消息关联起来.
		messages[name] = message
	}
	return messages, nil
}

func init() {
	// 在包初始化时设置随机数生成器种子
	rand.Seed(time.Now().UnixNano())
}

// randomFormat 返回一个随机格式化的问候消息.
func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	// 返回一个随机格式.
	return formats[rand.Intn(len(formats))]
}
