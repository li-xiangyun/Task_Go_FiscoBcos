pragma solidity^0.6.0;

interface IUSER {
    //注册接口设计
    function register(string calldata username, string calldata passwd) external;
    //登陆接口设计
    function login(string calldata username, string calldata passwd) external view returns(bool);
}