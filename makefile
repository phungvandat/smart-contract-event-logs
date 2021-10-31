gen:
	abigen --abi abigenBindings/abi/NFT.abi --pkg abinft --out contracts/abinft/NFT.go
	abigen --abi abigenBindings/abi/NFTMarket.abi --pkg abimarket --out contracts/abimarket/NFTMarket.go