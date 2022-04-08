// Copyright 2022 VMware Tanzu Community Edition contributors. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package pinniped_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestPinniped(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Pinniped Addons Templates Suite")
}
