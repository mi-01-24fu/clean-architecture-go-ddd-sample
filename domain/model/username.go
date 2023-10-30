package model

import "fmt"

type UserName struct {
	value string
}

// デリファレンスの * は、実際のデータにアクセスする動作をする
// ポインタ型の * は、型の定義や関数のシグネチャにおいて、データがポインタであることを示している
func NewUserName(value string) (*UserName, error) {
	nameLen := len([]rune(value))

	if nameLen == 0 {
		return nil, fmt.Errorf("user name is not input")
	}
	if nameLen < 3 {
		return nil, fmt.Errorf("user name 3 or more characters and please input length %v", nameLen)
	}
	if nameLen > 10 {
		return nil, fmt.Errorf("user name 10 characters or less and please input length %v", nameLen)
	}

	return &UserName{value: value}, nil
}

func (u UserName) String() string {
	return u.value
}
