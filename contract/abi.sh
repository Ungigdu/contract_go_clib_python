cd ~/workspace/go/src/github.com/Ungigdu/contract_go_clib_python/contract/
rm ./SimpleDemo.abi
rm ./SimpleDemo.go
solc ./DemoContract.sol --abi -o ./
abigen --abi SimpleDemo.abi --pkg contract --type SimpleDemo --out SimpleDemo.go