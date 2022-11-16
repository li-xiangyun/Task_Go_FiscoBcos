pragma solidity^0.6.0;

import "./IERC2000.sol";
import "./strings.sol";
import "./SafeMath.sol";

contract Token is IERC2000 {
    uint256 _totalSupply;
    address admin;
    mapping(string=>uint256) _balances;
    using strings for string;
    using SafeMath for uint256;
    constructor() public {
        _totalSupply = 0;
        admin = msg.sender;
    }
    //总的发行量
    function totalSupply() override external view returns (uint256) {
        return _totalSupply;
    }
    //用户余额
    function balanceOf(string calldata _owner) override external view returns (uint256 balance) {
        return _balances[_owner];
    }
    //转账
    function transfer(string calldata _from, string calldata _to, uint256 _value) override external returns (bool success) {
        require(!_to.isEqual(""), "to must exists");
        require(_value > 0, "_value must > 0");
        require(_balances[_from] >= _value, "balance must enough");
        _balances[_from] = _balances[_from].sub(_value);
        _balances[_to] = _balances[_to].add(_value);
        
        emit Transfer(_from, _to, _value);
    }
    
    function mint(string calldata _to, uint256 _value) external {
        require(admin == msg.sender, "only admin can do");
        require(!_to.isEqual(""), "to must exists");
        require(_value > 0, "_value must > 0");
        _totalSupply = _totalSupply.add(_value);
        _balances[_to] = _balances[_to].add(_value);
        emit Transfer("system", _to, _value);
    }
    //事件通知
    event Transfer(string indexed _from, string indexed _to, uint256 _value);
}