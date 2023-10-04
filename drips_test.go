package drips

import (
	"context"
	"fmt"
	"log"
	"net"
	"testing"

	"github.com/mathyourlife/drips/model"
	"github.com/mathyourlife/drips/proto"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	pb "github.com/mathyourlife/drips/proto"
)

type testHarness struct {
	t      *testing.T
	lis    *bufconn.Listener
	server *grpc.Server
	conn   *grpc.ClientConn
	client pb.DripsServiceClient
	db     *gorm.DB
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

	db, err := gorm.Open(sqlite.Open(fmt.Sprintf("file:memdb%s?mode=memory&cache=shared", th.t.Name())), &gorm.Config{})
	if err != nil {
		return err
	}

	// Migrate the schema
	db.Debug().AutoMigrate(
		&model.User{},
		&model.Modifier{},
		&model.ExerciseClass{},
		&model.Exercise{},
		&model.Routine{},
	)

	th.db = db
	return nil
}

func (th *testHarness) Close() {
	th.conn.Close()
	db, err := th.db.DB()
	if err == nil { // If err == nil instead of != nil
		db.Close()
	}
}

func (th *testHarness) bufDialer(context.Context, string) (net.Conn, error) {
	return th.lis.Dial()
}

func TestUser(t *testing.T) {
	th, err := newTestHarness(t)
	assert.NoError(t, err)
	t.Cleanup(th.Close)

	th.db.Create(&model.User{DisplayName: "person 1"})
	th.db.Create(&model.User{DisplayName: "person 2"})
	u := model.User{DisplayName: "person 3"}
	th.db.Create(&u)

	resp, err := th.client.User(context.Background(), &pb.UserRequest{UserId: int32(u.ID)})
	assert.NoError(t, err)
	assert.Equal(t, int32(u.ID), resp.User.UserId)
	assert.Equal(t, "person 3", resp.User.DisplayName)
}

func TestRoutine(t *testing.T) {
	th, err := newTestHarness(t)
	assert.NoError(t, err)
	t.Cleanup(th.Close)

	resp, err := th.client.Routine(context.Background(), &pb.RoutineRequest{RoutineId: 4})
	assert.NoError(t, err)
	return
	assert.Equal(t, "my workout", resp.Routine.Name)

	want := `my workout
#5

1: lunge (right,rear step) for 60 seconds then rest for 30 seconds
2: romanian dead lift (staggered) for 60 seconds`

	assert.Equal(t, want, PrintRoutine(resp.Routine))
}

func TestRoutines(t *testing.T) {
	th, err := newTestHarness(t)
	assert.NoError(t, err)
	t.Cleanup(th.Close)

	resp, err := th.client.Routines(context.Background(), &pb.RoutinesRequest{Name: "work"})
	assert.NoError(t, err)
	return
	assert.Equal(t, 1, len(resp.Routines))
	assert.Equal(t, "my workout", resp.Routines[0].Name)

	resp, err = th.client.Routines(context.Background(), &pb.RoutinesRequest{})
	assert.NoError(t, err)
	assert.Equal(t, 2, len(resp.Routines))
}
