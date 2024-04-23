// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract DelegateCaller {
    uint256 i = 0;
    address callee;

    constructor() {
        callee = address(new DelegateCallee());
    }

    function call() public payable {
        // First, let's increase our frame count using CREATE operation to 5.
        new DelegateCallee();
        new DelegateCallee();
        new DelegateCallee();
        new DelegateCallee();
        new DelegateCallee();

        // Then we call delegatecall many times. Which will incur 7100 gas penalty every time in L2.
        for (uint256 j = 0; j < 256; j++) {
            callee.delegatecall(abi.encodeWithSelector(DelegateCallee.called.selector));
        }
    }
}

contract DelegateCallee {
    uint256 i = 0;
    constructor() {}
    function called() public {
        i += 1;
    }
}
