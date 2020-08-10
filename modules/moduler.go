package modules

const (
	ModBootlog       = "bootlog"
	ModBotnet        = "botnet"
	ModCargo         = "cargo"
	ModCc            = "cc"
	ModComposer      = "composer"
	ModCryptomining  = "cryptomining"
	ModDocker        = "docker"
	ModDownload      = "download"
	ModGomod         = "gomod"
	ModKernelCompile = "kernel_compile"
	ModMemdump       = "memdump"
	ModWeblog        = "weblog"
)

var Registry = [...]Moduler{
	NewBootlogModule(),
	NewBotnetModule(),
	NewCargoModule(),
	NewCcModule(),
	NewComposerModule(),
	NewCryptominingModule(),
	NewDockerModule(),
	NewDownloadModule(),
	NewGomodModule(),
	NewKernelCompileModule(),
	NewMemdumpModule(),
	NewWeblogModule(),
}

type Moduler interface {
	Name() string
	Display()
}
