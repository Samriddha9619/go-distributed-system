package standalone_storage

import (
	"path/filepath"
	"github.com/pingcap-incubator/tinykv/kv/config"
	"github.com/pingcap-incubator/tinykv/kv/storage"
	"github.com/Connor1996/badger"
	"github.com/pingcap-incubator/tinykv/kv/util/engine_util"
	"github.com/pingcap-incubator/tinykv/proto/pkg/kvrpcpb"
)

// StandAloneStorage is an implementation of `Storage` for a single-node TinyKV instance. It does not
// communicate with other nodes and all data is stored locally.
type StandAloneStorage struct {
	// Your Data Here (1).
	db *badger.DB
}
type StandAloneReader struct{
	kvTxn *badger.Txn// KV Txn is term for the data transaction (user data)
	// while Raft Txn is consensus log transactioin (internal data)
}
func NewStandAloneStorage(conf *config.Config) *StandAloneStorage {
	// Your Code Here (1).
	dbPath := conf.DBPath
	kvPath := filepath.Join(dbPath,"kv")
	//raftpath := path.Join(dbPath,"raft")
	db := engine_util.CreateDB(kvPath,false)
	return &StandAloneStorage{
		db: db,
	}
}

func (s *StandAloneStorage) Start() error {
	// Your Code Here (1).
	return nil
}

func (s *StandAloneStorage) Stop() error {
	// Your Code Here (1).

	return s.db.Close()
}

func (s *StandAloneStorage) Reader(ctx *kvrpcpb.Context) (storage.StorageReader, error) {
	// Your Code Here (1).
	txn := s.db.NewTransaction(false)

	return &StandAloneReader{
		kvTxn:txn,
	},nil
}

func (s *StandAloneStorage) Write(ctx *kvrpcpb.Context, batch []storage.Modify) error {
    txn := s.db.NewTransaction(true)
    defer txn.Discard()

    for _, m := range batch {
        switch m.Data.(type) {
        case storage.Put:
            put := m.Data.(storage.Put)
            key := engine_util.KeyWithCF(put.Cf, put.Key)
            
            err := txn.Set(key, put.Value)
            if err != nil {       
                return err       
            }                     
            
        case storage.Delete:
            del := m.Data.(storage.Delete)
            key := engine_util.KeyWithCF(del.Cf, del.Key)
            
            err := txn.Delete(key)
            if err != nil {       
                return err        
            }                     
        }
    }
    
    return txn.Commit()
}

func (r *StandAloneReader) GetCF(cf string, key []byte) ([]byte, error){
	val,err := engine_util.GetCFFromTxn(r.kvTxn,cf,key)
	if err == badger.ErrKeyNotFound{
		return nil, nil
	}
	return val,err
}

func (r *StandAloneReader) IterCF(cf string) engine_util.DBIterator{
	return engine_util.NewCFIterator(cf,r.kvTxn)
}

func (r *StandAloneReader) Close(){
	r.kvTxn.Discard()
}

