package bus

import (
	"github.com/pazzabec/go-box/common"
)

//Bus defines systerm'bus
type Bus struct {
	common.BaseSkill

	skills map[string]Skill
}

// Skill defines a skill can be used
type Skill interface {
	common.Ability

	SetBus(Bus)
}
