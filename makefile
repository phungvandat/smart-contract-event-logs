gen:
	abigen --abi abigenBindings/abi/NFT.abi --pkg abinft --out contracts/abinft/NFT.go
	abigen --abi abigenBindings/abi/NFTMarket.abi --pkg abimarket --out contracts/abimarket/NFTMarket.go
	abigen --abi abigenBindings/abi/NFTMarket2.abi --pkg abimarket2 --out contracts/abimarket2/NFTMarket.go
	abigen --abi abigenBindings/abi/NFTMarket3.abi --pkg abimarket3 --out contracts/abimarket3/NFTMarket.go