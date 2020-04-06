package grpc

import (
	"context"
	"net"

	pbgameengine "github.com/Chans321/m-apis/m-game-engine/v1"
	"github.com/Chans321/m-game-engine/internal/server/logic"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

type Grpc struct {
	address string
	srv     *grpc.Server
}

func NewServer(address string) *Grpc {
	return &Grpc{
		address: address,
	}
}

func (g *Grpc) GetSize(ctx context.Context, input *pbgameengine.GetSizeRequest) (*pbgameengine.GetSizeResponse, error) {
	log.Info().Msg("GetSize in m-game-engine called")
	return &pbgameengine.GetSizeResponse{
		Size: logic.GetSize(),
	}, nil
}

func (g *Grpc) SetScore(ctx context.Context, input *pbgameengine.SetScoreRequest) (*pbgameengine.SetScoreResponse, error) {
	log.Info().Msg("SetScore in m-game-engine called")
	return &pbgameengine.SetScoreResponse{
		Set: logic.SetScore(input.Score),
	}, nil
}

func (g *Grpc) ListenAndServe() error {
	lis, err := net.Listen("tcp", g.address)
	if err != nil {
		return errors.Wrap(err, "failed to open tcp port in m-highscore microoservice")
	}
	g.srv = grpc.NewServer()

	pbgameengine.RegisterGameEngineServer(g.srv, g)
	log.Info().Str("address", g.address).Msg("starting grpc service for m-game-engine microservice")

	err = g.srv.Serve(lis)

	if err != nil {
		return errors.Wrap(err, "failed to start grpc server for m-game-engine micoservice")
	}

	return nil

}
