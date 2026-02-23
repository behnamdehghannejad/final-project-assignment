# 🧪 Final Project Assignment

## Mini OMS – Order & Voucher Engine (In-Memory)

Design and implement a simplified Order Management System (OMS) entirely in memory.
Follow the required signatures strictly.

---

## 1️⃣ Order State Machine

### ✅ What You Must Implement

Define the state type and constants:

```go
type OrderState int

const (
	Created OrderState = iota
	Paid
	VendorAccepted
	Shipped
	Delivered
	Cancelled
)
```

Implement transition validation:

```go
func (o *Order) ChangeState(newState OrderState) error
```

### 🎯 Requirements

- Only valid transitions allowed.
- Cancelled is terminal.
- Use switch.
- Return descriptive errors.

---

## 2️⃣ Product and OrderItem

### Required Types

```go
type Product struct {
	ID    int
	Name  string
	Price int
}

type OrderItem struct {
	Product  Product
	Quantity int
}
```

### Required Methods

```go
func NewOrderItem(product Product, qty int) (OrderItem, error)

func (oi OrderItem) TotalPrice() int
```

### 🎯 Requirements

- Prevent zero/negative quantity.
- TotalPrice = Price × Quantity.
- No invalid construction.

---

## 3️⃣ Voucher System (Interface-Based Design)

### Interface

```go
type Voucher interface {
	Apply(total int) int
	Code() string
}
```

### Implementations

**Percentage Voucher:**

```go
type PercentageVoucher struct {
	code    string
	percent int
}

func NewPercentageVoucher(code string, percent int) (*PercentageVoucher, error)

func (v *PercentageVoucher) Apply(total int) int
func (v *PercentageVoucher) Code() string
```

**Fixed Amount Voucher:**

```go
type FixedAmountVoucher struct {
	code   string
	amount int
}

func NewFixedAmountVoucher(code string, amount int) (*FixedAmountVoucher, error)

func (v *FixedAmountVoucher) Apply(total int) int
func (v *FixedAmountVoucher) Code() string
```

### 🎯 Requirements

- Validate percent (1–100).
- Validate fixed amount (>0).
- Do not allow negative result totals.
- Handle nil interface correctly.

---

## 4️⃣ Order Aggregate

### Structure

```go
type Order struct {
	ID      int
	Items   []OrderItem
	Voucher Voucher
	State   OrderState
}
```

### Constructor

```go
func NewOrder(id int) *Order
```

### Required Methods

```go
func (o *Order) AddItem(product Product, qty int) error

func (o *Order) ApplyVoucher(v Voucher) error

func (o *Order) TotalAmount() (int, error)

func (o *Order) Pay() error

func (o *Order) Cancel() error

func (o *Order) SnapshotItems() []OrderItem
```

### 🎯 Requirements

- Use append for items.
- Use copy in SnapshotItems.
- Prevent applying voucher twice.
- Prevent zero/negative final totals.
- Enforce business rules.
- Use defer meaningfully in at least one method.

---

## 5️⃣ Repository Layer (In-Memory)

### Interface

```go
type OrderRepository interface {
	Save(order *Order) error
	FindByID(id int) (*Order, error)
	List() []Order
	Delete(id int) error
	Clear()
}
```

### Implementation

```go
type InMemoryOrderRepo struct {
	orders map[int]*Order
}

func NewInMemoryOrderRepo() *InMemoryOrderRepo

func (r *InMemoryOrderRepo) Save(order *Order) error

func (r *InMemoryOrderRepo) FindByID(id int) (*Order, error)

func (r *InMemoryOrderRepo) List() []Order

func (r *InMemoryOrderRepo) Delete(id int) error

func (r *InMemoryOrderRepo) Clear()
```

### 🎯 Requirements

- Initialize map properly.
- Use comma ok.
- Decide pointer vs value behavior consciously.
- Avoid exposing internal map directly.

---

## 6️⃣ Higher-Order Function

### Required Function

```go
func FilterOrders(orders []Order, fn func(Order) bool) []Order
```

### 🎯 Example Usage

```go
paidOrders := FilterOrders(allOrders, func(o Order) bool {
	return o.State == Paid
})
```

### 🎯 Requirements

- Must allocate a new slice.
- Use anonymous functions.
- Demonstrate closure (e.g., minAmount filter).

---

## 📏 Business Rules to Enforce

- Cannot ship before paid.
- Cannot deliver before shipped.
- Cannot transition after cancellation.
- Cannot apply voucher twice.
- Cannot allow total ≤ 0.
- All violations must return errors.

---

## ⚙️ Optional Advanced Signatures

**Custom error example:**

```go
type InvalidStateTransitionError struct {
	From OrderState
	To   OrderState
}

func (e InvalidStateTransitionError) Error() string
```

**Set implementation:**

```go
type StringSet map[string]struct{}

func (s StringSet) Add(val string)
func (s StringSet) Contains(val string) bool
func (s StringSet) Remove(val string)
```

---

## 📦 Deliverables

- Fully working Go module
- Clean structure
- README explaining:
  - Pointer vs value decisions
  - Interface reasoning
  - State validation logic
  - Edge cases handled
