package main

import "github.com/IshiniKiridena/SmartC/call"

func main() {
	abi := `[{"inputs":[],"name":"addDetails","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"getFormulaDetails","outputs":[{"internalType":"string[]","name":"","type":"string[]"}],"stateMutability":"view","type":"function"}]`
	bin := "0x60806040526040518060a001604052806040518060400160405280601c81526020017f363333363761663736356434346137623134326637445465737430320000000081525081526020016040518060400160405280601081526020017f436172626f6e20466f6f747072696e7400000000000000000000000000000000815250815260200160405180606001604052806024815260200162001f7e6024913981526020016005815260200160405180606001604052806038815260200162001f466038913981525060036000820151816000019081620000e19190620003c4565b506020820151816001019081620000f99190620003c4565b506040820151816002019081620001119190620003c4565b50606082015181600301556080820151816004019081620001339190620003c4565b5050503480156200014357600080fd5b50620004ab565b600081519050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680620001cc57607f821691505b602082108103620001e257620001e162000184565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b6000600883026200024c7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826200020d565b6200025886836200020d565b95508019841693508086168417925050509392505050565b6000819050919050565b6000819050919050565b6000620002a56200029f620002998462000270565b6200027a565b62000270565b9050919050565b6000819050919050565b620002c18362000284565b620002d9620002d082620002ac565b8484546200021a565b825550505050565b600090565b620002f0620002e1565b620002fd818484620002b6565b505050565b5b81811015620003255762000319600082620002e6565b60018101905062000303565b5050565b601f82111562000374576200033e81620001e8565b6200034984620001fd565b8101602085101562000359578190505b620003716200036885620001fd565b83018262000302565b50505b505050565b600082821c905092915050565b6000620003996000198460080262000379565b1980831691505092915050565b6000620003b4838362000386565b9150826002028217905092915050565b620003cf826200014a565b67ffffffffffffffff811115620003eb57620003ea62000155565b5b620003f78254620001b3565b6200040482828562000329565b600060209050601f8311600181146200043c576000841562000427578287015190505b620004338582620003a6565b865550620004a3565b601f1984166200044c86620001e8565b60005b8281101562000476578489015182556001820191506020850194506020810190506200044f565b8683101562000496578489015162000492601f89168262000386565b8355505b6001600288020188555050505b505050505050565b611a8b80620004bb6000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80639d8f5c441461003b578063ab11bdb714610059575b600080fd5b610043610063565b6040516100509190611486565b60405180910390f35b61006161013c565b005b60606002805480602002602001604051908101604052809291908181526020016000905b828210156101335783829060005260206000200180546100a6906114d7565b80601f01602080910402602001604051908101604052809291908181526020018280546100d2906114d7565b801561011f5780601f106100f45761010080835404028352916020019161011f565b820191906000526020600020905b81548152906001019060200180831161010257829003601f168201915b505050505081526020019060010190610087565b50505050905090565b61021860486040518060600160405280602181526020016119e36021913973a59c88c8d9aaf6322a6959978258bf1c9c1859c560016040518060400160405280601881526020017f36333336376531643231386565363835633565316362306200000000000000008152506040518060400160405280602081526020017f35333666363936633230353037323635373036313732363137343639366636658152506040518060400160405280602081526020017f653132323331303934383163663733396431396536373335643032333635373781525061107a565b60028060018154018082558091505060019003906000526020600020016000604051806080016040528060528152602001611813605291399091909150908161026191906116ed565b506103fc6040518060400160405280602081526020017f65313232333130393438316366373339643139653637333564303233363537378152506040518060400160405280600581526020017f77617465720000000000000000000000000000000000000000000000000000008152506040518060400160405280601881526020017f36313337333238386439646236333633393036633735313200000000000000008152506040518060400160405280600381526020017f3130300000000000000000000000000000000000000000000000000000000000815250604051806060016040528060408152602001611982604091396040518060800160405280604281526020016118ee604291396040518060400160405280600181526020017f350000000000000000000000000000000000000000000000000000000000000081525060016040518060200160405280600081525060405180602001604052806000815250604051806020016040528060008152506040518060200160405280600081525060405180602001604052806000815250611195565b6104d860476040518060600160405280602181526020016118656021913973fd7a1e0456e2b1520dc3591995e420f31dff227160016040518060400160405280601881526020017f36333336376531643231386565363835633565316362306200000000000000008152506040518060400160405280602081526020017f35333666363936633230353037323635373036313732363137343639366636658152506040518060400160405280602081526020017f653132323331303934383163663733396431396536373335643032333635373781525061107a565b60028060018154018082558091505060019003906000526020600020016000604051806080016040528060528152602001611930605291399091909150908161052191906116ed565b5061077a6040518060400160405280602081526020017f65313232333130393438316366373339643139653637333564303233363537378152506040518060400160405280600581526020017f77617465720000000000000000000000000000000000000000000000000000008152506040518060400160405280601881526020017f36313337333238386439646236333633393036633735313200000000000000008152506040518060400160405280600381526020017f3130300000000000000000000000000000000000000000000000000000000000815250604051806060016040528060408152602001611982604091396040518060a0016040528060688152602001611886606891396040518060400160405280600181526020017f350000000000000000000000000000000000000000000000000000000000000081525060026040518060400160405280601881526020017f36323239396533346363363162326436343630353631343700000000000000008152506040518060400160405280600681526020017f74657374494400000000000000000000000000000000000000000000000000008152506040518060400160405280601481526020017f34363635373237343639366336393761363537320000000000000000000000008152506040518060400160405280601c81526020017f36363635373237343639366336393761363537323534373937303635000000008152506040518060400160405280601e81526020017f3436363537323734363936633639376136353732323035343739373036350000815250611195565b61085660466040518060600160405280602181526020016119c26021913973ee51e2f43880da936ea6615b20a06fd04e93d70560016040518060400160405280601881526020017f36333336376531643231386565363835633565316362306200000000000000008152506040518060400160405280602081526020017f35333666363936633230353037323635373036313732363137343639366636658152506040518060400160405280602081526020017f653132323331303934383163663733396431396536373335643032333635373781525061107a565b60028060018154018082558091505060019003906000526020600020016000604051806080016040528060528152602001611a04605291399091909150908161089f91906116ed565b50610af86040518060400160405280602081526020017f65313232333130393438316366373339643139653637333564303233363537378152506040518060400160405280600581526020017f77617465720000000000000000000000000000000000000000000000000000008152506040518060400160405280601881526020017f36313337333238386439646236333633393036633735313200000000000000008152506040518060400160405280600381526020017f3130300000000000000000000000000000000000000000000000000000000000815250604051806060016040528060408152602001611982604091396040518060a0016040528060688152602001611886606891396040518060400160405280600181526020017f350000000000000000000000000000000000000000000000000000000000000081525060026040518060400160405280601881526020017f36323239396533346363363162326436343630353631343700000000000000008152506040518060400160405280600681526020017f74657374494400000000000000000000000000000000000000000000000000008152506040518060400160405280601481526020017f34363635373237343639366336393761363537320000000000000000000000008152506040518060400160405280601c81526020017f36363635373237343639366336393761363537323534373937303635000000008152506040518060400160405280601e81526020017f3436363537323734363936633639376136353732323035343739373036350000815250611195565b610bd460486040518060600160405280602181526020016119e36021913973a59c88c8d9aaf6322a6959978258bf1c9c1859c560016040518060400160405280601881526020017f36333336376531643231386565363835633565316362306200000000000000008152506040518060400160405280602081526020017f35333666363936633230353037323635373036313732363137343639366636658152506040518060400160405280602081526020017f653132323331303934383163663733396431396536373335643032333635373781525061107a565b600280600181540180825580915050600190039060005260206000200160006040518060800160405280605281526020016118136052913990919091509081610c1d91906116ed565b50610db86040518060400160405280602081526020017f65313232333130393438316366373339643139653637333564303233363537378152506040518060400160405280600581526020017f77617465720000000000000000000000000000000000000000000000000000008152506040518060400160405280601881526020017f36313337333238386439646236333633393036633735313200000000000000008152506040518060400160405280600381526020017f3130300000000000000000000000000000000000000000000000000000000000815250604051806060016040528060408152602001611982604091396040518060800160405280604281526020016118ee604291396040518060400160405280600181526020017f350000000000000000000000000000000000000000000000000000000000000081525060016040518060200160405280600081525060405180602001604052806000815250604051806020016040528060008152506040518060200160405280600081525060405180602001604052806000815250611195565b610e9460486040518060600160405280602181526020016119e36021913973a59c88c8d9aaf6322a6959978258bf1c9c1859c560016040518060400160405280601881526020017f36333336376531643231386565363835633565316362306200000000000000008152506040518060400160405280602081526020017f35333666363936633230353037323635373036313732363137343639366636658152506040518060400160405280602081526020017f653132323331303934383163663733396431396536373335643032333635373781525061107a565b600280600181540180825580915050600190039060005260206000200160006040518060800160405280605281526020016118136052913990919091509081610edd91906116ed565b506110786040518060400160405280602081526020017f65313232333130393438316366373339643139653637333564303233363537378152506040518060400160405280600581526020017f77617465720000000000000000000000000000000000000000000000000000008152506040518060400160405280601881526020017f36313337333238386439646236333633393036633735313200000000000000008152506040518060400160405280600381526020017f3130300000000000000000000000000000000000000000000000000000000000815250604051806060016040528060408152602001611982604091396040518060800160405280604281526020016118ee604291396040518060400160405280600181526020017f350000000000000000000000000000000000000000000000000000000000000081525060016040518060200160405280600081525060405180602001604052806000815250604051806020016040528060008152506040518060200160405280600081525060405180602001604052806000815250611195565b565b6040518060e001604052808881526020018781526020018673ffffffffffffffffffffffffffffffffffffffff16815260200185815260200184815260200183815260200182815250600160008981526020019081526020016000206000820151816000015560208201518160010190816110f591906116ed565b5060408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160030155608082015181600401908161115c91906116ed565b5060a082015181600501908161117291906116ed565b5060c082015181600601908161118891906116ed565b5090505050505050505050565b604051806101a001604052808e81526020018d81526020018c81526020018b81526020018a81526020018981526020018881526020018781526020018681526020018581526020018481526020018381526020018281525060008e6040516111fd91906117fb565b9081526020016040518091039020600082015181600001908161122091906116ed565b50602082015181600101908161123691906116ed565b50604082015181600201908161124c91906116ed565b50606082015181600301908161126291906116ed565b50608082015181600401908161127891906116ed565b5060a082015181600501908161128e91906116ed565b5060c08201518160060190816112a491906116ed565b5060e082015181600701556101008201518160080190816112c591906116ed565b506101208201518160090190816112dc91906116ed565b5061014082015181600a0190816112f391906116ed565b5061016082015181600b01908161130a91906116ed565b5061018082015181600c01908161132191906116ed565b5090505050505050505050505050505050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b600081519050919050565b600082825260208201905092915050565b60005b8381101561139a57808201518184015260208101905061137f565b60008484015250505050565b6000601f19601f8301169050919050565b60006113c282611360565b6113cc818561136b565b93506113dc81856020860161137c565b6113e5816113a6565b840191505092915050565b60006113fc83836113b7565b905092915050565b6000602082019050919050565b600061141c82611334565b611426818561133f565b93508360208202850161143885611350565b8060005b85811015611474578484038952815161145585826113f0565b945061146083611404565b925060208a0199505060018101905061143c565b50829750879550505050505092915050565b600060208201905081810360008301526114a08184611411565b905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806114ef57607f821691505b602082108103611502576115016114a8565b5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b6000600883026115997fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8261155c565b6115a3868361155c565b95508019841693508086168417925050509392505050565b6000819050919050565b6000819050919050565b60006115ea6115e56115e0846115bb565b6115c5565b6115bb565b9050919050565b6000819050919050565b611604836115cf565b611618611610826115f1565b848454611569565b825550505050565b600090565b61162d611620565b6116388184846115fb565b505050565b5b8181101561165c57611651600082611625565b60018101905061163e565b5050565b601f8211156116a15761167281611537565b61167b8461154c565b8101602085101561168a578190505b61169e6116968561154c565b83018261163d565b50505b505050565b600082821c905092915050565b60006116c4600019846008026116a6565b1980831691505092915050565b60006116dd83836116b3565b9150826002028217905092915050565b6116f682611360565b67ffffffffffffffff81111561170f5761170e611508565b5b61171982546114d7565b611724828285611660565b600060209050601f8311600181146117575760008415611745578287015190505b61174f85826116d1565b8655506117b7565b601f19841661176586611537565b60005b8281101561178d57848901518255600182019150602085019450602081019050611768565b868310156117aa57848901516117a6601f8916826116b3565b8355505b6001600288020188555050505b505050505050565b600081905092915050565b60006117d582611360565b6117df81856117bf565b93506117ef81856020860161137c565b80840191505092915050565b600061180782846117ca565b91508190509291505056fe7b22666f726d756c614964223a2037322c2022636f6e747261637441646472657373223a2022307841353943383863386439414166363332326136393539393738323538426631633943313835396335227d36333161306234616439323431613933373466396631303844657654657374303265306236623465306237396365306237383465306237396365306236626265306237383065306236626265306237386165306236396365306236626132666530616538396530616562306530616561656530616638646530616562356530616539356530616638386530623662346530623738616530623662626530623662386530623738666530623661626530623662613266353336663639366335313735363136653734363937397b22666f726d756c614964223a2037312c2022636f6e747261637441646472657373223a2022307846443741316530343536653262313532306463333539313939356534323066333164466632323731227d353336663639366332303530373236353730363137323631373436393666366532306535396339666535613338636533383161656536626139366535383239393633316130623461643932343161393337346639663130384465765465737430313633316130623461643932343161393337346639663130384465765465737430337b22666f726d756c614964223a2037302c2022636f6e747261637441646472657373223a2022307845653531453266343338383064613933366541363631356232306130366644303465393344373035227da26469706673582212205a59934947c86756807455d8613a54fabcecfb505d4f46b7dd725792eafd3ade64736f6c6343000811003347424f4b354c4134464142564f4b3736584f5a4759485844463558594d533546514a37585a4e503654544147364e5953373636544f37454637376365376162302d633737652d313165632d623666662d343131616434323133396364"
	contractAddress := "0x29057e2F029524A618a38954c7FB63BaB26aAF31"
	call.CreateInstance(abi, bin, contractAddress)
}
