/**
 * Description:
 * Author: Yihen.Liu
 * Create: 2020-02-03
 */
package conf

import (
	"v2ray.com/core/app/p2p"
)

func DefaultP2PConfig() *p2p.Config {
	return &p2p.Config{
		Ip:       "127.0.0.1",
		Port:     55155,
		Protocol: "tcp",
		PubId:    "default-p2p-pub-id",
		Seedlist: []*p2p.Seed{
			{
				Protocol: "tcp",
				Ip:       "127.0.0.1",
				Port:     55255,
				PubId:    "seed1.pub.id",
			},
			{
				Protocol: "tcp",
				Ip:       "127.0.0.1",
				Port:     55355,
				PubId:    "seed2.pub.id",
			},
		},
	}
}

type seed struct {
	Protocol string `json:"protocol"`
	IP       string `json:"ip"`
	Port     int32  `json:"port"`
	PubID    string `json:"pub_id"`
}
type P2PConfig struct {
	Protocol string `json:"protocol"`
	IP       string `json:"ip"`
	Port     int32  `json:"port"`
	PubID    string `json:"pub_id"`
	SeedList []seed `json:"seedlist"`
}

func (v *P2PConfig) Build() *p2p.Config {
	var seeds []*p2p.Seed
	for _, value := range v.SeedList {
		seed := p2p.Seed{
			Ip:       value.IP,
			Port:     value.Port,
			Protocol: value.Protocol,
			PubId:    value.PubID}
		seeds = append(seeds, &seed)
	}
	return &p2p.Config{
		Ip:       v.IP,
		Port:     v.Port,
		Protocol: v.Protocol,
		PubId:    v.PubID,
		Seedlist: seeds,
	}
}
