/**
 * Description:
 * Author: Yihen.Liu
 * Create: 2020-02-11
 */
package wire

const CHANNEL_MAX_SIXE uint16 = 256

var SeedlistNotice chan HelloSeedList

func init() {
	SeedlistNotice = make(chan HelloSeedList, CHANNEL_MAX_SIXE)
}
