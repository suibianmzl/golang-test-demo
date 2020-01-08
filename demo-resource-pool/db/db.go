package db

import (
	"io"
	"log"
	"sync/atomic"
)

// 给每一个连接分配独一无二的id
var idCounter int32

// 资源数据库连接
type DBConnection struct {
	ID int32
}

// dbConnection 实现了 io.Closer 接口
// 关闭资源
func (conn *DBConnection) Close() error  {
	log.Println("conn Closed")
	return nil
}

// 创建一个资源 - dbConnection 
func CreateConn()  (io.Closer, error){
	id := atomic.AddInt32(&idCounter, 1)
	log.Println("create conn, id:", id)
	return &DBConnection{
		ID:id,
	},nil
}





