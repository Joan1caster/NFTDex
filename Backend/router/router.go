package routes

// import (
//     "github.com/gin-gonic/gin"
//     "nftPlantform/handlers"
//     "nftPlantform/middleware"
// )

// func SetupRouter(h *handlers.Handler) *gin.Engine {
//     router := gin.Default()

//     // 使用自定义的中间件
//     router.Use(middleware.CORSMiddleware())

//     // 公开路由
//     public := router.Group("/api/v1")
//     {
//         public.POST("/auth", h.AuthenticateUser)
//         public.GET("/nfts", h.GetNFTs)
//         public.GET("/nfts/:id", h.GetNFTByID)
//         public.GET("/orders", h.GetOpenOrders)
//     }

//     // 需要认证的路由
//     authenticated := router.Group("/api/v1")
//     authenticated.Use(middleware.AuthMiddleware())
//     {
//         // NFT 相关路由
//         authenticated.POST("/nfts", h.CreateNFT)
//         authenticated.PUT("/nfts/:id", h.UpdateNFT)
//         authenticated.DELETE("/nfts/:id", h.DeleteNFT)

//         // 订单相关路由
//         authenticated.POST("/orders", h.CreateOrder)
//         authenticated.PUT("/orders/:id", h.UpdateOrder)
//         authenticated.DELETE("/orders/:id", h.CancelOrder)
//         authenticated.POST("/orders/:id/buy", h.BuyNFT)

//         // 用户相关路由
//         authenticated.GET("/user", h.GetUserProfile)
//         authenticated.PUT("/user", h.UpdateUserProfile)
//     }

//     return router
// }