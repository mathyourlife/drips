# Drips

Nothing like a little sweat.

## Server

Start the local gRPC server with CRUD operations.

```bash
go run ./cmd/server
```

## Client

Run a cli client

```bash
go run ./cmd/client exercise-class list
```

## Example Usage

Start a server.

```bash
go run ./cmd/server start --db-file exercises.db
```

Create a user.

```bash
go run ./cmd/client user create --display-name someuser
# user_id:1  display_name:"someuser"
go run ./cmd/client user create --display-name otherperson
# user_id:2  display_name:"otherperson"
```

List current users

```bash
go run ./cmd/client user list
# 1 user_id:1  display_name:"someuser"
# 2 user_id:2  display_name:"otherperson"
```

Create exercise modifiers

```bash
go run ./cmd/client modifier create --name 'right'
# name:"right"
go run ./cmd/client modifier create --name 'left'
# name:"left"
go run ./cmd/client modifier create --name 'back step'
# name:"back step"
```

List modifiers

```bash
$ go run ./cmd/client modifier list
# 1 modifier_id:1  name:"right"
# 2 modifier_id:2  name:"left"
# 3 modifier_id:3  name:"back step"
```

Create exercise classes

```bash
$ go run ./cmd/client exercise-class create --name lunge
# exercise_class_id:1  name:"lunge"
```

List exercise classes

```bash
go run ./cmd/client exercise-class list
# 1 exercise_class_id:1  name:"lunge"
```

Create exercises

```bash
go run ./cmd/client exercise create --sequence 5 --exercise-class-id 1 --modifier-id 2 --modifier-id 3 --duration 66s  --rest 45s
# exercise_id:1  sequence:5  class:{exercise_class_id:1  name:"lunge"}  modifiers:{modifier_id:2  name:"left"}  modifiers:{modifier_id:3  name:"back step"}  duration:{seconds:66}  rest:{seconds:45}
go run ./cmd/client exercise create --sequence 6 --exercise-class-id 1 --modifier-id 1 --modifier-id 3 --duration 1s
# exercise_id:2  sequence:6  class:{exercise_class_id:1  name:"lunge"}  modifiers:{modifier_id:1  name:"right"}  modifiers:{modifier_id:3  name:"back step"}  duration:{seconds:1}  rest:{}
```

List exercises

```bash
go run ./cmd/client exercise list
1 exercise_id:1  sequence:5  class:{exercise_class_id:1  name:"lunge"}  modifiers:{modifier_id:2  name:"left"}  modifiers:{modifier_id:3  name:"back step"}  duration:{seconds:66}  rest:{seconds:45}
2 exercise_id:2  sequence:6  class:{exercise_class_id:1  name:"lunge"}  modifiers:{modifier_id:1  name:"right"}  modifiers:{modifier_id:3  name:"back step"}  duration:{seconds:1}  rest:{}
```

Create an exercise routine

```bash
go run ./cmd/client routine create --name myroutine --sequence 7 --exercise-id 1 --exercise-id 2
# routine_id:1  name:"myroutine"  sequence:7  exercises:{exercise_id:1  sequence:5  class:{}  duration:{seconds:66}  rest:{seconds:45}}  exercises:{exercise_id:2  sequence:6  class:{}  duration:{seconds:1}  rest:{}}
```

List routines

```bash
go run ./cmd/client routine list
# 1 routine_id:1  name:"myroutine1"  sequence:7  exercises:{exercise_id:1  sequence:5  class:{}  duration:{seconds:66}  rest:{seconds:45}}  exercises:{exercise_id:2  sequence:6  class:{}  duration:{seconds:1}  rest:{}}
```

## Development

Compile Golang language bindings.

```bash
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    proto/drips.proto
```

## To Do

* Show ID on client create call.
* Log user activity.
* Create subclasses of exercises based on sets of modifiers.
