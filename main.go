package main

import call "github.com/IshiniKiridena/SmartC/newCall1"

func main() {
	abi := `[{"inputs":[],"name":"getCOMPOST_QUANTITY","outputs":[{"components":[{"internalType":"string","name":"valueID","type":"string"},{"internalType":"string","name":"valueName","type":"string"},{"internalType":"string","name":"formulaID","type":"string"},{"internalType":"string","name":"workflowID","type":"string"},{"internalType":"string","name":"stageID","type":"string"},{"internalType":"string","name":"stageName","type":"string"},{"internalType":"string","name":"keyName","type":"string"},{"internalType":"string","name":"tdpType","type":"string"},{"internalType":"int256","name":"bindingType","type":"int256"},{"internalType":"string","name":"artifactID","type":"string"},{"internalType":"string","name":"primaryKeyRowID","type":"string"},{"internalType":"string","name":"artifactTemplateName","type":"string"},{"internalType":"string","name":"fieldKey","type":"string"},{"internalType":"string","name":"fieldName","type":"string"}],"internalType":"struct Metric2.Value","name":"","type":"tuple"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"getFERTILIZER_EMISSION","outputs":[{"components":[{"internalType":"string","name":"valueID","type":"string"},{"internalType":"string","name":"valueName","type":"string"},{"internalType":"string","name":"formulaID","type":"string"},{"internalType":"string","name":"workflowID","type":"string"},{"internalType":"string","name":"stageID","type":"string"},{"internalType":"string","name":"stageName","type":"string"},{"internalType":"string","name":"keyName","type":"string"},{"internalType":"string","name":"tdpType","type":"string"},{"internalType":"int256","name":"bindingType","type":"int256"},{"internalType":"string","name":"artifactID","type":"string"},{"internalType":"string","name":"primaryKeyRowID","type":"string"},{"internalType":"string","name":"artifactTemplateName","type":"string"},{"internalType":"string","name":"fieldKey","type":"string"},{"internalType":"string","name":"fieldName","type":"string"}],"internalType":"struct Metric2.Value","name":"","type":"tuple"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"getFertilizer_Emission","outputs":[{"internalType":"string","name":"","type":"string"},{"internalType":"address","name":"","type":"address"},{"internalType":"uint256","name":"","type":"uint256"},{"internalType":"string","name":"","type":"string"},{"internalType":"string","name":"","type":"string"},{"internalType":"string","name":"","type":"string"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"getWATER","outputs":[{"components":[{"internalType":"string","name":"valueID","type":"string"},{"internalType":"string","name":"valueName","type":"string"},{"internalType":"string","name":"formulaID","type":"string"},{"internalType":"string","name":"workflowID","type":"string"},{"internalType":"string","name":"stageID","type":"string"},{"internalType":"string","name":"stageName","type":"string"},{"internalType":"string","name":"keyName","type":"string"},{"internalType":"string","name":"tdpType","type":"string"},{"internalType":"int256","name":"bindingType","type":"int256"},{"internalType":"string","name":"artifactID","type":"string"},{"internalType":"string","name":"primaryKeyRowID","type":"string"},{"internalType":"string","name":"artifactTemplateName","type":"string"},{"internalType":"string","name":"fieldKey","type":"string"},{"internalType":"string","name":"fieldName","type":"string"}],"internalType":"struct Metric2.Value","name":"","type":"tuple"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"getWater_Emission","outputs":[{"internalType":"string","name":"","type":"string"},{"internalType":"address","name":"","type":"address"},{"internalType":"uint256","name":"","type":"uint256"},{"internalType":"string","name":"","type":"string"},{"internalType":"string","name":"","type":"string"},{"internalType":"string","name":"","type":"string"}],"stateMutability":"view","type":"function"}]`
	bin := "60806040526040518060a001604052806040518060400160405280601e81526020017f363333363761663736356434346137623134326637446576546573743031000081525081526020016040518060400160405280601081526020017f436172626f6e20466f6f747072696e7400000000000000000000000000000000815250815260200160405180606001604052806024815260200162003657602491398152602001600281526020016040518060600160405280603881526020016200344f60389139815250600080820151816000019081620000e0919062001221565b506020820151816001019081620000f8919062001221565b50604082015181600201908162000110919062001221565b5060608201518160030155608082015181600401908162000132919062001221565b5050506040518060c001604052806040518060400160405280601b81526020017f36333161306234616439323431613933373466436f6e66696730310000000000815250815260200173b861ab9ca8cabe266078e9f25ba5ab99b3d2f67a73ffffffffffffffffffffffffffffffffffffffff168152602001600281526020016040518060400160405280601881526020017f363333363765316432313865653638356335653161303031000000000000000081525081526020016040518060800160405280604a815260200162003789604a913981526020016040518060800160405280604281526020016200367b604291398152506005600082015181600001908162000242919062001221565b5060208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550604082015181600201556060820151816003019081620002ab919062001221565b506080820151816004019081620002c3919062001221565b5060a0820151816005019081620002db919062001221565b5050506040518060c001604052806040518060400160405280601881526020017f3633363062633163326532653037613730346261626661360000000000000000815250815260200173aa1adc9cf23f234a3927df60bd020a328aeb903173ffffffffffffffffffffffffffffffffffffffff168152602001600181526020016040518060400160405280601881526020017f36333336376531643231386565363835633565316130303200000000000000008152508152602001604051806060016040528060408152602001620035856040913981526020016040518060400160405280602081526020017f6531323233313039343831636637333964313965363733356430323336353738815250815250600b600082015181600001908162000407919062001221565b5060208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160020155606082015181600301908162000470919062001221565b50608082015181600401908162000488919062001221565b5060a0820151816005019081620004a0919062001221565b505050604051806101c001604052806040518060400160405280602081526020017f653132323331303934383163663733396431396536373335643032333635373781525081526020016040518060400160405280601381526020017f46455254494c495a45525f454d495353494f4e0000000000000000000000000081525081526020016040518060400160405280601b81526020017f36333161306234616439323431613933373466436f6e6669673031000000000081525081526020016040518060400160405280601881526020017f363133373332383864396462363336333930366337353132000000000000000081525081526020016040518060400160405280600381526020017f313030000000000000000000000000000000000000000000000000000000000081525081526020016040518060c0016040528060948152602001620036f5609491398152602001604051806060016040528060388152602001620036bd6038913981526020016040518060400160405280600181526020017f35000000000000000000000000000000000000000000000000000000000000008152508152602001600281526020016040518060400160405280601881526020017f363232393965333463633631623264363436303536313437000000000000000081525081526020016040518060400160405280600681526020017f746573744944000000000000000000000000000000000000000000000000000081525081526020016040518060400160405280601481526020017f343636353732373436393663363937613635373200000000000000000000000081525081526020016040518060400160405280601c81526020017f363536643639373337333639366636653436363136333734366637320000000081525081526020016040518060400160405280601e81526020017f34363635373237343639366336393761363537323230353437393730363500008152508152506011600082015181600001908162000794919062001221565b506020820151816001019081620007ac919062001221565b506040820151816002019081620007c4919062001221565b506060820151816003019081620007dc919062001221565b506080820151816004019081620007f4919062001221565b5060a08201518160050190816200080c919062001221565b5060c082015181600601908162000824919062001221565b5060e08201518160070190816200083c919062001221565b50610100820151816008015561012082015181600901908162000860919062001221565b5061014082015181600a01908162000879919062001221565b5061016082015181600b01908162000892919062001221565b5061018082015181600c019081620008ab919062001221565b506101a082015181600d019081620008c4919062001221565b505050604051806101c001604052806040518060400160405280602081526020017f653132323331303934383163663733396431396536373335643032333635373881525081526020016040518060400160405280601381526020017f46455254494c495a45525f454d495353494f4e0000000000000000000000000081525081526020016040518060400160405280601b81526020017f36333161306234616439323431613933373466436f6e6669673031000000000081525081526020016040518060400160405280601881526020017f363133373332383864396462363336333930366337353132000000000000000081525081526020016040518060400160405280600381526020017f313030000000000000000000000000000000000000000000000000000000000081525081526020016040518060c0016040528060948152602001620036f56094913981526020016040518060a0016040528060808152602001620035056080913981526020016040518060400160405280600181526020017f350000000000000000000000000000000000000000000000000000000000000081525081526020016001815260200160405180602001604052806000815250815260200160405180602001604052806000815250815260200160405180602001604052806000815250815260200160405180602001604052806000815250815260200160405180602001604052806000815250815250601f600082015181600001908162000afa919062001221565b50602082015181600101908162000b12919062001221565b50604082015181600201908162000b2a919062001221565b50606082015181600301908162000b42919062001221565b50608082015181600401908162000b5a919062001221565b5060a082015181600501908162000b72919062001221565b5060c082015181600601908162000b8a919062001221565b5060e082015181600701908162000ba2919062001221565b50610100820151816008015561012082015181600901908162000bc6919062001221565b5061014082015181600a01908162000bdf919062001221565b5061016082015181600b01908162000bf8919062001221565b5061018082015181600c01908162000c11919062001221565b506101a082015181600d01908162000c2a919062001221565b505050604051806101c001604052806040518060400160405280602081526020017f653132323331303934383163663733396431396536373335643032333635373981525081526020016040518060400160405280600581526020017f576174657200000000000000000000000000000000000000000000000000000081525081526020016040518060400160405280601881526020017f363336306263316332653265303761373034626162666136000000000000000081525081526020016040518060400160405280601881526020017f363133373332383864396462363336333930366337353132000000000000000081525081526020016040518060400160405280600381526020017f313030000000000000000000000000000000000000000000000000000000000081525081526020016040518060a00160405280607e815260200162003487607e913981526020016040518060c0016040528060928152602001620035c56092913981526020016040518060400160405280600181526020017f350000000000000000000000000000000000000000000000000000000000000081525081526020016001815260200160405180602001604052806000815250815260200160405180602001604052806000815250815260200160405180602001604052806000815250815260200160405180602001604052806000815250815260200160405180602001604052806000815250815250602d600082015181600001908162000e60919062001221565b50602082015181600101908162000e78919062001221565b50604082015181600201908162000e90919062001221565b50606082015181600301908162000ea8919062001221565b50608082015181600401908162000ec0919062001221565b5060a082015181600501908162000ed8919062001221565b5060c082015181600601908162000ef0919062001221565b5060e082015181600701908162000f08919062001221565b50610100820151816008015561012082015181600901908162000f2c919062001221565b5061014082015181600a01908162000f45919062001221565b5061016082015181600b01908162000f5e919062001221565b5061018082015181600c01908162000f77919062001221565b506101a082015181600d01908162000f90919062001221565b50505034801562000fa057600080fd5b5062001308565b600081519050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806200102957607f821691505b6020821081036200103f576200103e62000fe1565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302620010a97fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826200106a565b620010b586836200106a565b95508019841693508086168417925050509392505050565b6000819050919050565b6000819050919050565b600062001102620010fc620010f684620010cd565b620010d7565b620010cd565b9050919050565b6000819050919050565b6200111e83620010e1565b620011366200112d8262001109565b84845462001077565b825550505050565b600090565b6200114d6200113e565b6200115a81848462001113565b505050565b5b8181101562001182576200117660008262001143565b60018101905062001160565b5050565b601f821115620011d1576200119b8162001045565b620011a6846200105a565b81016020851015620011b6578190505b620011ce620011c5856200105a565b8301826200115f565b50505b505050565b600082821c905092915050565b6000620011f660001984600802620011d6565b1980831691505092915050565b6000620012118383620011e3565b9150826002028217905092915050565b6200122c8262000fa7565b67ffffffffffffffff81111562001248576200124762000fb2565b5b62001254825462001010565b6200126182828562001186565b600060209050601f83116001811462001299576000841562001284578287015190505b62001290858262001203565b86555062001300565b601f198416620012a98662001045565b60005b82811015620012d357848901518255600182019150602085019450602081019050620012ac565b86831015620012f35784890151620012ef601f891682620011e3565b8355505b6001600288020188555050505b505050505050565b61213780620013186000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c80632592cb651461005c578063747a4eec1461007f5780638ddfa8b91461009d578063ba1fc7bb146100bb578063f9e565cc146100de575b600080fd5b6100646100fc565b60405161007696959493929190611e1d565b60405180910390f35b61008761038a565b604051610094919061207f565b60405180910390f35b6100a5610b19565b6040516100b2919061207f565b60405180910390f35b6100c36112a8565b6040516100d596959493929190611e1d565b60405180910390f35b6100e6611535565b6040516100f3919061207f565b60405180910390f35b60606000806060806060600b600001600b60010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600b60020154600b600301600b600401600b600501858054610152906120d0565b80601f016020809104026020016040519081016040528092919081815260200182805461017e906120d0565b80156101cb5780601f106101a0576101008083540402835291602001916101cb565b820191906000526020600020905b8154815290600101906020018083116101ae57829003601f168201915b505050505095508280546101de906120d0565b80601f016020809104026020016040519081016040528092919081815260200182805461020a906120d0565b80156102575780601f1061022c57610100808354040283529160200191610257565b820191906000526020600020905b81548152906001019060200180831161023a57829003601f168201915b5050505050925081805461026a906120d0565b80601f0160208091040260200160405190810160405280929190818152602001828054610296906120d0565b80156102e35780601f106102b8576101008083540402835291602001916102e3565b820191906000526020600020905b8154815290600101906020018083116102c657829003601f168201915b505050505091508080546102f6906120d0565b80601f0160208091040260200160405190810160405280929190818152602001828054610322906120d0565b801561036f5780601f106103445761010080835404028352916020019161036f565b820191906000526020600020905b81548152906001019060200180831161035257829003601f168201915b50505050509050955095509550955095509550909192939495565b610392611cc4565b6011604051806101c00160405290816000820180546103b0906120d0565b80601f01602080910402602001604051908101604052809291908181526020018280546103dc906120d0565b80156104295780601f106103fe57610100808354040283529160200191610429565b820191906000526020600020905b81548152906001019060200180831161040c57829003601f168201915b50505050508152602001600182018054610442906120d0565b80601f016020809104026020016040519081016040528092919081815260200182805461046e906120d0565b80156104bb5780601f10610490576101008083540402835291602001916104bb565b820191906000526020600020905b81548152906001019060200180831161049e57829003601f168201915b505050505081526020016002820180546104d4906120d0565b80601f0160208091040260200160405190810160405280929190818152602001828054610500906120d0565b801561054d5780601f106105225761010080835404028352916020019161054d565b820191906000526020600020905b81548152906001019060200180831161053057829003601f168201915b50505050508152602001600382018054610566906120d0565b80601f0160208091040260200160405190810160405280929190818152602001828054610592906120d0565b80156105df5780601f106105b4576101008083540402835291602001916105df565b820191906000526020600020905b8154815290600101906020018083116105c257829003601f168201915b505050505081526020016004820180546105f8906120d0565b80601f0160208091040260200160405190810160405280929190818152602001828054610624906120d0565b80156106715780601f1061064657610100808354040283529160200191610671565b820191906000526020600020905b81548152906001019060200180831161065457829003601f168201915b5050505050815260200160058201805461068a906120d0565b80601f01602080910402602001604051908101604052809291908181526020018280546106b6906120d0565b80156107035780601f106106d857610100808354040283529160200191610703565b820191906000526020600020905b8154815290600101906020018083116106e657829003601f168201915b5050505050815260200160068201805461071c906120d0565b80601f0160208091040260200160405190810160405280929190818152602001828054610748906120d0565b80156107955780601f1061076a57610100808354040283529160200191610795565b820191906000526020600020905b81548152906001019060200180831161077857829003601f168201915b505050505081526020016007820180546107ae906120d0565b80601f01602080910402602001604051908101604052809291908181526020018280546107da906120d0565b80156108275780601f106107fc57610100808354040283529160200191610827565b820191906000526020600020905b81548152906001019060200180831161080a57829003601f168201915b505050505081526020016008820154815260200160098201805461084a906120d0565b80601f0160208091040260200160405190810160405280929190818152602001828054610876906120d0565b80156108c35780601f10610898576101008083540402835291602001916108c3565b820191906000526020600020905b8154815290600101906020018083116108a657829003601f168201915b50505050508152602001600a820180546108dc906120d0565b80601f0160208091040260200160405190810160405280929190818152602001828054610908906120d0565b80156109555780601f1061092a57610100808354040283529160200191610955565b820191906000526020600020905b81548152906001019060200180831161093857829003601f168201915b50505050508152602001600b8201805461096e906120d0565b80601f016020809104026020016040519081016040528092919081815260200182805461099a906120d0565b80156109e75780601f106109bc576101008083540402835291602001916109e7565b820191906000526020600020905b8154815290600101906020018083116109ca57829003601f168201915b50505050508152602001600c82018054610a00906120d0565b80601f0160208091040260200160405190810160405280929190818152602001828054610a2c906120d0565b8015610a795780601f10610a4e57610100808354040283529160200191610a79565b820191906000526020600020905b815481529060010190602001808311610a5c57829003601f168201915b50505050508152602001600d82018054610a92906120d0565b80601f0160208091040260200160405190810160405280929190818152602001828054610abe906120d0565b8015610b0b5780601f10610ae057610100808354040283529160200191610b0b565b820191906000526020600020905b815481529060010190602001808311610aee57829003601f168201915b505050505081525050905090565b610b21611cc4565b601f604051806101c0016040529081600082018054610b3f906120d0565b80601f0160208091040260200160405190810160405280929190818152602001828054610b6b906120d0565b8015610bb85780601f10610b8d57610100808354040283529160200191610bb8565b820191906000526020600020905b815481529060010190602001808311610b9b57829003601f168201915b50505050508152602001600182018054610bd1906120d0565b80601f0160208091040260200160405190810160405280929190818152602001828054610bfd906120d0565b8015610c4a5780601f10610c1f57610100808354040283529160200191610c4a565b820191906000526020600020905b815481529060010190602001808311610c2d57829003601f168201915b50505050508152602001600282018054610c63906120d0565b80601f0160208091040260200160405190810160405280929190818152602001828054610c8f906120d0565b8015610cdc5780601f10610cb157610100808354040283529160200191610cdc565b820191906000526020600020905b815481529060010190602001808311610cbf57829003601f168201915b50505050508152602001600382018054610cf5906120d0565b80601f0160208091040260200160405190810160405280929190818152602001828054610d21906120d0565b8015610d6e5780601f10610d4357610100808354040283529160200191610d6e565b820191906000526020600020905b815481529060010190602001808311610d5157829003601f168201915b50505050508152602001600482018054610d87906120d0565b80601f0160208091040260200160405190810160405280929190818152602001828054610db3906120d0565b8015610e005780601f10610dd557610100808354040283529160200191610e00565b820191906000526020600020905b815481529060010190602001808311610de357829003601f168201915b50505050508152602001600582018054610e19906120d0565b80601f0160208091040260200160405190810160405280929190818152602001828054610e45906120d0565b8015610e925780601f10610e6757610100808354040283529160200191610e92565b820191906000526020600020905b815481529060010190602001808311610e7557829003601f168201915b50505050508152602001600682018054610eab906120d0565b80601f0160208091040260200160405190810160405280929190818152602001828054610ed7906120d0565b8015610f245780601f10610ef957610100808354040283529160200191610f24565b820191906000526020600020905b815481529060010190602001808311610f0757829003601f168201915b50505050508152602001600782018054610f3d906120d0565b80601f0160208091040260200160405190810160405280929190818152602001828054610f69906120d0565b8015610fb65780601f10610f8b57610100808354040283529160200191610fb6565b820191906000526020600020905b815481529060010190602001808311610f9957829003601f168201915b5050505050815260200160088201548152602001600982018054610fd9906120d0565b80601f0160208091040260200160405190810160405280929190818152602001828054611005906120d0565b80156110525780601f1061102757610100808354040283529160200191611052565b820191906000526020600020905b81548152906001019060200180831161103557829003601f168201915b50505050508152602001600a8201805461106b906120d0565b80601f0160208091040260200160405190810160405280929190818152602001828054611097906120d0565b80156110e45780601f106110b9576101008083540402835291602001916110e4565b820191906000526020600020905b8154815290600101906020018083116110c757829003601f168201915b50505050508152602001600b820180546110fd906120d0565b80601f0160208091040260200160405190810160405280929190818152602001828054611129906120d0565b80156111765780601f1061114b57610100808354040283529160200191611176565b820191906000526020600020905b81548152906001019060200180831161115957829003601f168201915b50505050508152602001600c8201805461118f906120d0565b80601f01602080910402602001604051908101604052809291908181526020018280546111bb906120d0565b80156112085780601f106111dd57610100808354040283529160200191611208565b820191906000526020600020905b8154815290600101906020018083116111eb57829003601f168201915b50505050508152602001600d82018054611221906120d0565b80601f016020809104026020016040519081016040528092919081815260200182805461124d906120d0565b801561129a5780601f1061126f5761010080835404028352916020019161129a565b820191906000526020600020905b81548152906001019060200180831161127d57829003601f168201915b505050505081525050905090565b606060008060608060606005600001600560010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1660056002015460056003016005600401600580018580546112fd906120d0565b80601f0160208091040260200160405190810160405280929190818152602001828054611329906120d0565b80156113765780601f1061134b57610100808354040283529160200191611376565b820191906000526020600020905b81548152906001019060200180831161135957829003601f168201915b50505050509550828054611389906120d0565b80601f01602080910402602001604051908101604052809291908181526020018280546113b5906120d0565b80156114025780601f106113d757610100808354040283529160200191611402565b820191906000526020600020905b8154815290600101906020018083116113e557829003601f168201915b50505050509250818054611415906120d0565b80601f0160208091040260200160405190810160405280929190818152602001828054611441906120d0565b801561148e5780601f106114635761010080835404028352916020019161148e565b820191906000526020600020905b81548152906001019060200180831161147157829003601f168201915b505050505091508080546114a1906120d0565b80601f01602080910402602001604051908101604052809291908181526020018280546114cd906120d0565b801561151a5780601f106114ef5761010080835404028352916020019161151a565b820191906000526020600020905b8154815290600101906020018083116114fd57829003601f168201915b50505050509050955095509550955095509550909192939495565b61153d611cc4565b602d604051806101c001604052908160008201805461155b906120d0565b80601f0160208091040260200160405190810160405280929190818152602001828054611587906120d0565b80156115d45780601f106115a9576101008083540402835291602001916115d4565b820191906000526020600020905b8154815290600101906020018083116115b757829003601f168201915b505050505081526020016001820180546115ed906120d0565b80601f0160208091040260200160405190810160405280929190818152602001828054611619906120d0565b80156116665780601f1061163b57610100808354040283529160200191611666565b820191906000526020600020905b81548152906001019060200180831161164957829003601f168201915b5050505050815260200160028201805461167f906120d0565b80601f01602080910402602001604051908101604052809291908181526020018280546116ab906120d0565b80156116f85780601f106116cd576101008083540402835291602001916116f8565b820191906000526020600020905b8154815290600101906020018083116116db57829003601f168201915b50505050508152602001600382018054611711906120d0565b80601f016020809104026020016040519081016040528092919081815260200182805461173d906120d0565b801561178a5780601f1061175f5761010080835404028352916020019161178a565b820191906000526020600020905b81548152906001019060200180831161176d57829003601f168201915b505050505081526020016004820180546117a3906120d0565b80601f01602080910402602001604051908101604052809291908181526020018280546117cf906120d0565b801561181c5780601f106117f15761010080835404028352916020019161181c565b820191906000526020600020905b8154815290600101906020018083116117ff57829003601f168201915b50505050508152602001600582018054611835906120d0565b80601f0160208091040260200160405190810160405280929190818152602001828054611861906120d0565b80156118ae5780601f10611883576101008083540402835291602001916118ae565b820191906000526020600020905b81548152906001019060200180831161189157829003601f168201915b505050505081526020016006820180546118c7906120d0565b80601f01602080910402602001604051908101604052809291908181526020018280546118f3906120d0565b80156119405780601f1061191557610100808354040283529160200191611940565b820191906000526020600020905b81548152906001019060200180831161192357829003601f168201915b50505050508152602001600782018054611959906120d0565b80601f0160208091040260200160405190810160405280929190818152602001828054611985906120d0565b80156119d25780601f106119a7576101008083540402835291602001916119d2565b820191906000526020600020905b8154815290600101906020018083116119b557829003601f168201915b50505050508152602001600882015481526020016009820180546119f5906120d0565b80601f0160208091040260200160405190810160405280929190818152602001828054611a21906120d0565b8015611a6e5780601f10611a4357610100808354040283529160200191611a6e565b820191906000526020600020905b815481529060010190602001808311611a5157829003601f168201915b50505050508152602001600a82018054611a87906120d0565b80601f0160208091040260200160405190810160405280929190818152602001828054611ab3906120d0565b8015611b005780601f10611ad557610100808354040283529160200191611b00565b820191906000526020600020905b815481529060010190602001808311611ae357829003601f168201915b50505050508152602001600b82018054611b19906120d0565b80601f0160208091040260200160405190810160405280929190818152602001828054611b45906120d0565b8015611b925780601f10611b6757610100808354040283529160200191611b92565b820191906000526020600020905b815481529060010190602001808311611b7557829003601f168201915b50505050508152602001600c82018054611bab906120d0565b80601f0160208091040260200160405190810160405280929190818152602001828054611bd7906120d0565b8015611c245780601f10611bf957610100808354040283529160200191611c24565b820191906000526020600020905b815481529060010190602001808311611c0757829003601f168201915b50505050508152602001600d82018054611c3d906120d0565b80601f0160208091040260200160405190810160405280929190818152602001828054611c69906120d0565b8015611cb65780601f10611c8b57610100808354040283529160200191611cb6565b820191906000526020600020905b815481529060010190602001808311611c9957829003601f168201915b505050505081525050905090565b604051806101c0016040528060608152602001606081526020016060815260200160608152602001606081526020016060815260200160608152602001606081526020016000815260200160608152602001606081526020016060815260200160608152602001606081525090565b600081519050919050565b600082825260208201905092915050565b60005b83811015611d6d578082015181840152602081019050611d52565b60008484015250505050565b6000601f19601f8301169050919050565b6000611d9582611d33565b611d9f8185611d3e565b9350611daf818560208601611d4f565b611db881611d79565b840191505092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000611dee82611dc3565b9050919050565b611dfe81611de3565b82525050565b6000819050919050565b611e1781611e04565b82525050565b600060c0820190508181036000830152611e378189611d8a565b9050611e466020830188611df5565b611e536040830187611e0e565b8181036060830152611e658186611d8a565b90508181036080830152611e798185611d8a565b905081810360a0830152611e8d8184611d8a565b9050979650505050505050565b600082825260208201905092915050565b6000611eb682611d33565b611ec08185611e9a565b9350611ed0818560208601611d4f565b611ed981611d79565b840191505092915050565b6000819050919050565b611ef781611ee4565b82525050565b60006101c0830160008301518482036000860152611f1b8282611eab565b91505060208301518482036020860152611f358282611eab565b91505060408301518482036040860152611f4f8282611eab565b91505060608301518482036060860152611f698282611eab565b91505060808301518482036080860152611f838282611eab565b91505060a083015184820360a0860152611f9d8282611eab565b91505060c083015184820360c0860152611fb78282611eab565b91505060e083015184820360e0860152611fd18282611eab565b915050610100830151611fe8610100860182611eee565b506101208301518482036101208601526120028282611eab565b91505061014083015184820361014086015261201e8282611eab565b91505061016083015184820361016086015261203a8282611eab565b9150506101808301518482036101808601526120568282611eab565b9150506101a08301518482036101a08601526120728282611eab565b9150508091505092915050565b600060208201905081810360008301526120998184611efd565b905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806120e857607f821691505b6020821081036120fb576120fa6120a1565b5b5091905056fea26469706673582212202f2db040d6890877be278d4774a0066498a21953630c43e80a8268c0dd4be01764736f6c6343000811003347424f4b354c4134464142564f4b3736584f5a4759485844463558594d533546514a37585a4e503654544147364e5953373636544f374546353336353635363436393665363732303236323034373635366436393665363137343639366636653230373736313734363537323230373537333631363736353230653639326164653761386165653338336262653739396261653838616264653662306234653338316165653462646266653739346138653938373866653062366234653062373963653062373834653062373963653062366262653062366234653062373861653238303864653062366262653062366238653062373866653062366162653062366261326665306165383965306165623065306165616565306166386465306165383565306165623365306165623565306166383135333635363536343639366536373230323632303437363536643639366536313734363936663665323037373631373436353732323037353733363136373635653062366132653062366264653062366234653062366262653062373932653062366237653062373964653062366132653062366231653062366261326665306165613465306165613365306166386465306165613365306166383065306165623065306166386465306165613865306166383165306165393565306165623065306166386465306165623565306166383137376365376162302d633737652d313165632d623666662d34313161643432313339636465313232333130393438316366373339643139653637333564303233363537372c2065313232333130393438316366373339643139653637333564303233363537376530623662346530623739636530623738346530623739636530623662623266653061653839653061656230653061656165653061663864353336353635363436393665363732303236323034373635366436393665363137343639366636653230363636353732373436393663363937613635373232303735373336313637363532306537613861656533383162656533383138646533383362626537393962616538386162646538383261356536393639396533383161656534626462666533383138346536393662393533363536353634363936653637323032363230343736353664363936653631373436393666366532303636363537323734363936633639376136353732323037353733363136373635"
	contractAddress := "0x006F517BECB4C4200ED6742f97E4Be7c251d1C96"
	call.CreateInstance1(abi, bin, contractAddress)

	//deploy.DeployContract()
}
