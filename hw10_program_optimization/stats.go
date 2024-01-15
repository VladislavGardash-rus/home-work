//nolint:all
package hw10programoptimization

import (
	"bufio"
	"fmt"
	"github.com/mailru/easyjson"
	"io"
	"strings"
)

//easyjson:json
type User struct {
	ID       int
	Name     string
	Username string
	Email    string
	Phone    string
	Password string
	Address  string
}

type DomainStat map[string]int

func GetDomainStat(r io.Reader, domain string) (DomainStat, error) {
	u, err := getUsers(r)
	if err != nil {
		return nil, fmt.Errorf("get users error: %w", err)
	}
	return countDomains(u, domain)
}

type users func() (*User, bool, error)

func getUsers(r io.Reader) (result users, err error) {
	scanner := bufio.NewScanner(r)
	user := new(User)

	result = func() (*User, bool, error) {
		ok := scanner.Scan()
		if scanner.Err() != nil {
			return nil, false, err
		}

		if !ok {
			return nil, false, nil
		}

		err = easyjson.Unmarshal(scanner.Bytes(), user)
		if err != nil {
			return nil, false, err
		}

		return user, ok, nil
	}

	return
}

func countDomains(users users, domain string) (DomainStat, error) {
	result := make(DomainStat)

	for {
		user, ok, err := users()
		if err != nil {
			return nil, err
		}

		if !ok {
			break
		}

		if strings.HasSuffix(user.Email, "."+domain) {
			result[strings.ToLower(strings.SplitN(user.Email, "@", 2)[1])]++
		}
	}

	return result, nil
}
