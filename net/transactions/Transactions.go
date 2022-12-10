package transactions

type Transactions struct {
	list []*Transaction
}

func NewTransactions() *Transactions {
	return &Transactions{
		list: make([]*Transaction, 0),
	}
}

func (t *Transactions) Next() *Transaction {
	if len(t.list) == 0 {
		return nil
	}
	tr := t.list[0]
	t.list = t.list[1:]
	return tr
}

func (t *Transactions) Post(tr *Transaction) {
	t.list = append(t.list, tr)
}

func (t *Transactions) Delete(tr *Transaction) {
	for i, v := range t.list {
		if v == tr {
			t.list = append(t.list[:i], t.list[i+1:]...)
			return
		}
	}
}
