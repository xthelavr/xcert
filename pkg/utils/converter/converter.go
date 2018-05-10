package converter

import (
	"fmt"
	"time"

	"github.com/spacelavr/pandora/pkg/pb"
	"github.com/spacelavr/pandora/pkg/types"
)

// func TPBMB(mb *types.MasterBlock) *pb.MasterBlock {
// 	return &pb.MasterBlock{}
// }
//
// func FPBMB(mb *pb.MasterBlock) *types.MasterBlock {
// 	return &types.MasterBlock{}
// }
//
// func FPBCB(cb *pb.CertBlock) *types.CertBlock {
// 	return &types.CertBlock{}
// }

func FPBMC(pbmc *pb.MasterChain) types.MasterChain {
	mc := types.MasterChain{}

	fmt.Printf("%#v\n", pbmc.MasterBlock)

	for _, mb := range pbmc.MasterBlock {
		cc := types.CertChain{}
		for _, cb := range mb.CertChain.CertBlock {
			cc = append(cc, &types.CertBlock{
				Block: &types.Block{
					Hash:      cb.Block.Hash,
					Timestamp: time.Unix(cb.Block.Timestamp, 0),
					PrevHash:  cb.Block.PrevHash,
					Index:     int(cb.Block.Index),
					Key:       cb.Block.Hash,
				},
				PublicKey: cb.PublicKey.PublicKey,
			})
		}

		mc = append(mc, &types.MasterBlock{
			CertChain: cc,
			PublicKey: mb.PublicKey.PublicKey,
			Block: &types.Block{
				Hash:      mb.Block.Hash,
				Key:       mb.Block.Hash,
				Index:     int(mb.Block.Index),
				PrevHash:  mb.Block.PrevHash,
				Timestamp: time.Unix(mb.Block.Timestamp, 0),
			},
		})
	}

	return mc
}

func FPBCC(cc *pb.CertChain) types.CertChain {
	return types.CertChain{}
}

func TPBMC(mc types.MasterChain) *pb.MasterChain {
	pbmc := &pb.MasterChain{}

	for _, mb := range mc {
		pbmc.MasterBlock = append(pbmc.MasterBlock, &pb.MasterBlock{
			PublicKey: TPBPK(mb.PublicKey),
			CertChain: TPBCC(mb.CertChain),
			Block:     TPBB(mb.Block),
		})
	}

	return pbmc
}

func TPBPK(publicKey string) *pb.PublicKey {
	return &pb.PublicKey{
		PublicKey: publicKey,
	}
}

func TPBB(b *types.Block) *pb.Block {
	return &pb.Block{
		Hash:      b.Hash,
		Timestamp: b.Timestamp.Unix(),
		PrevHash:  b.PrevHash,
		Index:     int64(b.Index),
	}
}

func TPBCC(cc types.CertChain) *pb.CertChain {
	ch := &pb.CertChain{}

	for _, cb := range cc {
		ch.CertBlock = append(ch.CertBlock, &pb.CertBlock{
			Block:     TPBB(cb.Block),
			PublicKey: TPBPK(cb.PublicKey),
		})
	}

	return ch
}
