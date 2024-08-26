// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ERC20drink

import (
	"math/big"
	"strings"

	"github.com/FISCO-BCOS/go-sdk/abi"
	"github.com/FISCO-BCOS/go-sdk/abi/bind"
	"github.com/FISCO-BCOS/go-sdk/core/types"
	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
)

// ERC20drinkABI is the input ABI used to generate the binding from.
const ERC20drinkABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"delegate\",\"type\":\"address\"},{\"name\":\"numTokens\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"drinkname\",\"type\":\"string\"},{\"name\":\"num\",\"type\":\"uint256\"}],\"name\":\"_buyDrink\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"buyer\",\"type\":\"address\"},{\"name\":\"numTokens\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"add\",\"type\":\"address\"}],\"name\":\"calEgg\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"add\",\"type\":\"address\"}],\"name\":\"calBuyerDrink\",\"outputs\":[{\"name\":\"drinkname\",\"type\":\"string\"},{\"name\":\"drinknumber\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenOwner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\"},{\"name\":\"numTokens\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"eggNum\",\"type\":\"uint256\"}],\"name\":\"PointSwap\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"add\",\"type\":\"address\"}],\"name\":\"calPoint\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"delegate\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"price\",\"type\":\"uint256\"},{\"name\":\"number\",\"type\":\"uint256\"}],\"name\":\"_updateDrink\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"drinkna\",\"type\":\"string\"}],\"name\":\"calDrinkNum\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"drinkna1\",\"type\":\"string\"}],\"name\":\"calDrinkPrice\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"drinkname\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"num\",\"type\":\"uint256\"}],\"name\":\"_drinkBuy\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

// ERC20drinkBin is the compiled bytecode used for deploying new contracts.
var ERC20drinkBin = "0x6080604052620f42406000557360e9635e225335a41a4e616dfa7c352fee641d8b600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506040805190810160405280600781526020017f6d6169446f6e6700000000000000000000000000000000000000000000000000815250600a9080519060200190620000ad929190620009c6565b506040805190810160405280600981526020017f77616e674c616f4a690000000000000000000000000000000000000000000000815250600b9080519060200190620000fb929190620009c6565b506040805190810160405280600581526020017f6c76436861000000000000000000000000000000000000000000000000000000815250600c908051906020019062000149929190620009c6565b506040805190810160405280600781526020017f686f6e6743686100000000000000000000000000000000000000000000000000815250600d908051906020019062000197929190620009c6565b506040805190810160405280600c81526020017f7869616e4368656e675a68690000000000000000000000000000000000000000815250600e9080519060200190620001e5929190620009c6565b506040805190810160405280600581526020017f796942616f000000000000000000000000000000000000000000000000000000815250600f908051906020019062000233929190620009c6565b506040805190810160405280601081526020017f79696e6759616e674b7561695869616e000000000000000000000000000000008152506010908051906020019062000281929190620009c6565b503480156200028f57600080fd5b5060005460026000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555060646004600a60405180828054600181600116156101000203166002900480156200035d5780601f106200033a5761010080835404028352918201916200035d565b820191906000526020600020905b81548152906001019060200180831162000348575b505091505090815260200160405180910390208190555060646004600b6040518082805460018160011615610100020316600290048015620003d95780601f10620003b6576101008083540402835291820191620003d9565b820191906000526020600020905b815481529060010190602001808311620003c4575b505091505090815260200160405180910390208190555060646004600c6040518082805460018160011615610100020316600290048015620004555780601f106200043257610100808354040283529182019162000455565b820191906000526020600020905b81548152906001019060200180831162000440575b505091505090815260200160405180910390208190555060646004600d6040518082805460018160011615610100020316600290048015620004d15780601f10620004ae576101008083540402835291820191620004d1565b820191906000526020600020905b815481529060010190602001808311620004bc575b505091505090815260200160405180910390208190555060646004600e60405180828054600181600116156101000203166002900480156200054d5780601f106200052a5761010080835404028352918201916200054d565b820191906000526020600020905b81548152906001019060200180831162000538575b505091505090815260200160405180910390208190555060646004600f6040518082805460018160011615610100020316600290048015620005c95780601f10620005a6576101008083540402835291820191620005c9565b820191906000526020600020905b815481529060010190602001808311620005b4575b50509150509081526020016040518091039020819055506064600460106040518082805460018160011615610100020316600290048015620006455780601f106200062257610100808354040283529182019162000645565b820191906000526020600020905b81548152906001019060200180831162000630575b505091505090815260200160405180910390208190555060286005600a6040518082805460018160011615610100020316600290048015620006c15780601f106200069e576101008083540402835291820191620006c1565b820191906000526020600020905b815481529060010190602001808311620006ac575b5050915050908152602001604051809103902081905550601e6005600b60405180828054600181600116156101000203166002900480156200073d5780601f106200071a5761010080835404028352918201916200073d565b820191906000526020600020905b81548152906001019060200180831162000728575b5050915050908152602001604051809103902081905550601e6005600c6040518082805460018160011615610100020316600290048015620007b95780601f1062000796576101008083540402835291820191620007b9565b820191906000526020600020905b815481529060010190602001808311620007a4575b5050915050908152602001604051809103902081905550601e6005600d6040518082805460018160011615610100020316600290048015620008355780601f106200081257610100808354040283529182019162000835565b820191906000526020600020905b81548152906001019060200180831162000820575b505091505090815260200160405180910390208190555060146005600e6040518082805460018160011615610100020316600290048015620008b15780601f106200088e576101008083540402835291820191620008b1565b820191906000526020600020905b8154815290600101906020018083116200089c575b5050915050908152602001604051809103902081905550600a6005600f60405180828054600181600116156101000203166002900480156200092d5780601f106200090a5761010080835404028352918201916200092d565b820191906000526020600020905b81548152906001019060200180831162000918575b50509150509081526020016040518091039020819055506028600560106040518082805460018160011615610100020316600290048015620009a95780601f1062000986576101008083540402835291820191620009a9565b820191906000526020600020905b81548152906001019060200180831162000994575b505091505090815260200160405180910390208190555062000a75565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1062000a0957805160ff191683800117855562000a3a565b8280016001018555821562000a3a579182015b8281111562000a3957825182559160200191906001019062000a1c565b5b50905062000a49919062000a4d565b5090565b62000a7291905b8082111562000a6e57600081600090555060010162000a54565b5090565b90565b611d558062000a856000396000f3006080604052600436106100f1576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806306fdde03146100f6578063095ea7b3146101865780630aa39505146101eb57806318160ddd1461027657806323b872dd146102a1578063313ce5671461032657806337cc77f6146103575780636390b023146103ae57806370a082311461047157806395d89b41146104c8578063a9059cbb14610558578063c1fbffad146105bd578063cc2a1eb214610602578063dd62ed3e14610659578063e095c25c146106d0578063e29d93be14610765578063ee5db636146107e2575b600080fd5b34801561010257600080fd5b5061010b61085f565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561014b578082015181840152602081019050610130565b50505050905090810190601f1680156101785780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561019257600080fd5b506101d1600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610898565b604051808215151515815260200191505060405180910390f35b3480156101f757600080fd5b5061025c600480360381019080803590602001908201803590602001908080601f01602080910402602001604051908101604052809392919081815260200183838082843782019150505050505091929192908035906020019092919050505061098a565b604051808215151515815260200191505060405180910390f35b34801561028257600080fd5b5061028b610e92565b6040518082815260200191505060405180910390f35b3480156102ad57600080fd5b5061030c600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610e9b565b604051808215151515815260200191505060405180910390f35b34801561033257600080fd5b5061033b6111ef565b604051808260ff1660ff16815260200191505060405180910390f35b34801561036357600080fd5b50610398600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506111f4565b6040518082815260200191505060405180910390f35b3480156103ba57600080fd5b506103ef600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061123d565b6040518080602001838152602001828103825284818151815260200191508051906020019080838360005b8381101561043557808201518184015260208101905061041a565b50505050905090810190601f1680156104625780820380516001836020036101000a031916815260200191505b50935050505060405180910390f35b34801561047d57600080fd5b506104b2600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506115db565b6040518082815260200191505060405180910390f35b3480156104d457600080fd5b506104dd611624565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561051d578082015181840152602081019050610502565b50505050905090810190601f16801561054a5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561056457600080fd5b506105a3600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919050505061165d565b604051808215151515815260200191505060405180910390f35b3480156105c957600080fd5b506105e860048036038101908080359060200190929190505050611826565b604051808215151515815260200191505060405180910390f35b34801561060e57600080fd5b50610643600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061198f565b6040518082815260200191505060405180910390f35b34801561066557600080fd5b506106ba600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506119d8565b6040518082815260200191505060405180910390f35b3480156106dc57600080fd5b5061074b600480360381019080803590602001908201803590602001908080601f01602080910402602001604051908101604052809392919081815260200183838082843782019150505050505091929192908035906020019092919080359060200190929190505050611a5f565b604051808215151515815260200191505060405180910390f35b34801561077157600080fd5b506107cc600480360381019080803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050611b9a565b6040518082815260200191505060405180910390f35b3480156107ee57600080fd5b50610849600480360381019080803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050611c0f565b6040518082815260200191505060405180910390f35b6040805190810160405280600981526020017f4472696e6b5f473132000000000000000000000000000000000000000000000081525081565b600081600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925846040518082815260200191505060405180910390a36001905092915050565b6000816004846040518082805190602001908083835b6020831015156109c557805182526020820191506020810190506020830392506109a0565b6001836020036101000a03801982511681845116808217855250505050505090500191505090815260200160405180910390205410151515610a0657600080fd5b6005836040518082805190602001908083835b602083101515610a3e5780518252602082019150602081019050602083039250610a19565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040518091039020548202600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205410151515610ac157600080fd5b610b4c7360e9635e225335a41a4e616dfa7c352fee641d8b6005856040518082805190602001908083835b602083101515610b115780518252602082019150602081019050602083039250610aec565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902054840261165d565b50816004846040518082805190602001908083835b602083101515610b865780518252602082019150602081019050602083039250610b61565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902054036004846040518082805190602001908083835b602083101515610bf35780518252602082019150602081019050602083039250610bce565b6001836020036101000a03801982511681845116808217855250505050505090500191505090815260200160405180910390208190555082600660003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209080519060200190610c7d929190611c84565b5081600760003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020846040518082805190602001908083835b602083101515610cf45780518252602082019150602081019050602083039250610ccf565b6001836020036101000a03801982511681845116808217855250505050505090500191505090815260200160405180910390208190555081600860003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205401600860003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055507f3d793e2ba72604006e0c2d8656f1af9b2acece635ad25c0c9787877239f33f31338484604051808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200180602001838152602001828103825284818151815260200191508051906020019080838360005b83811015610e4c578082015181840152602081019050610e31565b50505050905090810190601f168015610e795780820380516001836020036101000a031916815260200191505b5094505050505060405180910390a16001905092915050565b60008054905090565b6000600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020548211151515610eeb57600080fd5b600360008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020548211151515610f7657600080fd5b81600260008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205403600260008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555081600360008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205403600360008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555081600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205401600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040518082815260200191505060405180910390a3600190509392505050565b600281565b6000600960008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b60606000600660008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156113145780601f106112e957610100808354040283529160200191611314565b820191906000526020600020905b8154815290600101906020018083116112f757829003601f168201915b50505050509150600760008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600660008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060405180828054600181600116156101000203166002900480156113f45780601f106113d25761010080835404028352918201916113f4565b820191906000526020600020905b8154815290600101906020018083116113e0575b50509150509081526020016040518091039020549050600660008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600760008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600660008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060405180828054600181600116156101000203166002900480156115225780601f10611500576101008083540402835291820191611522565b820191906000526020600020905b81548152906001019060200180831161150e575b5050915050908152602001604051809103902054818054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156115cb5780601f106115a0576101008083540402835291602001916115cb565b820191906000526020600020905b8154815290600101906020018083116115ae57829003601f168201915b5050505050915091509150915091565b6000600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b6040805190810160405280600481526020017f445231320000000000000000000000000000000000000000000000000000000081525081565b6000600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205482111515156116ad57600080fd5b81600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205403600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555081600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205401600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040518082815260200191505060405180910390a36001905092915050565b6000600a8202600860003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541015151561187957600080fd5b600a8202600860003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205403600860003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555081600960003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205401600960003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555060019050919050565b6000600860008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b6000600360008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b60007360e9635e225335a41a4e616dfa7c352fee641d8b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515611aaf57600080fd5b826005856040518082805190602001908083835b602083101515611ae85780518252602082019150602081019050602083039250611ac3565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902081905550816004856040518082805190602001908083835b602083101515611b585780518252602082019150602081019050602083039250611b33565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902081905550600190509392505050565b60006004826040518082805190602001908083835b602083101515611bd45780518252602082019150602081019050602083039250611baf565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040518091039020549050919050565b60006005826040518082805190602001908083835b602083101515611c495780518252602082019150602081019050602083039250611c24565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040518091039020549050919050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10611cc557805160ff1916838001178555611cf3565b82800160010185558215611cf3579182015b82811115611cf2578251825591602001919060010190611cd7565b5b509050611d009190611d04565b5090565b611d2691905b80821115611d22576000816000905550600101611d0a565b5090565b905600a165627a7a72305820948a671c18a978199015441740c1a4822fc4cad83cbb175d9eb6981d9b5b8ab30029"

// DeployERC20drink deploys a new contract, binding an instance of ERC20drink to it.
func DeployERC20drink(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ERC20drink, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20drinkABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC20drinkBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20drink{ERC20drinkCaller: ERC20drinkCaller{contract: contract}, ERC20drinkTransactor: ERC20drinkTransactor{contract: contract}, ERC20drinkFilterer: ERC20drinkFilterer{contract: contract}}, nil
}

func AsyncDeployERC20drink(auth *bind.TransactOpts, handler func(*types.Receipt, error), backend bind.ContractBackend) (*types.Transaction, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20drinkABI))
	if err != nil {
		return nil, err
	}

	tx, err := bind.AsyncDeployContract(auth, handler, parsed, common.FromHex(ERC20drinkBin), backend)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

// ERC20drink is an auto generated Go binding around a Solidity contract.
type ERC20drink struct {
	ERC20drinkCaller     // Read-only binding to the contract
	ERC20drinkTransactor // Write-only binding to the contract
	ERC20drinkFilterer   // Log filterer for contract events
}

// ERC20drinkCaller is an auto generated read-only Go binding around a Solidity contract.
type ERC20drinkCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20drinkTransactor is an auto generated write-only Go binding around a Solidity contract.
type ERC20drinkTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20drinkFilterer is an auto generated log filtering Go binding around a Solidity contract events.
type ERC20drinkFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20drinkSession is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type ERC20drinkSession struct {
	Contract     *ERC20drink       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20drinkCallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type ERC20drinkCallerSession struct {
	Contract *ERC20drinkCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ERC20drinkTransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type ERC20drinkTransactorSession struct {
	Contract     *ERC20drinkTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ERC20drinkRaw is an auto generated low-level Go binding around a Solidity contract.
type ERC20drinkRaw struct {
	Contract *ERC20drink // Generic contract binding to access the raw methods on
}

// ERC20drinkCallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type ERC20drinkCallerRaw struct {
	Contract *ERC20drinkCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20drinkTransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type ERC20drinkTransactorRaw struct {
	Contract *ERC20drinkTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20drink creates a new instance of ERC20drink, bound to a specific deployed contract.
func NewERC20drink(address common.Address, backend bind.ContractBackend) (*ERC20drink, error) {
	contract, err := bindERC20drink(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20drink{ERC20drinkCaller: ERC20drinkCaller{contract: contract}, ERC20drinkTransactor: ERC20drinkTransactor{contract: contract}, ERC20drinkFilterer: ERC20drinkFilterer{contract: contract}}, nil
}

// NewERC20drinkCaller creates a new read-only instance of ERC20drink, bound to a specific deployed contract.
func NewERC20drinkCaller(address common.Address, caller bind.ContractCaller) (*ERC20drinkCaller, error) {
	contract, err := bindERC20drink(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20drinkCaller{contract: contract}, nil
}

// NewERC20drinkTransactor creates a new write-only instance of ERC20drink, bound to a specific deployed contract.
func NewERC20drinkTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20drinkTransactor, error) {
	contract, err := bindERC20drink(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20drinkTransactor{contract: contract}, nil
}

// NewERC20drinkFilterer creates a new log filterer instance of ERC20drink, bound to a specific deployed contract.
func NewERC20drinkFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20drinkFilterer, error) {
	contract, err := bindERC20drink(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20drinkFilterer{contract: contract}, nil
}

// bindERC20drink binds a generic wrapper to an already deployed contract.
func bindERC20drink(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20drinkABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20drink *ERC20drinkRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20drink.Contract.ERC20drinkCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20drink *ERC20drinkRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _ERC20drink.Contract.ERC20drinkTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20drink *ERC20drinkRaw) TransactWithResult(opts *bind.TransactOpts, result interface{}, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _ERC20drink.Contract.ERC20drinkTransactor.contract.TransactWithResult(opts, result, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20drink *ERC20drinkCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20drink.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20drink *ERC20drinkTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _ERC20drink.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20drink *ERC20drinkTransactorRaw) TransactWithResult(opts *bind.TransactOpts, result interface{}, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _ERC20drink.Contract.contract.TransactWithResult(opts, result, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address delegate) constant returns(uint256)
func (_ERC20drink *ERC20drinkCaller) Allowance(opts *bind.CallOpts, owner common.Address, delegate common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20drink.contract.Call(opts, out, "allowance", owner, delegate)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address delegate) constant returns(uint256)
func (_ERC20drink *ERC20drinkSession) Allowance(owner common.Address, delegate common.Address) (*big.Int, error) {
	return _ERC20drink.Contract.Allowance(&_ERC20drink.CallOpts, owner, delegate)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address delegate) constant returns(uint256)
func (_ERC20drink *ERC20drinkCallerSession) Allowance(owner common.Address, delegate common.Address) (*big.Int, error) {
	return _ERC20drink.Contract.Allowance(&_ERC20drink.CallOpts, owner, delegate)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address tokenOwner) constant returns(uint256)
func (_ERC20drink *ERC20drinkCaller) BalanceOf(opts *bind.CallOpts, tokenOwner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20drink.contract.Call(opts, out, "balanceOf", tokenOwner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address tokenOwner) constant returns(uint256)
func (_ERC20drink *ERC20drinkSession) BalanceOf(tokenOwner common.Address) (*big.Int, error) {
	return _ERC20drink.Contract.BalanceOf(&_ERC20drink.CallOpts, tokenOwner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address tokenOwner) constant returns(uint256)
func (_ERC20drink *ERC20drinkCallerSession) BalanceOf(tokenOwner common.Address) (*big.Int, error) {
	return _ERC20drink.Contract.BalanceOf(&_ERC20drink.CallOpts, tokenOwner)
}

// CalBuyerDrink is a free data retrieval call binding the contract method 0x6390b023.
//
// Solidity: function calBuyerDrink(address add) constant returns(string drinkname, uint256 drinknumber)
func (_ERC20drink *ERC20drinkCaller) CalBuyerDrink(opts *bind.CallOpts, add common.Address) (struct {
	Drinkname   string
	Drinknumber *big.Int
}, error) {
	ret := new(struct {
		Drinkname   string
		Drinknumber *big.Int
	})
	out := ret
	err := _ERC20drink.contract.Call(opts, out, "calBuyerDrink", add)
	return *ret, err
}

// CalBuyerDrink is a free data retrieval call binding the contract method 0x6390b023.
//
// Solidity: function calBuyerDrink(address add) constant returns(string drinkname, uint256 drinknumber)
func (_ERC20drink *ERC20drinkSession) CalBuyerDrink(add common.Address) (struct {
	Drinkname   string
	Drinknumber *big.Int
}, error) {
	return _ERC20drink.Contract.CalBuyerDrink(&_ERC20drink.CallOpts, add)
}

// CalBuyerDrink is a free data retrieval call binding the contract method 0x6390b023.
//
// Solidity: function calBuyerDrink(address add) constant returns(string drinkname, uint256 drinknumber)
func (_ERC20drink *ERC20drinkCallerSession) CalBuyerDrink(add common.Address) (struct {
	Drinkname   string
	Drinknumber *big.Int
}, error) {
	return _ERC20drink.Contract.CalBuyerDrink(&_ERC20drink.CallOpts, add)
}

// CalDrinkNum is a free data retrieval call binding the contract method 0xe29d93be.
//
// Solidity: function calDrinkNum(string drinkna) constant returns(uint256)
func (_ERC20drink *ERC20drinkCaller) CalDrinkNum(opts *bind.CallOpts, drinkna string) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20drink.contract.Call(opts, out, "calDrinkNum", drinkna)
	return *ret0, err
}

// CalDrinkNum is a free data retrieval call binding the contract method 0xe29d93be.
//
// Solidity: function calDrinkNum(string drinkna) constant returns(uint256)
func (_ERC20drink *ERC20drinkSession) CalDrinkNum(drinkna string) (*big.Int, error) {
	return _ERC20drink.Contract.CalDrinkNum(&_ERC20drink.CallOpts, drinkna)
}

// CalDrinkNum is a free data retrieval call binding the contract method 0xe29d93be.
//
// Solidity: function calDrinkNum(string drinkna) constant returns(uint256)
func (_ERC20drink *ERC20drinkCallerSession) CalDrinkNum(drinkna string) (*big.Int, error) {
	return _ERC20drink.Contract.CalDrinkNum(&_ERC20drink.CallOpts, drinkna)
}

// CalDrinkPrice is a free data retrieval call binding the contract method 0xee5db636.
//
// Solidity: function calDrinkPrice(string drinkna1) constant returns(uint256)
func (_ERC20drink *ERC20drinkCaller) CalDrinkPrice(opts *bind.CallOpts, drinkna1 string) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20drink.contract.Call(opts, out, "calDrinkPrice", drinkna1)
	return *ret0, err
}

// CalDrinkPrice is a free data retrieval call binding the contract method 0xee5db636.
//
// Solidity: function calDrinkPrice(string drinkna1) constant returns(uint256)
func (_ERC20drink *ERC20drinkSession) CalDrinkPrice(drinkna1 string) (*big.Int, error) {
	return _ERC20drink.Contract.CalDrinkPrice(&_ERC20drink.CallOpts, drinkna1)
}

// CalDrinkPrice is a free data retrieval call binding the contract method 0xee5db636.
//
// Solidity: function calDrinkPrice(string drinkna1) constant returns(uint256)
func (_ERC20drink *ERC20drinkCallerSession) CalDrinkPrice(drinkna1 string) (*big.Int, error) {
	return _ERC20drink.Contract.CalDrinkPrice(&_ERC20drink.CallOpts, drinkna1)
}

// CalEgg is a free data retrieval call binding the contract method 0x37cc77f6.
//
// Solidity: function calEgg(address add) constant returns(uint256)
func (_ERC20drink *ERC20drinkCaller) CalEgg(opts *bind.CallOpts, add common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20drink.contract.Call(opts, out, "calEgg", add)
	return *ret0, err
}

// CalEgg is a free data retrieval call binding the contract method 0x37cc77f6.
//
// Solidity: function calEgg(address add) constant returns(uint256)
func (_ERC20drink *ERC20drinkSession) CalEgg(add common.Address) (*big.Int, error) {
	return _ERC20drink.Contract.CalEgg(&_ERC20drink.CallOpts, add)
}

// CalEgg is a free data retrieval call binding the contract method 0x37cc77f6.
//
// Solidity: function calEgg(address add) constant returns(uint256)
func (_ERC20drink *ERC20drinkCallerSession) CalEgg(add common.Address) (*big.Int, error) {
	return _ERC20drink.Contract.CalEgg(&_ERC20drink.CallOpts, add)
}

// CalPoint is a free data retrieval call binding the contract method 0xcc2a1eb2.
//
// Solidity: function calPoint(address add) constant returns(uint256)
func (_ERC20drink *ERC20drinkCaller) CalPoint(opts *bind.CallOpts, add common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20drink.contract.Call(opts, out, "calPoint", add)
	return *ret0, err
}

// CalPoint is a free data retrieval call binding the contract method 0xcc2a1eb2.
//
// Solidity: function calPoint(address add) constant returns(uint256)
func (_ERC20drink *ERC20drinkSession) CalPoint(add common.Address) (*big.Int, error) {
	return _ERC20drink.Contract.CalPoint(&_ERC20drink.CallOpts, add)
}

// CalPoint is a free data retrieval call binding the contract method 0xcc2a1eb2.
//
// Solidity: function calPoint(address add) constant returns(uint256)
func (_ERC20drink *ERC20drinkCallerSession) CalPoint(add common.Address) (*big.Int, error) {
	return _ERC20drink.Contract.CalPoint(&_ERC20drink.CallOpts, add)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_ERC20drink *ERC20drinkCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _ERC20drink.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_ERC20drink *ERC20drinkSession) Decimals() (uint8, error) {
	return _ERC20drink.Contract.Decimals(&_ERC20drink.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_ERC20drink *ERC20drinkCallerSession) Decimals() (uint8, error) {
	return _ERC20drink.Contract.Decimals(&_ERC20drink.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_ERC20drink *ERC20drinkCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ERC20drink.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_ERC20drink *ERC20drinkSession) Name() (string, error) {
	return _ERC20drink.Contract.Name(&_ERC20drink.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_ERC20drink *ERC20drinkCallerSession) Name() (string, error) {
	return _ERC20drink.Contract.Name(&_ERC20drink.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_ERC20drink *ERC20drinkCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ERC20drink.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_ERC20drink *ERC20drinkSession) Symbol() (string, error) {
	return _ERC20drink.Contract.Symbol(&_ERC20drink.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_ERC20drink *ERC20drinkCallerSession) Symbol() (string, error) {
	return _ERC20drink.Contract.Symbol(&_ERC20drink.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20drink *ERC20drinkCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20drink.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20drink *ERC20drinkSession) TotalSupply() (*big.Int, error) {
	return _ERC20drink.Contract.TotalSupply(&_ERC20drink.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20drink *ERC20drinkCallerSession) TotalSupply() (*big.Int, error) {
	return _ERC20drink.Contract.TotalSupply(&_ERC20drink.CallOpts)
}

// PointSwap is a paid mutator transaction binding the contract method 0xc1fbffad.
//
// Solidity: function PointSwap(uint256 eggNum) returns(bool)
func (_ERC20drink *ERC20drinkTransactor) PointSwap(opts *bind.TransactOpts, eggNum *big.Int) (bool, *types.Transaction, *types.Receipt, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	transaction, receipt, err := _ERC20drink.contract.TransactWithResult(opts, out, "PointSwap", eggNum)
	return *ret0, transaction, receipt, err
}

func (_ERC20drink *ERC20drinkTransactor) AsyncPointSwap(handler func(*types.Receipt, error), opts *bind.TransactOpts, eggNum *big.Int) (*types.Transaction, error) {
	return _ERC20drink.contract.AsyncTransact(opts, handler, "PointSwap", eggNum)
}

// PointSwap is a paid mutator transaction binding the contract method 0xc1fbffad.
//
// Solidity: function PointSwap(uint256 eggNum) returns(bool)
func (_ERC20drink *ERC20drinkSession) PointSwap(eggNum *big.Int) (bool, *types.Transaction, *types.Receipt, error) {
	return _ERC20drink.Contract.PointSwap(&_ERC20drink.TransactOpts, eggNum)
}

func (_ERC20drink *ERC20drinkSession) AsyncPointSwap(handler func(*types.Receipt, error), eggNum *big.Int) (*types.Transaction, error) {
	return _ERC20drink.Contract.AsyncPointSwap(handler, &_ERC20drink.TransactOpts, eggNum)
}

// PointSwap is a paid mutator transaction binding the contract method 0xc1fbffad.
//
// Solidity: function PointSwap(uint256 eggNum) returns(bool)
func (_ERC20drink *ERC20drinkTransactorSession) PointSwap(eggNum *big.Int) (bool, *types.Transaction, *types.Receipt, error) {
	return _ERC20drink.Contract.PointSwap(&_ERC20drink.TransactOpts, eggNum)
}

func (_ERC20drink *ERC20drinkTransactorSession) AsyncPointSwap(handler func(*types.Receipt, error), eggNum *big.Int) (*types.Transaction, error) {
	return _ERC20drink.Contract.AsyncPointSwap(handler, &_ERC20drink.TransactOpts, eggNum)
}

// BuyDrink is a paid mutator transaction binding the contract method 0x0aa39505.
//
// Solidity: function _buyDrink(string drinkname, uint256 num) returns(bool)
func (_ERC20drink *ERC20drinkTransactor) BuyDrink(opts *bind.TransactOpts, drinkname string, num *big.Int) (bool, *types.Transaction, *types.Receipt, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	transaction, receipt, err := _ERC20drink.contract.TransactWithResult(opts, out, "_buyDrink", drinkname, num)
	return *ret0, transaction, receipt, err
}

func (_ERC20drink *ERC20drinkTransactor) AsyncBuyDrink(handler func(*types.Receipt, error), opts *bind.TransactOpts, drinkname string, num *big.Int) (*types.Transaction, error) {
	return _ERC20drink.contract.AsyncTransact(opts, handler, "_buyDrink", drinkname, num)
}

// BuyDrink is a paid mutator transaction binding the contract method 0x0aa39505.
//
// Solidity: function _buyDrink(string drinkname, uint256 num) returns(bool)
func (_ERC20drink *ERC20drinkSession) BuyDrink(drinkname string, num *big.Int) (bool, *types.Transaction, *types.Receipt, error) {
	return _ERC20drink.Contract.BuyDrink(&_ERC20drink.TransactOpts, drinkname, num)
}

func (_ERC20drink *ERC20drinkSession) AsyncBuyDrink(handler func(*types.Receipt, error), drinkname string, num *big.Int) (*types.Transaction, error) {
	return _ERC20drink.Contract.AsyncBuyDrink(handler, &_ERC20drink.TransactOpts, drinkname, num)
}

// BuyDrink is a paid mutator transaction binding the contract method 0x0aa39505.
//
// Solidity: function _buyDrink(string drinkname, uint256 num) returns(bool)
func (_ERC20drink *ERC20drinkTransactorSession) BuyDrink(drinkname string, num *big.Int) (bool, *types.Transaction, *types.Receipt, error) {
	return _ERC20drink.Contract.BuyDrink(&_ERC20drink.TransactOpts, drinkname, num)
}

func (_ERC20drink *ERC20drinkTransactorSession) AsyncBuyDrink(handler func(*types.Receipt, error), drinkname string, num *big.Int) (*types.Transaction, error) {
	return _ERC20drink.Contract.AsyncBuyDrink(handler, &_ERC20drink.TransactOpts, drinkname, num)
}

// UpdateDrink is a paid mutator transaction binding the contract method 0xe095c25c.
//
// Solidity: function _updateDrink(string name, uint256 price, uint256 number) returns(bool)
func (_ERC20drink *ERC20drinkTransactor) UpdateDrink(opts *bind.TransactOpts, name string, price *big.Int, number *big.Int) (bool, *types.Transaction, *types.Receipt, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	transaction, receipt, err := _ERC20drink.contract.TransactWithResult(opts, out, "_updateDrink", name, price, number)
	return *ret0, transaction, receipt, err
}

func (_ERC20drink *ERC20drinkTransactor) AsyncUpdateDrink(handler func(*types.Receipt, error), opts *bind.TransactOpts, name string, price *big.Int, number *big.Int) (*types.Transaction, error) {
	return _ERC20drink.contract.AsyncTransact(opts, handler, "_updateDrink", name, price, number)
}

// UpdateDrink is a paid mutator transaction binding the contract method 0xe095c25c.
//
// Solidity: function _updateDrink(string name, uint256 price, uint256 number) returns(bool)
func (_ERC20drink *ERC20drinkSession) UpdateDrink(name string, price *big.Int, number *big.Int) (bool, *types.Transaction, *types.Receipt, error) {
	return _ERC20drink.Contract.UpdateDrink(&_ERC20drink.TransactOpts, name, price, number)
}

func (_ERC20drink *ERC20drinkSession) AsyncUpdateDrink(handler func(*types.Receipt, error), name string, price *big.Int, number *big.Int) (*types.Transaction, error) {
	return _ERC20drink.Contract.AsyncUpdateDrink(handler, &_ERC20drink.TransactOpts, name, price, number)
}

// UpdateDrink is a paid mutator transaction binding the contract method 0xe095c25c.
//
// Solidity: function _updateDrink(string name, uint256 price, uint256 number) returns(bool)
func (_ERC20drink *ERC20drinkTransactorSession) UpdateDrink(name string, price *big.Int, number *big.Int) (bool, *types.Transaction, *types.Receipt, error) {
	return _ERC20drink.Contract.UpdateDrink(&_ERC20drink.TransactOpts, name, price, number)
}

func (_ERC20drink *ERC20drinkTransactorSession) AsyncUpdateDrink(handler func(*types.Receipt, error), name string, price *big.Int, number *big.Int) (*types.Transaction, error) {
	return _ERC20drink.Contract.AsyncUpdateDrink(handler, &_ERC20drink.TransactOpts, name, price, number)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address delegate, uint256 numTokens) returns(bool)
func (_ERC20drink *ERC20drinkTransactor) Approve(opts *bind.TransactOpts, delegate common.Address, numTokens *big.Int) (bool, *types.Transaction, *types.Receipt, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	transaction, receipt, err := _ERC20drink.contract.TransactWithResult(opts, out, "approve", delegate, numTokens)
	return *ret0, transaction, receipt, err
}

func (_ERC20drink *ERC20drinkTransactor) AsyncApprove(handler func(*types.Receipt, error), opts *bind.TransactOpts, delegate common.Address, numTokens *big.Int) (*types.Transaction, error) {
	return _ERC20drink.contract.AsyncTransact(opts, handler, "approve", delegate, numTokens)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address delegate, uint256 numTokens) returns(bool)
func (_ERC20drink *ERC20drinkSession) Approve(delegate common.Address, numTokens *big.Int) (bool, *types.Transaction, *types.Receipt, error) {
	return _ERC20drink.Contract.Approve(&_ERC20drink.TransactOpts, delegate, numTokens)
}

func (_ERC20drink *ERC20drinkSession) AsyncApprove(handler func(*types.Receipt, error), delegate common.Address, numTokens *big.Int) (*types.Transaction, error) {
	return _ERC20drink.Contract.AsyncApprove(handler, &_ERC20drink.TransactOpts, delegate, numTokens)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address delegate, uint256 numTokens) returns(bool)
func (_ERC20drink *ERC20drinkTransactorSession) Approve(delegate common.Address, numTokens *big.Int) (bool, *types.Transaction, *types.Receipt, error) {
	return _ERC20drink.Contract.Approve(&_ERC20drink.TransactOpts, delegate, numTokens)
}

func (_ERC20drink *ERC20drinkTransactorSession) AsyncApprove(handler func(*types.Receipt, error), delegate common.Address, numTokens *big.Int) (*types.Transaction, error) {
	return _ERC20drink.Contract.AsyncApprove(handler, &_ERC20drink.TransactOpts, delegate, numTokens)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address receiver, uint256 numTokens) returns(bool)
func (_ERC20drink *ERC20drinkTransactor) Transfer(opts *bind.TransactOpts, receiver common.Address, numTokens *big.Int) (bool, *types.Transaction, *types.Receipt, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	transaction, receipt, err := _ERC20drink.contract.TransactWithResult(opts, out, "transfer", receiver, numTokens)
	return *ret0, transaction, receipt, err
}

func (_ERC20drink *ERC20drinkTransactor) AsyncTransfer(handler func(*types.Receipt, error), opts *bind.TransactOpts, receiver common.Address, numTokens *big.Int) (*types.Transaction, error) {
	return _ERC20drink.contract.AsyncTransact(opts, handler, "transfer", receiver, numTokens)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address receiver, uint256 numTokens) returns(bool)
func (_ERC20drink *ERC20drinkSession) Transfer(receiver common.Address, numTokens *big.Int) (bool, *types.Transaction, *types.Receipt, error) {
	return _ERC20drink.Contract.Transfer(&_ERC20drink.TransactOpts, receiver, numTokens)
}

func (_ERC20drink *ERC20drinkSession) AsyncTransfer(handler func(*types.Receipt, error), receiver common.Address, numTokens *big.Int) (*types.Transaction, error) {
	return _ERC20drink.Contract.AsyncTransfer(handler, &_ERC20drink.TransactOpts, receiver, numTokens)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address receiver, uint256 numTokens) returns(bool)
func (_ERC20drink *ERC20drinkTransactorSession) Transfer(receiver common.Address, numTokens *big.Int) (bool, *types.Transaction, *types.Receipt, error) {
	return _ERC20drink.Contract.Transfer(&_ERC20drink.TransactOpts, receiver, numTokens)
}

func (_ERC20drink *ERC20drinkTransactorSession) AsyncTransfer(handler func(*types.Receipt, error), receiver common.Address, numTokens *big.Int) (*types.Transaction, error) {
	return _ERC20drink.Contract.AsyncTransfer(handler, &_ERC20drink.TransactOpts, receiver, numTokens)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address owner, address buyer, uint256 numTokens) returns(bool)
func (_ERC20drink *ERC20drinkTransactor) TransferFrom(opts *bind.TransactOpts, owner common.Address, buyer common.Address, numTokens *big.Int) (bool, *types.Transaction, *types.Receipt, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	transaction, receipt, err := _ERC20drink.contract.TransactWithResult(opts, out, "transferFrom", owner, buyer, numTokens)
	return *ret0, transaction, receipt, err
}

func (_ERC20drink *ERC20drinkTransactor) AsyncTransferFrom(handler func(*types.Receipt, error), opts *bind.TransactOpts, owner common.Address, buyer common.Address, numTokens *big.Int) (*types.Transaction, error) {
	return _ERC20drink.contract.AsyncTransact(opts, handler, "transferFrom", owner, buyer, numTokens)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address owner, address buyer, uint256 numTokens) returns(bool)
func (_ERC20drink *ERC20drinkSession) TransferFrom(owner common.Address, buyer common.Address, numTokens *big.Int) (bool, *types.Transaction, *types.Receipt, error) {
	return _ERC20drink.Contract.TransferFrom(&_ERC20drink.TransactOpts, owner, buyer, numTokens)
}

func (_ERC20drink *ERC20drinkSession) AsyncTransferFrom(handler func(*types.Receipt, error), owner common.Address, buyer common.Address, numTokens *big.Int) (*types.Transaction, error) {
	return _ERC20drink.Contract.AsyncTransferFrom(handler, &_ERC20drink.TransactOpts, owner, buyer, numTokens)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address owner, address buyer, uint256 numTokens) returns(bool)
func (_ERC20drink *ERC20drinkTransactorSession) TransferFrom(owner common.Address, buyer common.Address, numTokens *big.Int) (bool, *types.Transaction, *types.Receipt, error) {
	return _ERC20drink.Contract.TransferFrom(&_ERC20drink.TransactOpts, owner, buyer, numTokens)
}

func (_ERC20drink *ERC20drinkTransactorSession) AsyncTransferFrom(handler func(*types.Receipt, error), owner common.Address, buyer common.Address, numTokens *big.Int) (*types.Transaction, error) {
	return _ERC20drink.Contract.AsyncTransferFrom(handler, &_ERC20drink.TransactOpts, owner, buyer, numTokens)
}

// ERC20drinkApproval represents a Approval event raised by the ERC20drink contract.
type ERC20drinkApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20drink *ERC20drinkFilterer) WatchApproval(fromBlock *uint64, handler func(int, []types.Log), owner common.Address, spender common.Address) (string, error) {
	return _ERC20drink.contract.WatchLogs(fromBlock, handler, "Approval", owner, spender)
}

func (_ERC20drink *ERC20drinkFilterer) WatchAllApproval(fromBlock *uint64, handler func(int, []types.Log)) (string, error) {
	return _ERC20drink.contract.WatchLogs(fromBlock, handler, "Approval")
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20drink *ERC20drinkFilterer) ParseApproval(log types.Log) (*ERC20drinkApproval, error) {
	event := new(ERC20drinkApproval)
	if err := _ERC20drink.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20drink *ERC20drinkSession) WatchApproval(fromBlock *uint64, handler func(int, []types.Log), owner common.Address, spender common.Address) (string, error) {
	return _ERC20drink.Contract.WatchApproval(fromBlock, handler, owner, spender)
}

func (_ERC20drink *ERC20drinkSession) WatchAllApproval(fromBlock *uint64, handler func(int, []types.Log)) (string, error) {
	return _ERC20drink.Contract.WatchAllApproval(fromBlock, handler)
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20drink *ERC20drinkSession) ParseApproval(log types.Log) (*ERC20drinkApproval, error) {
	return _ERC20drink.Contract.ParseApproval(log)
}

// ERC20drinkTransfer represents a Transfer event raised by the ERC20drink contract.
type ERC20drinkTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20drink *ERC20drinkFilterer) WatchTransfer(fromBlock *uint64, handler func(int, []types.Log), from common.Address, to common.Address) (string, error) {
	return _ERC20drink.contract.WatchLogs(fromBlock, handler, "Transfer", from, to)
}

func (_ERC20drink *ERC20drinkFilterer) WatchAllTransfer(fromBlock *uint64, handler func(int, []types.Log)) (string, error) {
	return _ERC20drink.contract.WatchLogs(fromBlock, handler, "Transfer")
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20drink *ERC20drinkFilterer) ParseTransfer(log types.Log) (*ERC20drinkTransfer, error) {
	event := new(ERC20drinkTransfer)
	if err := _ERC20drink.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20drink *ERC20drinkSession) WatchTransfer(fromBlock *uint64, handler func(int, []types.Log), from common.Address, to common.Address) (string, error) {
	return _ERC20drink.Contract.WatchTransfer(fromBlock, handler, from, to)
}

func (_ERC20drink *ERC20drinkSession) WatchAllTransfer(fromBlock *uint64, handler func(int, []types.Log)) (string, error) {
	return _ERC20drink.Contract.WatchAllTransfer(fromBlock, handler)
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20drink *ERC20drinkSession) ParseTransfer(log types.Log) (*ERC20drinkTransfer, error) {
	return _ERC20drink.Contract.ParseTransfer(log)
}

// ERC20drinkDrinkBuy represents a DrinkBuy event raised by the ERC20drink contract.
type ERC20drinkDrinkBuy struct {
	Buyer     common.Address
	Drinkname string
	Num       *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// WatchDrinkBuy is a free log subscription operation binding the contract event 0x3d793e2ba72604006e0c2d8656f1af9b2acece635ad25c0c9787877239f33f31.
//
// Solidity: event _drinkBuy(address buyer, string drinkname, uint256 num)
func (_ERC20drink *ERC20drinkFilterer) WatchDrinkBuy(fromBlock *uint64, handler func(int, []types.Log)) (string, error) {
	return _ERC20drink.contract.WatchLogs(fromBlock, handler, "_drinkBuy")
}

func (_ERC20drink *ERC20drinkFilterer) WatchAllDrinkBuy(fromBlock *uint64, handler func(int, []types.Log)) (string, error) {
	return _ERC20drink.contract.WatchLogs(fromBlock, handler, "_drinkBuy")
}

// ParseDrinkBuy is a log parse operation binding the contract event 0x3d793e2ba72604006e0c2d8656f1af9b2acece635ad25c0c9787877239f33f31.
//
// Solidity: event _drinkBuy(address buyer, string drinkname, uint256 num)
func (_ERC20drink *ERC20drinkFilterer) ParseDrinkBuy(log types.Log) (*ERC20drinkDrinkBuy, error) {
	event := new(ERC20drinkDrinkBuy)
	if err := _ERC20drink.contract.UnpackLog(event, "_drinkBuy", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WatchDrinkBuy is a free log subscription operation binding the contract event 0x3d793e2ba72604006e0c2d8656f1af9b2acece635ad25c0c9787877239f33f31.
//
// Solidity: event _drinkBuy(address buyer, string drinkname, uint256 num)
func (_ERC20drink *ERC20drinkSession) WatchDrinkBuy(fromBlock *uint64, handler func(int, []types.Log)) (string, error) {
	return _ERC20drink.Contract.WatchDrinkBuy(fromBlock, handler)
}

func (_ERC20drink *ERC20drinkSession) WatchAllDrinkBuy(fromBlock *uint64, handler func(int, []types.Log)) (string, error) {
	return _ERC20drink.Contract.WatchAllDrinkBuy(fromBlock, handler)
}

// ParseDrinkBuy is a log parse operation binding the contract event 0x3d793e2ba72604006e0c2d8656f1af9b2acece635ad25c0c9787877239f33f31.
//
// Solidity: event _drinkBuy(address buyer, string drinkname, uint256 num)
func (_ERC20drink *ERC20drinkSession) ParseDrinkBuy(log types.Log) (*ERC20drinkDrinkBuy, error) {
	return _ERC20drink.Contract.ParseDrinkBuy(log)
}
