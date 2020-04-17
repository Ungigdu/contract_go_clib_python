pragma solidity >= 0.5.0;

contract SimpleDemo{

    mapping(string=>string) public map;

    event Update(string key, address operator);

    function updateMap(string calldata key, string calldata value) external {
            map[key] = value;
            emit Update(key, msg.sender);
    }
}