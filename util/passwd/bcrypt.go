package passwd

import (
	"github.com/savfx/savgo/util/bcrypt"
)

// DefaultCost is the base 2 logarithm of the Bcrypt work factor
// and is exported here for documentation purposes only.
// Use SetCost() to change it.
const DefaultCost = 8

type config struct {
	Cost int
}

func register(c config) error {
	return Register(BCRYPT, &c)
}

func init() {
	err := register(config{DefaultCost})
	if err != nil {
		panic(err)
	}
}

// SetCost sets the cost parameter of the Bcrypt algorithm.
// The value is the base 2 logarithm of the work factor.
func SetCost(cost int) error {
	// punt and see if the underlying algorithm likes the new value!
	_, err := bcrypt.GenerateFromPassword([]byte("password"), cost)
	if err != nil {
		return err
	}
	return register(config{cost})
}

func (c *config) Id() []byte {
	return []byte("2a") //  hashes are prefixed with ..., not "bcrypt"
}

func (c *config) Create(plaintext []byte) (encoded []byte, err error) {
	return bcrypt.GenerateFromPassword(plaintext, c.Cost)
}

func (c *config) Verify(plaintext, encoded []byte) (isValid bool, err error) {
	err = bcrypt.CompareHashAndPassword(encoded, plaintext)
	isValid = err == nil
	if err == bcrypt.ErrMismatchedHashAndPassword {
		err = nil
	}
	return
}

func (c *config) IsCurrent(encoded []byte) (isCurrent bool, err error) {
	cost, err := bcrypt.Cost(encoded)
	if err == nil {
		isCurrent = cost >= c.Cost
	}
	return
}
