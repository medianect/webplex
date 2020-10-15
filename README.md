# Usage

Run and browse.

# Contributing

Every source file should mention the copyright holder and the granted license
in top of the file comments, e.g.

```
# Copyright © 2018 Xvdfe <xvdfe@medianect.com>
# SPDX-License-Identifier: BSD-2-Clause
```

## Go modules

This project uses vendored Go modules. Not sure if this way of managing Go
dependancies is here to stay or not but that's the choice I made.

For a developper I so recommend:

### Adding a dependancy

In one of the file where you intend to use the dependancy import it:

```
import (
	"a/lot/of"
	"various"
	"modules/v2"
	"gitprovider.org/author/module-to-add"
)
```

Then run `go mod vendor` which should download the module from
`gitprovider.org` into the vendor directory and register it accordingly into
`go.{sum,mod}`.

