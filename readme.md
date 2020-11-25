## Palette-Tool

It is a tool to help generate palette node and build palette network.

## genesis network
in order to build genesis network. first, you should prepare items as follow:
* several node keys
* genesis file
* static nodes json file

1.run command like this:
```bash
./bin/palette-tool-darwin-amd64 setup --nodes --num 4 --verbose --save
```
This command will generate 4 points nodekey and validator files, these files are stored in the directory of `setup/nodekeys`.

2.generate accounts:
first make a new direction to store you account json files.
```bash
mkdir keystore
```
generate admin account, run command like this:
```bash
ubuntu@10-23-76-181:~$ geth account new --keystore keystore
```
input your password of admin account, and DONT forget it, it's very important!!!
```dat
INFO [11-25|11:35:15.316] Maximum peer count                       ETH=50 LES=0 total=50
Your new account is locked with a password. Please give a password. Do not forget this password.
Password: 
Repeat password: 

Your new key was generated

Public address of the key:   0xA******1
Path of the secret key file: .../keystore/UTC--******1

- You can share your public address with anyone. Others need it to interact with you.
- You must NEVER share the secret key with anyone! The key controls access to your funds!
- You must BACKUP your key file! Without the key, it's impossible to access account funds!
- You must REMEMBER your password! Without the password, it's impossible to decrypt the key!
```

generate base reward pool account, the procedure is just like how to generate admin account.

3.modify genesis.json
open the file of `setup/genesis.json`, replace the `admin` public address and the `base reward pool` public address.

4.modify static-nodes.json
open the file of `setup/static-nodes.json`, set the right host and ip as you like.
e.g: update the first node info, modify the host to 10.162.75.23, and modify the p2p port to 30304.
```dat
"enode://1****6@10.162.75.23:30304?discport=0",
```

5.execute the script
login someone linux machine and copy the correct `nodekey` file, `genesis.json` file and `static-nodes.json` file into direction of `~/setup`, create the director if not exist:

init genesis node:
```bash
./run.sh 0 /home/ubuntu/palette/setup init 
```
run this command for other 3 nodes on different machine, you can also run them in same machine too.

start genesis node:
```bash
./run.sh 0 /home/ubuntu/palette/setup start 4 18892 20200 30300
```
the 1st param `0` denote the index of node. e.g: `node0`'s index is 0;
the 2nd param `/home/ubuntu/palette/setup` is the absolute location of `setup` which generated before;
the 3rd param `start` denote run this node;
the 4th param `4` denote log level, 4: debug, 5: info;
the 5th param `18892` denote palette chain networkID;
the 6th param `20200` denote the RPC port for this node;
the 7th param `30300` denote the P2P port for this node;

if you want to creat all genesis nodes in the same machine:
```bash
./run.sh setup init node0
./run.sh setup init node1
./run.sh setup init node2
./run.sh setup init node3
./run.sh setup init node4

./run.sh setup start node0 4 103 20200 30300
./run.sh setup start node1 4 103 20201 30301
./run.sh setup start node2 4 103 20202 30302
./run.sh setup start node3 4 103 20203 30303
./run.sh setup start node4 4 103 20204 30304
```