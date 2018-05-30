package checker

func (self *Checker) Lgt(length, min int) *Checker {
	if length > min {
		return self
	}
	panic(ErrRuleLgt)
}

func (self *Checker) Lgte(length, min int) *Checker {
	if length >= min {
		return self
	}
	panic(ErrRuleLgt)
}

func (self *Checker) Llt(length, max int) *Checker {
	if length < max {
		return self
	}
	panic(ErrRuleLgt)
}

func (self *Checker) Llte(length, max int) *Checker {
	if length <= max {
		return self
	}
	panic(ErrRuleLgt)
}
