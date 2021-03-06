package p2pstore

import (
	"encoding/hex"
	"time"

	"github.com/33cn/chain33/types"
	"github.com/libp2p/go-libp2p-core/peer"
	kbt "github.com/libp2p/go-libp2p-kbucket"
)

func (p *Protocol) refreshLocalChunk() {
	//优先处理同步任务
	if reply, err := p.API.IsSync(); err != nil || !reply.IsOk {
		return
	}
	start := time.Now()
	var saveNum, syncNum, deleteNum int
	for chunkNum := int64(0); ; chunkNum++ {
		records, err := p.getChunkRecordFromBlockchain(&types.ReqChunkRecords{Start: chunkNum, End: chunkNum})
		if err != nil || len(records.Infos) != 1 {
			log.Info("refreshLocalChunk", "break at chunk num", chunkNum, "error", err)
			break
		}
		info := records.Infos[0]
		msg := &types.ChunkInfoMsg{
			ChunkHash: info.ChunkHash,
			Start:     info.Start,
			End:       info.End,
		}
		if p.shouldSave(info.ChunkHash) {
			saveNum++
			// 本地处理
			// 1. p2pStore已保存数据则直接更新p2pStore
			// 2. p2pStore未保存数据则从blockchain获取数据
			if err := p.storeChunk(msg); err == nil {
				continue
			}

			// 通过网络从其他节点获取数据
			p.chunkStatusCacheMutex.Lock()
			chunkHash := hex.EncodeToString(info.ChunkHash)
			if _, ok := p.chunkStatusCache[hex.EncodeToString(info.ChunkHash)]; ok {
				continue
			}
			p.chunkStatusCache[chunkHash] = ChunkStatus{info: msg, status: Waiting}
			p.chunkStatusCacheMutex.Unlock()
			syncNum++
			p.chunkToSync <- msg
		} else {
			deleteNum++
			p.chunkToDelete <- msg
		}
	}
	log.Info("refreshLocalChunk", "save num", saveNum, "sync num", syncNum, "delete num", deleteNum, "time cost", time.Since(start), "exRT size", p.getExtendRoutingTable().Size())
}

func (p *Protocol) shouldSave(chunkHash []byte) bool {
	if p.SubConfig.IsFullNode {
		return true
	}
	pids := p.getExtendRoutingTable().NearestPeers(genDHTID(chunkHash), backup)
	if len(pids) < backup || kbt.Closer(p.Host.ID(), pids[backup-1], genChunkNameSpaceKey(chunkHash)) {
		return true
	}
	return false
}

func (p *Protocol) getExtendRoutingTable() *kbt.RoutingTable {
	if p.extendRoutingTable.Size() < p.RoutingTable.Size() {
		return p.RoutingTable
	}
	return p.extendRoutingTable
}

func (p *Protocol) updateExtendRoutingTable() {
	key := []byte("temp")
	count := 100
	start := time.Now()
	for _, pid := range p.extendRoutingTable.ListPeers() {
		if p.RoutingTable.Find(pid) == "" {
			p.extendRoutingTable.Remove(pid)
		}
	}
	peers := p.RoutingTable.ListPeers()
	for _, pid := range peers {
		_, _ = p.extendRoutingTable.Update(pid)
	}
	if key != nil {
		peers = p.RoutingTable.NearestPeers(genDHTID(key), backup-1)
	}

	searchedPeers := make(map[peer.ID]struct{})
	for i, pid := range peers {
		// 保证 extendRoutingTable 至少有 300 个节点，且至少从 3 个节点上获取新节点，
		if i+1 > 3 && p.extendRoutingTable.Size() > p.SubConfig.MaxExtendRoutingTableSize {
			break
		}
		searchedPeers[pid] = struct{}{}
		extendPeers, err := p.fetchShardPeers(key, count, pid)
		if err != nil {
			log.Error("updateExtendRoutingTable", "fetchShardPeers error", err, "peer id", pid)
			continue
		}
		for _, cPid := range extendPeers {
			if cPid == p.Host.ID() {
				continue
			}
			_, _ = p.extendRoutingTable.Update(cPid)
		}
	}

	// 如果扩展路由表节点数小于200，则迭代查询增加节点
	var lastSize int //如果经过一轮迭代节点数没有增加则结束迭代，防止节点数不到200导致无法退出
	for p.extendRoutingTable.Size() < p.SubConfig.MinExtendRoutingTableSize && p.extendRoutingTable.Size() > lastSize {
		lastSize = p.extendRoutingTable.Size()
		for _, pid := range p.extendRoutingTable.ListPeers() {
			if _, ok := searchedPeers[pid]; ok {
				continue
			}
			searchedPeers[pid] = struct{}{}
			closerPeers, err := p.fetchShardPeers(key, count, pid)
			if err != nil {
				log.Error("updateExtendRoutingTable", "fetchShardPeers error", err, "peer id", pid)
				continue
			}
			for _, cPid := range closerPeers {
				if cPid == p.Host.ID() {
					continue
				}
				_, _ = p.extendRoutingTable.Update(cPid)
			}
		}
	}
	log.Info("updateExtendRoutingTable", "local peers count", p.RoutingTable.Size(), "extendRoutingTable peer count", p.extendRoutingTable.Size(), "time cost", time.Since(start), "origin count", count)
}
