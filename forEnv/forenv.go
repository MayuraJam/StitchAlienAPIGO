package forenv

import "os"

func EnvVatiable(key string) string {
	os.Setenv(key, "root:MayuSQL3310@tcp(127.0.0.1:3306)/stitchdb")
	return os.Getenv(key)
}

// Db, err = sql.Open("mysql", "root:MayuSQL3310@tcp(127.0.0.1:3306)/stitchdb")
