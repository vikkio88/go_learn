package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	units := make(map[string]int)
	units["dozen"] = 12
	units["half_of_a_dozen"] = 6
	units["quarter_of_a_dozen"] = 3
	units["great_gross"] = 1728
	units["gross"] = 144
	units["small_gross"] = 120
	return units
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return make(map[string]int)
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	q, exists := units[unit]

	if !exists {
		return false
	}

	bill[item] += q
	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	q, exists := bill[item]
	qUnit, isUnit := units[unit]
	newQ := q - qUnit
	if !exists || !isUnit || (newQ < 0) {
		return false
	}

	if newQ == 0 {
		delete(bill, item)
		return true
	}

	bill[item] = newQ
	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	q, exists := bill[item]

	return q, exists
}
