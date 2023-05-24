package buildpack_test

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	"github.com/buildpacks/lifecycle/api"
	"github.com/heroku/color"
	"github.com/sclevine/spec"
	"github.com/sclevine/spec/report"

	ifakes "github.com/buildpacks/pack/internal/fakes"
	"github.com/buildpacks/pack/pkg/archive"
	"github.com/buildpacks/pack/pkg/buildpack"
	"github.com/buildpacks/pack/pkg/dist"
	"github.com/buildpacks/pack/pkg/logging"
	h "github.com/buildpacks/pack/testhelpers"
)

func TestBuildModuleWriter(t *testing.T) {
	color.Disable(true)
	defer color.Disable(false)
	spec.Run(t, "testBuildModuleWriter", testBuildModuleWriter, spec.Parallel(), spec.Report(report.Terminal{}))
}

func testBuildModuleWriter(t *testing.T, when spec.G, it spec.S) {
	var (
		outBuf            bytes.Buffer
		logger            logging.Logger
		buildModuleWriter *buildpack.BuildModuleWriter
		bp1v1             buildpack.BuildModule
		bp1v2             buildpack.BuildModule
		bp2v1             buildpack.BuildModule
		bp3v1             buildpack.BuildModule
		tmpDir            string
		err               error
	)

	it.Before(func() {
		logger = logging.NewLogWithWriters(&outBuf, &outBuf, logging.WithVerbose())
		buildModuleWriter = buildpack.NewBuildModuleWriter(logger, archive.DefaultTarWriterFactory())
		tmpDir, err = os.MkdirTemp("", "test_build_module_writer")
		h.AssertNil(t, err)

		bp1v1, err = ifakes.NewFakeBuildpack(dist.BuildpackDescriptor{
			WithAPI: api.MustParse("0.2"),
			WithInfo: dist.ModuleInfo{
				ID:      "buildpack-1-id",
				Version: "buildpack-1-version-1",
			},
			WithStacks: []dist.Stack{{
				ID: "*",
			}},
		}, 0644)
		h.AssertNil(t, err)

		bp1v2, err = ifakes.NewFakeBuildpack(dist.BuildpackDescriptor{
			WithAPI: api.MustParse("0.2"),
			WithInfo: dist.ModuleInfo{
				ID:      "buildpack-1-id",
				Version: "buildpack-1-version-2",
			},
			WithStacks: []dist.Stack{{
				ID: "*",
			}},
		}, 0644)
		h.AssertNil(t, err)

		bp2v1, err = ifakes.NewFakeBuildpack(dist.BuildpackDescriptor{
			WithAPI: api.MustParse("0.2"),
			WithInfo: dist.ModuleInfo{
				ID:      "buildpack-2-id",
				Version: "buildpack-2-version-1",
			},
			WithStacks: []dist.Stack{{
				ID: "*",
			}},
		}, 0644)
		h.AssertNil(t, err)

		bp3v1, err = ifakes.NewFakeBuildpack(dist.BuildpackDescriptor{
			WithAPI: api.MustParse("0.2"),
			WithInfo: dist.ModuleInfo{
				ID:      "buildpack-3-id",
				Version: "buildpack-3-version-1",
			},
			WithStacks: []dist.Stack{{
				ID: "*",
			}},
		}, 0644)
		h.AssertNil(t, err)
	})

	it.After(func() {
		err := os.RemoveAll(tmpDir)
		h.AssertNil(t, err)
	})

	when("#NToLayerTar", func() {
		when("there are not duplicated buildpacks", func() {
			it("creates a tar", func() {
				bpModules := []buildpack.BuildModule{bp1v1, bp2v1, bp3v1}
				tarFile, err := buildModuleWriter.NToLayerTar(tmpDir, "test-file-1", bpModules)

				h.AssertNil(t, err)
				h.AssertNotNil(t, tarFile)
				assertBuildpackModuleWritten(t, tarFile, bpModules)
			})
		})

		when("there are duplicated buildpacks", func() {
			it("creates a tar skipping root folder from duplicated buildpacks", func() {
				expectedMessage := fmt.Sprintf("folder '%s' was already added, skipping it", "/cnb/buildpacks/buildpack-1-id")
				bpModules := []buildpack.BuildModule{bp1v1, bp1v2, bp2v1, bp3v1}
				tarFile, err := buildModuleWriter.NToLayerTar(tmpDir, "test-file-2", bpModules)
				h.AssertNil(t, err)
				h.AssertNotNil(t, tarFile)
				assertBuildpackModuleWritten(t, tarFile, bpModules)
				h.AssertContains(t, outBuf.String(), expectedMessage)
			})
		})
	})
}

func assertBuildpackModuleWritten(t *testing.T, path string, modules []buildpack.BuildModule) {
	t.Helper()
	for _, module := range modules {
		dirPath := fmt.Sprintf("/cnb/buildpacks/%s/%s", module.Descriptor().Info().ID, module.Descriptor().Info().Version)
		h.AssertOnTarEntry(t, path, dirPath,
			h.IsDirectory(),
		)
	}
}
