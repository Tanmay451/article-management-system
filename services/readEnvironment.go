package services

import "os"

// GetAppPORT will return app port on which this web-app is supposed to run
func GetAppPORT() string {
	PORT := os.Getenv("APP_PORT")
	if len(PORT) > 0 {
		return PORT
	}
	return "8000"
}
