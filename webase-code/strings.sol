pragma solidity^0.6.0;

library strings {
    function isEqual(string memory a, string memory b) internal pure returns (bool) {
        //借助hash函数防碰撞特性
        // 如果hash(a) != hash(b), 那么认为a != b
        bytes32 hashA = keccak256(abi.encode(a));
        bytes32 hashB = keccak256(abi.encode(b));
        return uint256(hashA) == uint256(hashB);
    }
}