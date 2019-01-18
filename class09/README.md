# 09 OOP in Go

物件導向基本三個特性：

1. 封裝 (Encapsulation)
1. 繼承 (Inheritance)
1. 多型 (Polymorphism)

封裝的部分，與 struct/method 方式相同。Go 並沒有繼承的特性，但是我們可以利用 **Struct Embedding and Anonymous Fields** 與 interface 來達到繼承的效果。

以下我們用遊戲角色設計當範例，來說明如何在 Go 實作 OOP。

## 初步實作

### Role

```go {.line-numbers}
// Role ...
type Role struct {
    hp    float64
    mp    float64
    skill string
}

// HP ...
func (r *Role) HP() float64 {
    return r.hp
}

// MP ...
func (r *Role) MP() float64 {
    return r.mp
}

// Skill ...
func (r *Role) Skill() string {
    return r.skill
}

func (r *Role) String() string {
    return "simple role"
}

// NewRole ...
func NewRole(hp, mp float64, skill string) *Role {
    return &Role{
        hp:    hp,
        mp:    mp,
        skill: skill,
    }
}
```

### Magician

繼承 **Role**

```go {.line-numbers}
// Magician ...
type Magician struct {
    *Role
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

m := NewMagican(100, 200, "fireball")

fmt.Println("hp:", m.HP())       // hp: 100
fmt.Println("mp:", m.MP())       // mp: 200
fmt.Println("skill:", m.Skill()) // skill: fireball
fmt.Println("role is ", m)       // role is  magicain has fireball
```

利用 **Struct Embedding and Anonymous Fields** 將 `*Role` 當作是 `Magician` 的變數。如此，我們可以直接呼叫 `*Role` 的 methods:

1. `m.HP()` 等同 `m.Role.HP()`
1. `m.MP()` 等同 `m.Role.MP()`
1. `m.Skill()` 等同 `m.Role.Skill()`

在 Go 也可以 override method. 在這邊，`Magician` override `Role` 的 `String()` method.

## 為什麼 Go 沒有繼承

先看以下的範例

```go {.line-numbers}
// WhoIs ...
func WhoIs(r *Role) string {
    return r.String()
}

fmt.Println("it is ", WhoIs(m))  // compile error: cannot use m (type *Magician) as type *Role in argument to WhoIs
```

在一般 OOP 的程式語言，`fmt.Println("it is ", WhoIs(m))` 是對的。但是 Go 並沒有繼承，只是使用 **Struct Embedding and Anonymous Fields** 的方式，讓程式語法有繼承的效果。在 `func WhoIs(r *Role) string`, parameter type 是 `*Role`，因此不能傳 `*Magician`。

1. `*Role` 是 `Magician` 的 member data。
1. Go 是 strong type，`*Role` 與 `*Magician` 是不同的 data type.
1. Go 壓根就沒有繼承。

## 解決 Go Strong Type 問題

因為 Go 是 strong type, 因此不同 struct 都會被視為不同的 data type。要讓 Go 有完整繼承的效果，就需要用到 interface。

### 定義 Role interface

```go {.line-numbers}
// Role ...
type Role interface {
    HP() float64
    MP() float64
    Skill() string
    fmt.Stringer
}
```

`Role` 也加入 `fmt.Stringer` interface, 之後實作 `Role` 的 struct 也要實作 `String() string` method.

### 定義 RoleImpl 實作 Role

主要方便之後實作 `Role` interface.

```go {.line-numbers}
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
```

### 定義 Magician 內含 Role

```go {.line-numbers}
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
```

如此, `Magician` 實作了 `Role`。

### 測試

```go {.line-numbers}
// WhoIs ...
func WhoIs(r Role) string {
    return r.String()
}

m := NewMagican(100, 200, "fireball")

fmt.Println("hp:", m.HP())       // hp: 100
fmt.Println("mp:", m.MP())       // mp: 200
fmt.Println("skill:", m.Skill()) // skill: fireball
fmt.Println("role is", m)        // role is magicain has fireball
fmt.Println("it is", WhoIs(m))   // it is magicain has fireball
```

## 多重繼承與 Ambiguous

Go 可以在 struct 內含多個 struct 或 interface 來達到多重繼承的效果。

### 定義 Flyer 與 FlyerImpl

```go {.line-number}
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
```

### 定義 Bahamut 繼承 Role 與 Flyer

```go {.line-numbers}
// Bahamut ...
type Bahamut struct {
    Role
    Flyer
}

// NewBahamut ...
func NewBahamut() *Bahamut {
    return &Bahamut{
        Role:  NewRole(10000, 10000, "fireball"),
        Flyer: NewFlyer("fast"),
    }
}

bahamut := NewBahamut()
fmt.Println(bahamut)         // &{simple role simple flyer}
fmt.Println(bahamut.Skill()) // compile error: ambiguous selector bahamut.Skill
fmt.Println(bahamut.Role.Skill())  // fireball
fmt.Println(bahamut.Flyer.Skill()) // fly fast
```

### 用 Override 修正 Ambiguous 問題

在 `Bahamut` 實作 `String()` 與 `Skill()`。

```go {.line-numbers}
func (b *Bahamut) String() string {
    return "bahamut"
}

// Skill ...
func (b *Bahamut) Skill() string {
    return "bahamut has " + b.Role.Skill() + " and " + b.Flyer.Skill()
}

fmt.Println(bahamut)         // bahamut
fmt.Println(bahamut.Skill()) // bahamut has fireball and fly fast
```

## Visible

在 Go 沒有 `public`, `protected`, 及 `private` 等關鍵字，是用名稱**第一個字母大小寫**，來分 public 還是 private. **大寫** 是 public, 小寫是 **private**。

重點整理：

1. 大寫是 public
1. 小寫是 private
1. 在同 package 下，可以存取不同 struct 內的 private 變數
1. 不同 package 只能存取 public 變數。

常見的誤用：

```go {.line-numbers}
type supplier struct {
    ID   int
    Name string
}

// GetSupplier ...
func GetSupplier(id int) *supplier { // warning: exported func GetSupplier returns unexported type *visible.supplier, which can be annoying to use
    return &supplier{id, fmt.Sprintf("test-%d", id)}
}

s := visible.GetSupplier(1) // it is fine
fmt.Println(s.ID)
```

在 Go 以上的寫法是可以過 compile 的，也可以存取 private struct 內的 public variable. 但是無法用 `var s *visible.supplier = GetSupplier(1)`。