// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: protobuf/catan.proto

package pb

import (
	context "context"
	models "github.com/vulpes-ferrilata/catan-service-proto/pb/models"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CatanClient is the client API for Catan service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CatanClient interface {
	FindGamePaginationByLimitByOffset(ctx context.Context, in *models.FindGamePaginationByLimitByOffsetRequest, opts ...grpc.CallOption) (*models.GamePagination, error)
	GetGameDetailByIDByUserID(ctx context.Context, in *models.GetGameDetailByIDByUserIDRequest, opts ...grpc.CallOption) (*models.GameDetail, error)
	CreateGame(ctx context.Context, in *models.CreateGameRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	JoinGame(ctx context.Context, in *models.JoinGameRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	StartGame(ctx context.Context, in *models.StartGameRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	BuildSettlementAndRoad(ctx context.Context, in *models.BuildSettlementAndRoadRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	RollDices(ctx context.Context, in *models.RollDicesRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DiscardResourceCards(ctx context.Context, in *models.DiscardResourceCardsRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	MoveRobber(ctx context.Context, in *models.MoveRobberRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	EndTurn(ctx context.Context, in *models.EndTurnRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	BuildSettlement(ctx context.Context, in *models.BuildSettlementRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	BuildRoad(ctx context.Context, in *models.BuildRoadRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UpgradeCity(ctx context.Context, in *models.UpgradeCityRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	BuyDevelopmentCard(ctx context.Context, in *models.BuyDevelopmentCardRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ToggleResourceCards(ctx context.Context, in *models.ToggleResourceCardsRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	MaritimeTrade(ctx context.Context, in *models.MaritimeTradeRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	SendTradeOffer(ctx context.Context, in *models.SendTradeOfferRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ConfirmTradeOffer(ctx context.Context, in *models.ConfirmTradeOfferRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	CancelTradeOffer(ctx context.Context, in *models.CancelTradeOfferRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	PlayKnightCard(ctx context.Context, in *models.PlayKnightCardRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	PlayRoadBuildingCard(ctx context.Context, in *models.PlayRoadBuildingCardRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	PlayYearOfPlentyCard(ctx context.Context, in *models.PlayYearOfPlentyCardRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	PlayMonopolyCard(ctx context.Context, in *models.PlayMonopolyCardRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	PlayVictoryPointCard(ctx context.Context, in *models.PlayVictoryPointCardRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type catanClient struct {
	cc grpc.ClientConnInterface
}

func NewCatanClient(cc grpc.ClientConnInterface) CatanClient {
	return &catanClient{cc}
}

func (c *catanClient) FindGamePaginationByLimitByOffset(ctx context.Context, in *models.FindGamePaginationByLimitByOffsetRequest, opts ...grpc.CallOption) (*models.GamePagination, error) {
	out := new(models.GamePagination)
	err := c.cc.Invoke(ctx, "/pb.Catan/FindGamePaginationByLimitByOffset", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catanClient) GetGameDetailByIDByUserID(ctx context.Context, in *models.GetGameDetailByIDByUserIDRequest, opts ...grpc.CallOption) (*models.GameDetail, error) {
	out := new(models.GameDetail)
	err := c.cc.Invoke(ctx, "/pb.Catan/GetGameDetailByIDByUserID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catanClient) CreateGame(ctx context.Context, in *models.CreateGameRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/pb.Catan/CreateGame", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catanClient) JoinGame(ctx context.Context, in *models.JoinGameRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/pb.Catan/JoinGame", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catanClient) StartGame(ctx context.Context, in *models.StartGameRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/pb.Catan/StartGame", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catanClient) BuildSettlementAndRoad(ctx context.Context, in *models.BuildSettlementAndRoadRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/pb.Catan/BuildSettlementAndRoad", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catanClient) RollDices(ctx context.Context, in *models.RollDicesRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/pb.Catan/RollDices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catanClient) DiscardResourceCards(ctx context.Context, in *models.DiscardResourceCardsRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/pb.Catan/DiscardResourceCards", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catanClient) MoveRobber(ctx context.Context, in *models.MoveRobberRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/pb.Catan/MoveRobber", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catanClient) EndTurn(ctx context.Context, in *models.EndTurnRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/pb.Catan/EndTurn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catanClient) BuildSettlement(ctx context.Context, in *models.BuildSettlementRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/pb.Catan/BuildSettlement", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catanClient) BuildRoad(ctx context.Context, in *models.BuildRoadRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/pb.Catan/BuildRoad", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catanClient) UpgradeCity(ctx context.Context, in *models.UpgradeCityRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/pb.Catan/UpgradeCity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catanClient) BuyDevelopmentCard(ctx context.Context, in *models.BuyDevelopmentCardRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/pb.Catan/BuyDevelopmentCard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catanClient) ToggleResourceCards(ctx context.Context, in *models.ToggleResourceCardsRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/pb.Catan/ToggleResourceCards", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catanClient) MaritimeTrade(ctx context.Context, in *models.MaritimeTradeRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/pb.Catan/MaritimeTrade", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catanClient) SendTradeOffer(ctx context.Context, in *models.SendTradeOfferRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/pb.Catan/SendTradeOffer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catanClient) ConfirmTradeOffer(ctx context.Context, in *models.ConfirmTradeOfferRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/pb.Catan/ConfirmTradeOffer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catanClient) CancelTradeOffer(ctx context.Context, in *models.CancelTradeOfferRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/pb.Catan/CancelTradeOffer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catanClient) PlayKnightCard(ctx context.Context, in *models.PlayKnightCardRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/pb.Catan/PlayKnightCard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catanClient) PlayRoadBuildingCard(ctx context.Context, in *models.PlayRoadBuildingCardRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/pb.Catan/PlayRoadBuildingCard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catanClient) PlayYearOfPlentyCard(ctx context.Context, in *models.PlayYearOfPlentyCardRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/pb.Catan/PlayYearOfPlentyCard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catanClient) PlayMonopolyCard(ctx context.Context, in *models.PlayMonopolyCardRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/pb.Catan/PlayMonopolyCard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catanClient) PlayVictoryPointCard(ctx context.Context, in *models.PlayVictoryPointCardRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/pb.Catan/PlayVictoryPointCard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CatanServer is the server API for Catan service.
// All implementations must embed UnimplementedCatanServer
// for forward compatibility
type CatanServer interface {
	FindGamePaginationByLimitByOffset(context.Context, *models.FindGamePaginationByLimitByOffsetRequest) (*models.GamePagination, error)
	GetGameDetailByIDByUserID(context.Context, *models.GetGameDetailByIDByUserIDRequest) (*models.GameDetail, error)
	CreateGame(context.Context, *models.CreateGameRequest) (*emptypb.Empty, error)
	JoinGame(context.Context, *models.JoinGameRequest) (*emptypb.Empty, error)
	StartGame(context.Context, *models.StartGameRequest) (*emptypb.Empty, error)
	BuildSettlementAndRoad(context.Context, *models.BuildSettlementAndRoadRequest) (*emptypb.Empty, error)
	RollDices(context.Context, *models.RollDicesRequest) (*emptypb.Empty, error)
	DiscardResourceCards(context.Context, *models.DiscardResourceCardsRequest) (*emptypb.Empty, error)
	MoveRobber(context.Context, *models.MoveRobberRequest) (*emptypb.Empty, error)
	EndTurn(context.Context, *models.EndTurnRequest) (*emptypb.Empty, error)
	BuildSettlement(context.Context, *models.BuildSettlementRequest) (*emptypb.Empty, error)
	BuildRoad(context.Context, *models.BuildRoadRequest) (*emptypb.Empty, error)
	UpgradeCity(context.Context, *models.UpgradeCityRequest) (*emptypb.Empty, error)
	BuyDevelopmentCard(context.Context, *models.BuyDevelopmentCardRequest) (*emptypb.Empty, error)
	ToggleResourceCards(context.Context, *models.ToggleResourceCardsRequest) (*emptypb.Empty, error)
	MaritimeTrade(context.Context, *models.MaritimeTradeRequest) (*emptypb.Empty, error)
	SendTradeOffer(context.Context, *models.SendTradeOfferRequest) (*emptypb.Empty, error)
	ConfirmTradeOffer(context.Context, *models.ConfirmTradeOfferRequest) (*emptypb.Empty, error)
	CancelTradeOffer(context.Context, *models.CancelTradeOfferRequest) (*emptypb.Empty, error)
	PlayKnightCard(context.Context, *models.PlayKnightCardRequest) (*emptypb.Empty, error)
	PlayRoadBuildingCard(context.Context, *models.PlayRoadBuildingCardRequest) (*emptypb.Empty, error)
	PlayYearOfPlentyCard(context.Context, *models.PlayYearOfPlentyCardRequest) (*emptypb.Empty, error)
	PlayMonopolyCard(context.Context, *models.PlayMonopolyCardRequest) (*emptypb.Empty, error)
	PlayVictoryPointCard(context.Context, *models.PlayVictoryPointCardRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedCatanServer()
}

// UnimplementedCatanServer must be embedded to have forward compatible implementations.
type UnimplementedCatanServer struct {
}

func (UnimplementedCatanServer) FindGamePaginationByLimitByOffset(context.Context, *models.FindGamePaginationByLimitByOffsetRequest) (*models.GamePagination, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindGamePaginationByLimitByOffset not implemented")
}
func (UnimplementedCatanServer) GetGameDetailByIDByUserID(context.Context, *models.GetGameDetailByIDByUserIDRequest) (*models.GameDetail, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGameDetailByIDByUserID not implemented")
}
func (UnimplementedCatanServer) CreateGame(context.Context, *models.CreateGameRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGame not implemented")
}
func (UnimplementedCatanServer) JoinGame(context.Context, *models.JoinGameRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinGame not implemented")
}
func (UnimplementedCatanServer) StartGame(context.Context, *models.StartGameRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartGame not implemented")
}
func (UnimplementedCatanServer) BuildSettlementAndRoad(context.Context, *models.BuildSettlementAndRoadRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BuildSettlementAndRoad not implemented")
}
func (UnimplementedCatanServer) RollDices(context.Context, *models.RollDicesRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RollDices not implemented")
}
func (UnimplementedCatanServer) DiscardResourceCards(context.Context, *models.DiscardResourceCardsRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DiscardResourceCards not implemented")
}
func (UnimplementedCatanServer) MoveRobber(context.Context, *models.MoveRobberRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MoveRobber not implemented")
}
func (UnimplementedCatanServer) EndTurn(context.Context, *models.EndTurnRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EndTurn not implemented")
}
func (UnimplementedCatanServer) BuildSettlement(context.Context, *models.BuildSettlementRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BuildSettlement not implemented")
}
func (UnimplementedCatanServer) BuildRoad(context.Context, *models.BuildRoadRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BuildRoad not implemented")
}
func (UnimplementedCatanServer) UpgradeCity(context.Context, *models.UpgradeCityRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpgradeCity not implemented")
}
func (UnimplementedCatanServer) BuyDevelopmentCard(context.Context, *models.BuyDevelopmentCardRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BuyDevelopmentCard not implemented")
}
func (UnimplementedCatanServer) ToggleResourceCards(context.Context, *models.ToggleResourceCardsRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ToggleResourceCards not implemented")
}
func (UnimplementedCatanServer) MaritimeTrade(context.Context, *models.MaritimeTradeRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MaritimeTrade not implemented")
}
func (UnimplementedCatanServer) SendTradeOffer(context.Context, *models.SendTradeOfferRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendTradeOffer not implemented")
}
func (UnimplementedCatanServer) ConfirmTradeOffer(context.Context, *models.ConfirmTradeOfferRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfirmTradeOffer not implemented")
}
func (UnimplementedCatanServer) CancelTradeOffer(context.Context, *models.CancelTradeOfferRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelTradeOffer not implemented")
}
func (UnimplementedCatanServer) PlayKnightCard(context.Context, *models.PlayKnightCardRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PlayKnightCard not implemented")
}
func (UnimplementedCatanServer) PlayRoadBuildingCard(context.Context, *models.PlayRoadBuildingCardRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PlayRoadBuildingCard not implemented")
}
func (UnimplementedCatanServer) PlayYearOfPlentyCard(context.Context, *models.PlayYearOfPlentyCardRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PlayYearOfPlentyCard not implemented")
}
func (UnimplementedCatanServer) PlayMonopolyCard(context.Context, *models.PlayMonopolyCardRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PlayMonopolyCard not implemented")
}
func (UnimplementedCatanServer) PlayVictoryPointCard(context.Context, *models.PlayVictoryPointCardRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PlayVictoryPointCard not implemented")
}
func (UnimplementedCatanServer) mustEmbedUnimplementedCatanServer() {}

// UnsafeCatanServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CatanServer will
// result in compilation errors.
type UnsafeCatanServer interface {
	mustEmbedUnimplementedCatanServer()
}

func RegisterCatanServer(s grpc.ServiceRegistrar, srv CatanServer) {
	s.RegisterService(&Catan_ServiceDesc, srv)
}

func _Catan_FindGamePaginationByLimitByOffset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(models.FindGamePaginationByLimitByOffsetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatanServer).FindGamePaginationByLimitByOffset(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Catan/FindGamePaginationByLimitByOffset",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatanServer).FindGamePaginationByLimitByOffset(ctx, req.(*models.FindGamePaginationByLimitByOffsetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Catan_GetGameDetailByIDByUserID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(models.GetGameDetailByIDByUserIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatanServer).GetGameDetailByIDByUserID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Catan/GetGameDetailByIDByUserID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatanServer).GetGameDetailByIDByUserID(ctx, req.(*models.GetGameDetailByIDByUserIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Catan_CreateGame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(models.CreateGameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatanServer).CreateGame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Catan/CreateGame",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatanServer).CreateGame(ctx, req.(*models.CreateGameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Catan_JoinGame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(models.JoinGameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatanServer).JoinGame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Catan/JoinGame",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatanServer).JoinGame(ctx, req.(*models.JoinGameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Catan_StartGame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(models.StartGameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatanServer).StartGame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Catan/StartGame",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatanServer).StartGame(ctx, req.(*models.StartGameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Catan_BuildSettlementAndRoad_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(models.BuildSettlementAndRoadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatanServer).BuildSettlementAndRoad(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Catan/BuildSettlementAndRoad",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatanServer).BuildSettlementAndRoad(ctx, req.(*models.BuildSettlementAndRoadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Catan_RollDices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(models.RollDicesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatanServer).RollDices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Catan/RollDices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatanServer).RollDices(ctx, req.(*models.RollDicesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Catan_DiscardResourceCards_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(models.DiscardResourceCardsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatanServer).DiscardResourceCards(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Catan/DiscardResourceCards",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatanServer).DiscardResourceCards(ctx, req.(*models.DiscardResourceCardsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Catan_MoveRobber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(models.MoveRobberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatanServer).MoveRobber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Catan/MoveRobber",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatanServer).MoveRobber(ctx, req.(*models.MoveRobberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Catan_EndTurn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(models.EndTurnRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatanServer).EndTurn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Catan/EndTurn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatanServer).EndTurn(ctx, req.(*models.EndTurnRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Catan_BuildSettlement_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(models.BuildSettlementRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatanServer).BuildSettlement(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Catan/BuildSettlement",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatanServer).BuildSettlement(ctx, req.(*models.BuildSettlementRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Catan_BuildRoad_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(models.BuildRoadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatanServer).BuildRoad(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Catan/BuildRoad",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatanServer).BuildRoad(ctx, req.(*models.BuildRoadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Catan_UpgradeCity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(models.UpgradeCityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatanServer).UpgradeCity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Catan/UpgradeCity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatanServer).UpgradeCity(ctx, req.(*models.UpgradeCityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Catan_BuyDevelopmentCard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(models.BuyDevelopmentCardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatanServer).BuyDevelopmentCard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Catan/BuyDevelopmentCard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatanServer).BuyDevelopmentCard(ctx, req.(*models.BuyDevelopmentCardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Catan_ToggleResourceCards_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(models.ToggleResourceCardsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatanServer).ToggleResourceCards(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Catan/ToggleResourceCards",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatanServer).ToggleResourceCards(ctx, req.(*models.ToggleResourceCardsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Catan_MaritimeTrade_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(models.MaritimeTradeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatanServer).MaritimeTrade(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Catan/MaritimeTrade",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatanServer).MaritimeTrade(ctx, req.(*models.MaritimeTradeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Catan_SendTradeOffer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(models.SendTradeOfferRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatanServer).SendTradeOffer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Catan/SendTradeOffer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatanServer).SendTradeOffer(ctx, req.(*models.SendTradeOfferRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Catan_ConfirmTradeOffer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(models.ConfirmTradeOfferRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatanServer).ConfirmTradeOffer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Catan/ConfirmTradeOffer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatanServer).ConfirmTradeOffer(ctx, req.(*models.ConfirmTradeOfferRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Catan_CancelTradeOffer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(models.CancelTradeOfferRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatanServer).CancelTradeOffer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Catan/CancelTradeOffer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatanServer).CancelTradeOffer(ctx, req.(*models.CancelTradeOfferRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Catan_PlayKnightCard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(models.PlayKnightCardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatanServer).PlayKnightCard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Catan/PlayKnightCard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatanServer).PlayKnightCard(ctx, req.(*models.PlayKnightCardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Catan_PlayRoadBuildingCard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(models.PlayRoadBuildingCardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatanServer).PlayRoadBuildingCard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Catan/PlayRoadBuildingCard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatanServer).PlayRoadBuildingCard(ctx, req.(*models.PlayRoadBuildingCardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Catan_PlayYearOfPlentyCard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(models.PlayYearOfPlentyCardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatanServer).PlayYearOfPlentyCard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Catan/PlayYearOfPlentyCard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatanServer).PlayYearOfPlentyCard(ctx, req.(*models.PlayYearOfPlentyCardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Catan_PlayMonopolyCard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(models.PlayMonopolyCardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatanServer).PlayMonopolyCard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Catan/PlayMonopolyCard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatanServer).PlayMonopolyCard(ctx, req.(*models.PlayMonopolyCardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Catan_PlayVictoryPointCard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(models.PlayVictoryPointCardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatanServer).PlayVictoryPointCard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Catan/PlayVictoryPointCard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatanServer).PlayVictoryPointCard(ctx, req.(*models.PlayVictoryPointCardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Catan_ServiceDesc is the grpc.ServiceDesc for Catan service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Catan_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Catan",
	HandlerType: (*CatanServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindGamePaginationByLimitByOffset",
			Handler:    _Catan_FindGamePaginationByLimitByOffset_Handler,
		},
		{
			MethodName: "GetGameDetailByIDByUserID",
			Handler:    _Catan_GetGameDetailByIDByUserID_Handler,
		},
		{
			MethodName: "CreateGame",
			Handler:    _Catan_CreateGame_Handler,
		},
		{
			MethodName: "JoinGame",
			Handler:    _Catan_JoinGame_Handler,
		},
		{
			MethodName: "StartGame",
			Handler:    _Catan_StartGame_Handler,
		},
		{
			MethodName: "BuildSettlementAndRoad",
			Handler:    _Catan_BuildSettlementAndRoad_Handler,
		},
		{
			MethodName: "RollDices",
			Handler:    _Catan_RollDices_Handler,
		},
		{
			MethodName: "DiscardResourceCards",
			Handler:    _Catan_DiscardResourceCards_Handler,
		},
		{
			MethodName: "MoveRobber",
			Handler:    _Catan_MoveRobber_Handler,
		},
		{
			MethodName: "EndTurn",
			Handler:    _Catan_EndTurn_Handler,
		},
		{
			MethodName: "BuildSettlement",
			Handler:    _Catan_BuildSettlement_Handler,
		},
		{
			MethodName: "BuildRoad",
			Handler:    _Catan_BuildRoad_Handler,
		},
		{
			MethodName: "UpgradeCity",
			Handler:    _Catan_UpgradeCity_Handler,
		},
		{
			MethodName: "BuyDevelopmentCard",
			Handler:    _Catan_BuyDevelopmentCard_Handler,
		},
		{
			MethodName: "ToggleResourceCards",
			Handler:    _Catan_ToggleResourceCards_Handler,
		},
		{
			MethodName: "MaritimeTrade",
			Handler:    _Catan_MaritimeTrade_Handler,
		},
		{
			MethodName: "SendTradeOffer",
			Handler:    _Catan_SendTradeOffer_Handler,
		},
		{
			MethodName: "ConfirmTradeOffer",
			Handler:    _Catan_ConfirmTradeOffer_Handler,
		},
		{
			MethodName: "CancelTradeOffer",
			Handler:    _Catan_CancelTradeOffer_Handler,
		},
		{
			MethodName: "PlayKnightCard",
			Handler:    _Catan_PlayKnightCard_Handler,
		},
		{
			MethodName: "PlayRoadBuildingCard",
			Handler:    _Catan_PlayRoadBuildingCard_Handler,
		},
		{
			MethodName: "PlayYearOfPlentyCard",
			Handler:    _Catan_PlayYearOfPlentyCard_Handler,
		},
		{
			MethodName: "PlayMonopolyCard",
			Handler:    _Catan_PlayMonopolyCard_Handler,
		},
		{
			MethodName: "PlayVictoryPointCard",
			Handler:    _Catan_PlayVictoryPointCard_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protobuf/catan.proto",
}
