package main

type TLDRDBCached struct {

}

func (t *TLDRDBCached) Retrieve(key string) string {
	
}

func (t *TLDRDBCached) List() []string {

}

func NewTLDRDBCached(nonCachedProvider TLDRProvider) TLDRProvider {
	return nil
}
