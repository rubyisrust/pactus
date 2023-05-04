package http

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/pactus-project/pactus/consensus"
	"github.com/pactus-project/pactus/crypto"
	"github.com/pactus-project/pactus/crypto/bls"
	"github.com/pactus-project/pactus/state"
	"github.com/pactus-project/pactus/sync"
	"github.com/pactus-project/pactus/www/grpc"
	"github.com/stretchr/testify/assert"
)

var tMockState *state.MockState
var tMockSync *sync.MockSync
var tMockConsMgr consensus.Manager
var tGRPCServer *grpc.Server
var tHTTPServer *Server

func setup(t *testing.T) {
	if tHTTPServer != nil {
		return
	}

	tMockState = state.MockingState()
	tMockSync = sync.MockingSync()
	tMockConsMgr, _ = consensus.MockingManager([]crypto.Signer{
		bls.GenerateTestSigner(), bls.GenerateTestSigner(),
	})

	grpcConf := &grpc.Config{
		Enable: true,
		Listen: "[::]:0",
	}
	httpConf := &Config{
		Enable: true,
		Listen: "[::]:0",
	}

	tGRPCServer = grpc.NewServer(grpcConf, tMockState, tMockSync, tMockConsMgr)
	assert.NoError(t, tGRPCServer.StartServer())

	tHTTPServer = NewServer(httpConf)
	assert.NoError(t, tHTTPServer.StartServer(tGRPCServer.Address()))
}

func TestRootHandler(t *testing.T) {
	setup(t)

	w := httptest.NewRecorder()
	r := new(http.Request)
	tHTTPServer.RootHandler(w, r)
	assert.Equal(t, w.Code, 200)
	fmt.Println(w.Body)
}
