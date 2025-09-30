package data_model

import (
	"cocogame-max/chickenroad2_srv/pb_chickenroad2"
	"github.com/globalsign/mgo/bson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

type ChickenRoad2Model struct {
	ParentId           int64   `json:"parent_id" bson:"parent_id"`
	SubId              int64   `json:"sub_id" bson:"sub_id"`
	IsFinished         bool    `json:"is_finished" bson:"is_finished"`
	IsWin              bool    `json:"is_win" bson:"is_win"`
	Currency           string  `json:"currency" bson:"currency"`
	BetAmount          float64 `json:"bet_amount" bson:"bet_amount"`
	Coefficients       float64 `json:"coefficients" bson:"coefficients"`
	WinAmount          float64 `json:"win_amount" bson:"win_amount"`
	Difficulty         string  `json:"difficulty" bson:"difficulty"`
	LineNumber         int32   `json:"line_number" bson:"line_number"`
	CollisionPositions []int32 `json:"collision_positions" bson:"collision_positions"`
	ClientSeed         string  `json:"client_seed" bson:"client_seed"`
	ServerSeed         string  `json:"server_seed" bson:"server_seed"`
	RoleId             string  `json:"role_id" bson:"role_id"`
}

func (r *ChickenRoad2Model) ToPB() proto.Message {
	return &pb_chickenroad2.ChickenRoad2Role{
		ParentId:           r.ParentId,
		SubId:              r.SubId,
		IsFinished:         r.IsFinished,
		IsWin:              r.IsWin,
		Currency:           r.Currency,
		BetAmount:          r.BetAmount,
		Coefficients:       r.Coefficients,
		WinAmount:          r.WinAmount,
		Difficulty:         r.Difficulty,
		LineNumber:         r.LineNumber,
		CollisionPositions: r.CollisionPositions,
		ClientSeed:         r.ClientSeed,
		ServerSeed:         r.ServerSeed,
		RoleId:             r.RoleId,
	}
}

func (r *ChickenRoad2Model) FromPB(pb proto.Message) error {
	if ap, ok := pb.(*anypb.Any); ok {
		pb, _ = ap.UnmarshalNew()
	}
	if role, ok := pb.(*pb_chickenroad2.ChickenRoad2Role); ok {
		r.ParentId = role.ParentId
		r.SubId = role.SubId
		r.IsFinished = role.IsFinished
		r.IsWin = role.IsWin
		r.Currency = role.Currency
		r.BetAmount = role.BetAmount
		r.Coefficients = role.Coefficients
		r.WinAmount = role.WinAmount
		r.Difficulty = role.Difficulty
		r.LineNumber = role.LineNumber
		r.CollisionPositions = role.CollisionPositions
		r.ClientSeed = role.ClientSeed
		r.ServerSeed = role.ServerSeed
		r.RoleId = role.RoleId
		//_ = r.CommonRoleModel.FromPB(role.CR)

		return nil
	}
	return nil
}

func (r *ChickenRoad2Model) Marshal() []byte {
	out, _ := bson.Marshal(r)
	return out
}

func (r *ChickenRoad2Model) Unmarshal(in []byte) {
	_ = bson.Unmarshal(in, r)
}

func (r *ChickenRoad2Model) GetId() int32           { return 300 }
func (r *ChickenRoad2Model) GetGameId() int32       { return 200 }
func (r *ChickenRoad2Model) GetMathVersion() string { return "" }
func (r *ChickenRoad2Model) GetMesh() []int32       { return nil }
func (r *ChickenRoad2Model) GetRoleId() string      { return r.RoleId }
func (r *ChickenRoad2Model) GetBet() float64        { return r.BetAmount }
func (r *ChickenRoad2Model) GetParentId() int64     { return r.ParentId }
func (r *ChickenRoad2Model) GetSubId() int64        { return r.SubId }
func (r *ChickenRoad2Model) GetScore() float64      { return r.WinAmount }
func (r *ChickenRoad2Model) GetGameState() int32    { return 0 }
func (r *ChickenRoad2Model) GetTotalScore() float64 { return 0 }
