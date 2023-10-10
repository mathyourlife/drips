#!/bin/bash

go run ./cmd/client exercise-class create --name squat
go run ./cmd/client modifier create --name suitcase
go run ./cmd/client exercise create --sequence 1 --exercise-class squat --modifier suitcase --duration 60s --rest 30s
go run ./cmd/client exercise create --sequence 2 --exercise-class squat --modifier suitcase --duration 60s --rest 30s
go run ./cmd/client exercise-class create --name lunge
go run ./cmd/client modifier create --name static
go run ./cmd/client modifier create --name left
go run ./cmd/client modifier create --name right
go run ./cmd/client exercise create --sequence 3 --exercise-class lunge --modifier static --modifier left --duration 60s --rest 30s
go run ./cmd/client exercise create --sequence 4 --exercise-class lunge --modifier static --modifier right --duration 60s --rest 30s
go run ./cmd/client exercise create --sequence 5 --exercise-class lunge --modifier static --modifier left --duration 60s --rest 30s
go run ./cmd/client exercise create --sequence 6 --exercise-class lunge --modifier static --modifier right --duration 60s --rest 30s
go run ./cmd/client exercise-class create --name 'romanian dead lift'
go run ./cmd/client exercise create --sequence 7 --exercise-class 'romanian dead lift' --duration 60s --rest 30s
go run ./cmd/client exercise create --sequence 8 --exercise-class 'romanian dead lift' --duration 60s --rest 30s
go run ./cmd/client exercise create --sequence 9 --exercise-class 'romanian dead lift' --duration 60s --rest 30s
go run ./cmd/client modifier create --name 'rear step'
go run ./cmd/client exercise create --sequence 10 --exercise-class lunge --modifier 'rear step' --modifier left --duration 60s --rest 30s
go run ./cmd/client exercise create --sequence 11 --exercise-class lunge --modifier 'rear step' --modifier right --duration 60s --rest 30s
go run ./cmd/client exercise create --sequence 12 --exercise-class lunge --modifier 'rear step' --modifier left --duration 60s --rest 30s
go run ./cmd/client exercise create --sequence 13 --exercise-class lunge --modifier 'rear step' --modifier right --duration 60s --rest 30s
go run ./cmd/client modifier create --name 'goblet'
go run ./cmd/client modifier create --name 'pause at bottom'
go run ./cmd/client exercise create --sequence 14 --exercise-class squat --modifier goblet --modifier 'pause at bottom' --duration 60s --rest 30s
go run ./cmd/client exercise create --sequence 15 --exercise-class squat --modifier goblet --modifier 'pause at bottom' --duration 60s --rest 30s
go run ./cmd/client modifier create --name 'lateral'
go run ./cmd/client exercise create --sequence 16 --exercise-class lunge --modifier lateral --modifier left --duration 60s --rest 30s
go run ./cmd/client exercise create --sequence 17 --exercise-class lunge --modifier lateral --modifier right --duration 60s --rest 30s
go run ./cmd/client exercise create --sequence 18 --exercise-class lunge --modifier lateral --modifier left --duration 60s --rest 30s
go run ./cmd/client exercise create --sequence 19 --exercise-class lunge --modifier lateral --modifier right --duration 60s --rest 30s
go run ./cmd/client modifier create --name '1/2 rep'
go run ./cmd/client exercise create --sequence 20 --exercise-class squat --modifier goblet --modifier '1/2 rep' --duration 60s
go run ./cmd/client exercise create --sequence 21 --exercise-class squat --modifier goblet --duration 60s
go run ./cmd/client exercise create --sequence 22 --exercise-class squat --modifier goblet --modifier '1/2 rep' --duration 60s
go run ./cmd/client exercise create --sequence 23 --exercise-class squat --modifier goblet --duration 60s
go run ./cmd/client routine create --name 'iron day 1' --source 'https://www.youtube.com/playlist\?list\=PLhu1QCKrfgPWmStsg7imo5EQ0zmkxymJ2' \
  --sequence 1 --exercise-id 1 --exercise-id 2 --exercise-id 3 --exercise-id 4 --exercise-id 5 --exercise-id 6 --exercise-id 7  \
  --exercise-id 8 --exercise-id 9 --exercise-id 10 --exercise-id 11 --exercise-id 12 --exercise-id 13 --exercise-id 14 --exercise-id 15  \
  --exercise-id 16 --exercise-id 17 --exercise-id 18 --exercise-id 19 --exercise-id 20 --exercise-id 21 --exercise-id 22 --exercise-id 23

go run ./cmd/client exercise-class create --name chest-press
go run ./cmd/client exercise create --sequence 1 --exercise-class chest-press --duration 60s --rest 30s
go run ./cmd/client exercise create --sequence 2 --exercise-class chest-press --duration 60s --rest 30s
go run ./cmd/client exercise create --sequence 3 --exercise-class chest-press --duration 60s --rest 30s
go run ./cmd/client exercise-class create --name fly
go run ./cmd/client modifier create --name chest
go run ./cmd/client exercise create --sequence 4 --exercise-class fly --modifier chest --duration 60s --rest 30s
go run ./cmd/client exercise create --sequence 5 --exercise-class fly --modifier chest --duration 60s --rest 30s
go run ./cmd/client exercise create --sequence 6 --exercise-class fly --modifier chest --duration 60s --rest 30s
go run ./cmd/client exercise-class create --name push-up
go run ./cmd/client modifier create --name 'renegade row'
go run ./cmd/client exercise create --sequence 7 --exercise-class push-up --modifier 'renegade row' --duration 60s --rest 30s
go run ./cmd/client exercise create --sequence 8 --exercise-class push-up --modifier 'renegade row' --duration 60s --rest 30s
go run ./cmd/client exercise create --sequence 9 --exercise-class push-up --modifier 'renegade row' --duration 60s --rest 30s
go run ./cmd/client exercise-class create --name pullover
go run ./cmd/client exercise create --sequence 10 --exercise-class pullover --duration 60s --rest 30s
go run ./cmd/client exercise create --sequence 11 --exercise-class pullover --duration 60s --rest 30s
go run ./cmd/client exercise create --sequence 12 --exercise-class pullover --duration 60s --rest 30s
go run ./cmd/client exercise-class create --name 'shoulder press'
go run ./cmd/client exercise create --sequence 13 --exercise-class 'shoulder press' --duration 60s --rest 30s
go run ./cmd/client exercise create --sequence 14 --exercise-class 'shoulder press' --duration 60s --rest 30s
go run ./cmd/client exercise create --sequence 15 --exercise-class 'shoulder press' --duration 60s --rest 30s
go run ./cmd/client modifier create --name 'rear delt'
go run ./cmd/client exercise create --sequence 16 --exercise-class fly --modifier 'rear delt' --duration 60s --rest 30s
go run ./cmd/client exercise-class create --name 'arm raise'

go run ./cmd/client exercise create --sequence 17 --exercise-class 'arm raise' --modifier lateral --duration 60s --rest 30s
go run ./cmd/client exercise create --sequence 18 --exercise-class fly --modifier 'rear delt' --duration 60s --rest 30s
go run ./cmd/client exercise create --sequence 19 --exercise-class 'arm raise' --modifier lateral --duration 60s --rest 30s
go run ./cmd/client modifier create --name 'frontal'
go run ./cmd/client modifier create --name 'alternating'
go run ./cmd/client exercise create --sequence 20 --exercise-class 'arm raise' --modifier frontal --modifier alternating --duration 45s
go run ./cmd/client exercise create --sequence 20 --exercise-class 'arm raise' --modifier frontal --duration 45s

go run ./cmd/client routine create --name 'iron day 2' --source 'https://www.youtube.com/playlist\?list\=PLhu1QCKrfgPWmStsg7imo5EQ0zmkxymJ2' \
  --sequence 2  --exercise-id 24 --exercise-id 25 --exercise-id 26 --exercise-id 27 --exercise-id 28 --exercise-id 29 --exercise-id 30 \
  --exercise-id 31 --exercise-id 32 --exercise-id 33 --exercise-id 34 --exercise-id 35 --exercise-id 36 --exercise-id 37 --exercise-id 38 \
  --exercise-id 39 --exercise-id 40 --exercise-id 41 --exercise-id 42 --exercise-id 43 --exercise-id 44