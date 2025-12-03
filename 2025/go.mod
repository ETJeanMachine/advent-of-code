module advent-of-code/main

go 1.25.5

require (
	advent-of-code/solutions v0.0.0-00010101000000-000000000000
	github.com/joho/godotenv v1.5.1
)

require github.com/deckarep/golang-set/v2 v2.8.0 // indirect

replace advent-of-code/solutions => ./solutions
