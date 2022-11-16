package blocks

import (
	"fmt"
	"log"
	"math/big"

	"github.com/FISCO-BCOS/go-sdk/client"
	"github.com/FISCO-BCOS/go-sdk/conf"
	"github.com/ethereum/go-ethereum/common"
)

var instance *Task
var contractAddr string = "0xdefeb60d21837f6d6bfb5e260669cfd5a26104d9"
var Cli *client.Client

type TaskInfox struct {
	TaskID  string `json:"task_id"`
	Issuer  string `json:"issuer"`
	Worker  string `json:"task_user"`
	Desc    string `json:"task_name"`
	Bonus   int64  `json:"bonus"`
	Comment string `json:"comment"`
	Status  uint8  `json:"task_status"`
}

//init会自动触发运行
func init() {
	//0. 读取解析配置文件
	configs, err := conf.ParseConfigFile("config.toml")
	if err != nil {
		log.Fatal(err)
	}
	config := &configs[0]

	//1. 连接到节点
	cli, err := client.Dial(config)
	if err != nil {
		log.Fatal(err)
	}
	Cli = cli
	task, err := NewTask(common.HexToAddress(contractAddr), cli)
	if err != nil {
		log.Panic("failed to NewTask", err)
	}
	instance = task
}

func Register(username, passwd string) error {

	_, _, err := instance.Register(Cli.GetTransactOpts(), username, passwd)
	if err != nil {
		fmt.Println("failed to Register", err)
		return err
	}
	return nil
}

func Login(username, passwd string) (bool, error) {

	return instance.Login(Cli.GetCallOpts(), username, passwd)
}

func Issue(username, passwd, desc string, bonus int64) error {

	//合约调用
	_, _, err := instance.Issue(Cli.GetTransactOpts(), username, passwd, desc, big.NewInt(bonus))
	if err != nil {
		fmt.Println("failed to Issue", err)
		return err
	}
	return nil
}

//status : 1- 接受 2- 提交 3- 确认 4- 退回
func Update(username, passwd, comment string, status uint8, taskID int64) error {

	//合约调用
	if status == 1 {
		_, _, err := instance.Take(Cli.GetTransactOpts(), username, passwd, big.NewInt(taskID))
		if err != nil {
			fmt.Println("failed to Take task", err)
			return err
		}
	} else if status == 2 {
		_, _, err := instance.Commit(Cli.GetTransactOpts(), username, passwd, big.NewInt(taskID))
		if err != nil {
			fmt.Println("failed to Take Commit", err)
			return err
		}
	} else if status == 3 || status == 4 {
		if status == 4 {
			status = 1
		}
		_, _, err := instance.Confirm(Cli.GetTransactOpts(), username, passwd, big.NewInt(taskID), comment, status)
		if err != nil {
			fmt.Println("failed to Take Confirm", err)
			return err
		}
	}

	return nil
}

func Tasklist() ([]TaskInfox, error) {
	var tasks []TaskInfox
	_issuer, err := instance.QryAllIssuers(Cli.GetCallOpts())
	if err != nil {
		fmt.Println("failed to QryAllIssuers", err)
		return nil, err
	}
	_worker, err := instance.QryAllWorkers(Cli.GetCallOpts())
	if err != nil {
		fmt.Println("failed to QryAllWorkers", err)
		return nil, err
	}
	_desc, err := instance.QryAllDesc(Cli.GetCallOpts())
	if err != nil {
		fmt.Println("failed to QryAllDesc", err)
		return nil, err
	}
	_comment, err := instance.QryAllComments(Cli.GetCallOpts())
	if err != nil {
		fmt.Println("failed to QryAllComments", err)
		return nil, err
	}
	_bonus, err := instance.QryAllBonus(Cli.GetCallOpts())
	if err != nil {
		fmt.Println("failed to QryAllBonus", err)
		return nil, err
	}
	_status, err := instance.QryAllStatus(Cli.GetCallOpts())
	if err != nil {
		fmt.Println("failed to QryAllStatus", err)
		return nil, err
	}

	var task TaskInfox

	for k, v := range _issuer {
		task.TaskID = fmt.Sprintf("%d", k)
		task.Bonus = _bonus[k].Int64()
		task.Comment = _comment[k]
		task.Desc = _desc[k]
		task.Issuer = v
		task.Status = _status[k]
		task.Worker = _worker[k]
		tasks = append(tasks, task)
	}

	return tasks, nil
}
