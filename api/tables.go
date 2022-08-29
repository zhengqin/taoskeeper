package api

import "strconv"

type Report struct {
	Ts          string       `json:"ts"`
	DnodeID     int          `json:"dnode_id"`
	DnodeEp     string       `json:"dnode_ep"`
	ClusterID   string       `json:"cluster_id"`
	Protocol    int          `json:"protocol"`
	ClusterInfo ClusterInfo  `json:"cluster_info"`
	StbInfos    []StbInfo    `json:"stb_infos"`
	VgroupInfos []VgroupInfo `json:"vgroup_infos"`
	GrantInfo   GrantInfo    `json:"grant_info"`
	DnodeInfo   DnodeInfo    `json:"dnode_info"`
	DiskInfos   DiskInfo     `json:"disk_infos"`
	LogInfos    LogInfo      `json:"log_infos"`
}

type ClusterInfo struct {
	FirstEp          string  `json:"first_ep"`
	FirstEpDnodeID   int     `json:"first_ep_dnode_id"`
	Version          string  `json:"version"`
	MasterUptime     float32 `json:"master_uptime"`
	MonitorInterval  int     `json:"monitor_interval"`
	DbsTotal         int     `json:"dbs_total"`
	TbsTotal         int     `json:"tbs_total"`
	StbsTotal        int     `json:"stbs_total"`
	VgroupsTotal     int     `json:"vgroups_total"`
	VgroupsAlive     int     `json:"vgroups_alive"`
	VnodesTotal      int     `json:"vnodes_total"`
	VnodesAlive      int     `json:"vnodes_alive"`
	ConnectionsTotal int     `json:"connections_total"`
	Dnodes           []Dnode `json:"dnodes"`
	Mnodes           []Mnode `json:"mnodes"`
}

var dnodeEpLen = strconv.Itoa(256)

var CreateClusterInfoSql = "create table if not exists cluster_info (" +
	"ts timestamp, " +
	"first_ep binary(134), " +
	"first_ep_dnode_id int, " +
	"version binary(12), " +
	"master_uptime float, " +
	"monitor_interval int, " +
	"dbs_total int, " +
	"tbs_total int, " +
	"stbs_total int, " +
	"dnodes_total int, " +
	"dnodes_alive int, " +
	"mnodes_total int, " +
	"mnodes_alive int, " +
	"vgroups_total int, " +
	"vgroups_alive int, " +
	"vnodes_total int, " +
	"vnodes_alive int, " +
	"connections_total int, " +
	"protocol int " +
	") tags (cluster_id nchar(32))"

type Dnode struct {
	DnodeID int    `json:"dnode_id"`
	DnodeEp string `json:"dnode_ep"`
	Status  string `json:"status"`
}

var CreateDnodeSql = "create table if not exists d_info (" +
	"ts timestamp, " +
	"status binary(10)" +
	") tags (dnode_id int, dnode_ep nchar(" + dnodeEpLen + "), cluster_id nchar(32))"

type Mnode struct {
	MnodeID int    `json:"mnode_id"`
	MnodeEp string `json:"mnode_ep"`
	Role    string `json:"role"`
}

var CreateMnodeSql = "create table if not exists m_info (" +
	"ts timestamp, " +
	"role binary(10)" +
	") tags (mnode_id int, mnode_ep nchar(" + dnodeEpLen + "), cluster_id nchar(32))"

type DnodeInfo struct {
	Uptime                float32 `json:"uptime"`
	CPUEngine             float32 `json:"cpu_engine"`
	CPUSystem             float32 `json:"cpu_system"`
	CPUCores              float32 `json:"cpu_cores"`
	MemEngine             int     `json:"mem_engine"`
	MemSystem             int     `json:"mem_system"`
	MemTotal              int     `json:"mem_total"`
	DiskEngine            int     `json:"disk_engine"`
	DiskUsed              int64   `json:"disk_used"`
	DiskTotal             int64   `json:"disk_total"`
	NetIn                 float32 `json:"net_in"`
	NetOut                float32 `json:"net_out"`
	IoRead                float32 `json:"io_read"`
	IoWrite               float32 `json:"io_write"`
	IoReadDisk            float32 `json:"io_read_disk"`
	IoWriteDisk           float32 `json:"io_write_disk"`
	ReqSelect             int     `json:"req_select"`
	ReqSelectRate         float32 `json:"req_select_rate"`
	ReqInsert             int     `json:"req_insert"`
	ReqInsertSuccess      int     `json:"req_insert_success"`
	ReqInsertRate         float32 `json:"req_insert_rate"`
	ReqInsertBatch        int     `json:"req_insert_batch"`
	ReqInsertBatchSuccess int     `json:"req_insert_batch_success"`
	ReqInsertBatchRate    float32 `json:"req_insert_batch_rate"`
	Errors                int     `json:"errors"`
	VnodesNum             int     `json:"vnodes_num"`
	Masters               int     `json:"masters"`
	HasMnode              int8    `json:"has_mnode"`
	HasQnode              int8    `json:"has_qnode"`
	HasSnode              int8    `json:"has_snode"`
	HasBnode              int8    `json:"has_bnode"`
}

var CreateDnodeInfoSql = "create table if not exists dnodes_info (" +
	"ts timestamp, " +
	"uptime float, " +
	"cpu_engine float, " +
	"cpu_system float, " +
	"cpu_cores float, " +
	"mem_engine int, " +
	"mem_system int, " +
	"mem_total int, " +
	"disk_engine int, " +
	"disk_used bigint, " +
	"disk_total bigint, " +
	"net_in float, " +
	"net_out float, " +
	"io_read float, " +
	"io_write float, " +
	"io_read_disk float, " +
	"io_write_disk float, " +
	"req_select int, " +
	"req_select_rate float, " +
	"req_insert int, " +
	"req_insert_success int, " +
	"req_insert_rate float, " +
	"req_insert_batch int, " +
	"req_insert_batch_success int, " +
	"req_insert_batch_rate float, " +
	"errors int, " +
	"vnodes_num int, " +
	"masters int, " +
	"has_mnode int, " +
	"has_qnode int, " +
	"has_snode int, " +
	"has_bnode int " +
	") tags (dnode_id int, dnode_ep nchar(" + dnodeEpLen + "), cluster_id nchar(32))"

type DiskInfo struct {
	Datadir []DataDir `json:"datadir"`
	Logdir  LogDir    `json:"logdir"`
	Tempdir TempDir   `json:"tempdir"`
}

type DataDir struct {
	Name  string `json:"name"`
	Level int    `json:"level"`
	Avail int64  `json:"avail"`
	Used  int64  `json:"used"`
	Total int64  `json:"total"`
}

var CreateDataDirSql = "create table if not exists data_dir (" +
	"ts timestamp, " +
	"name nchar(200), " +
	"level int, " +
	"avail bigint, " +
	"used bigint, " +
	"total bigint" +
	") tags (dnode_id int, dnode_ep nchar(" + dnodeEpLen + "), cluster_id nchar(32))"

type LogDir struct {
	Name  string `json:"name"`
	Avail int64  `json:"avail"`
	Used  int64  `json:"used"`
	Total int64  `json:"total"`
}

var CreateLogDirSql = "create table if not exists log_dir (" +
	"ts timestamp, " +
	"name nchar(200), " +
	"avail bigint, " +
	"used bigint, " +
	"total bigint" +
	") tags (dnode_id int, dnode_ep nchar(" + dnodeEpLen + "), cluster_id nchar(32))"

type TempDir struct {
	Name  string `json:"name"`
	Avail int64  `json:"avail"`
	Used  int64  `json:"used"`
	Total int64  `json:"total"`
}

var CreateTempDirSql = "create table if not exists temp_dir(" +
	"ts timestamp, " +
	"name nchar(200), " +
	"avail bigint, " +
	"used bigint, " +
	"total bigint " +
	") tags (dnode_id int, dnode_ep nchar(" + dnodeEpLen + "), cluster_id nchar(32))"

type StbInfo struct {
	StbName      string `json:"stb_name"`
	DataBaseName string `json:"database_name"`
}

type VgroupInfo struct {
	VgroupID     int     `json:"vgroup_id"`
	DatabaseName string  `json:"database_name"`
	TablesNum    int     `json:"tables_num"`
	Status       string  `json:"status"`
	Vnodes       []Vnode `json:"vnodes"`
}

var CreateVgroupsInfoSql = "create table if not exists vgroups_info (" +
	"ts timestamp, " +
	"vgroup_id int, " +
	"database_name binary(33), " +
	"tables_num int, " +
	"status binary(512) " +
	") tags (dnode_id int, dnode_ep nchar(" + dnodeEpLen + "), cluster_id nchar(32))"

type Vnode struct {
	DnodeID   int    `json:"dnode_id"`
	VnodeRole string `json:"vnode_role"`
}

var CreateVnodeRoleSql = "create table if not exists vnodes_role (" +
	"ts timestamp, " +
	"vnode_role binary(10) " +
	") tags (dnode_id int, dnode_ep nchar(" + dnodeEpLen + "), cluster_id nchar(32))"

type LogInfo struct {
	Logs    []Log     `json:"logs"`
	Summary []Summary `json:"summary"`
}

type Log struct {
	Ts      string `json:"ts"`
	Level   string `json:"level"`
	Content string `json:"content"`
}

var CreateLogSql = "create table if not exists logs (" +
	"ts timestamp, " +
	"level binary(10), " +
	"content nchar(1024)" +
	") tags (dnode_id int, dnode_ep nchar(" + dnodeEpLen + "), cluster_id nchar(32))"

type Summary struct {
	Level string `json:"level"`
	Total int    `json:"total"`
}

var CreateSummarySql = "create table if not exists log_summary(" +
	"ts timestamp, " +
	"error int, " +
	"info int, " +
	"debug int, " +
	"trace int " +
	") tags (dnode_id int, dnode_ep nchar(" + dnodeEpLen + "), cluster_id nchar(32))"

type GrantInfo struct {
	ExpireTime      int `json:"expire_time"`
	TimeseriesUsed  int `json:"timeseries_used"`
	TimeseriesTotal int `json:"timeseries_total"`
}

var CreateGrantInfoSql = "create table if not exists grants_info(" +
	"ts timestamp, " +
	"expire_time int, " +
	"timeseries_used int, " +
	"timeseries_total int " +
	") tags (dnode_id int, dnode_ep nchar(" + dnodeEpLen + "), cluster_id nchar(32))"

var CreateKeeperSql = "create table if not exists keeper_monitor (" +
	"ts timestamp, " +
	"cpu float, " +
	"mem float, " +
	"total_reports int " +
	") tags (identify nchar(50))"
