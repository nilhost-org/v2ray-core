/**
 * Description:
 * Author: Yihen.Liu
 * Create: 2019-08-01
 */
package account

import (
	"fmt"

	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/syndtr/goleveldb/leveldb"
	"v2ray.com/core/store"
)

const DB_DIR = "./db"

type SocksDB struct {
	db *store.LevelDBStore
}

type Account struct {
	pubKey  []byte
	privKey []byte
}

func GetAccount() *Account {
	socksdb, err := OpenSocksDB(DB_DIR)
	if err != nil {
		fmt.Println("in GetAccount open leveldb err:", err.Error())
	}

	account, err := socksdb.QueryDefaultAccount()
	if account != nil && err == nil {
		fmt.Println("get default account from socksdb")
		return account
	}

	if account == nil && err == nil {
		priv, pubk, _ := crypto.GenerateKeyPair(crypto.Secp256k1, 256)
		err, acnt := socksdb.InsertOrUpdateAccount(pubk, priv)
		if err != nil {
			fmt.Println("in GetAccount, insert or update accout err:", err.Error())
			return nil
		}
		fmt.Println("no account, create default one successfully.")
		return acnt
	}

	return nil
}

func OpenSocksDB(path string) (*SocksDB, error) {
	if ldb, err := store.NewLevelDBStore(path); err == nil {
		return NewSocksDB(ldb), nil
	} else {
		fmt.Printf("open socks leveldb error:%s, setting path:%s", err.Error(), path)
		return nil, err
	}
}

func NewSocksDB(d *store.LevelDBStore) *SocksDB {
	p := &SocksDB{
		db: d,
	}
	return p
}

func (this *SocksDB) Close() {
	this.db.Close()
}

func (this *SocksDB) InsertOrUpdateAccount(pubKey crypto.PubKey, privKey crypto.PrivKey) (error, *Account) {
	pubBytes, err := pubKey.Raw()
	if err != nil {
		fmt.Println("in InsertOrUpdateAccount, pubKey.Bytes() err:", err.Error())
		return err, nil
	}

	privBytes, err := privKey.Raw()
	if err != nil {
		fmt.Println("in InsertOrUpdateAccount, privKey.Bytes() err:", err.Error())
		return err, nil
	}
	info := &Account{pubKey: pubBytes, privKey: privBytes}
	/*	buf, err := json.Marshal(info)
		if err != nil {
			return err, nil
		}*/
	return this.db.Put([]byte("default"), privBytes), info
}

func (this *SocksDB) QueryDefaultAccount() (*Account, error) {
	return this.QueryAccount("default")
}

func (this *SocksDB) QueryAccount(label string) (*Account, error) {
	key := []byte(label)
	value, err := this.db.Get(key)
	if err != nil {
		if err != leveldb.ErrNotFound {
			return nil, err
		}
	}
	if len(value) == 0 {
		return nil, nil
	}

	info := &Account{privKey: value}
	/*	err = json.Unmarshal(value, info)
		if err != nil {
			return nil, err
		}*/
	return info, nil
}

func (this *SocksDB) DeleteAccount(pubKey crypto.PubKey) error {
	if rawKey, err := pubKey.Raw(); err == nil {
		return this.db.Delete(rawKey)
	} else {
		return err
	}
}

func (this *Account) GetPrivKey() crypto.PrivKey {
	if this == nil {
		fmt.Println("account is nil, please use an valid one.")
		return nil
	}
	if privKey, err := crypto.UnmarshalSecp256k1PrivateKey(this.privKey); err == nil {
		return privKey
	} else {
		fmt.Println("unmarshal secp256k1 privKey err:", err.Error())
		return nil
	}
}
