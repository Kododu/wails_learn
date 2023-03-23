package api

import (
	"changeme/internal/connection"
	"changeme/internal/define"
	"context"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// Startup is called when the app starts. The context is saved,
// so we can call the runtime methods
func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
}

// ConnectList 返回配置列表
func (a *App) ConnectList() Response {
	connectionList, err := connection.List()
	if err != nil {
		return ReplayBasicError(err)
	}
	return BuildResponse("list", connectionList).Replay()
}

// DBConnectList 返回数据库配置列表
func (a *App) DBConnectList() Response {
	connectionList, err := connection.List()
	if err != nil {
		return ReplayBasicError(err)
	}
	return BuildResponse("list", connectionList).Replay()
}

// ConnectionCreate 创建一个连接
func (a *App) ConnectionCreate(con *define.Connection) Response {
	err := connection.Create(con)
	if err != nil {
		return ReplayBasicError(err)
	}
	return ReplaySuccessEmpty()
}

// ConnectionDelete 删除指定连接
func (a *App) ConnectionDelete(id string) Response {
	err := connection.Delete(id)
	if err != nil {
		return ReplayBasicError(err)
	}
	return ReplaySuccessEmpty()
}

func (a *App) DialDB(id string) Response {
	err := connection.DBConnection(id)
	if err != nil {
		return ReplayBasicError(err)
	}
	return ReplaySuccessEmpty()
}
