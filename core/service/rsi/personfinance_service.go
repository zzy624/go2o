/**
 * Copyright 2015 @ z3q.net.
 * name : personfinance_service
 * author : jarryliu
 * date : 2016-04-01 09:41
 * description :
 * history :
 */
package rsi

import (
	"errors"
	"github.com/jsix/gof/log"
	"go2o/core/domain/interface/member"
	"go2o/core/domain/interface/personfinance"
	"go2o/core/infrastructure/domain"
	"go2o/core/variable"
)

type personFinanceService struct {
	_rep     personfinance.IPersonFinanceRepository
	_accRepo member.IMemberRepo
}

func NewPersonFinanceService(rep personfinance.IPersonFinanceRepository,
	accRepo member.IMemberRepo) *personFinanceService {
	return &personFinanceService{
		_rep:     rep,
		_accRepo: accRepo,
	}
}

func (p *personFinanceService) GetRiseInfo(personId int64) (
	personfinance.RiseInfoValue, error) {
	pf := p._rep.GetPersonFinance(personId)
	return pf.GetRiseInfo().Value()
}

// 开通增利服务
func (p *personFinanceService) OpenRiseService(personId int64) error {
	m := p._accRepo.GetMember(personId)
	if m == nil {
		return member.ErrNoSuchMember
	}
	if m.GetValue().Level < int32(variable.PersonFinanceMinLevelLimit) {
		return errors.New("会员等级不够,请升级后再开通理财账户！")
	}
	pf := p._rep.GetPersonFinance(personId)
	return pf.CreateRiseInfo()
}

// 提交转入/转出日志
func (p *personFinanceService) CommitTransfer(personId int64, logId int32) error {
	pf := p._rep.GetPersonFinance(personId)
	rs := pf.GetRiseInfo()
	if rs == nil {
		return personfinance.ErrNoSuchRiseInfo
	}
	return rs.CommitTransfer(logId)
}

// 转入(业务放在service,是为person_finance解耦)
func (p *personFinanceService) RiseTransferIn(personId int64,
	transferWith personfinance.TransferWith, amount float32) (err error) {
	//return errors.New("服务暂时不可用")
	pf := p._rep.GetPersonFinance(personId)
	r := pf.GetRiseInfo()
	return r.TransferIn(amount, transferWith)

	//if amount < personfinance.RiseMinTransferInAmount {
	//	//金额不足最低转入金额
	//	return errors.New(fmt.Sprintf(personfinance.ErrLessThanMinTransferIn.Error(),
	//		format.FormatFloat(personfinance.RiseMinTransferInAmount)))
	//}
	//m := p._accRepo.GetMember(personId)
	//if m == nil {
	//	return member.ErrNoSuchMember
	//}
	//acc := m.GetAccount()
	//if transferWith == personfinance.TransferFromWithBalance {
	//	//从余额转入
	//	if err = acc.DiscountBalance("理财转入",
	//		domain.NewTradeNo(10000), amount, member.DefaultRelateUser); err != nil {
	//		return err
	//	}
	//	if err = r.TransferIn(amount, transferWith); err != nil {
	//		//转入
	//		return err
	//	}
	//	return pf.SyncToAccount() //同步到会员账户
	//}
	//
	//if transferWith == personfinance.TransferFromWithWallet {
	//	//从奖金转入
	//	if err := acc.DiscountWallet("理财转入", domain.NewTradeNo(10000),
	//		amount, member.DefaultRelateUser, true); err != nil {
	//		return err
	//	}
	//	if err = r.TransferIn(amount, transferWith); err != nil {
	//		//转入
	//		return err
	//	}
	//	return pf.SyncToAccount() //同步到会员账户
	//}
	//
	//return errors.New("暂时无法提供服务")
}

// 转出
func (p *personFinanceService) RiseTransferOut(personId int64,
	transferWith personfinance.TransferWith, amount float32) (err error) {
	//return errors.New("系统正在升级，暂停服务!")

	pf := p._rep.GetPersonFinance(personId)
	r := pf.GetRiseInfo()

	m := p._accRepo.GetMember(personId)
	if m == nil {
		return member.ErrNoSuchMember
	}
	acc := m.GetAccount()

	if transferWith == personfinance.TransferOutWithBalance {
		//转入余额
		if err = r.TransferOut(amount, transferWith, personfinance.RiseStateOk); err == nil {
			err = acc.Charge(member.AccountBalance,
				member.KindBalanceSystemCharge, variable.AliasGrowAccount+"转出",
				domain.NewTradeNo(10000), amount, member.DefaultRelateUser)
			if err != nil {
				log.Println("[ TransferOut][ Error]:", err.Error())
			}
			err = pf.SyncToAccount()
		}
		return err
	}

	if transferWith == personfinance.TransferFromWithWallet {
		//转入钱包
		if err = r.TransferOut(amount, transferWith, personfinance.RiseStateOk); err == nil {
			err = acc.Charge(member.AccountWallet,
				member.KindWalletAdd, variable.AliasGrowAccount+"转出",
				domain.NewTradeNo(10000), amount, member.DefaultRelateUser)
			if err != nil {
				log.Println("[ TransferOut][ Error]:", err.Error())
			}
			err = pf.SyncToAccount()
		}
		return err
	}

	if transferWith == personfinance.TransferOutWithBank {
		if b := m.Profile().GetBank(); !b.Right() || !b.Locked() {
			return member.ErrNoSuchBankInfo
		}
		if err = r.TransferOut(amount, transferWith,
			personfinance.RiseStateOk); err == nil {
			err = pf.SyncToAccount()
		}
		return err
	}

	return errors.New("暂时无法提供服务")
}

// 结算收益(按日期每天结息)
func (p *personFinanceService) RiseSettleByDay(personId int64,
	settleUnix int64, dayRatio float32) (err error) {
	pf := p._rep.GetPersonFinance(personId)
	r := pf.GetRiseInfo()
	if err = r.RiseSettleByDay(settleUnix, dayRatio); err != nil {
		return err
	}
	return pf.SyncToAccount() //同步到会员账户
}
