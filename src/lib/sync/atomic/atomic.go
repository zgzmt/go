package atomic

import "sync/atomic"

type Atomicbool uint32

func (c *Atomicbool)Get()bool{
	return atomic.LoadUint32((*uint32)(c)) !=0
}
func (c *Atomicbool)SetAtomicUnit32(v bool){
	if v{
		atomic.StoreUint32((*uint32)(c),1)
	}else{
		atomic.StoreUint32((*uint32)(c),0)
	}
}

