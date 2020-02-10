pragma solidity >=0.4.22 <0.7.0;


contract Flytrap {

    mapping(address => bool) public admins;

    mapping(address => mapping(bytes32 => bool)) public pub;
    mapping(address => mapping(bytes32 => bool)) public sub;

    constructor() public {
        admins[msg.sender] = true;
     }
    function addAdmin(address person) public {
        require(admins[msg.sender] == true);
        admins[person] = true;
    }
    function addPub(address person, bytes32 topic) public {
        require(admins[msg.sender] == true);
        pub[person][topic] = true;
    }
    function addSub(address person, bytes32 topic) public {
        require(admins[msg.sender] == true);
        sub[person][topic] = true;
    }
}
