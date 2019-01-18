package main

import "fmt"

// Role ...
type Role interface {
	HP() float64
	MP() float64
	Skill() string
	fmt.Stringer
}

// RoleImpl ...
type RoleImpl struct {
	hp    float64
	mp    float64
	skill string
}

// HP ...
func (r *RoleImpl) HP() float64 {
	return r.hp
}

// MP ...
func (r *RoleImpl) MP() float64 {
	return r.mp
}

// Skill ...
func (r *RoleImpl) Skill() string {
	return r.skill
}

func (r *RoleImpl) String() string {
	return "simple role"
}

// NewRole ...
func NewRole(hp, mp float64, skill string) Role {
	return &RoleImpl{
		hp:    hp,
		mp:    mp,
		skill: skill,
	}
}

// Flyer ...
type Flyer interface {
	Skill() string
	fmt.Stringer
}

// FlyerImpl ...
type FlyerImpl struct {
	skill string
}

// Skill ...
func (f *FlyerImpl) Skill() string {
	return f.skill
}

func (f *FlyerImpl) String() string {
	return "simple flyer"
}

// NewFlyer ...
func NewFlyer(speed string) Flyer {
	return &FlyerImpl{
		skill: "fly " + speed,
	}
}

// Bahamut ...
type Bahamut struct {
	Role
	Flyer
}

func (b *Bahamut) String() string {
	return "bahamut"
}

// Skill ...
func (b *Bahamut) Skill() string {
	return "bahamut has " + b.Role.Skill() + " and " + b.Flyer.Skill()
}

// NewBahamut ...
func NewBahamut() *Bahamut {
	return &Bahamut{
		Role:  NewRole(10000, 10000, "fireball"),
		Flyer: NewFlyer("fast"),
	}
}

func main() {
	bahamut := NewBahamut()
	fmt.Println(bahamut)         // bahamut
	fmt.Println(bahamut.Skill()) // bahamut has fireball and fly fast
}
