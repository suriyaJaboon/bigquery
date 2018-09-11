package hashs

import "golang.org/x/crypto/bcrypt"

func HashPassWord(password []byte) ([]byte, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func CheckPassWordHash(password, hash []byte) bool {
	return bcrypt.CompareHashAndPassword(hash, password) == nil
}