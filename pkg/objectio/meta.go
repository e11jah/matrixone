// Copyright 2021 Matrix Origin
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package objectio

// +---------------------------------------------------------------------------------------------+
// |                                           Header                                            |
// +-------------+---------------+--------------+---------------+---------------+----------------+
// | TableID(8B) | SegmentID(8B) | BlockID(8B)  | ColumnCnt(2B) | Reserved(34B) |  Chksum(4B)    |
// +-------------+---------------+--------------+---------------+---------------+----------------+
// |                                         ColumnMeta                                          |
// +---------------------------------------------------------------------------------------------+
// |                                         ColumnMeta                                          |
// +---------------------------------------------------------------------------------------------+
// |                                         ColumnMeta                                          |
// +---------------------------------------------------------------------------------------------+
// |                                         ..........                                          |
// +---------------------------------------------------------------------------------------------+
// Header Size = 64B
// TableID = Table ID for Block
// SegmentID = Segment ID
// BlockID = Block ID
// ColumnCnt = The number of column in the block
// Chksum = Block metadata checksum
// Reserved = 34 bytes reserved space
type BlockMeta struct {
	header BlockHeader
}

func (bm *BlockMeta) GetHeader() BlockHeader {
	return bm.header
}

type BlockHeader struct {
	tableId     uint64
	segmentId   uint64
	blockId     uint64
	columnCount uint16
	dummy       [34]byte
	checksum    uint32
}

func (bh *BlockHeader) GetTableId() uint64 {
	return bh.tableId
}
func (bh *BlockHeader) GetSegmentId() uint64 {
	return bh.segmentId
}
func (bh *BlockHeader) GetBlockId() uint64 {
	return bh.blockId
}
func (bh *BlockHeader) GetColumnCount() uint16 {
	return bh.columnCount
}

// +---------------------------------------------------------------------------------------------------------------+
// |                                                    ColumnMeta                                                 |
// +--------+-------+----------+--------+---------+--------+--------+------------+---------+----------+------------+
// |Type(1B)|Idx(2B)| Algo(1B) |Offset(4B)|Size(4B)|oSize(4B)|Min(32B)|Max(32B)|BFoffset(4b)|BFlen(4b)|BFoSize(4B) |
// +--------+-------+----------+--------+---------+--------+--------+------------+---------+----------+------------+
// |                                        Reserved(32B)                                             | Chksum(4B) |
// +---------------------------------------------------------------------------------------------------------------+
// ColumnMeta Size = 128B
// Type = Metadata type, always 0, representing column meta, used for extension.
// Idx = Column index
// Algo = Type of compression algorithm for column data
// Offset = Offset of column data
// Size = Size of column data
// oSize = Original data size
// Min = Column min value
// Max = Column Max value
// BFoffset = Bloomfilter data offset
// Bflen = Bloomfilter data size
// BFoSize = Bloomfilter original data size
// Chksum = Data checksum
// Reserved = 32 bytes reserved space
type ColumnMeta struct {
	typ         uint8
	idx         uint16
	alg         uint8
	location    Extent
	zoneMap     ZoneMap
	bloomFilter Extent
	dummy       [32]byte
	checksum    uint32
}

func (cm *ColumnMeta) GetType() uint8 {
	return cm.typ
}

func (cm *ColumnMeta) GetIdx() uint16 {
	return cm.idx
}

func (cm *ColumnMeta) GetAlg() uint8 {
	return cm.alg
}

func (cm *ColumnMeta) GetLocation() Extent {
	return cm.location
}

func (cm *ColumnMeta) GetZoneMap() ZoneMap {
	return cm.zoneMap
}

func (cm *ColumnMeta) GetBloomFilter() Extent {
	return cm.bloomFilter
}

type Header struct {
	magic   uint64
	version uint16
	dummy   [22]byte
}
