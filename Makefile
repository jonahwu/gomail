# This how we want to name the binary output
# #BINARY=glustercluster
# # # Builds the project
build:
	#go build nxgk8sinit.go nks.go util.go common.go webserver.go
	go build sm.go
install:
	cp db /root/../db
clean:
	rm sm
	rm nxgk8sinit

