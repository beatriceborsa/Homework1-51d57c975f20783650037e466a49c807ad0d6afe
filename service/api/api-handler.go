package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	rt.router.POST("/login", rt.wrap(rt.doLogin))
	rt.router.PUT("/user/:username/setusername", rt.wrap(rt.setMyUserName))
	rt.router.GET("/user/:username/streamuser", rt.wrap(rt.getMyStream))
	rt.router.GET("/users/:username/userprofile", rt.wrap(rt.getUserProfile))
	rt.router.PUT("/users/:username/photo/:phid", rt.wrap(rt.uploadPhoto))
	rt.router.DELETE("/users/:username/photo/:phid", rt.wrap(rt.deletePhoto))
	rt.router.GET("/users/:username/phuser", rt.wrap(rt.getUserPhotos))
	rt.router.PUT("/users/:username/ban/:banid", rt.wrap(rt.banUser))
	rt.router.DELETE("/users/:username/ban/:banid", rt.wrap(rt.unbanUser))
	rt.router.GET("/users/:username/userban", rt.wrap(rt.getBans))
	rt.router.PUT("/users/:username/follow/:followid", rt.wrap(rt.followUser))
	rt.router.DELETE("/users/:username/follow/:followid", rt.wrap(rt.unfollowUser))
	rt.router.GET("/users/:username/userfollow", rt.wrap(rt.getFollowers))
	rt.router.PUT("/users/:username/photo/:phid/like/:likeid", rt.wrap(rt.likePhoto))
	rt.router.DELETE("/users/:username/photo/:phid/like/:likeid", rt.wrap(rt.unlikePhoto))
	rt.router.GET("/users/:username/photo/:phid/like", rt.wrap(rt.getLikes))
	rt.router.PUT("/users/:username/photo/:phid/comment/:commentid", rt.wrap(rt.commentPhoto))
	rt.router.DELETE("/users/:username/photo/:phid/comment/:commentid", rt.wrap(rt.uncommentPhoto))
	rt.router.GET("/users/:username/photo/:phid/comment", rt.wrap(rt.getComments))
	rt.router.GET("/liveness", rt.liveness)
	return rt.router
}
