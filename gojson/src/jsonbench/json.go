package jsonbench

type SessionData struct {
	SessionRestored  []int64 `json:"sessionRestored"`
	FirstPaint       []int64 `json:"firstPaint"`
	Main             []int64 `json:"main"`
	CleanTotalTime   []int64 `json:"cleanTotalTime"`
	CleanActiveTicks []int64 `json:"cleanActiveTicks"`
	V                int64   `json:"_v"`
}

type Session map[string]SessionData

type Data struct {
	Days map[string]Session `json:"days"`
	Last LastData           `json:"last"`
}

type LastData struct {
	SysInfo SysInfo          `json:"org.mozilla.sysinfo.sysinfo"`
	Current Current          `json:"org.mozilla.appSessions.current"`
	Age     Age              `json:"org.mozilla.profile.age"`
	AppInfo GeckoAppInfo     `json:"org.mozilla.appInfo.appinfo"`
	Addons  map[string]Addon `json:"org.mozilla.addons.active`
}

type Addon struct {
	AppDisabled         bool
	ForeignInstall      bool
	HasBinaryComponents bool
	InstallDay          int
	Scope               int
	Type                string
	UpdateDay           int
	UserDisabled        bool
	Version             string
}

type SysInfo struct {
	Version  string `json:"version"`
	Name     string `json:"name"`
	Arch     string `json:"architecture"`
	MemoryMB uint64 `json:"memoryMB"`
	CPUCount uint   `json:"cpuCount"`
	V        int    `json:"_v"`
}

type Current struct {
	SessionRestored int64 `json:"sessionRestored"`
	FirstPaint      int64 `json:"firstPaint"`
	Main            int64 `json:"main"`
	TotalTime       int64 `json:"totalTime"`
	ActiveTicks     int64 `json:"activeTicks"`
	StartDay        int64 `json:"startDay"`
	V               int   `json:"_v"`
}

type Age struct {
	ProfileCreation int64 `json:"profileCreation"`
	V               int   `json:"_v"`
}

type GeckoAppInfo struct {
	UpdateChannel   string `json:"updateChannel"`
	XPComABI        string `json:"xpcomabi"`
	OS              string `json:"os"`
	V               int    `json:"_v"`
	Vendor          string `json:"vendor"`
	Name            string `json:"name"`
	ID              string `json:"id"`
	Version         string `json:"version"`
	AppBuildId      string `json:"appBuildID"`
	PlatformVersion string `json:"platformVersion"`
	PlatformBuildID string `json:"platformBuildID"`
}

type ChromeHangs struct {
	Durations []int `json:"durations"`
	MemoryMap []int `json:"memoryMap"`
	Stacks    []int `json:"stacks"`
}

type Metadata struct {
	LastPingDate string       `json:"lastPingDate"`
	Data         Data         `json:"data"`
	GeckoAppInfo GeckoAppInfo `json:"geckoAppInfo"`
	ThisPingDate string       `json:"thisPingDate"`
	Version      uint         `json:"version"`
}

type Histogram struct {
	BucketCount   uint  `json:"bucket_count"`
	HistogramType int   `json:"histogram_type"`
	Range         []int `json:"range"`
	LogSum        int   `json:"log_sum"`
	LogSumSquares int   `json:"log_sum_squares"`
	Sum           uint  `json:"sum"`
	SumSquaresHi  uint  `json:"sum_squares_hi"`
	SumSquaresLo  uint  `json:"sum_squares_lo"`
	Values        map[string]int
}

type Info struct {
	OS                   string `json:"OS"`
	AdapterDescription   string `json:"adapterDescription"`
	AdapterDeviceID      string `json:"adapterDescription"`
	AdapterDriverVersion string `json:"adapterDriverVersion"`
	AdapterVendorID      string `json:"adapterVendorID"`
	Addons               string `json:"addons"`
	AppBuildID           string `json:"appBuildID"`
	AppID                string `json:"appID"`
	AppName              string `json:"appName"`
	AppUpdateChannel     string `json:"appUpdateChannel"`
	AppVersion           string `json:"appVersion"`
	Arch                 string `json:"arch"`
	FlashVersion         string `json:"flashVersion"`
	HasARMv6             bool   `json:"hasARMv6"`
	HasARMv7             bool   `json:"hasARMv7"`
	HasEDSP              bool   `json:"hasEDSP"`
	HasMMX               bool   `json:"hasMMX"`
	HasNEON              bool   `json:"hasNEON"`
	HasSSE               bool   `json:"hasSSE"`
	HasSSE2              bool   `json:"hasSSE2"`
	HasSSE3              bool   `json:"hasSSE3"`
	HasSSE4A             bool   `json:"hasSSE4A"`
	HasSSE4_1            bool   `json:"hasSSE4_1"`
	HasSSE4_2            bool   `json:"hasSSE4_2"`
	HasSSSE3             bool   `json:"hasSSSE3"`
	Locale               string `json:"locale"`
	Memsize              int    `json:"memsize"`
	PlatformBuildID      string `json:"platformBuildID"`
	Reason               string `json:"reason"`
	Revision             string `json:"revision"`
	Version              string `json:"version"`
}

type AddonManager struct {
	InstalledUnpacked        int `json:"installedUnpacked"`
	ModifiedExceptInstallRDF int `json:"modifiedExceptInstallRDF"`
	ModifiedUnpacked         int `json:"modifiedUnpacked"`
	ModifiedXPI              int `json:"modifiedXPI"`
}

type SimpleMeasurement struct {
	AMI_startup_begin                int            `json:"AMI_startup_begin"`
	AMI_startup_end                  int            `json:"AMI_startup_end"`
	XPI_bootstrap_addons_begin       int            `json:"XPI_bootstrap_addons_begin"`
	XPI_bootstrap_addons_end         int            `json:"XPI_bootstrap_addons_end"`
	XPI_startup_begin                int            `json:"XPI_startup_begin"`
	XPI_startup_end                  int            `json:"XPI_startup_end"`
	AddonManager                     AddonManager   `json:"addonManager"`
	CreateTopLevelWindow             int            `json:"createTopLevelWindow"`
	DebuggerAttached                 int            `json:"debuggerAttached"`
	DelayedStartupFinished           int            `json:"delayedStartupFinished"`
	DelayedStartupStarted            int            `json:"delayedStartupStarted"`
	FirstLoadURI                     int            `json:"firstLoadURI"`
	FirstPaint                       int            `json:"firstPaint"`
	JS                               map[string]int `json:"js"`
	Main                             int            `json:"main"`
	MaximalNumberOfConcurrentThreads int            `json:"maximalNumberOfConcurrentThreads"`
	SessionRestoreInitialized        int            `json:"sessionRestoreInitialized"`
	SessionRestored                  int            `json:"sessionRestored"`
	Start                            int            `json:"start"`
	StartupCrashDetectionBegin       int            `json:"startupCrashDetectionBegin"`
	StartupCrashDetectionEnd         int            `json:"startupCrashDetectionEnd"`
	StartupInterrupted               int            `json:"startupInterrupted"`
	Uptime                           int            `json:"uptime"`
}

type Telemetry struct {
	AddonHistograms    struct{}             `json:"addonHistograms"`
	ChromeHangs        ChromeHangs          `json:"chromeHangs"`
	Histograms         map[string]Histogram `json:"histogram"`
	Info               Info                 `json:"info"`
	LateWrites         map[string][]int     `json:"lateWrites"`
	SimpleMeasurements SimpleMeasurement    `json:"simpleMeasurements"`
}

type JSONBlob struct {
	Metadata  Metadata    `json:"metadata"`
	Telemetry []Telemetry `json:"telemetry"`
}
