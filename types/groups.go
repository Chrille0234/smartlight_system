package types

type Group struct {
    IDs map[string]string `json:"ids"`
}

type Groups struct {
    Groups map[string]Group `json:"groups"`
}
