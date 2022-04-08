// Copyright 2022 VMware Tanzu Community Edition contributors. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package pinniped_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	// . "github.com/vmware-tanzu/community-edition/addons/packages/test/matchers"
	"github.com/vmware-tanzu/community-edition/addons/packages/test/pkg/repo"
	"github.com/vmware-tanzu/community-edition/addons/packages/test/pkg/ytt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Pinniped Ytt Templates", func() {
	var (
		packageRoot = filepath.Join(repo.RootDir(), "addons/packages/pinniped/0.12.1/")
		configDir   = filepath.Join(packageRoot, "bundle/config/")
		testDir     = filepath.Join(packageRoot, "test/unit/")
		fixtureDir  = filepath.Join(testDir, "fixtures/")
		valuesDir   = filepath.Join(fixtureDir, "values/")
		expectedDir = filepath.Join(fixtureDir, "expected/")
	)

	ValuesFromFile := func(filename string) string {
		data, err := ioutil.ReadFile(filepath.Join(valuesDir, filename))
		Expect(err).NotTo(HaveOccurred())

		return string(data)
	}

	OutputFromFile := func(filename string) string {
		data, err := ioutil.ReadFile(filepath.Join(expectedDir, filename))
		Expect(err).NotTo(HaveOccurred())

		return string(data)
	}

	yttRender := func(valuesFromInputFile string) string {
		var filePaths []string

		for _, p := range []string{
			"*.star",
			"*.yaml",
			"libs/constants.lib.yaml",
			"upstream/_ytt_lib/*.yaml",
			"upstream/*.yaml",
			"upstream/_ytt_lib/concierge/*.yaml",
			"upstream/_ytt_lib/supervisor/*.yaml",
			"overlay/*.yaml",
			// configDir, ytt is unhappy with the specified list for some reason and has trouble loading files properly
		} {
			matches, err := filepath.Glob(filepath.Join(configDir, p))
			Expect(err).NotTo(HaveOccurred())
			filePaths = append(filePaths, matches...)
		}

		fmt.Printf("🦄 🦄 🦄 all paths? \n\n\n %v \n", filePaths)

		result, err := ytt.RenderYTTTemplate(ytt.CommandOptions{}, filePaths, strings.NewReader(valuesFromInputFile))
		Expect(err).NotTo(HaveOccurred())
		return result
	}

	// Context("No configuration", func() {
	// 	It("renders with an error", func() {
	// 		Expect(err).To(ContainSubstring("configuration is required for pinniped"))
	// 	})
	// })

	Context("Providing an management cluster oicd configuration", func() {
		It("renders the correct set of resources to deploy on a management cluster", func() {

			filenameForInputAndOutput := "mc-oidc.yaml"
			valz := ValuesFromFile(filenameForInputAndOutput)
			expectedz := OutputFromFile(filenameForInputAndOutput)
			resultz := yttRender(valz)

			err := os.WriteFile(filepath.Join(fixtureDir, "hack-exectedz.yaml"), []byte(expectedz), 0644)
			if err != nil {
				panic(err)
			}
			err = os.WriteFile(filepath.Join(fixtureDir, "hack-resultz.yaml"), []byte(resultz), 0644)
			if err != nil {
				panic(err)
			}

			Expect(expectedz).To(BeEquivalentTo(resultz))
		})
	})
})
