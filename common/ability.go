package common

import "github.com/pazzabec/go-box/log"

// Ability defines a ability can be used
type Ability interface {
	//Activate the ability
	Activate() error

	//Close the ability
	Close() error

	//Reset the ability
	Reset() error

	//IsActive return true if ability is active
	IsActive() bool

	//Shut returns a channel, which is closed once ability is shut down.
	Shut() <-chan struct{}

	//Recorder set a recorder
	Recorder(log.Recorder)

	//Name return skill's name
	Name() string
}

//BaseSkill classical-inheritance-style ability declarations
type BaseSkill struct {
	active   bool
	name     string
	shut     chan struct{}
	recorder log.Recorder
}

//NewBaseSkill create a new BaseSkill
func NewBaseSkill(r log.Recorder) *BaseSkill {
	return &BaseSkill{
		active:   false,
		shut:     make(chan struct{}),
		recorder: r,
	}
}

//Activate the ability
func (s *BaseSkill) Activate() error {
	return nil
}

//Close the ability
func (s *BaseSkill) Close() error {
	return nil
}

//Reset the ability
func (s *BaseSkill) Reset() error {
	return nil
}

//IsActive return true if ability is active
func (s *BaseSkill) IsActive() bool {
	return s.active
}

//Shut returns a channel, which is closed once ability is shut down.
func (s *BaseSkill) Shut() <-chan struct{} {
	return s.shut
}

//Recorder set a recorder
func (s *BaseSkill) Recorder(r log.Recorder) {
	s.recorder = r
}

//Name return skill's name
func (s *BaseSkill) Name() string {
	return s.name
}
