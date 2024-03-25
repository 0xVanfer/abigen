# abigen

## no auto generate yet

solc --abi temp/GeneratedInterface.sol -o build
cd build
abigen --abi GeneratedInterface.abi --pkg yyy --type yyy --out yyy.go

#### notes

head

```solidity
// SPDX-License-Identifier: GPL-3.0-or-later
pragma solidity ^0.8.0;
```

## Next big version:

aave aaveUiPoolDataProviderV3 merge together, use the `aaveUiPoolDataProviderV3Arbitrum`
