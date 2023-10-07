package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/durationpb"

	pb "github.com/mathyourlife/drips/proto"
)

func main() {
	fmt.Println("starting client")
	conn, err := grpc.Dial("localhost:5050",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewDripsServiceClient(conn)

	app := cli.NewApp()
	app.Name = "drips-cli"
	app.Commands = []cli.Command{
		{
			Name:  "exercise",
			Usage: "exercise",
			Subcommands: []cli.Command{
				{
					Name:  "list",
					Usage: "show the current list of exercises",
					Flags: []cli.Flag{
						cli.IntFlag{Name: "exercise-class-id"},
					},
					Action: func(c *cli.Context) error {
						resp, err := client.Exercises(context.Background(), &pb.ExercisesRequest{
							ExerciseClassId: int32(c.Int("exercise-class-id")),
						})
						if err != nil {
							log.Fatal(err)
						}
						for _, e := range resp.Exercises {
							fmt.Printf("%d %s\n", e.ExerciseId, e.String())
						}

						return nil
					},
				}, {
					Name:  "create",
					Usage: "create an exercise",
					Flags: []cli.Flag{
						cli.IntFlag{Name: "sequence"},
						cli.IntFlag{Name: "exercise-class-id"},
						cli.IntSliceFlag{Name: "modifier-id"},
						cli.DurationFlag{Name: "duration"},
						cli.DurationFlag{Name: "rest"},
						cli.IntFlag{Name: "reps"},
					},
					Action: func(c *cli.Context) error {
						e := &pb.Exercise{
							Sequence: int32(c.Int("sequence")),
							Duration: durationpb.New(c.Duration("duration")),
							Rest:     durationpb.New(c.Duration("rest")),
							Reps:     int32(c.Int("reps")),
						}
						ecResp, err := client.ExerciseClass(context.Background(), &pb.ExerciseClassRequest{ExerciseClassId: int32(c.Int("exercise-class-id"))})
						if err != nil {
							log.Fatal(err)
						}
						e.Class = ecResp.GetExerciseClass()
						for _, mID := range c.IntSlice("modifier-id") {
							mResp, err := client.Modifier(context.Background(), &pb.ModifierRequest{ModiferId: int32(mID)})
							if err != nil {
								log.Fatal(err)
							}
							e.Modifiers = append(e.Modifiers, mResp.GetModifier())
						}
						resp, err := client.ExerciseCreate(context.Background(), &pb.ExerciseCreateRequest{Exercise: e})
						if err != nil {
							log.Fatal(err)
						}
						fmt.Println(resp.Exercise)

						return nil
					},
				}, {
					Name:  "delete",
					Usage: "delete a exercise",
					Flags: []cli.Flag{
						cli.IntFlag{
							Name:  "exercise-id",
							Usage: "ID for the exercise",
						},
					},
					Action: func(c *cli.Context) error {
						_, err := client.ExerciseDelete(context.Background(), &pb.ExerciseDeleteRequest{ExerciseId: int32(c.Int("exercise-id"))})
						if err != nil {
							log.Fatal(err)
						}

						return nil
					},
				},
			},
		}, {
			Name:  "exercise-class",
			Usage: "exercise class",
			Subcommands: []cli.Command{
				{
					Name:  "list",
					Usage: "show the current list of exercise-classes",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "name",
							Usage: "Substring search",
						},
					},
					Action: func(c *cli.Context) error {
						resp, err := client.ExerciseClasses(context.Background(), &pb.ExerciseClassesRequest{Name: c.String("name")})
						if err != nil {
							log.Fatal(err)
						}
						for _, ec := range resp.ExerciseClasses {
							fmt.Printf("%d %s\n", ec.ExerciseClassId, ec.String())
						}

						return nil
					},
				}, {
					Name:  "create",
					Usage: "create an exercise-class",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "name",
							Usage: "Exercise class name",
						},
					},
					Action: func(c *cli.Context) error {
						ec := &pb.ExerciseClass{
							Name: c.String("name"),
						}
						resp, err := client.ExerciseClassCreate(context.Background(), &pb.ExerciseClassCreateRequest{ExerciseClass: ec})
						if err != nil {
							log.Fatal(err)
						}
						fmt.Println(resp.ExerciseClass)

						return nil
					},
				}, {
					Name:  "delete",
					Usage: "delete a exercise-class",
					Flags: []cli.Flag{
						cli.IntFlag{
							Name:  "exercise-class-id",
							Usage: "ID for the exercise-class",
						},
					},
					Action: func(c *cli.Context) error {
						_, err := client.ExerciseClassDelete(context.Background(), &pb.ExerciseClassDeleteRequest{ExerciseClassId: int32(c.Int("exercise-class-id"))})
						if err != nil {
							log.Fatal(err)
						}

						return nil
					},
				},
			},
		}, {
			Name:  "modifier",
			Usage: "exercise modifiers",
			Subcommands: []cli.Command{
				{
					Name:  "list",
					Usage: "show the current list of modifiers",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "name",
							Usage: "Substring search",
						},
					},
					Action: func(c *cli.Context) error {
						resp, err := client.Modifiers(context.Background(), &pb.ModifiersRequest{Name: c.String("name")})
						if err != nil {
							log.Fatal(err)
						}
						for _, m := range resp.Modifiers {
							fmt.Printf("%d %s\n", m.ModifierId, m.String())
						}

						return nil
					},
				}, {
					Name:  "create",
					Usage: "create a modifiers",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "name",
							Usage: "Modifier name",
						},
					},
					Action: func(c *cli.Context) error {
						mod := &pb.Modifier{
							Name: c.String("name"),
						}
						resp, err := client.ModifierCreate(context.Background(), &pb.ModifierCreateRequest{Modifier: mod})
						if err != nil {
							log.Fatal(err)
						}
						fmt.Println(resp.Modifier)

						return nil
					},
				}, {
					Name:  "delete",
					Usage: "delete a modifier",
					Flags: []cli.Flag{
						cli.IntFlag{
							Name:  "modifier-id",
							Usage: "ID for the modifier",
						},
					},
					Action: func(c *cli.Context) error {
						_, err := client.ModifierDelete(context.Background(), &pb.ModifierDeleteRequest{ModifierId: int32(c.Int("modifier-id"))})
						if err != nil {
							log.Fatal(err)
						}

						return nil
					},
				},
			},
		}, {
			Name:  "user",
			Usage: "user accounts",
			Subcommands: []cli.Command{
				{
					Name:  "list",
					Usage: "show the current list of users",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "name",
							Usage: "Substring search",
						},
					},
					Action: func(c *cli.Context) error {
						resp, err := client.Users(context.Background(), &pb.UsersRequest{DisplayName: c.String("name")})
						if err != nil {
							log.Fatal(err)
						}
						for _, u := range resp.Users {
							fmt.Printf("%d %s\n", u.UserId, u.String())
						}

						return nil
					},
				}, {
					Name:  "create",
					Usage: "create a user",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "display-name",
							Usage: "User's display name",
						},
					},
					Action: func(c *cli.Context) error {
						u := &pb.User{
							DisplayName: c.String("display-name"),
						}
						resp, err := client.UserCreate(context.Background(), &pb.UserCreateRequest{User: u})
						if err != nil {
							log.Fatal(err)
						}
						fmt.Println(resp.User)

						return nil
					},
				}, {
					Name:  "delete",
					Usage: "delete a user",
					Flags: []cli.Flag{
						cli.IntFlag{
							Name:  "user-id",
							Usage: "ID for the user",
						},
					},
					Action: func(c *cli.Context) error {
						_, err := client.UserDelete(context.Background(), &pb.UserDeleteRequest{UserId: int32(c.Int("user-id"))})
						if err != nil {
							log.Fatal(err)
						}

						return nil
					},
				},
			},
		}, {
			Name:  "routine",
			Usage: "routine",
			Subcommands: []cli.Command{
				{
					Name:  "list",
					Usage: "show the current list of routines",
					Flags: []cli.Flag{
						cli.StringFlag{Name: "name"},
					},
					Action: func(c *cli.Context) error {
						resp, err := client.Routines(context.Background(), &pb.RoutinesRequest{Name: c.String("name")})
						if err != nil {
							log.Fatal(err)
						}
						for _, r := range resp.Routines {
							fmt.Printf("%d %s\n", r.RoutineId, r.String())
						}

						return nil
					},
				}, {
					Name:  "create",
					Usage: "create a routine",
					Flags: []cli.Flag{
						cli.StringFlag{Name: "name"},
						cli.StringFlag{Name: "source"},
						cli.IntFlag{Name: "sequence"},
						cli.IntSliceFlag{Name: "exercise-id"},
					},
					Action: func(c *cli.Context) error {
						r := &pb.Routine{
							Name:     c.String("name"),
							Source:   c.String("source"),
							Sequence: int32(c.Int("sequence")),
						}
						for _, eID := range c.IntSlice("exercise-id") {
							eResp, err := client.Exercise(context.Background(), &pb.ExerciseRequest{ExerciseId: int32(eID)})
							if err != nil {
								log.Fatal(err)
							}
							r.Exercises = append(r.Exercises, eResp.GetExercise())
						}
						resp, err := client.RoutineCreate(context.Background(), &pb.RoutineCreateRequest{Routine: r})
						if err != nil {
							log.Fatal(err)
						}
						fmt.Println(resp.Routine)

						return nil
					},
				}, {
					Name:  "delete",
					Usage: "delete a routine",
					Flags: []cli.Flag{
						cli.IntFlag{
							Name:  "routine-id",
							Usage: "ID for the routine",
						},
					},
					Action: func(c *cli.Context) error {
						_, err := client.RoutineDelete(context.Background(), &pb.RoutineDeleteRequest{RoutineId: int32(c.Int("routine-id"))})
						if err != nil {
							log.Fatal(err)
						}

						return nil
					},
				},
			},
		},
	}

	err = app.Run(os.Args)
}
