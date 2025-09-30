package logic

import (
	"cocogame-max/chickenroad2_srv/proto/login_srv/pb_login"
	"cocogame-max/chickenroad2_srv/proto/pb_common"
	"context"
	"time"

	"cocogame-max/chickenroad2_srv/internal/svc"
	"cocogame-max/chickenroad2_srv/pb_chickenroad2"

	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/logx"
)

const gameID = 300

type AuthLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

type CustomClaims struct {
	UserId       string      `json:"userId"`
	Nickname     string      `json:"nickname"`
	Balance      string      `json:"balance"`
	Currency     string      `json:"currency"`
	Operator     string      `json:"operator"`
	OperatorId   string      `json:"operatorId"`
	Meta         interface{} `json:"meta"`
	GameAvatar   interface{} `json:"gameAvatar"`
	SessionToken string      `json:"sessionToken"`
	AuthToken    string      `json:"authToken"`
	jwt.RegisteredClaims
}

func NewAuthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AuthLogic {
	return &AuthLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AuthLogic) Auth(in *pb_chickenroad2.ChickenRoad2_Auth) (*pb_chickenroad2.ChickenRoad2_AuthResp, error) {
	lr, err := l.svcCtx.LoginSrv.Login(l.ctx, &pb_login.Req_Login{
		GameId:              gameID,
		OperatorToken:       in.Operator,
		OperatorPlayerToken: in.AuthToken,
		Ip:                  in.Ip,
	})
	if err != nil {
		return nil, err
	}
	//userInfo, err := l.svcCtx.PlayerCenterSrv.GetRole(context.Background(), &playercentersrv.Req_GetRole{
	//	AuthToken: lr.Token,
	//	GameId:    gameID,
	//})
	//if err != nil {
	//	return nil, err
	//}
	//token, err := l.GenToken(userInfo.OperatorRole, lr.Token)
	//if err != nil {
	//	l.Logger.Errorf("jwt err: %+v", err)
	//	return nil, err
	//}
	return &pb_chickenroad2.ChickenRoad2_AuthResp{
		Token: lr.Token,
	}, nil
}

func (l *AuthLogic) GenToken(operatorRole *pb_common.MOperatorRole, authToken string) (string, error) {
	now := time.Now()
	claims := CustomClaims{
		UserId:       operatorRole.OperatorPlayerId,
		Nickname:     operatorRole.OperatorNickName,
		Currency:     operatorRole.OperatorCurrencyCode,
		Operator:     operatorRole.OperatorToken,
		OperatorId:   operatorRole.OperatorToken,
		Balance:      "5000000",
		Meta:         nil,
		GameAvatar:   nil,
		SessionToken: "4g7fw3",
		AuthToken:    authToken,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(time.Hour * 24)), // 过期时间
			IssuedAt:  jwt.NewNumericDate(now),                     // 签发时间
		},
	}

	// 创建token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(l.svcCtx.Config.JwtAuth.AccessSecret))
}
