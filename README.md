# Orb Badge [![CircleCI Build Status](https://circleci.com/gh/felicianotech/orb-badge.svg?style=shield)](https://circleci.com/gh/felicianotech/orb-badge) [![GitHub License](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/felicianotech/orb-badge/master/LICENSE)

Orb Badge is a mini Go server designed to provide a JSON endpoint for Shields.io's [custom badges](https://shields.io/#/endpoint).
This allows Orb developers to use a status badge on their repo's readme file.

This software is currently considered a proof of concept.
It may go away, and it may break as the Shields.io service is in beta.


## Using a Badge

You show an `orb version` status badge with the following URL:

`https://img.shields.io/badge/endpoint.svg?url=https://badges.circleci.io/orb/<namespace>/<orb-name>`

where `<namespace>` is your Orb's namespace and `<orb-name>` is the name of the specific Orb.

For example, here's a badge for the CircleCI AWS CLI Orb in Markdown with a link back to the Orbs page:

```
[![CircleCI Orb Version](https://img.shields.io/badge/endpoint.svg?url=https://badges.circleci.io/orb/circleci/aws-cli)](https://circleci.com/orbs/registry/orb/circleci/aws-cli)
```

and here is how it render's: [![CircleCI Orb Version](https://img.shields.io/badge/endpoint.svg?url=https://badges.circleci.io/orb/circleci/aws-cli)](https://circleci.com/orbs/registry/orb/circleci/aws-cli)


## Development

This server is written in Go.
Please use Go v1.11 or later.
It may work with earlier versions but it's untested.

```
go get -u github.com/felicianotech/orb-badge
go run .
```

The server will be available at `http://localhost:1107/`.

## Production

There really isn't a production server right now.
I am running a test/beta at https://badges.circleci.io.

To run your own, place the binary on a server and run it.
If you use port 80, `sudo` will be needed for non-root users.

```
go build ./...
./orb-badge <port>
```

Running the binary in something like tmux will make your life easier as well.


## License

This repository is licensed under the MIT license.
This repo's license can be found [here](./LICENSE).
