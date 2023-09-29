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
		`INSERT INTO user (user_id, display_name) VALUES (1, 'Someone'), (2, 'Anyone');`,
		`CREATE TABLE routine (routine_id BIGSERIAL PRIMARY KEY, name TEXT, source TEXT, sequence INT);`,
		`INSERT INTO routine (routine_id, name, source, sequence) VALUES (4, 'my workout', 'https://localhost', 4);`,
		`CREATE TABLE class (class_id BIGSERIAL PRIMARY KEY, name TEXT, short_name TEXT);`,
		`INSERT INTO class (class_id, name, short_name) VALUES (2, 'lunge', ''), (3, 'romanian dead lift', 'rdl');`,
		`CREATE TABLE exercise (exercise_id BIGSERIAL PRIMARY KEY, sequence INT, class_id INT, duration_seconds INT, rest_seconds INT);`,
		`INSERT INTO exercise (exercise_id, sequence, class_id, duration_seconds, rest_seconds) VALUES (6, 0, 2, 60, 30), (7, 1, 3, 60, 0);`,
		`CREATE TABLE routine_exercise (routine_exercise_id BIGSERIAL PRIMARY KEY, routine_id INT, exercise_id INT);`,
		`INSERT INTO routine_exercise (routine_exercise_id, routine_id, exercise_id) VALUES (11, 4, 6), (12, 4, 7);`,
		`CREATE TABLE modifier (modifier_id BIGSERIAL PRIMARY KEY, name TEXT);`,
		`INSERT INTO modifier (modifier_id, name) VALUES (8, 'right'), (9, 'left'), (10, 'rear step'), (11, 'staggered');`,
		`CREATE TABLE exercise_modifier (exercise_modifier_id BIGSERIAL PRIMARY KEY, exercise_id INT, modifier_id INT);`,
		`INSERT INTO exercise_modifier (exercise_modifier_id, exercise_id, modifier_id) VALUES (5, 6, 8), (6, 6, 10), (7, 7, 11);`,
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

	resp, err := th.client.Routine(context.Background(), &pb.RoutineRequest{RoutineId: 4})
	assert.NoError(t, err)
	assert.Equal(t, "my workout", resp.Routine.Name)

	want := `my workout
#5

1: lunge (right,rear step) for 60 seconds then rest for 30 seconds
2: romanian dead lift (staggered) for 60 seconds`

	assert.Equal(t, want, PrintRoutine(resp.Routine))
}
