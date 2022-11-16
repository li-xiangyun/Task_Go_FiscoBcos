pragma solidity^0.6.0;

import "./interface/IUSER.sol";
import "./strings.sol";

contract User is IUSER {
    mapping(string=>string) _users;
    address admin;
    using strings for string;//string类型的变量可以使用strings库的函数
    constructor() public {
        admin = msg.sender;
    }
    //注册接口设计
    function register(string calldata username, string calldata passwd) override external {
        //检测：用户没有注册过；用户不能为空；密码不能为空
        require(!username.isEqual(""), "username must not null");
        require(!passwd.isEqual(""), "passwd must not null");
        require(_users[username].isEqual(""), "user must not");
        _users[username] = passwd;
    }
    //登陆接口设计
    function login(string calldata username, string calldata passwd) override external view returns(bool) {
        //require(!username.isEqual(""), "username must not null");
        //require(!passwd.isEqual(""), "passwd must not null");
        if(username.isEqual("") || passwd.isEqual("")) {
            return false;
        }
        return _users[username].isEqual(passwd);
    }
}