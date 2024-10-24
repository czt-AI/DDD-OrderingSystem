package query

import (
	"DDD-OrderingSystem/Domain/Model"
)

// UserDetailsQuery 用户详情查询
type UserDetailsQuery struct {
	UserID uint64
}

// NewUserDetailsQuery 创建用户详情查询实例
func NewUserDetailsQuery(userID uint64) *UserDetailsQuery {
	return &UserDetailsQuery{
		UserID: userID,
	}
}

// Execute 执行查询
func (q *UserDetailsQuery) Execute(ctx context.Context) (*Model.User, error) {
	// 这里应该实现查询用户详情的逻辑，通常是通过仓库来获取数据
	// 以下代码仅为示例，实际应用中需要根据具体逻辑来获取用户详情

	user, err := Model.FindUserById(ctx, q.UserID)
	if err != nil {
		return nil, err
	}

	return user, nil
}
