package api

import (
	"DDD-OrderingSystem/OrderingApplication/Domain/Command"
	"DDD-OrderingSystem/OrderingApplication/Domain/Model"
	"DDD-OrderingSystem/OrderingApplication/Infrastructure/Repository"
	"encoding/json"
	"net/http"
)

// UserController 用户控制器
type UserController struct {
	userRepository Repository.UserRepository
	commandHandler Command.Handler
}

// NewUserController 创建用户控制器实例
func NewUserController(userRepository Repository.UserRepository, commandHandler Command.Handler) *UserController {
	return &UserController{
		userRepository: userRepository,
		commandHandler: commandHandler,
	}
}

// CreateUser 创建用户
func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user Model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// 创建命令
	createUserCommand := Command.NewCreateUserCommand(user.Username, user.Email, user.Password, user.Role)

	// 处理命令
	if err := uc.commandHandler.Handle(context.Background(), createUserCommand); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 返回成功响应
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createUserCommand.GetUser())
}

// GetUser 获取用户
func (uc *UserController) GetUser(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("id")
	if userID == "" {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}

	user, err := uc.userRepository.FindById(context.Background(), userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

// UpdateUser 更新用户
func (uc *UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user Model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// 创建命令
	updateUserCommand := Command.NewUpdateUserCommand(user.ID, user.Username, user.Email, user.Password, user.Role)

	// 处理命令
	if err := uc.commandHandler.Handle(context.Background(), updateUserCommand); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 返回成功响应
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updateUserCommand.GetUser())
}

// DeleteUser 删除用户
func (uc *UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("id")
	if userID == "" {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}

	// 创建命令
	deleteUserCommand := Command.NewDeleteUserCommand(userID)

	// 处理命令
	if err := uc.commandHandler.Handle(context.Background(), deleteUserCommand); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 返回成功响应
	w.WriteHeader(http.StatusOK)
}
