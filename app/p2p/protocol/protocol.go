/**
 * Description:
 * Author: Yihen.Liu
 * Create: 2019-09-06
 */
package protocol

import "v2ray.com/core/app/p2p/protocol/seedlist"

type Protocol struct {
	*seedlist.SeedListProtocol
}
