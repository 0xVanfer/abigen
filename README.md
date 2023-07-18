# abigen

## no auto generate yet

solc --abi temp/xxx.sol -o build
cd build
abigen --abi xxx.abi --pkg yyy --type yyy --out yyy.go
