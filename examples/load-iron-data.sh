#!/bin/bash

go run ./cmd/client exercise-class create --name squat
go run ./cmd/client modifier create --name suitcase
go run ./cmd/client exercise create --sequence 1 --exercise-class squat --modifier-id 1 --duration 60s --rest 30s
go run ./cmd/client exercise create --sequence 2 --exercise-class squat --modifier-id 1 --duration 60s --rest 30s
go run ./cmd/client exercise-class create --name lunge
go run ./cmd/client modifier create --name static
go run ./cmd/client modifier create --name left
go run ./cmd/client modifier create --name right
go run ./cmd/client exercise create --sequence 3 --exercise-class lunge --modifier-id 2 --modifier-id 3 --duration 60s --rest 30s
go run ./cmd/client exercise create --sequence 4 --exercise-class lunge --modifier-id 2 --modifier-id 4 --duration 60s --rest 30s
go run ./cmd/client exercise create --sequence 5 --exercise-class lunge --modifier-id 2 --modifier-id 3 --duration 60s --rest 30s
go run ./cmd/client exercise create --sequence 6 --exercise-class lunge --modifier-id 2 --modifier-id 4 --duration 60s --rest 30s
go run ./cmd/client exercise-class create --name 'romanian dead lift'
go run ./cmd/client exercise create --sequence 7 --exercise-class 'romanian dead lift' --duration 60s --rest 30s
go run ./cmd/client exercise create --sequence 8 --exercise-class 'romanian dead lift' --duration 60s --rest 30s
go run ./cmd/client exercise create --sequence 9 --exercise-class 'romanian dead lift' --duration 60s --rest 30s
go run ./cmd/client modifier create --name 'rear step'
go run ./cmd/client exercise create --sequence 10 --exercise-class lunge --modifier-id 5 --modifier-id 3 --duration 60s --rest 30s
go run ./cmd/client exercise create --sequence 11 --exercise-class lunge --modifier-id 5 --modifier-id 4 --duration 60s --rest 30s
go run ./cmd/client exercise create --sequence 12 --exercise-class lunge --modifier-id 5 --modifier-id 3 --duration 60s --rest 30s
go run ./cmd/client exercise create --sequence 13 --exercise-class lunge --modifier-id 5 --modifier-id 4 --duration 60s --rest 30s
go run ./cmd/client modifier create --name 'goblet'
go run ./cmd/client modifier create --name 'pause at bottom'
go run ./cmd/client exercise create --sequence 14 --exercise-class squat --modifier-id 6 --modifier-id 7 --duration 60s --rest 30s
go run ./cmd/client exercise create --sequence 15 --exercise-class squat --modifier-id 6 --modifier-id 7 --duration 60s --rest 30s
go run ./cmd/client modifier create --name 'lateral'
go run ./cmd/client exercise create --sequence 16 --exercise-class lunge --modifier-id 8 --modifier-id 3 --duration 60s --rest 30s
go run ./cmd/client exercise create --sequence 17 --exercise-class lunge --modifier-id 8 --modifier-id 4 --duration 60s --rest 30s
go run ./cmd/client exercise create --sequence 18 --exercise-class lunge --modifier-id 8 --modifier-id 3 --duration 60s --rest 30s
go run ./cmd/client exercise create --sequence 19 --exercise-class lunge --modifier-id 8 --modifier-id 4 --duration 60s --rest 30s
go run ./cmd/client modifier create --name '1/2 rep'
go run ./cmd/client exercise create --sequence 20 --exercise-class squat --modifier-id 6 --modifier-id 9 --duration 60s
go run ./cmd/client exercise create --sequence 21 --exercise-class squat --modifier-id 6 --duration 60s
go run ./cmd/client exercise create --sequence 22 --exercise-class squat --modifier-id 6 --modifier-id 9 --duration 60s
go run ./cmd/client exercise create --sequence 23 --exercise-class squat --modifier-id 6 --duration 60s
go run ./cmd/client routine create --name iron --source 'https://www.youtube.com/playlist\?list\=PLhu1QCKrfgPWmStsg7imo5EQ0zmkxymJ2' --sequence 1  --exercise-id 1  --exercise-id 2  --exercise-id 3  --exercise-id 4  --exercise-id 5  --exercise-id 6  --exercise-id 7  --exercise-id 8  --exercise-id 9  --exercise-id 10  --exercise-id 11  --exercise-id 12  --exercise-id 13  --exercise-id 14  --exercise-id 15  --exercise-id 16  --exercise-id 17  --exercise-id 18  --exercise-id 19  --exercise-id 20  --exercise-id 21  --exercise-id 22  --exercise-id 23
