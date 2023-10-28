package htpasswd

import (
	"errors"
	"os"
)

// Accounts name => hash
type Accounts map[string]string

// Bytes bytes representation
func (a Accounts) Bytes() (passwordBytes []byte) {
	passwordBytes = []byte{}
	for name, hash := range a {
		passwordBytes = append(passwordBytes, []byte(name+PasswordSeparator+hash+LineSeparator)...)
	}
	return passwordBytes
}

// WriteToFile put them to a file will be overwritten or created
func (a Accounts) WriteToFile(file string) error {
	return os.WriteFile(file, a.Bytes(), 0644)
}

// SetPassword set a password for a user with a hashing algo
func (a Accounts) SetPassword(name, password string, hashAlgorithm HashAlgorithm) (err error) {
	if len(password) == 0 {
		return errors.New("passwords must not be empty, if you want to delete a user call RemoveUser")
	}
	var hash string
	var prefix string
	switch hashAlgorithm {
	case HashAPR1:
		hash, err = hashApr1(password)
	case HashBCrypt:
		hash, err = hashBcrypt(password)
	case HashSHA:
		prefix = "{SHA}"
		hash = hashSha(password)

	}
	if err != nil {
		return err
	}
	a[name] = prefix + hash
	return nil
}
