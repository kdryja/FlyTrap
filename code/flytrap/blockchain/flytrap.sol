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
        bytes2 country;
        mapping(address => bool) publishers;
        mapping(address => bool) subscribers;
    }

    enum ActionPerformed { 
      AddTopic,
      AddPub,
      AddSub,
      RevokePub,
      RevokeSub,
      WrongCountry,
      FiveFailures
    }

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

    function addTopic(bytes32 topic, bytes2 country, uint addPubCost, uint addSubCost) public payable {
        require(topics[topic].isValue == false);
        require(msg.value >= addTopicCost);
        topics[topic] = Topic(true, topic, addPubCost, addSubCost, msg.sender, country);

        if (addTopicCost > 0) {
          owner.transfer(msg.value);
        }

        emit ACLChange(msg.sender, msg.sender, ActionPerformed.AddTopic, topic, block.timestamp);
    }
    
    function addPub(address person, bytes32 topic) public payable {
        require(topics[topic].owner == msg.sender || topics[topic].addPubCost > 0);
        require(msg.value >= topics[topic].addPubCost);
        topics[topic].publishers[person] = true;

        if (topics[topic].addPubCost > 0) {
          topics[topic].owner.transfer(msg.value);
        }

        emit ACLChange(msg.sender, person, ActionPerformed.AddPub, topic, block.timestamp);
    }
    function addSub(address person, bytes32 topic) public payable{
        require(topics[topic].owner == msg.sender || topics[topic].addSubCost > 0);
        require(msg.value >= topics[topic].addSubCost);
        topics[topic].subscribers[person] = true;
        if (topics[topic].addSubCost > 0) {
          topics[topic].owner.transfer(msg.value);
        }
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

    function verifyPub(address person, bytes32 topic) public view returns (bool, bytes2) {
        return (topics[topic].publishers[person], topics[topic].country);
    }
    function verifySub(address person, bytes32 topic) public view returns (bool, bytes2) {
        return (topics[topic].subscribers[person], topics[topic].country);
    }
    function logAlert(address person, ActionPerformed alert, bytes32 desc) public {
        emit ACLChange(msg.sender, person, alert, desc, block.timestamp);
    }
}
