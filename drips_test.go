package drips

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"

	"github.com/mathyourlife/drips/proto"
	_ "github.com/proullon/ramsql/driver"
	"github.com/stretchr/testify/assert"

	pb "github.com/mathyourlife/drips/proto"
)

type testHarness struct {
	t      *testing.T
	lis    *bufconn.Listener
	server *grpc.Server
	conn   *grpc.ClientConn
	client pb.DripsServiceClient
	db     *sql.DB
}

func newTestHarness(t *testing.T) (*testHarness, error) {
	th := &testHarness{
		t: t,
	}

	if err := th.setupDB(); err != nil {
		return nil, err
	}

	svc := NewService(th.db)
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

	return th, nil
}

func (th *testHarness) setupDB() error {
	stmts := []string{
		`CREATE TABLE user (user_id BIGSERIAL PRIMARY KEY, display_name TEXT);`,
		`INSERT INTO user (user_id, display_name) VALUES (1, 'Someone');`,
		`INSERT INTO user (user_id, display_name) VALUES (2, 'Anyone');`,
	}

	var err error
	th.db, err = sql.Open("ramsql", th.t.Name())
	if err != nil {
		return fmt.Errorf("sql.Open : Error : %w", err)
	}

	for _, stmt := range stmts {
		_, err = th.db.Exec(stmt)
		if err != nil {
			return fmt.Errorf("sql.Exec: %w", err)
		}
	}
	return nil
}

func (th *testHarness) Close() {
	th.conn.Close()
	th.db.Close()
}

func (th *testHarness) bufDialer(context.Context, string) (net.Conn, error) {
	return th.lis.Dial()
}

func TestUser(t *testing.T) {
	th, err := newTestHarness(t)
	assert.NoError(t, err)
	t.Cleanup(th.Close)

	resp, err := th.client.User(context.Background(), &pb.UserRequest{UserId: 2})
	assert.NoError(t, err)
	assert.Equal(t, int32(2), resp.User.UserId)
	assert.Equal(t, "Anyone", resp.User.DisplayName)
}

func TestGetRoutine(t *testing.T) {
	th, err := newTestHarness(t)
	assert.NoError(t, err)
	t.Cleanup(th.Close)

	resp, err := th.client.Routine(context.Background(), &pb.RoutineRequest{Name: "Caroline Girvan - Iron Series"})
	assert.NoError(t, err)
	assert.Equal(t, "Caroline Girvan - Iron Series", resp.Routine.Name)

	want := `Caroline Girvan - Iron Series
#1

1: squat (suitcase) for 60 seconds then rest for 30 seconds
2: squat (suitcase) for 60 seconds then rest for 30 seconds
3: lunge (static,left) for 60 seconds then rest for 30 seconds
4: lunge (static,right) for 60 seconds then rest for 30 seconds
5: lunge (static,left) for 60 seconds then rest for 30 seconds
6: lunge (static,right) for 60 seconds then rest for 30 seconds
7: romanian dead lift () for 60 seconds then rest for 30 seconds
8: romanian dead lift () for 60 seconds then rest for 30 seconds
9: romanian dead lift () for 60 seconds then rest for 30 seconds
10: lunge (rear step,left) for 60 seconds then rest for 30 seconds
11: lunge (rear step,right) for 60 seconds then rest for 30 seconds
12: lunge (rear step,left) for 60 seconds then rest for 30 seconds
13: lunge (rear step,right) for 60 seconds then rest for 30 seconds
14: squat (goblet,pause at bottom) for 60 seconds then rest for 30 seconds
15: squat (goblet,pause at bottom) for 60 seconds then rest for 30 seconds
16: lunge (lateral,left) for 60 seconds then rest for 30 seconds
17: lunge (lateral,right) for 60 seconds then rest for 30 seconds
18: lunge (lateral,left) for 60 seconds then rest for 30 seconds
19: lunge (lateral,right) for 60 seconds then rest for 30 seconds
20: squat (goblet,1/2 rep) for 60 seconds
21: squat (goblet) for 60 seconds
22: squat (goblet,1/2 rep) for 60 seconds
23: squat (goblet) for 60 seconds`

	assert.Equal(t, want, PrintRoutine(resp.Routine))
}
