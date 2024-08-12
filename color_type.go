package main

type OnStatus struct {
	On bool `json:"on"`
}

type XYStatus struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type Color struct {
	XY XYStatus `json:"xy"`
	// Other fields omitted for brevity
}

type Dimming struct {
	Brightness   float64 `json:"brightness"`
	MinDimLevel  float64 `json:"min_dim_level"`
}

type ColorTemperature struct {
	Mirek       *int  `json:"mirek"`
	MirekValid  bool  `json:"mirek_valid"`
	MirekSchema struct {
		MirekMinimum int `json:"mirek_minimum"`
		MirekMaximum int `json:"mirek_maximum"`
	} `json:"mirek_schema"`
}

type Powerup struct {
	Preset    string `json:"preset"`
	Configured bool `json:"configured"`
	On        struct {
		Mode string `json:"mode"`
		On   OnStatus `json:"on"`
	} `json:"on"`
	Dimming struct {
		Mode    string  `json:"mode"`
		Dimming struct {
			Brightness float64 `json:"brightness"`
		} `json:"dimming"`
	} `json:"dimming"`
	Color struct {
		Mode            string            `json:"mode"`
		ColorTemperature struct {
			Mirek int `json:"mirek"`
		} `json:"color_temperature"`
	} `json:"color"`
}

type Light struct {
	ID               string           `json:"id"`
	IDV1             string           `json:"id_v1"`
	Owner            struct {
		RID   string `json:"rid"`
		RType string `json:"rtype"`
	} `json:"owner"`
	Metadata         struct {
		Name      string `json:"name"`
		Archetype string `json:"archetype"`
		Function  string `json:"function"`
	} `json:"metadata"`
	ProductData      struct {
		Function string `json:"function"`
	} `json:"product_data"`
	Identify         struct{}      `json:"identify"`
	ServiceID        int           `json:"service_id"`
	On               OnStatus      `json:"on"`
	Dimming          Dimming       `json:"dimming"`
	DimmingDelta     struct{}      `json:"dimming_delta"`
	ColorTemperature ColorTemperature `json:"color_temperature"`
	ColorTemperatureDelta struct{} `json:"color_temperature_delta"`
	Color            Color         `json:"color"`
	Dynamics         struct {
		Status       string   `json:"status"`
		StatusValues []string `json:"status_values"`
		Speed        float64  `json:"speed"`
		SpeedValid   bool     `json:"speed_valid"`
	} `json:"dynamics"`
	Alert struct {
		ActionValues []string `json:"action_values"`
	} `json:"alert"`
	Signaling struct {
		SignalValues []string `json:"signal_values"`
	} `json:"signaling"`
	Mode    string `json:"mode"`
	Effects struct {
		StatusValues []string `json:"status_values"`
		Status       string   `json:"status"`
		EffectValues []string `json:"effect_values"`
	} `json:"effects"`
	Powerup Powerup `json:"powerup"`
	Type    string  `json:"type"`
}

type LightResponse struct {
	Errors []interface{} `json:"errors"`
	Data   []Light       `json:"data"`
}
