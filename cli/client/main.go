package main

import (
	"context"
	"flag"
	"time"

	pbgameengine "github.com/Chans321/m-apis/m-game-engine/v1"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

func main() {
	var addressPtr = flag.String("address", "localhost:60051", "address to connect")
	flag.Parse()
	timeoutCtx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	con, err := grpc.Dial(*addressPtr, grpc.WithInsecure())
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to dial establish connection with m-game-engine server")
	}
	c := pbgameengine.NewGameEngineClient(con)
	if c == nil {
		log.Info().Msg("Client Nil")
	}
	r, err := c.GetSize(timeoutCtx, &pbgameengine.GetSizeRequest{})
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to fetch size from m-game-engine server")
	}
	if r != nil {
		log.Info().Interface("size", r.GetSize()).Msg("size from m-game-engine microservice")
	} else {
		log.Fatal().Err(err).Msg("Failed to fetch size from m-game-engine server")
	}
	defer func() {
		err := con.Close()
		if err != nil {
			log.Info().Msg("Failed to close connection")
		}
	}()
}
