pragma solidity >=0.4.22 <0.7.0;


contract Authorizer {
    address payable owner;

    mapping(bytes32 => address) public tokens;
    mapping(address => bool) public addresses;

    constructor(address payable sender) public {
        owner = sender;
    }

    function registerToken(bytes32 token, address person) public {
        require(msg.sender == owner);
        require(addresses[person] == false);
        tokens[token] = person;
        addresses[person] = true;
    }
    function retrieveKey(bytes32 token) public view returns (address) {
        return tokens[token];
    }
}
