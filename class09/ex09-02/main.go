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

// NewNPC ...
func NewNPC() Role {
	return &RoleImpl{
		hp:    -1,
		mp:    -1,
		skill: "talk",
	}
}

// Magician ...
type Magician struct {
	Role
}

func (m *Magician) String() string {
	return "magicain has " + m.Skill()
}

// NewMagican ...
func NewMagican(hp, mp float64, skill string) *Magician {
	return &Magician{
		Role: NewRole(hp, mp, skill),
	}
}

// WhoIs ...
func WhoIs(r Role) string {
	return r.String()
}

func main() {
	m := NewMagican(100, 200, "fireball")

	fmt.Println("hp:", m.HP())       // hp: 100
	fmt.Println("mp:", m.MP())       // mp: 200
	fmt.Println("skill:", m.Skill()) // skill: fireball
	fmt.Println("role is", m)        // role is magicain has fireball
	fmt.Println("it is", WhoIs(m))   // it is magicain has fireball
}
