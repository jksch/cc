build:
	GOARCH=wasm GOOS=js go build -o courses/web/app.wasm courses/main.go 

pack:
	go install github.com/mjibson/esc
	esc \
	-o="static/static.go" \
	-pkg="static" \
       	-prefix="courses" \
       	courses/web
	
assable: build pack
