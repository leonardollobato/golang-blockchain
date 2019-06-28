package blockchain

// BlockChain block chain
type BlockChain struct {
	Blocks []*Block
}

// Block a simple block
type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
	Nonce    int
}

// CreateBlock create a new block
func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash, 0}
	pow := NewProof(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

// AddBlock attach new block to the chain
func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)
	chain.Blocks = append(chain.Blocks, new)
}

// Genesis first chain block
func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

// InitBlockChain initialize the first block
func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}
