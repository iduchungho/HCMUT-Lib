package config

import (
    "context"
    "github.com/cloudinary/cloudinary-go/v2"
    // "github.com/cloudinary/cloudinary-go/v2/api/admin"
    // "github.com/cloudinary/cloudinary-go/v2/api/uploader"
    // "log"
)

type Cloud struct {
	Cld *cloudinary.Cloudinary
	Ctx context.Context
}

func credentials() (*cloudinary.Cloudinary, context.Context) {
    // Add your Cloudinary credentials, set configuration parameter 
    // Secure=true to return "https" URLs, and create a context
    //===================
    cld, _ := cloudinary.New()
    cld.Config.URL.Secure = true
    ctx := context.Background()
    return cld, ctx
}

func (c *Cloud) Connect(){
	c.Cld, c.Ctx = credentials()
}
