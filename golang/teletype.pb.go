// Code generated by protoc-gen-go.
// source: teletype.proto
// DO NOT EDIT!

/*
Package teletype is a generated protocol buffer package.

It is generated from these files:
	teletype.proto

It has these top-level messages:
	Telecast
*/
package teletype

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type TelecastDeviceType int32

const (
	Telecast_UNKNOWN_DEVICE_TYPE TelecastDeviceType = 0
	Telecast_BGEIGIE_NANO        TelecastDeviceType = 1
	Telecast_SOLARCAST           TelecastDeviceType = 2
	Telecast_TTAPP               TelecastDeviceType = 3
	Telecast_TTNODE              TelecastDeviceType = 4
	Telecast_TTGATE              TelecastDeviceType = 5
	Telecast_TTSERVE             TelecastDeviceType = 6
)

var TelecastDeviceType_name = map[int32]string{
	0: "UNKNOWN_DEVICE_TYPE",
	1: "BGEIGIE_NANO",
	2: "SOLARCAST",
	3: "TTAPP",
	4: "TTNODE",
	5: "TTGATE",
	6: "TTSERVE",
}
var TelecastDeviceType_value = map[string]int32{
	"UNKNOWN_DEVICE_TYPE": 0,
	"BGEIGIE_NANO":        1,
	"SOLARCAST":           2,
	"TTAPP":               3,
	"TTNODE":              4,
	"TTGATE":              5,
	"TTSERVE":             6,
}

func (x TelecastDeviceType) Enum() *TelecastDeviceType {
	p := new(TelecastDeviceType)
	*p = x
	return p
}
func (x TelecastDeviceType) String() string {
	return proto.EnumName(TelecastDeviceType_name, int32(x))
}
func (x *TelecastDeviceType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(TelecastDeviceType_value, data, "TelecastDeviceType")
	if err != nil {
		return err
	}
	*x = TelecastDeviceType(value)
	return nil
}
func (TelecastDeviceType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type TelecastUnit int32

const (
	Telecast_UNKNOWN_UNIT TelecastUnit = 0
	Telecast_CPM          TelecastUnit = 1
)

var TelecastUnit_name = map[int32]string{
	0: "UNKNOWN_UNIT",
	1: "CPM",
}
var TelecastUnit_value = map[string]int32{
	"UNKNOWN_UNIT": 0,
	"CPM":          1,
}

func (x TelecastUnit) Enum() *TelecastUnit {
	p := new(TelecastUnit)
	*p = x
	return p
}
func (x TelecastUnit) String() string {
	return proto.EnumName(TelecastUnit_name, int32(x))
}
func (x *TelecastUnit) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(TelecastUnit_value, data, "TelecastUnit")
	if err != nil {
		return err
	}
	*x = TelecastUnit(value)
	return nil
}
func (TelecastUnit) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 1} }

type Telecast struct {
	DeviceType                      *TelecastDeviceType `protobuf:"varint,1,req,name=DeviceType,enum=teletype.TelecastDeviceType" json:"DeviceType,omitempty"`
	DEPRECATED2017FEBDeviceIDString *string             `protobuf:"bytes,2,opt,name=DEPRECATED2017FEBDeviceIDString" json:"DEPRECATED2017FEBDeviceIDString,omitempty"`
	DeviceID                        *uint32             `protobuf:"varint,3,opt,name=DeviceID" json:"DeviceID,omitempty"`
	Message                         *string             `protobuf:"bytes,4,opt,name=Message" json:"Message,omitempty"`
	CapturedAt                      *string             `protobuf:"bytes,5,opt,name=CapturedAt" json:"CapturedAt,omitempty"`
	DEPRECATED2017FEBUnit           *TelecastUnit       `protobuf:"varint,6,opt,name=DEPRECATED2017FEBUnit,enum=teletype.TelecastUnit" json:"DEPRECATED2017FEBUnit,omitempty"`
	DEPRECATED2017FEBValue          *uint32             `protobuf:"varint,7,opt,name=DEPRECATED2017FEBValue" json:"DEPRECATED2017FEBValue,omitempty"`
	Latitude                        *float32            `protobuf:"fixed32,8,opt,name=Latitude" json:"Latitude,omitempty"`
	Longitude                       *float32            `protobuf:"fixed32,9,opt,name=Longitude" json:"Longitude,omitempty"`
	Altitude                        *int32              `protobuf:"varint,10,opt,name=Altitude" json:"Altitude,omitempty"`
	BatteryVoltage                  *float32            `protobuf:"fixed32,11,opt,name=BatteryVoltage" json:"BatteryVoltage,omitempty"`
	BatterySOC                      *float32            `protobuf:"fixed32,12,opt,name=BatterySOC" json:"BatterySOC,omitempty"`
	WirelessSNR                     *float32            `protobuf:"fixed32,13,opt,name=WirelessSNR" json:"WirelessSNR,omitempty"`
	EnvTemperature                  *float32            `protobuf:"fixed32,14,opt,name=envTemperature" json:"envTemperature,omitempty"`
	EnvHumidity                     *float32            `protobuf:"fixed32,15,opt,name=envHumidity" json:"envHumidity,omitempty"`
	RelayDevice1                    *uint32             `protobuf:"varint,16,opt,name=RelayDevice1" json:"RelayDevice1,omitempty"`
	RelayDevice2                    *uint32             `protobuf:"varint,17,opt,name=RelayDevice2" json:"RelayDevice2,omitempty"`
	RelayDevice3                    *uint32             `protobuf:"varint,18,opt,name=RelayDevice3" json:"RelayDevice3,omitempty"`
	RelayDevice4                    *uint32             `protobuf:"varint,19,opt,name=RelayDevice4" json:"RelayDevice4,omitempty"`
	RelayDevice5                    *uint32             `protobuf:"varint,20,opt,name=RelayDevice5" json:"RelayDevice5,omitempty"`
	Cpm0                            *uint32             `protobuf:"varint,21,opt,name=cpm0" json:"cpm0,omitempty"`
	Cpm1                            *uint32             `protobuf:"varint,22,opt,name=cpm1" json:"cpm1,omitempty"`
	StatsUptimeMinutes              *uint32             `protobuf:"varint,23,opt,name=stats_uptime_minutes" json:"stats_uptime_minutes,omitempty"`
	StatsAppVersion                 *string             `protobuf:"bytes,24,opt,name=stats_app_version" json:"stats_app_version,omitempty"`
	StatsDeviceParams               *string             `protobuf:"bytes,25,opt,name=stats_device_params" json:"stats_device_params,omitempty"`
	StatsTransmittedBytes           *uint32             `protobuf:"varint,26,opt,name=stats_transmitted_bytes" json:"stats_transmitted_bytes,omitempty"`
	StatsReceivedBytes              *uint32             `protobuf:"varint,27,opt,name=stats_received_bytes" json:"stats_received_bytes,omitempty"`
	StatsOneshots                   *uint32             `protobuf:"varint,28,opt,name=stats_oneshots" json:"stats_oneshots,omitempty"`
	StatsCommsResets                *uint32             `protobuf:"varint,29,opt,name=stats_comms_resets" json:"stats_comms_resets,omitempty"`
	PmsPm01_0                       *uint32             `protobuf:"varint,30,opt,name=pms_pm01_0" json:"pms_pm01_0,omitempty"`
	PmsPm02_5                       *uint32             `protobuf:"varint,31,opt,name=pms_pm02_5" json:"pms_pm02_5,omitempty"`
	PmsPm10_0                       *uint32             `protobuf:"varint,32,opt,name=pms_pm10_0" json:"pms_pm10_0,omitempty"`
	PmsC00_30                       *uint32             `protobuf:"varint,33,opt,name=pms_c00_30" json:"pms_c00_30,omitempty"`
	PmsC00_50                       *uint32             `protobuf:"varint,34,opt,name=pms_c00_50" json:"pms_c00_50,omitempty"`
	PmsC01_00                       *uint32             `protobuf:"varint,35,opt,name=pms_c01_00" json:"pms_c01_00,omitempty"`
	PmsC02_50                       *uint32             `protobuf:"varint,36,opt,name=pms_c02_50" json:"pms_c02_50,omitempty"`
	PmsC05_00                       *uint32             `protobuf:"varint,37,opt,name=pms_c05_00" json:"pms_c05_00,omitempty"`
	PmsC10_00                       *uint32             `protobuf:"varint,38,opt,name=pms_c10_00" json:"pms_c10_00,omitempty"`
	PmsCsecs                        *uint32             `protobuf:"varint,39,opt,name=pms_csecs" json:"pms_csecs,omitempty"`
	OpcPm01_0                       *float32            `protobuf:"fixed32,40,opt,name=opc_pm01_0" json:"opc_pm01_0,omitempty"`
	OpcPm02_5                       *float32            `protobuf:"fixed32,41,opt,name=opc_pm02_5" json:"opc_pm02_5,omitempty"`
	OpcPm10_0                       *float32            `protobuf:"fixed32,42,opt,name=opc_pm10_0" json:"opc_pm10_0,omitempty"`
	OpcC00_38                       *uint32             `protobuf:"varint,43,opt,name=opc_c00_38" json:"opc_c00_38,omitempty"`
	OpcC00_54                       *uint32             `protobuf:"varint,44,opt,name=opc_c00_54" json:"opc_c00_54,omitempty"`
	OpcC01_00                       *uint32             `protobuf:"varint,45,opt,name=opc_c01_00" json:"opc_c01_00,omitempty"`
	OpcC02_10                       *uint32             `protobuf:"varint,46,opt,name=opc_c02_10" json:"opc_c02_10,omitempty"`
	OpcC05_00                       *uint32             `protobuf:"varint,47,opt,name=opc_c05_00" json:"opc_c05_00,omitempty"`
	OpcC10_00                       *uint32             `protobuf:"varint,48,opt,name=opc_c10_00" json:"opc_c10_00,omitempty"`
	OpcCsecs                        *uint32             `protobuf:"varint,49,opt,name=opc_csecs" json:"opc_csecs,omitempty"`
	EnvPressure                     *float32            `protobuf:"fixed32,50,opt,name=envPressure" json:"envPressure,omitempty"`
	StatsCommsPowerFails            *uint32             `protobuf:"varint,51,opt,name=stats_comms_power_fails" json:"stats_comms_power_fails,omitempty"`
	BatteryCurrent                  *float32            `protobuf:"fixed32,52,opt,name=BatteryCurrent" json:"BatteryCurrent,omitempty"`
	StatsCell                       *string             `protobuf:"bytes,53,opt,name=stats_cell" json:"stats_cell,omitempty"`
	StatsMotiondrops                *uint32             `protobuf:"varint,54,opt,name=stats_motiondrops" json:"stats_motiondrops,omitempty"`
	StatsDfu                        *string             `protobuf:"bytes,55,opt,name=stats_dfu" json:"stats_dfu,omitempty"`
	CapturedAtDate                  *uint32             `protobuf:"varint,56,opt,name=CapturedAtDate" json:"CapturedAtDate,omitempty"`
	CapturedAtTime                  *uint32             `protobuf:"varint,57,opt,name=CapturedAtTime" json:"CapturedAtTime,omitempty"`
	CapturedAtOffset                *uint32             `protobuf:"varint,58,opt,name=CapturedAtOffset" json:"CapturedAtOffset,omitempty"`
	StatsOneshotSeconds             *uint32             `protobuf:"varint,59,opt,name=stats_oneshot_seconds" json:"stats_oneshot_seconds,omitempty"`
	Stamp                           *uint32             `protobuf:"varint,60,opt,name=stamp" json:"stamp,omitempty"`
	StampVersion                    *uint32             `protobuf:"varint,61,opt,name=stamp_version" json:"stamp_version,omitempty"`
	XXX_unrecognized                []byte              `json:"-"`
}

func (m *Telecast) Reset()                    { *m = Telecast{} }
func (m *Telecast) String() string            { return proto.CompactTextString(m) }
func (*Telecast) ProtoMessage()               {}
func (*Telecast) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Telecast) GetDeviceType() TelecastDeviceType {
	if m != nil && m.DeviceType != nil {
		return *m.DeviceType
	}
	return Telecast_UNKNOWN_DEVICE_TYPE
}

func (m *Telecast) GetDEPRECATED2017FEBDeviceIDString() string {
	if m != nil && m.DEPRECATED2017FEBDeviceIDString != nil {
		return *m.DEPRECATED2017FEBDeviceIDString
	}
	return ""
}

func (m *Telecast) GetDeviceID() uint32 {
	if m != nil && m.DeviceID != nil {
		return *m.DeviceID
	}
	return 0
}

func (m *Telecast) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

func (m *Telecast) GetCapturedAt() string {
	if m != nil && m.CapturedAt != nil {
		return *m.CapturedAt
	}
	return ""
}

func (m *Telecast) GetDEPRECATED2017FEBUnit() TelecastUnit {
	if m != nil && m.DEPRECATED2017FEBUnit != nil {
		return *m.DEPRECATED2017FEBUnit
	}
	return Telecast_UNKNOWN_UNIT
}

func (m *Telecast) GetDEPRECATED2017FEBValue() uint32 {
	if m != nil && m.DEPRECATED2017FEBValue != nil {
		return *m.DEPRECATED2017FEBValue
	}
	return 0
}

func (m *Telecast) GetLatitude() float32 {
	if m != nil && m.Latitude != nil {
		return *m.Latitude
	}
	return 0
}

func (m *Telecast) GetLongitude() float32 {
	if m != nil && m.Longitude != nil {
		return *m.Longitude
	}
	return 0
}

func (m *Telecast) GetAltitude() int32 {
	if m != nil && m.Altitude != nil {
		return *m.Altitude
	}
	return 0
}

func (m *Telecast) GetBatteryVoltage() float32 {
	if m != nil && m.BatteryVoltage != nil {
		return *m.BatteryVoltage
	}
	return 0
}

func (m *Telecast) GetBatterySOC() float32 {
	if m != nil && m.BatterySOC != nil {
		return *m.BatterySOC
	}
	return 0
}

func (m *Telecast) GetWirelessSNR() float32 {
	if m != nil && m.WirelessSNR != nil {
		return *m.WirelessSNR
	}
	return 0
}

func (m *Telecast) GetEnvTemperature() float32 {
	if m != nil && m.EnvTemperature != nil {
		return *m.EnvTemperature
	}
	return 0
}

func (m *Telecast) GetEnvHumidity() float32 {
	if m != nil && m.EnvHumidity != nil {
		return *m.EnvHumidity
	}
	return 0
}

func (m *Telecast) GetRelayDevice1() uint32 {
	if m != nil && m.RelayDevice1 != nil {
		return *m.RelayDevice1
	}
	return 0
}

func (m *Telecast) GetRelayDevice2() uint32 {
	if m != nil && m.RelayDevice2 != nil {
		return *m.RelayDevice2
	}
	return 0
}

func (m *Telecast) GetRelayDevice3() uint32 {
	if m != nil && m.RelayDevice3 != nil {
		return *m.RelayDevice3
	}
	return 0
}

func (m *Telecast) GetRelayDevice4() uint32 {
	if m != nil && m.RelayDevice4 != nil {
		return *m.RelayDevice4
	}
	return 0
}

func (m *Telecast) GetRelayDevice5() uint32 {
	if m != nil && m.RelayDevice5 != nil {
		return *m.RelayDevice5
	}
	return 0
}

func (m *Telecast) GetCpm0() uint32 {
	if m != nil && m.Cpm0 != nil {
		return *m.Cpm0
	}
	return 0
}

func (m *Telecast) GetCpm1() uint32 {
	if m != nil && m.Cpm1 != nil {
		return *m.Cpm1
	}
	return 0
}

func (m *Telecast) GetStatsUptimeMinutes() uint32 {
	if m != nil && m.StatsUptimeMinutes != nil {
		return *m.StatsUptimeMinutes
	}
	return 0
}

func (m *Telecast) GetStatsAppVersion() string {
	if m != nil && m.StatsAppVersion != nil {
		return *m.StatsAppVersion
	}
	return ""
}

func (m *Telecast) GetStatsDeviceParams() string {
	if m != nil && m.StatsDeviceParams != nil {
		return *m.StatsDeviceParams
	}
	return ""
}

func (m *Telecast) GetStatsTransmittedBytes() uint32 {
	if m != nil && m.StatsTransmittedBytes != nil {
		return *m.StatsTransmittedBytes
	}
	return 0
}

func (m *Telecast) GetStatsReceivedBytes() uint32 {
	if m != nil && m.StatsReceivedBytes != nil {
		return *m.StatsReceivedBytes
	}
	return 0
}

func (m *Telecast) GetStatsOneshots() uint32 {
	if m != nil && m.StatsOneshots != nil {
		return *m.StatsOneshots
	}
	return 0
}

func (m *Telecast) GetStatsCommsResets() uint32 {
	if m != nil && m.StatsCommsResets != nil {
		return *m.StatsCommsResets
	}
	return 0
}

func (m *Telecast) GetPmsPm01_0() uint32 {
	if m != nil && m.PmsPm01_0 != nil {
		return *m.PmsPm01_0
	}
	return 0
}

func (m *Telecast) GetPmsPm02_5() uint32 {
	if m != nil && m.PmsPm02_5 != nil {
		return *m.PmsPm02_5
	}
	return 0
}

func (m *Telecast) GetPmsPm10_0() uint32 {
	if m != nil && m.PmsPm10_0 != nil {
		return *m.PmsPm10_0
	}
	return 0
}

func (m *Telecast) GetPmsC00_30() uint32 {
	if m != nil && m.PmsC00_30 != nil {
		return *m.PmsC00_30
	}
	return 0
}

func (m *Telecast) GetPmsC00_50() uint32 {
	if m != nil && m.PmsC00_50 != nil {
		return *m.PmsC00_50
	}
	return 0
}

func (m *Telecast) GetPmsC01_00() uint32 {
	if m != nil && m.PmsC01_00 != nil {
		return *m.PmsC01_00
	}
	return 0
}

func (m *Telecast) GetPmsC02_50() uint32 {
	if m != nil && m.PmsC02_50 != nil {
		return *m.PmsC02_50
	}
	return 0
}

func (m *Telecast) GetPmsC05_00() uint32 {
	if m != nil && m.PmsC05_00 != nil {
		return *m.PmsC05_00
	}
	return 0
}

func (m *Telecast) GetPmsC10_00() uint32 {
	if m != nil && m.PmsC10_00 != nil {
		return *m.PmsC10_00
	}
	return 0
}

func (m *Telecast) GetPmsCsecs() uint32 {
	if m != nil && m.PmsCsecs != nil {
		return *m.PmsCsecs
	}
	return 0
}

func (m *Telecast) GetOpcPm01_0() float32 {
	if m != nil && m.OpcPm01_0 != nil {
		return *m.OpcPm01_0
	}
	return 0
}

func (m *Telecast) GetOpcPm02_5() float32 {
	if m != nil && m.OpcPm02_5 != nil {
		return *m.OpcPm02_5
	}
	return 0
}

func (m *Telecast) GetOpcPm10_0() float32 {
	if m != nil && m.OpcPm10_0 != nil {
		return *m.OpcPm10_0
	}
	return 0
}

func (m *Telecast) GetOpcC00_38() uint32 {
	if m != nil && m.OpcC00_38 != nil {
		return *m.OpcC00_38
	}
	return 0
}

func (m *Telecast) GetOpcC00_54() uint32 {
	if m != nil && m.OpcC00_54 != nil {
		return *m.OpcC00_54
	}
	return 0
}

func (m *Telecast) GetOpcC01_00() uint32 {
	if m != nil && m.OpcC01_00 != nil {
		return *m.OpcC01_00
	}
	return 0
}

func (m *Telecast) GetOpcC02_10() uint32 {
	if m != nil && m.OpcC02_10 != nil {
		return *m.OpcC02_10
	}
	return 0
}

func (m *Telecast) GetOpcC05_00() uint32 {
	if m != nil && m.OpcC05_00 != nil {
		return *m.OpcC05_00
	}
	return 0
}

func (m *Telecast) GetOpcC10_00() uint32 {
	if m != nil && m.OpcC10_00 != nil {
		return *m.OpcC10_00
	}
	return 0
}

func (m *Telecast) GetOpcCsecs() uint32 {
	if m != nil && m.OpcCsecs != nil {
		return *m.OpcCsecs
	}
	return 0
}

func (m *Telecast) GetEnvPressure() float32 {
	if m != nil && m.EnvPressure != nil {
		return *m.EnvPressure
	}
	return 0
}

func (m *Telecast) GetStatsCommsPowerFails() uint32 {
	if m != nil && m.StatsCommsPowerFails != nil {
		return *m.StatsCommsPowerFails
	}
	return 0
}

func (m *Telecast) GetBatteryCurrent() float32 {
	if m != nil && m.BatteryCurrent != nil {
		return *m.BatteryCurrent
	}
	return 0
}

func (m *Telecast) GetStatsCell() string {
	if m != nil && m.StatsCell != nil {
		return *m.StatsCell
	}
	return ""
}

func (m *Telecast) GetStatsMotiondrops() uint32 {
	if m != nil && m.StatsMotiondrops != nil {
		return *m.StatsMotiondrops
	}
	return 0
}

func (m *Telecast) GetStatsDfu() string {
	if m != nil && m.StatsDfu != nil {
		return *m.StatsDfu
	}
	return ""
}

func (m *Telecast) GetCapturedAtDate() uint32 {
	if m != nil && m.CapturedAtDate != nil {
		return *m.CapturedAtDate
	}
	return 0
}

func (m *Telecast) GetCapturedAtTime() uint32 {
	if m != nil && m.CapturedAtTime != nil {
		return *m.CapturedAtTime
	}
	return 0
}

func (m *Telecast) GetCapturedAtOffset() uint32 {
	if m != nil && m.CapturedAtOffset != nil {
		return *m.CapturedAtOffset
	}
	return 0
}

func (m *Telecast) GetStatsOneshotSeconds() uint32 {
	if m != nil && m.StatsOneshotSeconds != nil {
		return *m.StatsOneshotSeconds
	}
	return 0
}

func (m *Telecast) GetStamp() uint32 {
	if m != nil && m.Stamp != nil {
		return *m.Stamp
	}
	return 0
}

func (m *Telecast) GetStampVersion() uint32 {
	if m != nil && m.StampVersion != nil {
		return *m.StampVersion
	}
	return 0
}

func init() {
	proto.RegisterType((*Telecast)(nil), "teletype.Telecast")
	proto.RegisterEnum("teletype.TelecastDeviceType", TelecastDeviceType_name, TelecastDeviceType_value)
	proto.RegisterEnum("teletype.TelecastUnit", TelecastUnit_name, TelecastUnit_value)
}

var fileDescriptor0 = []byte{
	// 833 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x94, 0x59, 0x53, 0xe3, 0x46,
	0x10, 0xc7, 0xc3, 0x61, 0xc0, 0x8d, 0xf1, 0x0a, 0x71, 0xf5, 0x1e, 0x2c, 0x2c, 0x39, 0x96, 0x5c,
	0x44, 0x36, 0x78, 0x77, 0x73, 0x3d, 0x18, 0x5b, 0x21, 0xae, 0xb0, 0xb6, 0xcb, 0x16, 0x6c, 0xe5,
	0x49, 0xa5, 0xd8, 0x03, 0x51, 0x95, 0xae, 0xd2, 0x8c, 0x9c, 0xf2, 0x07, 0xcb, 0xf7, 0xdb, 0x56,
	0xcb, 0x32, 0xd6, 0x2e, 0x6f, 0x9a, 0xdf, 0xbf, 0xd5, 0xdd, 0xff, 0x99, 0x9e, 0x81, 0xaa, 0x12,
	0x9e, 0x50, 0xd3, 0x48, 0x9c, 0x45, 0x71, 0xa8, 0x42, 0x7d, 0x23, 0x5f, 0x9f, 0xfc, 0x5f, 0x81,
	0x0d, 0x8b, 0x16, 0x23, 0x47, 0x2a, 0xbd, 0x06, 0xd0, 0x16, 0x13, 0x77, 0x24, 0x2c, 0x92, 0x70,
	0xe9, 0x78, 0xf9, 0xb4, 0x5a, 0x3f, 0x3c, 0x9b, 0xff, 0x9b, 0xc7, 0x9d, 0x8d, 0xe7, 0x41, 0xfa,
	0x6b, 0x38, 0x6a, 0x9b, 0xfd, 0x81, 0xd9, 0x6a, 0x5a, 0x66, 0xbb, 0x6e, 0xd4, 0xde, 0xfe, 0x61,
	0x5e, 0x66, 0x39, 0x3a, 0xed, 0xa1, 0x8a, 0xdd, 0xe0, 0x1e, 0x97, 0x8f, 0x97, 0x4e, 0xcb, 0xba,
	0x06, 0x1b, 0x39, 0xc7, 0x15, 0x22, 0x5b, 0xfa, 0x13, 0x58, 0x7f, 0x2f, 0xa4, 0x74, 0xee, 0x05,
	0xae, 0x72, 0x88, 0x0e, 0xd0, 0x72, 0x22, 0x95, 0xc4, 0x62, 0xdc, 0x54, 0x58, 0x62, 0xf6, 0x06,
	0xf6, 0x3e, 0xcb, 0x7f, 0x13, 0xb8, 0x0a, 0xd7, 0x48, 0xae, 0xd6, 0x0f, 0x1e, 0xe9, 0x2e, 0x21,
	0x59, 0x7f, 0x09, 0xfb, 0x9f, 0xfd, 0x77, 0xeb, 0x78, 0x89, 0xc0, 0x75, 0x2e, 0x4e, 0xed, 0x5c,
	0x3b, 0xca, 0x55, 0xc9, 0x58, 0xe0, 0x06, 0x91, 0x65, 0x7d, 0x1b, 0xca, 0xd7, 0x61, 0x70, 0x9f,
	0xa1, 0x32, 0x23, 0x0a, 0x6a, 0x7a, 0xb3, 0x20, 0x20, 0x52, 0xd2, 0xf7, 0xa1, 0x7a, 0xe9, 0x28,
	0x25, 0xe2, 0xe9, 0x6d, 0xe8, 0xa9, 0xb4, 0xf5, 0x4d, 0x8e, 0xa4, 0xd6, 0x67, 0x7c, 0xd8, 0x6b,
	0x61, 0x85, 0xd9, 0x0e, 0x6c, 0x7e, 0x70, 0x63, 0xea, 0x4a, 0xca, 0x61, 0x77, 0x80, 0x5b, 0x0c,
	0x29, 0x81, 0x08, 0x26, 0x96, 0xf0, 0x23, 0x11, 0x3b, 0xa9, 0x55, 0xac, 0xe6, 0xc1, 0xc4, 0xff,
	0x4c, 0x7c, 0x77, 0xec, 0xaa, 0x29, 0x3e, 0x61, 0xb8, 0x0b, 0x95, 0x81, 0xf0, 0x9c, 0x69, 0xb6,
	0x71, 0x35, 0xd4, 0xb8, 0xf5, 0x22, 0xad, 0xe3, 0xf6, 0x23, 0xf4, 0x1c, 0xf5, 0x47, 0xe8, 0x05,
	0xee, 0x3c, 0x42, 0x1b, 0xb8, 0xcb, 0xb4, 0x02, 0xab, 0xa3, 0xc8, 0x37, 0x70, 0x6f, 0x61, 0x55,
	0xc3, 0x7d, 0x5e, 0xbd, 0x80, 0x5d, 0xa9, 0x1c, 0x25, 0xed, 0x24, 0x52, 0xae, 0x2f, 0x6c, 0xdf,
	0x0d, 0x12, 0x25, 0x24, 0x1e, 0xb0, 0xfa, 0x14, 0xb6, 0x33, 0xd5, 0x89, 0x22, 0x7b, 0x22, 0x62,
	0xe9, 0x86, 0x01, 0x22, 0x9f, 0xdf, 0x73, 0xd8, 0xc9, 0xa4, 0x6c, 0x66, 0xec, 0xc8, 0x89, 0x1d,
	0x5f, 0xe2, 0x53, 0x16, 0x8f, 0xe0, 0x20, 0x13, 0x55, 0xec, 0x04, 0xd2, 0x77, 0x69, 0x03, 0xc7,
	0xf6, 0x3f, 0xd3, 0x34, 0xf1, 0xb3, 0x62, 0xd9, 0x58, 0x8c, 0x84, 0x3b, 0x99, 0xab, 0xcf, 0x59,
	0xa5, 0xbd, 0xcc, 0xd4, 0x30, 0x10, 0xf2, 0xdf, 0x50, 0x49, 0x7c, 0xc1, 0xfc, 0x19, 0xe8, 0x19,
	0x1f, 0x85, 0xbe, 0x9f, 0xfe, 0x2b, 0x05, 0x69, 0x87, 0xac, 0xd1, 0x41, 0x45, 0xc4, 0xc8, 0x67,
	0xcd, 0x36, 0xf0, 0xe5, 0x27, 0xac, 0x6e, 0x37, 0xf0, 0xa8, 0xc8, 0x6a, 0x06, 0xc5, 0x1d, 0x2f,
	0xb2, 0x91, 0x61, 0xd8, 0xe7, 0x06, 0xbe, 0xfa, 0x94, 0x35, 0x0c, 0x3c, 0x29, 0x32, 0x2a, 0x61,
	0xe0, 0x97, 0x45, 0x56, 0x4f, 0xe3, 0xbe, 0x2a, 0xb2, 0x46, 0x1a, 0xf7, 0x75, 0x81, 0xa5, 0x65,
	0x0d, 0xfc, 0x86, 0x19, 0x4d, 0x26, 0x33, 0x29, 0x46, 0x12, 0x5f, 0xe7, 0x61, 0x61, 0x34, 0xca,
	0x6d, 0x9c, 0xe6, 0x33, 0x38, 0x63, 0xa9, 0x8d, 0x6f, 0x8b, 0x8c, 0x6d, 0x7c, 0xb7, 0xc8, 0xd8,
	0xc6, 0x3b, 0xfc, 0x7e, 0x31, 0x1f, 0xdb, 0xb8, 0xc0, 0x1f, 0x8a, 0x8c, 0x6d, 0xfc, 0x58, 0x64,
	0x75, 0xbb, 0x66, 0xe0, 0x59, 0x91, 0xb1, 0x8d, 0x9f, 0x0a, 0x2c, 0xb3, 0x61, 0xe4, 0x36, 0x98,
	0xb1, 0x8d, 0x1a, 0xa3, 0x6c, 0xea, 0xfb, 0x74, 0x42, 0x32, 0xbd, 0x0a, 0x75, 0xee, 0x6f, 0x3e,
	0x15, 0xd9, 0xf1, 0x45, 0xe1, 0x7f, 0x22, 0xb6, 0xef, 0x1c, 0xd7, 0x93, 0x78, 0x9e, 0x9f, 0xfb,
	0xec, 0xb2, 0xb5, 0x92, 0x38, 0x16, 0x81, 0xc2, 0x8b, 0xdc, 0xd8, 0xec, 0x47, 0xe1, 0x79, 0xd8,
	0xe0, 0x11, 0x9b, 0x8f, 0xa6, 0x1f, 0x2a, 0x9a, 0xca, 0x71, 0x1c, 0x46, 0x12, 0xdf, 0xe4, 0xfd,
	0xcc, 0x46, 0xf3, 0x2e, 0xc1, 0xb7, 0x1c, 0x4d, 0x99, 0x1f, 0x5e, 0xa0, 0xb6, 0xa3, 0x04, 0xbe,
	0xcb, 0x2b, 0x3e, 0x70, 0x8b, 0x2e, 0x00, 0xfe, 0xcc, 0x1c, 0x41, 0x7b, 0xe0, 0xbd, 0xbb, 0x3b,
	0x1a, 0x34, 0xfc, 0x85, 0x95, 0x43, 0xd8, 0x2b, 0xcc, 0xa6, 0x4d, 0xae, 0xa9, 0xbc, 0xc4, 0x5f,
	0x59, 0xde, 0x82, 0x12, 0xc9, 0x7e, 0x84, 0xbf, 0xf1, 0x72, 0x0f, 0xb6, 0x78, 0x39, 0xbf, 0x3c,
	0xbf, 0xa7, 0xf8, 0x64, 0x02, 0xb0, 0xf0, 0xd4, 0x1e, 0xc0, 0xce, 0x4d, 0xf7, 0xaf, 0x6e, 0xef,
	0x43, 0xd7, 0x6e, 0x9b, 0xb7, 0x9d, 0x96, 0x69, 0x5b, 0x7f, 0xf7, 0x4d, 0xed, 0x0b, 0x7a, 0xa6,
	0x2a, 0x97, 0x57, 0x66, 0xe7, 0xaa, 0x63, 0xda, 0xdd, 0x66, 0xb7, 0xa7, 0x2d, 0x51, 0xfa, 0xf2,
	0xb0, 0x77, 0xdd, 0x1c, 0xb4, 0x9a, 0x43, 0x4b, 0x5b, 0xd6, 0xcb, 0x50, 0xb2, 0xac, 0x66, 0xbf,
	0xaf, 0xad, 0xd0, 0x16, 0xad, 0x59, 0x56, 0xb7, 0xd7, 0x36, 0xb5, 0xd5, 0xec, 0xfb, 0x8a, 0x1e,
	0x48, 0xad, 0xa4, 0x6f, 0xc2, 0xba, 0x65, 0x0d, 0xcd, 0xc1, 0xad, 0xa9, 0xad, 0x9d, 0xbc, 0x82,
	0x55, 0x7e, 0x44, 0x29, 0x71, 0x5e, 0xf1, 0xa6, 0xdb, 0xb1, 0xa8, 0xd4, 0x3a, 0xac, 0xb4, 0xfa,
	0xef, 0xb5, 0xa5, 0x8f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x77, 0xa7, 0x41, 0x30, 0x52, 0x06, 0x00,
	0x00,
}
