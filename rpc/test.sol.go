package rpc

import hex "github.com/tmthrgd/go-hex"

var Bytecode_HelloWorld = hex.MustDecodeString("608060405234801561001057600080fd5b5061011d806100206000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c8063bcdfe0d514602d575b600080fd5b603360ab565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101560715780820151818401526020810190506058565b50505050905090810190601f168015609d5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b60606040518060400160405280600c81526020017f48656c6c6f2c20576f726c64000000000000000000000000000000000000000081525090509056fea265627a7a723158200b2a5b7a53ba54371daa99f4d02346044f39e951d89df28ce64b4d36dd3fda0664736f6c634300050b0032")
var DeployedBytecode_HelloWorld = hex.MustDecodeString("6080604052348015600f57600080fd5b506004361060285760003560e01c8063bcdfe0d514602d575b600080fd5b603360ab565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101560715780820151818401526020810190506058565b50505050905090810190601f168015609d5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b60606040518060400160405280600c81526020017f48656c6c6f2c20576f726c64000000000000000000000000000000000000000081525090509056fea265627a7a723158200b2a5b7a53ba54371daa99f4d02346044f39e951d89df28ce64b4d36dd3fda0664736f6c634300050b0032")
var Abi_HelloWorld = []byte(`[{"constant":true,"inputs":[],"name":"Hello","outputs":[{"internalType":"string","name":"","type":"string"}],"payable":false,"stateMutability":"pure","type":"function"}]`)