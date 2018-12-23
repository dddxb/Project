.PHONY: build
build:
	go build -ldflags "-s -w" -gcflags="-trimpath ${GOPATH}" -v -i -o ./bin/project ./*.go
	upx -9 ./bin/seckill

.PHONY: run
run:
	go build -ldflags "-s -w" -gcflags="-trimpath ${GOPATH}" -v -i -o ./bin/project ./*.go
	./bin/project

.PHONY: mac
mac:
	mkdir -p ./dist/project
	rm -R ./dist/project
	go build -ldflags "-s -w" -gcflags="-trimpath ${GOPATH}" -v -i -o ./dist/mac/bin/project ./*.go
	upx -9 ./dist/mac/bin/project

.PHONY: linux
linux:
	mkdir -p ./dist/project
	rm -R ./dist/project
	docker run --rm \
		-e "GOPATH=/goroot" \
		-v "/Volumes/Document/workspace/GoLang":/goroot \
		-v "/Volumes/Document/workspace/mokai/simwallet":/simwallet \
		-w "/goroot/src/mokai/simwallet" \
		golang:1.9.2 \
		bash -c "go build -i -o ./dist/linux/bin/project" -v

.PHONY: win
win:
	mkdir ./dist/project
	rm -R ./dist/project
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -gcflags="-trimpath ${GOPATH}" -v -i -o ./dist/win/bin/seckill.exe ./*.go
	# upx -9 ./dist/win/bin/seckill
	mkdir -p ./dist/win/webui/frontend/dist
	cp -R ./webui/frontend/dist ./dist/win/webui/frontend
	mkdir -p ./dist/win/data
	cp -R ./data ./dist/win

.PHONY: mkOutside
mkOutside:
	docker run --rm \
		-e "GOPATH=/goroot" \
		-v "/Volumes/Document/workspace/GoLang":/goroot \
		-w "/goroot/src/mokai/simwallet" \
		golang:1.11 \
		bash -c "go build -i -o ./dist/linux/bin/simwallet ./*.go" -v

  # cp ./config.yaml ./dist/linux/bin/config.yaml
	ssh -p49175 -i /Volumes/Document/workspace/startup/MoKai/opts/mkserver root@219.239.243.149 /bin/systemctl stop simwallet.service
	scp -P 49175 -r -C -i /Volumes/Document/workspace/startup/MoKai/opts/mkserver /Volumes/Document/workspace/GoLang/src/mokai/simwallet/dist/linux root@219.239.243.149:/data/projects/simwallet
	ssh -p49175 -i /Volumes/Document/workspace/startup/MoKai/opts/mkserver root@219.239.243.149 /bin/systemctl start simwallet.service

.PHONY: cuc
cuc:
	docker run --rm \
		-e "GOPATH=/goroot" \
		-v "/Volumes/Document/workspace/GoLang":/goroot \
		-w "/goroot/src/mokai/simwallet" \
		golang:1.11 \
		bash -c "go build -i -o ./dist/linux/bin/simwallet ./*.go" -v

	# ssh -p48175 tecom@220.194.53.176 /bin/systemctl stop wallet.service
	scp -P 48175 -r -C /Volumes/Document/workspace/GoLang/src/mokai/simwallet/dist/linux tecom@220.194.53.176:/APP
	# ssh -p48175 tecom@220.194.53.176 /bin/systemctl start wallet.service
