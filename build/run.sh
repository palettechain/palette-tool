#!/bin/bash

# setup direction used to copy files
setupDir=$1;
# select function: init, start;
fn=$2;
# node index
nodeIdx=$3;

mkdir -p $setupDir;

node="node$nodeIdx";
genesisFile="$setupDir/genesis.json";
staticFile="$setupDir/static-nodes.json";
nodeFile="$setupDir/nodekeys/$node/nodekey";

function checksetup() 
{
    if [ ! -d "$setupDir" ]; then
        echo "direction of setup not exist!";
        return 0;
    fi

    if [ ! -x "$setupDir" ]; then
        echo "direction of setup can't be access!";
        return 0;
    fi

    if [ ! -f "$genesisFile" ]; then
        echo "file of genesis.json not exist in setup direction";
        return 0;
    fi

    if [ ! -f "$staticFile" ]; then
        echo "file of static-nodes.json not exist in setup direction";
        return 0;
    fi

    if [ ! -f "$nodeFile" ]; then
        echo "file of nodekey not exist or nodekey path invalid, correction path should be setup/nodekeys/node0/nodekey";
        return 0;
    fi

    return 1;
}

function checkNode()
{
    if [ ! -d "$node" ]; then
        echo "create $node direction";
        mkdir $node;
        mkdir -p $node/data/geth;
    fi
}

function execinit() 
{
    echo "init palette genesis node";
    checkNode;

    checksetup;
    if [[ $? -eq 0 ]]; then
        exit 1;
    fi

    cp -r $genesisFile $node;
    cp -r $staticFile $node/data;
    cp -r $nodeFile $node/data/geth;

    geth --datadir $node/data init $node/genesis.json;
}

networkID=$4;
loglevel=$5;
rpcPort=$6;
p2pPort=$7;
function execstart()
{
    echo "start palette genesis $node, networkID $networkID, rpc port $rpcPort, p2p port $p2pPort";
    
    cd $node;
    PRIVATE_CONFIG=ignore nohup geth \
    --datadir data \
    --nodiscover \
    --syncmode full \
    --mine --minerthreads 1 \
    --verbosity $loglevel \
    --networkid $networkID \
    --rpc --rpcaddr 0.0.0.0 \
    --rpcport $rpcPort \
    --rpcapi admin,db,eth,debug,miner,net,shh,txpool,personal,web3,quorum,istanbul \
    --emitcheckpoints \
    --port $p2pPort > node.log 2>&1 &
}

case $fn
in
    init) execinit ;;
    start) execstart ;;
    stop) execstop ;;
    *) echo "Nothing to do"
       exit ;;
esac