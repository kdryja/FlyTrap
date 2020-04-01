export BLOCKCHAIN_ADDRESS="http://localhost:7545"

CLIENT_ROOT="/home/kdryja/thesis/code/mqttclient/client.go"
BLOCKCHAIN_ROOT="/home/kdryja/thesis/code/blockchain/cmd/blockchain/main.go"
FLYTRAP_CONTRACT="0x77DCf90A3E9aCffec862A65c32Af934876775Eb2"

TOPIC="MyTestTopic"

generateContract () {
  FLYTRAP_CONTRACT=$(go run $BLOCKCHAIN_ROOT -new |& awk '{print $8}')
  echo "Contract generated: $FLYTRAP_CONTRACT"
  $(go run $BLOCKCHAIN_ROOT -new_topic="Creating Topic" -topic=$TOPIC -contract=$FLYTRAP_CONTRACT)
}

splitKeys () {
  input="./pubkeys.txt"
  output="./pubkeys/"
  i=1
  while IFS= read -r line
  do
    echo "$line" > "${output}pubkey${i}.txt";
    i=$((i+1))
  done < "$input"
}

allowPublish () {
  for i in $(seq $1 $2)
  do
    key=$(cat "pubkeys/pubkey$i.txt")
    $(go run $BLOCKCHAIN_ROOT -pub="Adding Publisher#${i}" -topic=$TOPIC -contract=$FLYTRAP_CONTRACT -client=$key)
    echo "Added client as publisher: $key"
  done
}

allowSubscribe () {
  for i in $(seq $1 $2)
  do
    key=$(cat "pubkeys/pubkey$i.txt")
    $(go run $BLOCKCHAIN_ROOT -sub="Adding Subscriber#${i}" -topic=$TOPIC -contract=$FLYTRAP_CONTRACT -client=$key)
    echo "Added client as subscriber: $key"
  done
}

startPublish () {
  cd "privkeys"
  for i in $(seq $1 $2)
  do
    go run $CLIENT_ROOT -priv "privkey$i.txt" -pub 20 -topic $TOPIC &
  done
  wait
}

startSubscribe () {
  cd "privkeys"
  for i in $(seq $1 $2)
  do
    go run $CLIENT_ROOT -priv "privkey$i.txt" -sub 5 -topic $TOPIC &
  done
  wait
}

# allowPublish 1 20
# allowSubscribe 1 20

startPublish 1 10 &
startSubscribe 11 20 &
