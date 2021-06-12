// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

contract Store {
	string public passphrase;
	mapping(bytes32 => bytes32) public items;

	event SetItem(bytes32 key, bytes32 value);

	constructor(string memory _passphrase) {
		passphrase = _passphrase;
	}

	function setItem(bytes32 key, bytes32 value) external {
		items[key] = value;
		emit SetItem(key, value);
	}
	
}