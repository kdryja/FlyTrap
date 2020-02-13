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

    enum ActionPerformed { AddTopic, AddPub, AddSub, RevokePub, RevokeSub }
    event ACLChange (
      address indexed _src,
      address indexed _target,
      ActionPerformed indexed _action,
      bytes32 _name,
      uint _timestamp
    );

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
        emit ACLChange(msg.sender, msg.sender, ActionPerformed.AddTopic, topic, block.timestamp);
    }
    
    function addPub(address person, bytes32 topic) public {
        require(topics[topic].owner == msg.sender);
        topics[topic].publishers[person] = true;
        emit ACLChange(msg.sender, person, ActionPerformed.AddPub, topic, block.timestamp);
    }
    function addSub(address person, bytes32 topic) public {
        require(topics[topic].owner == msg.sender);
        topics[topic].subscribers[person] = true;
        emit ACLChange(msg.sender, person, ActionPerformed.AddSub, topic, block.timestamp);
    }
    function revokePub(address person, bytes32 topic) public {
        require(topics[topic].owner == msg.sender);
        topics[topic].publishers[person] = false;
        emit ACLChange(msg.sender, person, ActionPerformed.RevokePub, topic, block.timestamp);
    }
    function revokeSub(address person, bytes32 topic) public {
        require(topics[topic].owner == msg.sender);
        topics[topic].subscribers[person] = false;
        emit ACLChange(msg.sender, person, ActionPerformed.RevokeSub, topic, block.timestamp);
    }
    function verifyPub(address person, bytes32 topic) public view returns (bool) {
        return topics[topic].publishers[person];
    }
    function verifySub(address person, bytes32 topic) public view returns (bool) {
        return topics[topic].subscribers[person];
    }
}
