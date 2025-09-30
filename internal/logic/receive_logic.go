package logic

import (
	"cocogame-max/chickenroad2_srv/internal/game_logic"
	"cocogame-max/chickenroad2_srv/internal/model"
	"cocogame-max/chickenroad2_srv/internal/utils/randx"
	"cocogame-max/chickenroad2_srv/proto/conn_gw/conngwservice"
	"cocogame-max/chickenroad2_srv/proto/operatorproxy_srv/pb_operatorproxy"
	"cocogame-max/chickenroad2_srv/proto/order_srv/pb_order"
	"cocogame-max/chickenroad2_srv/proto/pb_common"
	"cocogame-max/chickenroad2_srv/proto/playercenter_srv/pb_playercenter"
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"google.golang.org/protobuf/types/known/anypb"
	"strconv"
	"strings"
	"time"

	"cocogame-max/chickenroad2_srv/internal/svc"
	"cocogame-max/chickenroad2_srv/pb_chickenroad2"

	"github.com/zeromicro/go-zero/core/logx"
)

const (
	msgInHeader    = "42"
	msgOutHeader   = "43"
	heartBeatIn    = "3"
	initIn         = "40"
	gameService    = "gameService"
	historyService = "\"gameService-get-my-bets-history\""
	actionConfig   = "get-game-config"
	actionState    = "get-game-state"
	actionSeed     = "get-game-seeds"
	actionBet      = "bet"
	actionStep     = "step"
	actionWithdraw = "withdraw"
	null           = "[null]"
	seedDemo       = "[{\"userSeed\":\"34f1f665e945a12e\",\"hashedServerSeed\":\"ac1d379ea3ff6c239742725c2879dcab7aa6161cca48b1b0551089889052d6a3497a9300db93c2e27af5516b0e5926ca0ac7e86a53f1419078ceb43c97c94e60\",\"nonce\":\"0\"}]"
	historyDemo    = "[[]]"
	balanceChange  = "42[\"onBalanceChange\",{\"currency\":\"%s\",\"balance\":\"%s\"}]"
	errorMsg       = "[{\"error\":{\"message\":\"An error occurred while processing the bet\"}}]"
)

var (
	heartBeatOut = []byte("2")
	initOut1     = []byte("40{\"sid\":\"FatGk_OFR3LGrweGPVgn\"}")
	initOut2     = []byte("42[\"onBalanceChange\",{\"currency\":\"USD\",\"balance\":\"50000\"}]")
	initOut3     = []byte("42[\"betsRanges\",{\"USD\":[\"0.01\",\"200.00\"]}]")
	initOut4     = []byte("42[\"betsConfig\",{\"USD\":{\"betPresets\":[\"0.5\",\"1\",\"2\",\"7\"],\"minBetAmount\":\"0.01\",\"maxBetAmount\":\"200.00\",\"maxWinAmount\":\"20000.00\",\"defaultBetAmount\":\"0.60\",\"decimalPlaces\":null}}]")
	initOut5     = []byte("42[\"myData\",{\"userId\":\"5b381bfc-6e6f-4095-a7a6-561ee6391fa2\",\"nickname\":\"Plum Rear Coral\",\"gameAvatar\":null}]")
	initOut6     = []byte("42[\"currencies\",{\"ADA\":1.2463933582788993,\"AED\":3.6725,\"AFN\":70,\"ALL\":85.295,\"AMD\":383.82,\"ANG\":1.8022999999999998,\"AOA\":918.65,\"ARS\":1371.4821,\"AUD\":1.5559,\"AWG\":1.79,\"AZN\":1.7,\"BAM\":1.6673340771999998,\"BBD\":2.0181999999999998,\"BCH\":0.0017917458999608114,\"BDT\":122.24999999999999,\"BGN\":1.712,\"BHD\":0.377,\"BIF\":2981,\"BMD\":1,\"BNB\":0.0009894131139735745,\"BND\":1.2974999999999999,\"BOB\":6.907100000000001,\"BRL\":5.6015,\"BSD\":0.9997,\"BTC\":0.000008787276449358826,\"BTN\":88.6837706508,\"BUSD\":0.9996936638705801,\"BWP\":13.6553,\"BYN\":3.2712,\"BZD\":2.0078,\"CAD\":1.3858,\"CDF\":2641.8335286835,\"CHF\":0.8140000000000001,\"CLF\":0.0245760857,\"CLP\":972.65,\"COP\":4186.71,\"CRC\":505.29,\"CSC\":14672.743321885506,\"CUP\":23.990199999999998,\"CVE\":94.0045549397,\"CZK\":21.5136,\"DASH\":0.044536645157364316,\"DJF\":178.08,\"DKK\":6.5351,\"DLS\":33.333333333333336,\"DOGE\":4.278215429709092,\"DOP\":61,\"DZD\":130.923,\"EGP\":48.57,\"EOS\":1.2787330681036353,\"ERN\":15,\"ETB\":138.20000000000002,\"ETC\":0.05419504072411546,\"ETH\":0.00024008542131002082,\"EUR\":0.8755000000000001,\"FJD\":2.2723999999999998,\"FKP\":0.7444686628,\"GBP\":0.7571,\"GC\":1,\"GEL\":2.7035,\"GHS\":10.5,\"GIP\":0.7444686628,\"GMD\":72.815,\"GMS\":1,\"GNF\":8674.5,\"GTQ\":7.675,\"GYD\":209.299975191,\"HKD\":7.849799999999999,\"HNL\":26.2787,\"HRK\":6.4231188827,\"HTG\":131.16899999999998,\"HUF\":350.19,\"IDR\":16443.4,\"ILS\":3.3960999999999997,\"INR\":87.503,\"IQD\":1310,\"IRR\":42112.5,\"ISK\":124.46999999999998,\"JMD\":159.94400000000002,\"JOD\":0.709,\"JPY\":150.81,\"KES\":129.2,\"KGS\":87.45,\"KHR\":4015,\"KMF\":431.5,\"KPW\":899.9783434651,\"KRW\":1392.51,\"KWD\":0.30610000000000004,\"KYD\":0.8317894154000001,\"KZT\":540.8199999999999,\"LAK\":21580,\"LBP\":89550,\"LKR\":302.25,\"LRD\":180.5371034311,\"LSL\":18.2179,\"LTC\":0.009405582347999588,\"LYD\":5.415,\"MAD\":9.154300000000001,\"MDL\":17.08,\"MGA\":4430,\"MKD\":52.885000000000005,\"MMK\":3247.961,\"MNT\":3590,\"MOP\":8.089,\"MRU\":40.0019153095,\"MUR\":46.65,\"MVR\":15.459999999999999,\"MWK\":1733.67,\"MXN\":18.869,\"MYR\":4.265,\"MZN\":63.910000000000004,\"NAD\":18.2179,\"NGN\":1532.39,\"NIO\":36.75,\"NOK\":10.3276,\"NPR\":140.07,\"NZD\":1.6986,\"OMR\":0.385,\"PAB\":1.0009,\"PEN\":3.569,\"PGK\":4.1303,\"PHP\":58.27,\"PKR\":283.25,\"PLN\":3.7442,\"PYG\":7486.400000000001,\"QAR\":3.6408,\"R$\":476.1904761904762,\"RON\":4.440300000000001,\"RSD\":102.56500000000001,\"RUB\":79.87530000000001,\"RWF\":1440,\"SAR\":3.7513,\"SBD\":8.4976180797,\"SC\":1,\"SCR\":14.1448,\"SDG\":600.5,\"SEK\":9.7896,\"SGD\":1.2979,\"SHIB\":84033.61344537816,\"SHP\":0.7444686628,\"SLE\":22.7799892929,\"SOL\":0.0047611848038791055,\"SOS\":571.5,\"SRD\":37.6554720237,\"SSP\":130.26,\"SVC\":8.7464,\"SYP\":13005,\"SZL\":18.01,\"THB\":32.752,\"TND\":2.88,\"TON\":0.37226561631825705,\"TRX\":2.9781654607975443,\"TRY\":40.6684,\"TWD\":29.918000000000003,\"TZS\":2570,\"UAH\":41.6966,\"uBTC\":8.787276449358828,\"UGX\":3583.3,\"USD\":1,\"USDC\":1.00072715837509,\"USDT\":1,\"UYU\":40.0886,\"UZS\":12605,\"VEF\":17538792.500838745,\"VES\":123.7216,\"VND\":26199,\"XAF\":573.151,\"XLM\":2.689153150176188,\"XMR\":0.008457936691358008,\"XOF\":566.5,\"XRP\":0.3456560892873093,\"ZAR\":18.2178,\"ZEC\":0.014749559092204408,\"ZMW\":23.799624439699997,\"ZWL\":26.852999999999998}]")
)

type ReceiveLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewReceiveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReceiveLogic {
	return &ReceiveLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ReceiveLogic) Receive(in *pb_chickenroad2.ReceiveRequest) (e *pb_chickenroad2.Empty, err error) {
	payLoad := string(in.Payload)
	l.Infof("payLoad real %s", payLoad)
	messages := strings.Split(payLoad, "[")
	if len(messages) < 2 {
		if messages[0] == heartBeatIn {
			l.Unicast(in, heartBeatOut)
		} else if messages[0] == initIn {
			l.Unicast(in, initOut1)
			l.Unicast(in, initOut2)
			l.Unicast(in, initOut3)
			l.Unicast(in, initOut4)
			l.Unicast(in, initOut5)
			l.Unicast(in, initOut6)
		}
		return
	}
	msgHeader := messages[0]
	if msgHeader[:2] != msgInHeader {
		return
	}
	msgNo := msgOutHeader + msgHeader[2:] // 消息序号
	msgBody := messages[1]
	if strings.Contains(msgBody, actionConfig) { // 配置
		gc, _ := json.Marshal(game_logic.GameConfig)
		l.Unicast(in, []byte(msgNo+string(gc)))
	} else if strings.Contains(msgBody, actionState) { // 游戏状态 结束返回 null
		l.Unicast(in, []byte(msgNo+null))
	} else if strings.Contains(msgBody, actionSeed) {
		l.Unicast(in, []byte(msgNo+seedDemo))
	} else if strings.Contains(msgBody, historyService) {
		l.Unicast(in, []byte(msgNo+historyDemo))
	} else {
		l.play(in, msgNo, msgBody)
	}
	return
}

func (l *ReceiveLogic) play(in *pb_chickenroad2.ReceiveRequest, msgNo, body string) {
	bodies := strings.SplitN(body, "gameService\",", 2)
	if len(bodies) < 2 {
		return
	}
	var action = &pb_chickenroad2.Action{}
	err := json.Unmarshal([]byte(bodies[1][:len(bodies[1])-1]), action)
	if err != nil {
		logx.Errorf("unmarshal action err: %v", err)
		return
	}

	var (
		pcRole *pb_playercenter.Res_GetRole
		model  = data_model.ChickenRoad2Model{}
	)
	if pcRole, err = l.svcCtx.PlayerCenterSrv.GetRole(l.ctx, &pb_playercenter.Req_GetRole{
		AuthToken: in.UserInfo.Token,
		GameId:    gameID}); err != nil {
		l.Errorf("get role err: %v", err)
		return
	}

	_ = model.FromPB(pcRole.Role)
	var (
		order *pb_order.Res_OrderNo
	)
	if order, err = l.svcCtx.OrderSrv.ProduceOrderNo(l.ctx, &pb_order.Req_OrderNo{AuthToken: in.UserInfo.Token}); err != nil {
		return
	}
	defer l.svcCtx.OrderSrv.ConsumeOrderNo(l.ctx, &pb_order.Req_OrderNo{AuthToken: in.UserInfo.Token})
	switch action.Action {
	case actionBet:
		l.bet(in, &model, action.Payload, msgNo, pcRole.GetOperatorRole(), order.OrderNo)
	case actionStep:
		l.step(in, &model, action.Payload, msgNo, pcRole.GetOperatorRole(), order.OrderNo)
	case actionWithdraw:
		l.withdraw(in, &model, msgNo, pcRole.GetOperatorRole())
	default:
		return
	}
	// 历史账单
	role := model.ToPB()
	ay, _ := anypb.New(role)
	_, err = l.svcCtx.PlayerCenterSrv.SetRole(l.ctx, &pb_playercenter.Req_SetRole{
		AuthToken: in.UserInfo.Token,
		Role:      ay,
		GameId:    gameID,
	})
	if err != nil {
		l.Errorf("set role err: %v", err)
		return
	}
}

func (l *ReceiveLogic) bet(in *pb_chickenroad2.ReceiveRequest, model *data_model.ChickenRoad2Model, payload *pb_chickenroad2.ActionPayload, msgNo string, operatorRole *pb_common.MOperatorRole, orderNo int64) {
	//  {"action":"bet","payload":{"betAmount":"0.6","currency":"USD","difficulty":"EASY","countryCode":"CN"}}
	var (
		err error
	)
	if !model.IsFinished && model.ParentId != 0 { // 游戏中
		return
	}
	betAmount, _ := strconv.ParseFloat(payload.BetAmount, 64)
	bet := decimal.NewFromFloat(betAmount)
	if bet.GreaterThan(game_logic.MaxBet) || bet.LessThan(game_logic.MinBet) {
		return
	}

	model.ParentId = orderNo
	model.SubId = orderNo
	model.IsFinished = false
	model.Currency = payload.Currency
	model.Difficulty = payload.Difficulty
	model.LineNumber = -1
	model.BetAmount = bet.InexactFloat64()
	model.Coefficients = 1
	model.WinAmount = bet.InexactFloat64()
	// 扣钱 		err = l.cashSub(pcRole.GetOperatorRole(), bet, betLevel, model.ParentId, model.SubId, PLFSActivity, fsActivityAmount, freeSpinId, freeSpinCount)
	rsp, err := l.cashSub(operatorRole, betAmount, betAmount, model.ParentId, model.SubId)
	if err != nil {
		return
	}
	l.Unicast(in, []byte(fmt.Sprintf(balanceChange, model.Currency, fmt.Sprintf("%.2f", rsp.BalanceAmount)))) // 42["onBalanceChange",{"currency":"USD","balance":"1000000.00"}] // 更新账户
	// [{"isFinished":false,"currency":"USD","betAmount":"0.6","coeff":"1","winAmount":"0.60","difficulty":"EASY","lineNumber":-1}]
	step := &pb_chickenroad2.GameStep{
		IsFinished: model.IsFinished,
		Currency:   model.Currency,
		Difficulty: model.Difficulty,
		LineNumber: model.LineNumber,
		BetAmount:  fmt.Sprintf("%.2f", model.BetAmount),
		WinAmount:  fmt.Sprintf("%.2f", model.WinAmount),
		Coeff:      "1",
	}
	b, _ := json.Marshal(step)
	out := msgNo + "[" + string(b) + "]"
	l.Unicast(in, []byte(out))

}
func (l *ReceiveLogic) step(in *pb_chickenroad2.ReceiveRequest, model *data_model.ChickenRoad2Model, payload *pb_chickenroad2.ActionPayload, msgNo string, operatorRole *pb_common.MOperatorRole, orderNo int64) {
	if model.IsFinished { // 游戏完成
		return
	}
	if payload.LineNumber <= model.LineNumber || payload.LineNumber != model.LineNumber+1 { // 错误step
		return
	}
	coffs := game_logic.CoffMap[model.Difficulty]
	coffsLen := int32(len(coffs)) // [1.01 1.03 1.06]
	// 3 < 3
	if coffsLen == 0 || coffsLen <= payload.LineNumber {
		return
	}
	model.SubId = orderNo
	model.LineNumber = payload.LineNumber
	bet := decimal.NewFromFloat(model.BetAmount)
	coff := coffs[payload.LineNumber]
	step := &pb_chickenroad2.GameStep{
		BetAmount:  fmt.Sprintf("%.2f", model.BetAmount),
		Currency:   model.Currency,
		Difficulty: model.Difficulty,
		LineNumber: model.LineNumber,
	}
	flag := 99 - model.LineNumber
	if randx.Int31r(0, 100) > flag { // 爆炸
		model.IsFinished = true
		model.IsWin = false
		model.Coefficients = 0
		model.WinAmount = 0
		model.CollisionPositions = []int32{model.LineNumber}

		step.IsFinished = true
		step.IsWin = false
		step.Coeff = "0"
		step.WinAmount = "0.00"
		step.CollisionPositions = []int32{model.LineNumber}
	} else {
		score := bet.Mul(decimal.NewFromFloat(coff))
		model.Coefficients = coff
		step.Coeff = fmt.Sprintf("%.2f", coff)
		model.WinAmount = score.InexactFloat64()
		step.WinAmount = score.StringFixed(2)
		step.CollisionPositions = []int32{}
		model.CollisionPositions = []int32{}
		if payload.LineNumber == coffsLen-1 { //过关
			model.IsFinished = true
			step.IsFinished = true
			model.IsWin = true
			step.IsWin = true
			// 加钱 l.cashAdd(pcRole.GetOperatorRole(), model.Score, betLevel, model.ParentId, model.SubId, roundEnd,)
			rsp, err := l.cashAdd(operatorRole, score.InexactFloat64(), model.BetAmount, model.ParentId, model.SubId, true)
			if err != nil {
				return
			}
			l.Unicast(in, []byte(fmt.Sprintf(balanceChange, model.Currency, fmt.Sprintf("%.2f", rsp.BalanceAmount)))) // 42["onBalanceChange",{"currency":"USD","balance":"1000000.00"}] // 更新账户
		} else {
			model.IsFinished = false
			step.IsFinished = false
			model.IsWin = false
			step.IsWin = false
		}
	}
	b, _ := json.Marshal(step)
	out := msgNo + "[" + string(b) + "]"
	l.Unicast(in, []byte(out))
	// 4335[{"isFinished":false,"currency":"USD","betAmount":"1","coeff":"6.50","winAmount":"6.50","difficulty":"EASY","lineNumber":26}]
	// 4336[{"isFinished":true,"isWin":false,"currency":"USD","betAmount":"1","coeff":"0","winAmount":"0.00","difficulty":"EASY","lineNumber":27,"collisionPositions":[27]}]
}

func (l *ReceiveLogic) withdraw(in *pb_chickenroad2.ReceiveRequest, model *data_model.ChickenRoad2Model, msgNo string, operatorRole *pb_common.MOperatorRole) {
	if model.IsFinished {
		return
	}
	model.IsFinished = true
	model.IsWin = true
	model.CollisionPositions = []int32{21}
	// 结算
	rsp, err := l.cashAdd(operatorRole, model.WinAmount, model.BetAmount, model.ParentId, model.SubId, true)
	if err != nil {
		return
	}
	l.Unicast(in, []byte(fmt.Sprintf(balanceChange, model.Currency, fmt.Sprintf("%.2f", rsp.BalanceAmount))))
	//439[{"isFinished":true,"isWin":true,"currency":"USD","betAmount":"0.6","coeff":"1.03","winAmount":"0.61","difficulty":"EASY","lineNumber":1,"collisionPositions":[21]}]
	step := &pb_chickenroad2.GameStep{
		IsFinished:         true,
		IsWin:              true,
		Currency:           model.Currency,
		BetAmount:          fmt.Sprintf("%.2f", model.BetAmount),
		WinAmount:          fmt.Sprintf("%.2f", model.WinAmount),
		Difficulty:         model.Difficulty,
		LineNumber:         model.LineNumber,
		CollisionPositions: []int32{11},
	}
	b, _ := json.Marshal(step)
	out := msgNo + "[" + string(b) + "]"
	l.Unicast(in, []byte(out))
}

func (l *ReceiveLogic) Unicast(in *pb_chickenroad2.ReceiveRequest, payLoad []byte) {
	uu, _ := uuid.NewUUID()
	_, err := l.svcCtx.GW.Unicast(l.ctx, &conngwservice.UnicastRequest{
		RequestId: uu.String(),
		UserId:    in.UserInfo.Token,
		RoomId:    in.RoomId,
		SendType:  1,
		Payload:   payLoad,
	})
	if err != nil {
		l.Errorf("Unicast err: %v", err)
	}
}

func (l *ReceiveLogic) cashSub(operatorRole *pb_common.MOperatorRole, amount, betLevel float64, parentId, subId int64) (*pb_operatorproxy.Res_CashSub, error) {
	return l.svcCtx.OperatorProxySrv.CashSub(l.ctx, &pb_operatorproxy.Req_CashSub{
		OperatorRole: operatorRole,
		Amount:       amount,
		BetLevel:     betLevel,
		GameId:       l.svcCtx.GameID,
		BetId:        parentId,
		SubId:        subId,
		CreateTime:   time.Now().UnixNano(),
	})
}

func (l *ReceiveLogic) cashAdd(operatorRole *pb_common.MOperatorRole, amount, betLevel float64, parentId, subId int64, isEndRound bool) (*pb_operatorproxy.Res_CashAdd, error) {
	return l.svcCtx.OperatorProxySrv.CashAdd(l.ctx, &pb_operatorproxy.Req_CashAdd{
		OperatorRole: operatorRole,
		Amount:       amount,
		BetLevel:     betLevel,
		GameId:       l.svcCtx.GameID,
		BetId:        parentId,
		SubId:        subId,
		IsEndRound:   isEndRound,
		CreateTime:   time.Now().UnixNano(),
	})
}
