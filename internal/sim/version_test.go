package sim_test

import (
	"errors"
	"strings"
	"testing"

	"github.com/opencl-go/simulator-runtime/internal/sim"
)

func TestVersionOf(t *testing.T) {
	t.Parallel()
	type args struct {
		major int
		minor int
		patch int
	}
	tt := []struct {
		name string
		args args
		want sim.Version
	}{
		{name: "0.0.0", args: args{major: 0, minor: 0, patch: 0}, want: 0},
		{name: "1.0.0", args: args{major: 1, minor: 0, patch: 0}, want: 0x00400000},
		{name: "1.1.1", args: args{major: 1, minor: 1, patch: 1}, want: 0x00401001},
		{name: "4.5.6", args: args{major: 4, minor: 5, patch: 6}, want: 0x01005006},
		{name: "0.0.-1", args: args{major: 0, minor: 0, patch: -1}, want: 0x00000FFF},
		{name: "0.-1.0", args: args{major: 0, minor: -1, patch: 0}, want: 0x003FF000},
		{name: "-1.0.0", args: args{major: -1, minor: 0, patch: 0}, want: 0xFFC00000},
	}
	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			if got := sim.VersionOf(tc.args.major, tc.args.minor, tc.args.patch); got != tc.want {
				t.Errorf("VersionOf() = 0x%08X, want 0x%08X", got, tc.want)
			}
		})
	}
}

func TestVersionComponents(t *testing.T) {
	t.Parallel()
	tt := []struct {
		name  string
		ver   sim.Version
		major int
		minor int
		patch int
	}{
		{name: "VersionMin", ver: sim.VersionMin, major: 0, minor: 0, patch: 0},
		{name: "VersionMax", ver: sim.VersionMax, major: 0x3FF, minor: 0x3FF, patch: 0xFFF},
		{name: "1.2.3", ver: sim.VersionOf(1, 2, 3), major: 1, minor: 2, patch: 3},
	}
	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			if got := tc.ver.Major(); got != tc.major {
				t.Errorf("Major() = 0x%03X, want 0x%03X", got, tc.major)
			}
			if got := tc.ver.Minor(); got != tc.minor {
				t.Errorf("Minor() = 0x%03X, want 0x%03X", got, tc.minor)
			}
			if got := tc.ver.Patch(); got != tc.patch {
				t.Errorf("Patch() = 0x%03X, want 0x%03X", got, tc.patch)
			}
		})
	}
}

func TestVersionString(t *testing.T) {
	t.Parallel()
	tt := []struct {
		ver  sim.Version
		want string
	}{
		{ver: sim.VersionOf(0, 0, 0), want: "0.0.0"},
		{ver: sim.VersionOf(1, 2, 3), want: "1.2.3"},
		{ver: sim.VersionOf(10, 0, 100), want: "10.0.100"},
	}
	for _, tc := range tt {
		tc := tc
		t.Run(tc.want, func(t *testing.T) {
			t.Parallel()
			if got := tc.ver.String(); got != tc.want {
				t.Errorf("String() = '%s', want '%s'", got, tc.want)
			}
		})
	}
}

func TestNameVersionFromError(t *testing.T) {
	t.Parallel()
	tt := []struct {
		title   string
		ver     sim.Version
		name    string
		wantErr error
	}{
		{title: "too long name", ver: sim.VersionOf(0, 0, 0), name: strings.Repeat("a", sim.NameVersionNameMaxByteLength+1), wantErr: sim.ErrNameTooLong},
		{title: "name at limit", ver: sim.VersionOf(0, 0, 0), name: strings.Repeat("a", sim.NameVersionNameMaxByteLength), wantErr: nil},
		{title: "shortest name", ver: sim.VersionOf(0, 0, 0), name: "a", wantErr: nil},
		{title: "empty name", ver: sim.VersionOf(0, 0, 0), name: "", wantErr: sim.ErrNameEmpty},
	}
	for _, tc := range tt {
		tc := tc
		t.Run(tc.title, func(t *testing.T) {
			t.Parallel()
			if _, got := sim.NameVersionFrom(tc.name, tc.ver); !errors.Is(got, tc.wantErr) {
				t.Errorf("NameVersionFrom('%v', %v) = err '%v', want '%v'", tc.name, tc.ver, got, tc.wantErr)
			}
		})
	}
}
