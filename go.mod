module github.com/c-bata/kube-prompt

require (
	github.com/c-bata/go-prompt v0.2.5
	github.com/googleapis/gnostic v0.2.0 // indirect
	k8s.io/api v0.17.0
	k8s.io/apimachinery v0.17.0
	k8s.io/client-go v0.17.0
)

replace golang.org/x/crypto => golang.org/x/crypto v0.0.0-20201216223049-8b5274cf687f

go 1.15
