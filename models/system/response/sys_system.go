package response

import "leiserv/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
