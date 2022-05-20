# myGoOracle

## Description

It is a self-constructed Oracle service implemented by GO. It is mainly implemented through subscribing the event to accept the smart contract call, then our Oracle service calls the callback function provided by the smart contract. In the project, we call the Oracle service to get a random number. (ps: This project should be run in conjunction with the [DiceMyOracle](https://github.com/pinkyTseng/DiceMyOracle) project).

## How to start up
Consider simplicity and adjustability. Reduce most of the tedious operations and implement an automated deploy flow by setting one environment variable "MyGethOracleSetting" and a predefined setting structure (reference the "settingExample" directory). 

- Pre requirements
    - installed the Truffle
    - -installed the Ganache (the Ganache UI version seems to not support the ws protocol, so if you want to run on Ganache, don't use the UI version)

- The steps
    1. set up the "MyGethOracleSetting" environment variable. It defines the setting files directory path. (ps: it needs to end up with "/", both projects need the variable)
    2. copy all files and content in the "settingExample" directory to the "MyGethOracleSetting" directory    
    3. modify the files in the directory defined by "MyGethOracleSetting" according to your requirements.
    4. start the ganache, then modify the "DeployerPrivateKey" in the privateKey.json file. (only run with ganache need the step)
    5. cd to "DiceMyOracle" project, run truffle migrate --network XXX --reset, and wait for it to be done.
    6. cd to "myGoOracle" project, run the Oracle service  

- The directory structure of MyGethOracleSetting
    - ContractAddresses.json: it is read and written by both project programs automatically, no need to modify it.
    - privateKey.json: "DeployerPrivateKey" is your account's private key; "INFURA_PROJECT_ID" doesn't need to be used and modified when running with Ganache.
    - network/XXX.json: for each network setting file, only use the file you are running on which network (ex: run Ganache, only use network/ganache.json). You can extend new network setting files in the directory.
    - networkSetting.txt: defines what file in the "network" directory to use. The content in the file should be the same as the file name in the "network" directory.


