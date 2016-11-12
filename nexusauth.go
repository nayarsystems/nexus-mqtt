package main

import (
	"errors"

	"github.com/nayarsystems/nxgo"
	"github.com/nayarsystems/nxgo/nxcore"
	"github.com/surgemq/surgemq/auth"
)

type NexusAuthenticator struct {
	*nxcore.NexusConn
}

var nexusAuthenticator NexusAuthenticator

func init() {
	auth.Register("nexus", nexusAuthenticator)
}

func (nx NexusAuthenticator) Authenticate(user string, p interface{}) (err error) {
	password, ok := p.(string)
	if !ok {
		return errors.New("password must be a string")
	}

	nx.NexusConn, err = nxgo.Dial("tcp://localhost:1717", nil)
	if err != nil {
		return err
	}

	_, err = nx.Login(user, password)
	return err
}
