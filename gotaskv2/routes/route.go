package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"gotask/blocks"

	"github.com/gin-gonic/gin"
	"github.com/go-session/gin-session"
)

//通用响应消息格式
type RespData struct {
	Code string      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

//自定义一套错误编码
const (
	RESP_OK         = "0"
	RESP_PARAMERR   = "1"
	RESP_LOGINERR   = "2"
	RESP_BLOCKERR   = "3"
	RESP_UNKNOWNERR = "4"
)

var respMap map[string]string = map[string]string{
	RESP_OK:         "正常",
	RESP_PARAMERR:   "参数异常",
	RESP_LOGINERR:   "登陆失败",
	RESP_BLOCKERR:   "区块链操作失败",
	RESP_UNKNOWNERR: "未知异常",
}

func getRespMsg(code string) string {
	return respMap[code]
}

func respJsonMsg(c *gin.Context, r *RespData) {
	r.Msg = getRespMsg(r.Code)
	c.JSON(http.StatusOK, r)
}

type UserInfo struct {
	UserName string `json:"username"`
	PassWord string `json:"password"`
}

//注册功能函数
func Register(c *gin.Context) {
	//3. 组织响应消息
	var resp RespData
	resp.Code = RESP_OK

	defer respJsonMsg(c, &resp)
	//1. 解析请求消息
	var ui UserInfo
	err := c.ShouldBindJSON(&ui)
	if err != nil {
		resp.Code = RESP_PARAMERR
		fmt.Println("failed to ShouldBindJSON", err)
		return
	}
	fmt.Printf("Register:%+v\n", ui)
	//2. 业务规则处理
	err = blocks.Register(ui.UserName, ui.PassWord)
	if err != nil {
		resp.Code = RESP_BLOCKERR
		return
	}
}

//登陆请求
func Login(c *gin.Context) {
	//3. 组织响应消息
	var resp RespData
	resp.Code = RESP_OK

	defer respJsonMsg(c, &resp)
	//1. 解析请求消息
	var ui UserInfo
	err := c.ShouldBindJSON(&ui)
	if err != nil {
		resp.Code = RESP_PARAMERR
		fmt.Println("failed to ShouldBindJSON", err)
		return
	}
	fmt.Printf("Login:%+v\n", ui)
	//2. 业务规则处理
	ok, err := blocks.Login(ui.UserName, ui.PassWord)
	if err != nil {
		resp.Code = RESP_BLOCKERR
		fmt.Println("failed to blocks.Login", err)
		return
	}
	if !ok {
		fmt.Println("user not exists or password err")
		resp.Code = RESP_LOGINERR
		return
	}
	//4. 设置session
	store := ginsession.FromContext(c)
	store.Set("username", ui.UserName)
	store.Set("password", ui.PassWord)
	err = store.Save()
	if err != nil {
		//c.AbortWithError(500, err)
		resp.Code = RESP_UNKNOWNERR
		return
	}
}

//任务发布
func Issue(c *gin.Context) {
	//1. 组织响应消息
	var resp RespData
	resp.Code = RESP_OK

	defer respJsonMsg(c, &resp)
	//2. 解析请求消息
	var task blocks.TaskInfox
	err := c.ShouldBindJSON(&task)
	if err != nil {
		resp.Code = RESP_PARAMERR
		fmt.Println("failed to ShouldBindJSON", err)
		return
	}
	fmt.Printf("Issue:%+v\n", task)
	//3. 读取session信息
	store := ginsession.FromContext(c)
	username, ok := store.Get("username")
	if !ok {
		resp.Code = RESP_UNKNOWNERR
		fmt.Println("get session username err")
		return
	}
	password, ok := store.Get("password")
	if !ok {
		resp.Code = RESP_UNKNOWNERR
		fmt.Println("get session password err")
		return
	}
	//4. 业务规则处理
	err = blocks.Issue(username.(string), password.(string), task.Desc, task.Bonus)
	if err != nil {
		resp.Code = RESP_BLOCKERR
		fmt.Println("failed to blocks.Issue", err)
		return
	}
}

//任务修改
func Update(c *gin.Context) {
	//1. 组织响应消息
	var resp RespData
	resp.Code = RESP_OK

	defer respJsonMsg(c, &resp)
	//2. 解析请求消息
	var task blocks.TaskInfox
	err := c.ShouldBindJSON(&task)
	if err != nil {
		resp.Code = RESP_PARAMERR
		fmt.Println("failed to ShouldBindJSON", err)
		return
	}
	fmt.Printf("Update:%+v\n", task)
	//3. 读取session信息
	store := ginsession.FromContext(c)
	username, ok := store.Get("username")
	if !ok {
		resp.Code = RESP_UNKNOWNERR
		fmt.Println("get session username err")
		return
	}
	password, ok := store.Get("password")
	if !ok {
		resp.Code = RESP_UNKNOWNERR
		fmt.Println("get session password err")
		return
	}
	//4. 业务逻辑处理
	taskID, _ := strconv.Atoi(task.TaskID)
	err = blocks.Update(username.(string), password.(string), task.Comment, task.Status, int64(taskID))
	if err != nil {
		resp.Code = RESP_BLOCKERR
		fmt.Println("blocks.Update  err", err)
		return
	}
}

//任务列表请求
func Tasklist(c *gin.Context) {
	//1. 组织响应消息
	var resp RespData
	resp.Code = RESP_OK

	defer respJsonMsg(c, &resp)
	//2. 解析请求消息
	strpage := c.Query("page")
	ipage, _ := strconv.Atoi(strpage)

	fmt.Printf("Tasklist:ipage = %d\n", ipage)
	//3. 业务规则处理
	tasks, err := blocks.Tasklist()
	if err != nil {
		resp.Code = RESP_BLOCKERR
		fmt.Println("failed to blocks.Tasklist", err)
		return
	}
	//4. 重新响应消息中设置任务列表数据
	begin := (ipage - 1) * 10
	end := ipage * 10
	if end > len(tasks) {
		end = len(tasks)
	}
	data := struct {
		Total int         `json:"total"`
		Data  interface{} `json:"data"`
	}{
		Total: len(tasks),
		Data:  tasks[begin:end],
	}
	resp.Data = data

}
