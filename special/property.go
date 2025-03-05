package special

type HeadKind string

const (
	HeadKindSkeleton       HeadKind = "skeleton"
	HeadKindWitherSkeleton HeadKind = "wither_skeleton"
	HeadKindPlayer         HeadKind = "player"
	HeadKindZombie         HeadKind = "zombie"
	HeadKindCreeper        HeadKind = "creeper"
	HeadKindPiglin         HeadKind = "piglin"
	HeadKindDragon         HeadKind = "dragon"
)

func (h *HeadKind) IsValid(data HeadKind) bool {
	switch data {
	case HeadKindSkeleton, HeadKindWitherSkeleton, HeadKindPlayer, HeadKindZombie, HeadKindCreeper, HeadKindPiglin, HeadKindDragon:
		return true
	}
	return false
}
