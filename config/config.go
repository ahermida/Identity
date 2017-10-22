/*
        Config holds network and configuration data
*/
package config

var (
    Network = &NetworkConfig{
        port: "8080",
    }
    Params = &ParamsConfig{
        NodeID: "1337",
    }
)
type NetworkConfig struct {
    port string
}

type ParamsConfig struct {
    NodeID string
}
