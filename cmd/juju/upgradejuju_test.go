package main

import (
	"io/ioutil"
	. "launchpad.net/gocheck"
	"launchpad.net/juju-core/environs"
	envtesting "launchpad.net/juju-core/environs/testing"
	"launchpad.net/juju-core/environs/tools"
	"launchpad.net/juju-core/juju/testing"
	"launchpad.net/juju-core/state"
	coretesting "launchpad.net/juju-core/testing"
	"launchpad.net/juju-core/version"
)

type UpgradeJujuSuite struct {
	testing.JujuConnSuite
}

var _ = Suite(&UpgradeJujuSuite{})

var upgradeJujuTests = []struct {
	about          string
	private        []string
	public         []string
	currentVersion string
	agentVersion   string

	args              []string
	expectInitErr     string
	expectErr         string
	expectVersion     string
	expectDevelopment bool
	expectUploaded    string
}{{
	about:          "unwanted extra argument",
	currentVersion: "1.0.0-foo-bar",
	agentVersion:   "1.0.0",
	args:           []string{"foo"},
	expectInitErr:  "unrecognized args:.*",
}, {
	about:          "invalid --version value",
	currentVersion: "1.0.0-foo-bar",
	agentVersion:   "1.0.0",
	args:           []string{"--version", "invalid-version"},
	expectInitErr:  "invalid version .*",
}, {
	about:          "from private storage",
	private:        []string{"2.0.0-foo-bar", "2.0.2-foo-bletch", "2.0.3-foo-bar"},
	public:         []string{"2.0.0-foo-bar", "2.0.4-foo-bar", "2.0.5-foo-bar"},
	currentVersion: "2.0.0-foo-bar",
	agentVersion:   "2.0.0",
	expectVersion:  "2.0.2",
}, {
	about:          "current dev version, from private storage",
	private:        []string{"2.0.0-foo-bar", "2.0.2-foo-bar", "2.0.3-foo-bar", "3.0.1-foo-bar"},
	public:         []string{"2.0.0-foo-bar", "2.0.4-foo-bar", "2.0.5-foo-bar"},
	currentVersion: "2.0.1-foo-bar",
	agentVersion:   "2.0.1",
	expectVersion:  "2.0.3",
}, {
	about:             "dev version flag, from private storage",
	private:           []string{"2.0.0-foo-bar", "2.0.2-foo-bar", "2.0.3-foo-bar"},
	public:            []string{"2.0.0-foo-bar", "2.0.4-foo-bar", "2.0.5-foo-bar"},
	currentVersion:    "2.0.0-foo-bar",
	args:              []string{"--dev"},
	agentVersion:      "2.0.0",
	expectVersion:     "2.0.3",
	expectDevelopment: true,
}, {
	about:          "from public storage",
	public:         []string{"2.0.0-foo-bar", "2.0.2-arble-bletch", "2.0.3-foo-bar"},
	currentVersion: "2.0.0-foo-bar",
	agentVersion:   "2.0.0",
	expectVersion:  "2.0.2",
}, {
	about:          "current dev version, from public storage",
	public:         []string{"2.0.0-foo-bar", "2.0.2-arble-bletch", "2.0.3-foo-bar"},
	currentVersion: "2.0.1-foo-bar",
	agentVersion:   "2.0.1",
	expectVersion:  "2.0.3",
}, {
	about:             "dev version flag, from public storage",
	public:            []string{"2.0.0-foo-bar", "2.0.2-arble-bletch", "2.0.3-foo-bar"},
	currentVersion:    "2.0.0-foo-bar",
	args:              []string{"--dev"},
	agentVersion:      "2.0.0",
	expectVersion:     "2.0.3",
	expectDevelopment: true,
}, {
	about:          "specified version",
	public:         []string{"2.0.3-foo-bar"},
	currentVersion: "3.0.0-foo-bar",
	agentVersion:   "2.0.0",
	args:           []string{"--version", "2.0.3"},
	expectVersion:  "2.0.3",
}, {
	about:          "specified version missing",
	currentVersion: "3.0.0-foo-bar",
	agentVersion:   "2.0.0",
	args:           []string{"--version", "2.0.3"},
	expectVersion:  "2.0.3",
}, {
	about:          "major version change",
	currentVersion: "2.0.0-foo-bar",
	agentVersion:   "4.1.2",
	args:           []string{"--version", "3.1.2"},
	expectErr:      "cannot upgrade major versions yet",
}, {
	about:          "upload",
	currentVersion: "2.0.2-foo-bar",
	agentVersion:   "2.0.0",
	args:           []string{"--upload-tools"},
	expectVersion:  "2.0.2",
	expectUploaded: "2.0.2-foo-bar",
}, {
	about:          "upload dev version, currently on release version",
	currentVersion: "2.0.1-foo-bar",
	agentVersion:   "2.0.0",
	args:           []string{"--upload-tools"},
	expectVersion:  "2.0.1",
	expectUploaded: "2.0.1-foo-bar",
}, {
	about:          "upload and bump version",
	private:        []string{"2.4.6-foo-bar", "2.4.8-foo-bar"},
	public:         []string{"2.4.10-foo-bar"},
	currentVersion: "2.4.6-foo-bar",
	agentVersion:   "2.4.0",
	args:           []string{"--upload-tools", "--bump-version"},
	expectVersion:  "2.4.6.1",
	expectUploaded: "2.4.6.1-foo-bar",
}, {
	about:          "upload with previously bumped version",
	private:        []string{"2.4.6-foo-bar", "2.4.6.1-foo-bar", "2.4.8-foo-bar"},
	public:         []string{"2.4.10-foo-bar"},
	currentVersion: "2.4.6-foo-bar",
	agentVersion:   "2.4.6.1",
	args:           []string{"--upload-tools", "--bump-version"},
	expectVersion:  "2.4.6.2",
	expectUploaded: "2.4.6.2-foo-bar",
},
}

// mockUploadTools simulates the effect of tools.Upload, but skips the time-
// consuming build from source.
// TODO(fwereade) better factor environs/tools such that build logic is
// exposed and can itself be neatly mocked?
func mockUploadTools(putter tools.URLPutter, forceVersion *version.Number, fakeSeries ...string) (*state.Tools, error) {
	storage := putter.(environs.Storage)
	vers := version.Current
	if forceVersion != nil {
		vers.Number = *forceVersion
	}
	t := envtesting.MustUploadFakeToolsVersion(storage, vers)
	for _, series := range fakeSeries {
		vers.Series = series
		envtesting.MustUploadFakeToolsVersion(storage, vers)
	}
	return t, nil
}

func (s *UpgradeJujuSuite) TestUpgradeJuju(c *C) {
	oldVersion := version.Current
	uploadTools = mockUploadTools
	defer func() {
		version.Current = oldVersion
		uploadTools = tools.Upload
	}()

	for i, test := range upgradeJujuTests {
		c.Logf("\ntest %d: %s", i, test.about)
		// Set up the test preconditions.
		s.Reset(c)
		for _, v := range test.private {
			vers := version.MustParseBinary(v)
			envtesting.MustUploadFakeToolsVersion(s.Conn.Environ.Storage(), vers)
		}
		for _, v := range test.public {
			vers := version.MustParseBinary(v)
			storage := s.Conn.Environ.PublicStorage().(environs.Storage)
			envtesting.MustUploadFakeToolsVersion(storage, vers)
		}
		version.Current = version.MustParseBinary(test.currentVersion)
		err := SetAgentVersion(s.State, version.MustParse(test.agentVersion), false)
		c.Assert(err, IsNil)

		// Run the command
		com := &UpgradeJujuCommand{}
		err = coretesting.InitCommand(com, test.args)
		if test.expectInitErr != "" {
			c.Check(err, ErrorMatches, test.expectInitErr)
			continue
		}
		err = com.Run(coretesting.Context(c))
		if test.expectErr != "" {
			c.Check(err, ErrorMatches, test.expectErr)
			continue
		}
		c.Assert(err, IsNil)
		cfg, err := s.State.EnvironConfig()
		c.Check(err, IsNil)
		c.Check(cfg.AgentVersion(), Equals, version.MustParse(test.expectVersion))
		c.Check(cfg.Development(), Equals, test.expectDevelopment)

		if test.expectUploaded != "" {
			vers := version.MustParseBinary(test.expectUploaded)
			r, err := s.Conn.Environ.Storage().Get(tools.StorageName(vers))
			c.Assert(err, IsNil)
			data, err := ioutil.ReadAll(r)
			c.Check(err, IsNil)
			c.Check(string(data), Equals, test.expectUploaded)
			r.Close()
		}
	}
}

// JujuConnSuite very helpfully uploads some default
// tools to the environment's storage. We don't want
// 'em there.
func (s *UpgradeJujuSuite) Reset(c *C) {
	s.JujuConnSuite.Reset(c)
	envtesting.RemoveTools(c, s.Conn.Environ.Storage())
	envtesting.RemoveTools(c, s.Conn.Environ.PublicStorage().(environs.Storage))
}

func (s *UpgradeJujuSuite) TestUpgradeJujuWithRealPutTools(c *C) {
	s.Reset(c)
	_, err := coretesting.RunCommand(c, &UpgradeJujuCommand{}, []string{"--upload-tools", "--dev"})
	c.Assert(err, IsNil)
	name := tools.StorageName(version.Current)
	r, err := s.Conn.Environ.Storage().Get(name)
	c.Assert(err, IsNil)
	r.Close()
}
