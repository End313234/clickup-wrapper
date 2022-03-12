package helpers

import "github.com/joho/godotenv"

func GetEnv(filename string) map[string]string {
	envMap, _ := godotenv.Read(filename)
	return envMap
}
