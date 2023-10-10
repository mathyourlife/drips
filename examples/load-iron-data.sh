#!/bin/bash

set -e

go run ./cmd/client exercise-class create --name squat
go run ./cmd/client modifier create --name suitcase
go run ./cmd/client exercise create --exercise-class squat --modifier suitcase
go run ./cmd/client exercise-class create --name lunge
go run ./cmd/client modifier create --name static
go run ./cmd/client modifier create --name left
go run ./cmd/client modifier create --name right
go run ./cmd/client exercise create --exercise-class lunge --modifier static --modifier left
go run ./cmd/client exercise create --exercise-class lunge --modifier static --modifier right
go run ./cmd/client exercise-class create --name 'romanian dead lift'
go run ./cmd/client exercise create --exercise-class 'romanian dead lift'
go run ./cmd/client modifier create --name 'rear step'
go run ./cmd/client exercise create --exercise-class lunge --modifier 'rear step' --modifier left
go run ./cmd/client exercise create --exercise-class lunge --modifier 'rear step' --modifier right
go run ./cmd/client modifier create --name 'goblet'
go run ./cmd/client modifier create --name 'pause at bottom'
go run ./cmd/client exercise create --exercise-class squat --modifier goblet --modifier 'pause at bottom'
go run ./cmd/client modifier create --name 'lateral'
go run ./cmd/client exercise create --exercise-class lunge --modifier lateral --modifier left
go run ./cmd/client exercise create --exercise-class lunge --modifier lateral --modifier right
go run ./cmd/client modifier create --name '1/2 rep'
go run ./cmd/client exercise create --exercise-class squat --modifier goblet --modifier '1/2 rep'
go run ./cmd/client exercise create --exercise-class squat --modifier goblet
# go run ./cmd/client routine create --name 'iron day 1' --source 'https://www.youtube.com/playlist\?list\=PLhu1QCKrfgPWmStsg7imo5EQ0zmkxymJ2' \
#   --sequence 1 --exercise-id 1 --exercise-id 2 --exercise-id 3 --exercise-id 4 --exercise-id 5 --exercise-id 6 --exercise-id 7  \
#   --exercise-id 8 --exercise-id 9 --exercise-id 10 --exercise-id 11 --exercise-id 12 --exercise-id 13 --exercise-id 14 --exercise-id 15  \
#   --exercise-id 16 --exercise-id 17 --exercise-id 18 --exercise-id 19 --exercise-id 20 --exercise-id 21 --exercise-id 22 --exercise-id 23

go run ./cmd/client exercise-class create --name chest-press
go run ./cmd/client exercise create --exercise-class chest-press
go run ./cmd/client exercise-class create --name fly
go run ./cmd/client modifier create --name chest
go run ./cmd/client exercise create --exercise-class fly --modifier chest
go run ./cmd/client exercise-class create --name push-up
go run ./cmd/client modifier create --name 'renegade row'
go run ./cmd/client exercise create --exercise-class push-up --modifier 'renegade row'
go run ./cmd/client exercise-class create --name pullover
go run ./cmd/client exercise create --exercise-class pullover
go run ./cmd/client exercise-class create --name 'shoulder press'
go run ./cmd/client exercise create --exercise-class 'shoulder press'
go run ./cmd/client modifier create --name 'rear delt'
go run ./cmd/client exercise create --exercise-class fly --modifier 'rear delt'
go run ./cmd/client exercise-class create --name 'arm raise'
go run ./cmd/client exercise create --exercise-class 'arm raise' --modifier lateral
go run ./cmd/client modifier create --name 'frontal'
go run ./cmd/client modifier create --name 'alternating'
go run ./cmd/client exercise create --exercise-class 'arm raise' --modifier frontal --modifier alternating
go run ./cmd/client exercise create --exercise-class 'arm raise' --modifier frontal

# go run ./cmd/client routine create --name 'iron day 2' --source 'https://www.youtube.com/playlist\?list\=PLhu1QCKrfgPWmStsg7imo5EQ0zmkxymJ2' \
#   --sequence 2  --exercise-id 24 --exercise-id 25 --exercise-id 26 --exercise-id 27 --exercise-id 28 --exercise-id 29 --exercise-id 30 \
#   --exercise-id 31 --exercise-id 32 --exercise-id 33 --exercise-id 34 --exercise-id 35 --exercise-id 36 --exercise-id 37 --exercise-id 38 \
#   --exercise-id 39 --exercise-id 40 --exercise-id 41 --exercise-id 42 --exercise-id 43 --exercise-id 44