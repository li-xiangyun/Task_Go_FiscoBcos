pragma solidity^0.6.0;
pragma experimental ABIEncoderV2;

import "./Token.sol";
import "./User.sol";


contract Taskv2 {
    Token _token;
    User _user;
    //下标相同时，属于同一个任务
    string[] _issuer;
    string[] _worker;
    string[] _desc;
    uint256[] _bonus;
    uint8[]  _status;//0- 已发布 1- 已接受 2- 已提交 3- 已确认
    uint256[] _timestamp;
    string[] _comment;//任务执行效果评价
    //--------------
    address _admin;
    address _owner;
    using strings for string;
    constructor(address owner) public {
        _admin = msg.sender;
        _owner = owner;
        _token = new Token();
        _user  = new User();
    }
    
    function totalSupply()  external view returns (uint256) {
        return _token.totalSupply();
    }
    //用户余额
    function balanceOf(string calldata _owner)  external view returns (uint256 balance) {
        return _token.balanceOf(_owner);
    }
    //转账
    function transfer(string calldata _from, string calldata _to, uint256 _value)  external returns (bool success) {
        _token.transfer(_from, _to, _value);
    }
    
    function mint(string calldata _to, uint256 _value) external {
        require(_admin == msg.sender || _owner == msg.sender, "only admin can do");
        _token.mint(_to, _value);
    }
    //注册
    function register(string calldata username, string calldata passwd)  external {
        //检测：用户没有注册过；用户不能为空；密码不能为空
        _user.register(username, passwd);
    }
    //登陆接口设计
    function login(string calldata username, string calldata passwd)  external view returns(bool) {
        return _user.login(username, passwd);
    }
    //发布任务
    function issue(string calldata issuer, string calldata passwd, string calldata desc, uint256 bonus) external {
        require(_user.login(issuer, passwd), "issuer must have his right");
        require(bonus > 0, "bonus must > 0");
        require(_token.balanceOf(issuer) >= bonus, "issuer's balance must enough");
        _issuer.push(issuer);
        _worker.push("");
        _desc.push(desc);
        _bonus.push(bonus);
        _timestamp.push(now);
        _comment.push("");
        _status.push(0);
        
    }
    //接受任务
    function take(string calldata worker, string calldata passwd, uint256 taskID) external {
        require(_user.login(worker, passwd), "worker must have his right");
        require(_timestamp[taskID] > 0, "task must exists");
        require(_status[taskID] == 0, "task's status must equal 0");
        _worker[taskID] = worker;
        _status[taskID] = 1;
    }
    //提交任务
    function commit(string calldata worker, string calldata passwd, uint256 taskID) external {
        require(_user.login(worker, passwd), "worker must have his right");
        require(_timestamp[taskID] > 0, "task must exists");
        require(_status[taskID] == 1, "task's status must equal 1");
        require(worker.isEqual(_worker[taskID]), "task's works must equal worker");
        _status[taskID] = 2;//任务已提交
    }
    //确认任务
    function confirm(string calldata issuer, string calldata passwd, uint256 taskID, string calldata comment, uint8 status) external {
        require(_user.login(issuer, passwd), "issuer must have his right");
        require(_timestamp[taskID] > 0, "task must exists");
        require(_status[taskID] == 2, "task's status must equal 2");
        require(issuer.isEqual(_issuer[taskID]), "task's issuer must equal issuer");
        require(status == 3 || status == 1, "status must in (1,3)");
        _status[taskID] = status;
        _comment[taskID] = comment;
        if (status == 3) {
            //代表确认通过
            _token.transfer(issuer, _worker[taskID], _bonus[taskID]);
        }
    }
    //查询接口
    function qryOneTask(uint256 taskID) external view returns (string memory, string memory, uint256, string memory, uint8) {
        return (_issuer[taskID], _desc[taskID], _bonus[taskID], _worker[taskID], _status[taskID]);
    }
    //查询所有任务
    function qryAllIssuers() external view returns (string[] memory) {
        return _issuer;
    }
    function qryAllWorkers() external view returns (string[] memory) {
        return _worker;
    }
    function qryAllDesc() external view returns (string[] memory) {
        return _desc;
    }
    function qryAllComments() external view returns (string[] memory) {
        return _comment;
    }
    function qryAllBonus() external view returns (uint256[] memory) {
        return _bonus;
    }
    function qryAllStatus() external view returns (uint8[] memory) {
        return _status;
    }
    
    
}