pragma solidity >=0.4.22 <0.7.0;


contract Flytrap {
    address payable owner;
    uint addTopicCost;

    struct Topic {
        bool isValue;
        bytes32 name;
        uint addPubCost;
        uint addSubCost;
        address payable owner;
        mapping(address => bool) publishers;
        mapping(address => bool) subscribers;
    }

    mapping(bytes32 => Topic) public topics;

    constructor(uint cost) public {
        owner = msg.sender;
        addTopicCost = cost;
    }


    function addTopic(bytes32 topic, uint addPubCost, uint addSubCost) public payable {
        require(topics[topic].isValue == false);
        require(msg.value >= addTopicCost);
        topics[topic] = Topic(true, topic, addPubCost, addSubCost, msg.sender);
        owner.transfer(msg.value);
    }
    
    function addPub(address person, bytes32 topic) public {
        require(topics[topic].owner == msg.sender);
        topics[topic].publishers[person] = true;
    }
    function addSub(address person, bytes32 topic) public {
        require(topics[topic].owner == msg.sender);
        topics[topic].subscribers[person] = true;
    }
    function revokePub(address person, bytes32 topic) public {
        require(topics[topic].owner == msg.sender);
        topics[topic].publishers[person] = false;
    }
    function revokeSub(address person, bytes32 topic) public {
        require(topics[topic].owner == msg.sender);
        topics[topic].subscribers[person] = false;
    }
    function verifyPub(address person, bytes32 topic) public view returns (bool) {
        return topics[topic].publishers[person];
    }
    function verifySub(address person, bytes32 topic) public view returns (bool) {
        return topics[topic].subscribers[person];
    }
}
