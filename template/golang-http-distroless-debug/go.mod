module handler

go 1.13

replace handler/function => ./function

require (
	github.com/openfaas-incubator/go-function-sdk v0.0.0-20191017092257-70701da50a91
	github.com/sirupsen/logrus v1.4.2
	handler/function v0.0.0-00010101000000-000000000000
)
