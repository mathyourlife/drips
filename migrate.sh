
# sudo docker run -v $(pwd)/migrations:/migrations --network host migrate/migrate -path=/migrations/ -database postgres://localhost:5432/database up 2

sudo docker run -v $(pwd)/db/migrations:/migrations -v $(pwd)/db/data:/db/ --network host migrate/migrate -path=/migrations/ -database sqlite://db/drips.db up 2

# curl -v 'http://localhost:8080/api/exercise_class' -X POST --data-raw '{"name":"New Class","short_name":"nc"}'