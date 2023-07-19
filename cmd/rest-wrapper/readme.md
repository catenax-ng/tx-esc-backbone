

## Setup
```shell
W_HOME=wrapper-home
mkdir -p $W_HOME && cd $W_HOME
IMAGE_HASH=latest
docker run -it -e CONFIG=/wrapper/config/config.json -v $(pwd):/wrapper/config ghcr.io/catenax-ng/esc-res-sync-rest-wrapper:$IMAGE_HASH /bin/bash
cp default-config.json config/config.json
cd config/
echo "Create mnemonic, this is to be kept secret, from this the private key can be generated."
/wrapper/esc-backboned keys mnemonic --home $(pwd) > mnemonic
echo "Import private key with mnemonic as 'wrapper'"
cat mnemonic | /wrapper/esc-backboned keys add wrapper --keyring-backend test --home $(pwd) --recover
echo "Store public address for the faucet request."
/wrapper/esc-backboned keys show wrapper --home $(pwd) --keyring-backend test -a > pubaddr
exit
echo "Request funds at the test net."
DENOM=$(curl https://validator-webapp1-web-app.dev.demo.catena-x.net/chain/catenax-testnet-1-suggestion.json | jq ".feeCurrencies[0].coinMinimalDenom" -r)
curl -X POST --header "Content-Type: application/json" -d '{"address":"'$(cat pubaddr)'","denom":"'$DENOM'"}'  https://faucet-faucet.dev.demo.catena-x.net/
echo "Wait 10s and check the balance"
wait 10 && esc-backboned query bank balances --home $(pwd) --node https://validator2-tdmt-rpc.dev.demo.catena-x.net:443/  $(cat pubaddr)
echo "Start the wrapper"
docker run -p 7000:8080 -e CONFIG=/wrapper/config/config.json -v $(pwd):/wrapper/config ghcr.io/catenax-ng/esc-res-sync-rest-wrapper:$IMAGE_HASH
```

## Creating a resource
```shell
curl -X POST http://localhost:7000/resource -d '{"origResId":"localtest", "targetSystem":"http://localhost", "resourceKey":"my key" }'
```

## Generate openapi.json
```shell
docker run --rm -it  --user $(id -u):$(id -g) -e GOPATH=$(go env GOPATH) -v $HOME:$HOME -v $(go env GOCACHE):/.cache/go-build -w $(pwd) --entrypoint=/bin/sh quay.io/goswagger/swagger 
SWAGGER_GENERATE_EXTENSION=false swagger generate spec -m -c cmd/rest-wrapper/ -o cmd/rest-wrapper/openapi.json 
```

# Input
```
Hi Lars,
We’ve discussed the possible reliability issues with current synch solution:

- What should FC do when SD event happens, but Synch service is not available (it’s down or some network issue)
- What should Synch service do when it got SD event, but subscribed FC instance is not available (it’s down or some network issue)

To address them we’ll need a kind of reliable message broker solution supporting durable subscriptions. But then we’ll not need to communicate with Synch service via REST API at all.
Let’s agree on how will we resolve it.
```


# For discussion

## Structure
```plantuml
@startuml

title structure - current state

folder "OpCo 1"{
node val1 as "ESC-Backbone Validator 1"{
    component tm1 as "Tendermint"
    portout p2p1 as "P2P"
    tm1 -- p2p1 
    portin rpc1 as "RCP"
    tm1 -- rpc1
    database db1 as "Chain State"
    tm1 -- db1
}

node rwrap1 as "REST wrap 1"{
    portin rest1 as "REST"
    portout rpcc1 as "RCP-client"
    component restserv1 as "REST server"
    component escclient1 as "ESC Backbone client"
    restserv1 -down- escclient1
    rest1 -- restserv1
    escclient1 -- rpcc1
}
rpc1 -up- rpcc1
node FC1 as "FC 1"  
FC1 -down- rest1
}

folder "OpCo 2"{
node val2 as "ESC-Backbone Validator 2"{
    component tm2 as "Tendermint"
    portout p2p2 as "P2P"
    tm2 -- p2p2 
    portin rpc2 as "RCP"
    tm2 -- rpc2 
    database db2 as "Chain State"
    tm2 -- db2
}

node rwrap2 as "REST wrap 2"{
    portin rest2 as "REST"
    portout rpcc2 as "RCP-client"
    component restserv2 as "REST server"
    component escclient2 as "ESC Backbone client"
    restserv2 -down- escclient2
    rest2 -- restserv2
    escclient2 -- rpcc2
}
rpc2 -up- rpcc2
node FC2 as "FC 2"  
FC2 -down- rest2
}

folder "OpCo n"{
node valn as "ESC-Backbone Validator n"{
    component tmn as "Tendermint"
    portout p2pn as "P2P"
    tmn -- p2pn
    portin rpcn as "RCP"
    tmn -- rpcn
       
    database dbn as "Chain State"
    tmn -- dbn
}

node rwrapn as "REST wrap n"{
    portin restn as "REST"
    portout rpccn as "RCP-client"
    component restservn as "REST server"
    component escclientn as "ESC Backbone client"
    restservn -down- escclientn
    restn -- restservn
    escclientn -- rpccn
}
rpcn -up- rpccn
node FCn as "FC n"  
FCn -down- restn
}



p2p1 -down- p2p2
p2p2 -down- p2pn
p2p1 -down- p2pn


actor user as "FC user"
user -(0- FC1

@enduml
```

```plantuml
@startuml

title structure - with message broker 

folder "OpCo 1"{
node val1 as "ESC-Backbone Validator 1"{
    component tm1 as "Tendermint"
    portout p2p1 as "P2P"
    tm1 -- p2p1 
    portin rpc1 as "RCP"
    tm1 -- rpc1
    database db1 as "Chain State"
    tm1 -- db1
}

node mb1 as "Message Broker"{
    queue cqueue1 as "commands"
    queue equeue1 as "events"
}

node wrapper1 as "wrapper 1"{
    portout rpcc1 as "RCP-client"
    component queueclient1 as "Queue client"
    component escclient1 as "ESC Backbone client"
    queueclient1 -down- escclient1
    escclient1 -- rpcc1
}
queueclient1 <-up- cqueue1
queueclient1 -up-> equeue1

rpc1 -up- rpcc1
node FC1 as "FC 1"
FC1 -down-> cqueue1
FC1 <-down- equeue1
}

folder "OpCo 2"{
node val2 as "ESC-Backbone Validator 2"{
    component tm2 as "Tendermint"
    portout p2p2 as "P2P"
    tm2 -- p2p2 
    portin rpc2 as "RCP"
    tm2 -- rpc2
    database db2 as "Chain State"
    tm2 -- db2
}

node mb2 as "Message Broker"{
    queue cqueue2 as "commands"
    queue equeue2 as "events"
}

node wrapper2 as "wrapper 2"{
    portout rpcc2 as "RCP-client"
    component queueclient2 as "Queue client"
    component escclient2 as "ESC Backbone client"
    queueclient2 -down- escclient2
    escclient2 -- rpcc2
}
queueclient2 <-up- cqueue2
queueclient2 -up-> equeue2

rpc2 -up- rpcc2
node FC2 as "FC 2"
FC2 -down-> cqueue2
FC2 <-down- equeue2
}

folder "OpCo n"{
node valn as "ESC-Backbone Validator n"{
    component tmn as "Tendermint"
    portout p2pn as "P2P"
    tmn -- p2pn 
    portin rpcn as "RCP"
    tmn -- rpcn
    database dbn as "Chain State"
    tmn -- dbn
}

node mbn as "Message Broker"{
    queue cqueuen as "commands"
    queue equeuen as "events"
}

node wrappern as "wrapper n"{
    portout rpccn as "RCP-client"
    component queueclientn as "Queue client"
    component escclientn as "ESC Backbone client"
    queueclientn -down- escclientn
    escclientn -- rpccn
}
queueclientn <-up- cqueuen
queueclientn -up-> equeuen

rpcn -up- rpccn
node FCn as "FC n"
FCn -down-> cqueuen
FCn <-down- equeuen
}


p2p1 -down- p2p2
p2p2 -down- p2pn
p2p1 -down- p2pn


actor user as "FC user"
user -(0- FC1

@enduml
```

## Creating and processing change events

### current approach
```plantuml
@startuml
skinparam backgroundColor #CCCCCC
skinparam style strictuml
skinparam sequenceMessageAlign center
skinparam BoxPadding 0
hide footbox
!pragma teoz true
title current state


box "OpCo 1" #white
participant FC1 as "FC 1"
participant rwrap1 as "REST wrap 1"
participant wrapper1 as "wrapper 1"
end box

collections vals as "ESC-Backbone validators"
activate vals

box "OpCo 2" #white
participant wrapper2 as "wrapper 2"
participant rwrap2 as "REST wrap 2"
participant FC2 as "FC 2"
end box


== creating events ==
[--> FC1: "SD(DID)"
FC1 -> rwrap1 ++: "CHANGE(SD-HASH)"
rwrap1 -> wrapper1 ++: "REST(CHANGE(SD-HASH))"
wrapper1 --> vals: "TX(CHANGE(SD-HASH))"
return
return

vals -->] : "EVENT(SD-HASH)"

== processing events ==

activate FC1  
activate FC2
FC1 -> rwrap1 ++: watch for event
& FC2 -> rwrap2 ++: watch for event
rwrap1 -> wrapper1: subscribe on valdiator
& rwrap2 -> wrapper2: subscribe on valdiator

vals --> wrapper1 ++: "EVENT(SD-HASH)"
& vals --> wrapper2 ++: "EVENT(SD-HASH)"
wrapper1 --> rwrap1 ++: "EVENT(SD-HASH)"
& wrapper2 --> rwrap2 ++: "EVENT(SD-HASH)"
rwrap1 --> FC1 ++: "EVENT(SD-HASH)"
& rwrap2 --> FC2 ++: "EVENT(SD-HASH)"
deactivate rwrap1
deactivate rwrap2
FC1 -> FC1 ++ : knows the entry already
& FC2 -> FC2 ++: fetch the entry update
deactivate FC1
deactivate FC1
|||
alt Source FC available
FC2 -> FC1 ++: "GET(SD-HASH)" at targetsystem FC1
return "SD(DID)"

deactivate FC1
else Source FC unavailable


[--> FC1 !!: interuption

FC2 -> FC1: "GET(SD-HASH)" at targetsystem FC1
FC1 -> FC2: fail
note right: handle retry at FC
FC2 -> FC1: "GET(SD-HASH)" at targetsystem FC1
FC1 -> FC2: fail
note right: /!\\ data unavailable 
deactivate FC2
deactivate FC2

end 
```

### with message broker (not implemented)
same issus, if the `targetSystem` of the event is unavailable a sync is not possible.
```plantuml
@startuml
skinparam backgroundColor #CCCCCC
skinparam style strictuml
skinparam sequenceMessageAlign center
skinparam BoxPadding 0
hide footbox
!pragma teoz true
title with message broker
actor user as "FC user"
box "OpCo 1" #white
participant FC1 as "FC 1"
box "message brocker 1" #EEEEEE
queue equeue1 as "events 1"
queue cqueue1 as "commands 1"
end box
participant wrapper1 as "wrapper 1"
end box

collections vals as "ESC-Backbone validators"
activate vals

box "OpCo 2" #white
participant wrapper2 as "wrapper 2"
box "message brocker 2" #EEEEEE
queue cqueue2 as "commands 2"
queue equeue2 as "events 2"
end box
participant FC2 as "FC 2"

activate wrapper1
activate wrapper2
wrapper1 -> cqueue1: watch for command
& wrapper2 -> cqueue2: watch for command
activate FC1  
activate FC2
FC1 -> equeue1: watch for event
& FC2 -> equeue2: watch for event

== creating events ==
[--> FC1: "SD(DID)"
FC1 --> cqueue1: "CHANGE(SD-HASH)"
cqueue1 --> wrapper1 ++: "CHANGE(SD-HASH)"
wrapper1 --> vals: "TX(CHANGE(SD-HASH))"
deactivate wrapper1

vals -->] : "EVENT(SD-HASH)"

== processing events ==

vals --> wrapper1 ++: "EVENT(SD-HASH)"
& vals --> wrapper2 ++: "EVENT(SD-HASH)"
wrapper1 --> equeue1: "EVENT(SD-HASH)"
& wrapper2 --> equeue2: "EVENT(SD-HASH)"
equeue1 --> FC1 ++: "EVENT(SD-HASH)"
& equeue2 --> FC2 ++: "EVENT(SD-HASH)"

FC1 -> FC1 ++ : knows the entry already
& FC2 -> FC2 ++: fetch the entry update
deactivate FC1
deactivate FC1
alt Source FC available
FC2 -> FC1 ++: "GET(SD-HASH)" at targetsystem FC1
return "SD(DID)"

deactivate FC1
else Source FC unavailable


[--> FC1 !!: interuption

FC2 -> FC1: "GET(SD-HASH)" at targetsystem FC1
FC1 -> FC2: fail
deactivate FC2
FC2 -> equeue2: put back message
deactivate FC2
equeue2 --> FC2 ++: "EVENT(SD-HASH)"
note right: handle retry by reissuing the message
FC2 -> FC2 ++: fetch the entry update
FC2 -> FC1: "GET(SD-HASH)" at targetsystem FC1
FC1 -> FC2: fail
deactivate FC2
note right: /!\\ data unavailable
FC2 -> equeue2: put back message
deactivate FC2
end box


```