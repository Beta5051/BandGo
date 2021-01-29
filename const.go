package BandGo

const (
	baseUrl = "https://openapi.band.us"
	getProfileUrl = baseUrl + "/v2/profile"
	getBandsUrl = baseUrl + "/v2.1/bands"
	getPostsUrl = baseUrl + "/v2/band/posts"
	getPostUrl = baseUrl + "/v2.1/band/post"
	createPost = baseUrl + "/v2.2/band/post/create"
	removePost = baseUrl + "/v2/band/post/remove"
	getCommentUrl = baseUrl + "/v2/band/post/comments"
	createComment = baseUrl + "/v2/band/post/comment/create"
	removeComment = baseUrl + "/v2/band/post/comment/remove"
	getPermissionsUrl = baseUrl + "/v2/band/permissions"
	getAlbumsUrl = baseUrl + "/v2/band/albums"
	getPhotosUrl = baseUrl + "/v2/band/album/photos"
)