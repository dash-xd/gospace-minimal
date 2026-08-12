module github.com/dash-xd/gospace-minimal

go 1.23

require gospace.invalid/router v0.0.0-00010101000000-000000000000

require (
	github.com/dash-xd/github-device-auth v0.0.0-20260812043759-c7f083cc5ad9 // indirect
	github.com/go-chi/chi/v5 v5.3.1 // indirect
)

replace gospace.invalid/router => github.com/dash-xd/github-device-auth/router v0.0.0-20260812043840-c4d00f5e2ded
