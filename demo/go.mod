module github.com/ASouwn/PKI/demo

go 1.23.0

require (
    github.com/ASouwn/PKI/utils v0.0.0
)

replace (
    github.com/ASouwn/PKI/utils => ../utils
)