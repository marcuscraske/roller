package roller

type Action struct {
	Shell      string `json:"shell"`
	WorkingDir string `json:"working_dir"`
}
