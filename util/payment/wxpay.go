package payment

import (
	"github.com/bysir-zl/bygo/util/payment/core"
	"errors"
)

type WxPayClient struct {
	wxPay *core.WxPay
}

type WxPayNotify struct {
	TradeNO, OutTradeNO string
	Amount              string
}

// 微信调起支付准备
func (p *WxPayClient) CreateWxPayPayInfoJSAPI(tradeNo, subject, totalFee, clientIp, userOpenId string, wxNotifyUrl string) (i core.WXPayReqForJS, err error) {
	if wxNotifyUrl == "" {
		err = errors.New("wxNotifyUrl miss")
		return
	}

	o := p.wxPay.NewUnifiedOrderRequest()
	o.Body = subject
	o.OutTradeNo = tradeNo
	o.TotalFee = totalFee
	o.SpBillCreateIp = clientIp
	o.NotifyURL = wxNotifyUrl
	o.TradeType = "JSAPI"
	o.OpenId = userOpenId

	rsp, err := o.Post()
	if err != nil {
		return
	}
	err = rsp.Error()
	if err != nil {
		return
	}

	i = p.wxPay.NewWXPayReqForJS(rsp.PrePayId)
	return
}

// 检查微信回调
func (p *WxPayClient) CheckWxPayNotify(request []byte) (data WxPayNotify, err error) {
	n, err := p.wxPay.NewWXPayNotify(request)
	if err != nil {
		return
	}
	err = n.IsError()
	if err != nil {
		return
	}

	data = WxPayNotify{
		TradeNO:    n.TransactionId,
		Amount:     n.TotalFee,
		OutTradeNO: n.OutTradeNo,
	}

	return
}

func ResponseWxPayNotify(isSuccess bool, msg string) (rsp string) {
	rspB := core.WXPayResultResponse{
		ReturnCode: "FAIL",
		ReturnMsg:  msg,
	}
	if isSuccess {
		rspB.ReturnCode = "SUCCESS"
	}
	return rspB.ToXML()
}

func NewWxPayClient(appId string, mchId string, mchKey string) (*WxPayClient, error) {
	wxConfig := core.WXKeyConfig{
		APP_ID:  appId,
		MCH_ID:  mchId,
		MCH_KEY: mchKey,
	}
	if appId == "" {
		return nil, errors.New("AppId miss")
	}
	if mchId == "" {
		return nil, errors.New("MchId miss")
	}
	if mchKey == "" {
		return nil, errors.New("mchKey miss")
	}

	return &WxPayClient{
		wxPay: core.NewWxPay(wxConfig),
	}, nil
}
