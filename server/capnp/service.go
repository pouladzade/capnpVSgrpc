package capnp

import (
	"github.com/cryptobank/acm"
	crb "github.com/cryptobank/cryptobank"
	"github.com/cryptobank/server/cdb"
	"log"
	"zombiezen.com/go/capnproto2"
)

type Service struct {
	c  crb.CoreBanking
	db *cdb.CryptoDb
}

func (c *Service) SetDb(db *cdb.CryptoDb) {
	c.db = db
}

func (c *Service) CreateAccount(call crb.CoreBanking_createAccount) error {
	res, _ := crb.NewResponse(call.Results.Segment())
	defer log.Println(res)
	var acc acm.Account
	acc.Name, _ = call.Params.Name()
	acc.Balance = call.Params.Balance()
	acid, err := call.Params.AccountId()
	log.Printf("\n{request: CreateAccount\nparams[accountId = %v, name=%s, balance=%d]}\n",
		acid, acc.Name, acc.Balance)
	if err != nil {
		res.SetMessage(err.Error())
		res.SetCode(-1)
		return call.Results.SetRes(res)
	}
	acc.SetAccountId(acid)
	err = c.db.InsertAccount(acc)
	if err != nil {
		res.SetMessage(err.Error())
		res.SetCode(-10)
		return call.Results.SetRes(res)
	}
	res.SetMessage("Account succesfully created!")
	res.SetCode(0)
	c.db.Commit()
	return call.Results.SetRes(res)
}

func (c *Service) GetAccountInfo(call crb.CoreBanking_getAccountInfo) error {
	res, _ := crb.NewResponse(call.Results.Segment())
	defer log.Println(res)
	acid, err := call.Params.AccountId()
	log.Printf("\n{request: GetAccountInfo\nparams[accountId = %v]}\n",
		acid)
	if err != nil {
		res.SetMessage(err.Error())
		res.SetCode(-1)
		return call.Results.SetRes(res)
	}

	ac, err := c.db.LoadAccount(acid)
	if err != nil || ac == nil {
		res.SetMessage(err.Error())
		res.SetCode(-10)
		return call.Results.SetRes(res)
	}
	res.SetMessage("GetAccountInfo Succesfully done!")
	res.SetCode(0)
	call.Results.SetBalance(ac.Balance)
	call.Results.SetName(ac.Name)
	c.db.Commit()
	return call.Results.SetRes(res)
}

func (c *Service) createResponse(seg *capnp.Segment, code int32, message string) (crb.Response, error) {
	res, err := crb.NewResponse(seg)
	if err != nil {
		log.Println("Can not make the response!")
	}
	res.SetMessage(message)
	res.SetCode(code)
	return res, err
}
