// Copyright © 2015-2016 Platina Systems, Inc. All rights reserved.
// Use of this source code is governed by the GPL-2 license described in the
// LICENSE file.

package nodocker

import "github.com/platinasystems/go/internal/test"

var Suite = test.Suite{
	Name: "nodocker",
	Tests: test.Tests{
		&test.Unit{"twohost", twohost},
		&test.Unit{"onerouter", onerouter},
	},
}
