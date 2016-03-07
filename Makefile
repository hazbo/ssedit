all: ${GOPATH}/bin/ssedit

${GOPATH}/bin/ssedit: ssedit/ssedit.go
	cd ssedit && go install

clean: ${GOPATH}/bin/ssedit
	rm -f ${GOPATH}/bin/ssedit

.PHONY: clean
