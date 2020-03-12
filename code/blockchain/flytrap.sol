pragma solidity >=0.4.22 <0.7.0;

contract Flytrap {
    address payable owner;
    uint addTopicCost;

    struct Topic {
        bool isValue;
        bool sensitive;
        string name;
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
      Banned,
      Summary
    }

    event ACLChange (
      address indexed _src,
      address indexed _target,
      ActionPerformed indexed _action,
      string _name,
      string reason,
      uint _timestamp
    );

    mapping(string => Topic) public topics;

    constructor(uint cost) public {
        owner = msg.sender;
        addTopicCost = cost;
    }

    function addTopic(string memory topic, bytes2 country, uint addPubCost, uint addSubCost, string memory reason, bool sensitive) public payable {
        require(topics[topic].isValue == false);
        require(msg.value >= addTopicCost);
        topics[topic] = Topic(true, sensitive, topic, addPubCost, addSubCost, msg.sender, country);

        if (addTopicCost > 0) {
          owner.transfer(msg.value);
        }

        emit ACLChange(msg.sender, msg.sender, ActionPerformed.AddTopic, topic, reason, block.timestamp);
    }
    
    function addPub(address person, string memory topic, string memory reason) public payable {
        require(topics[topic].owner == msg.sender || topics[topic].addPubCost > 0);
        require(msg.value >= topics[topic].addPubCost);
        topics[topic].publishers[person] = true;

        if (topics[topic].addPubCost > 0) {
          topics[topic].owner.transfer(msg.value);
        }

        emit ACLChange(msg.sender, person, ActionPerformed.AddPub, topic, reason, block.timestamp);
    }
    function addSub(address person, string memory topic, string memory reason) public payable{
        require(topics[topic].owner == msg.sender || topics[topic].addSubCost > 0);
        require(msg.value >= topics[topic].addSubCost);
        topics[topic].subscribers[person] = true;
        if (topics[topic].addSubCost > 0) {
          topics[topic].owner.transfer(msg.value);
        }
        emit ACLChange(msg.sender, person, ActionPerformed.AddSub, topic, reason, block.timestamp);
    }
    function revokePub(address person, string memory topic, string memory reason) public {
        require(topics[topic].owner == msg.sender);
        topics[topic].publishers[person] = false;
        emit ACLChange(msg.sender, person, ActionPerformed.RevokePub, topic, reason, block.timestamp);
    }
    function revokeSub(address person, string memory topic, string memory reason) public {
        require(topics[topic].owner == msg.sender);
        topics[topic].subscribers[person] = false;
        emit ACLChange(msg.sender, person, ActionPerformed.RevokeSub, topic, reason, block.timestamp);
    }

    function verifyPub(address person, string memory topic) public view returns (bool, bytes2, bool) {
        return (topics[topic].publishers[person], topics[topic].country, topics[topic].sensitive);
    }
    function verifySub(address person, string memory topic) public view returns (bool, bytes2, bool) {
        return (topics[topic].subscribers[person], topics[topic].country, topics[topic].sensitive);
    }
    function logAlert(address person, ActionPerformed alert, string memory desc, string memory reason) public {
        emit ACLChange(msg.sender, person, alert, desc, reason, block.timestamp);
    }
}
