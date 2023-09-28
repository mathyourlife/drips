package drips

import (
	"context"
	"log"
	"net"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"

	"github.com/mathyourlife/drips/proto"
	pb "github.com/mathyourlife/drips/proto"
	"github.com/stretchr/testify/assert"
)

type testHarness struct {
	lis    *bufconn.Listener
	server *grpc.Server
	conn   *grpc.ClientConn
	client pb.DripsServiceClient
}

func newTestHarness() *testHarness {
	th := &testHarness{}
	svc := &Service{}
	server := grpc.NewServer()
	pb.RegisterDripsServiceServer(server, svc)

	th.lis = bufconn.Listen(1024 * 1024)
	go func() {
		if err := server.Serve(th.lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
	th.server = server

	var err error
	th.conn, err = grpc.DialContext(context.Background(), "memdialer",
		grpc.WithContextDialer(th.bufDialer),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatal(err)
	}

	th.client = proto.NewDripsServiceClient(th.conn)

	return th
}

func (th *testHarness) Close() {
	th.conn.Close()
}

func (th *testHarness) bufDialer(context.Context, string) (net.Conn, error) {
	return th.lis.Dial()
}

func TestUser(t *testing.T) {
	th := newTestHarness()

	resp, err := th.client.User(context.Background(), &pb.UserRequest{UserId: 2})
	assert.NoError(t, err)
	assert.Equal(t, int32(2), resp.User.UserId)
}
