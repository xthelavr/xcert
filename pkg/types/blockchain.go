package types

import (
	"fmt"
	"time"
)

type (
	CertBlock struct {
		*Block
		Cert *Certificate `json:"cert"`
	}

	MasterBlock struct {
		*block
		CertChain CertChain `json:"cert_chain"`
	}

	Block struct {
		Index     int       `json:"index"`
		Hash      string    `json:"hash"`
		PrevHash  string    `json:"prev_hash"`
		Timestamp time.Time `json:"timestamp"`
	}

	CertChain []*CertBlock
	MasterChain []*MasterBlock
)

func (b *MasterBlock) Bytes() []byte {
	return []byte(fmt.Sprint(b))
}

func (b *CertBlock) Bytes() []byte {
	return []byte(fmt.Sprint(b))
}
