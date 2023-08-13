# http-video-server
Simple video-streaming service


go version
go version go1.20.4 linux/amd64

go mod init github.com/michaelspinks/http-video-server
go mod tidy

make run
make build
make clean
make test

001 - Added /video and sample video
002 - Change / to /video.  Used http.ServeFile to stream video, moved video to /video folder
003 - 