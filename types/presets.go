package types

type Preset struct {
  Colors map[string]string `json:"colors"`
}

type Presets struct {
  Presets map[string]Preset `json:"presets"`
}
