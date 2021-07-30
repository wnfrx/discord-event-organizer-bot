package config

import "github.com/gin-gonic/gin"

func (c *Config) InitRouter() (err error) {
	if c.router != nil {
		return nil
	}

	c.router = gin.Default()

	return nil
}
