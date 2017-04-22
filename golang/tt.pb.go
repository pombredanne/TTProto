// Code generated by protoc-gen-go.
// source: tt.proto
// DO NOT EDIT!

/*
Package ttproto is a generated protocol buffer package.

It is generated from these files:
	tt.proto

It has these top-level messages:
	Telecast
*/
package ttproto

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

type TelecastReplyType int32

const (
	Telecast_NO_REPLY TelecastReplyType = 0
	Telecast_ALLOWED  TelecastReplyType = 1
)

var TelecastReplyType_name = map[int32]string{
	0: "NO_REPLY",
	1: "ALLOWED",
}
var TelecastReplyType_value = map[string]int32{
	"NO_REPLY": 0,
	"ALLOWED":  1,
}

func (x TelecastReplyType) Enum() *TelecastReplyType {
	p := new(TelecastReplyType)
	*p = x
	return p
}
func (x TelecastReplyType) String() string {
	return proto.EnumName(TelecastReplyType_name, int32(x))
}
func (x *TelecastReplyType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(TelecastReplyType_value, data, "TelecastReplyType")
	if err != nil {
		return err
	}
	*x = TelecastReplyType(value)
	return nil
}
func (TelecastReplyType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 1} }

type Telecast struct {
	DeviceType                      *TelecastDeviceType `protobuf:"varint,1,opt,name=device_type,enum=ttproto.TelecastDeviceType" json:"device_type,omitempty"`
	DEPRECATED2017FEBDeviceIDString *string             `protobuf:"bytes,2,opt,name=DEPRECATED2017FEBDeviceIDString" json:"DEPRECATED2017FEBDeviceIDString,omitempty"`
	// This field is actually REQUIRED by ttserve
	DeviceId               *uint32            `protobuf:"varint,3,opt,name=device_id" json:"device_id,omitempty"`
	Message                *string            `protobuf:"bytes,4,opt,name=message" json:"message,omitempty"`
	CapturedAt             *string            `protobuf:"bytes,5,opt,name=captured_at" json:"captured_at,omitempty"`
	ReplyType              *TelecastReplyType `protobuf:"varint,6,opt,name=reply_type,enum=ttproto.TelecastReplyType" json:"reply_type,omitempty"`
	DEPRECATED2017FEBValue *uint32            `protobuf:"varint,7,opt,name=DEPRECATED2017FEBValue" json:"DEPRECATED2017FEBValue,omitempty"`
	Latitude               *float32           `protobuf:"fixed32,8,opt,name=latitude" json:"latitude,omitempty"`
	Longitude              *float32           `protobuf:"fixed32,9,opt,name=longitude" json:"longitude,omitempty"`
	Altitude               *int32             `protobuf:"varint,10,opt,name=altitude" json:"altitude,omitempty"`
	BatVoltage             *float32           `protobuf:"fixed32,11,opt,name=bat_voltage" json:"bat_voltage,omitempty"`
	BatSoc                 *float32           `protobuf:"fixed32,12,opt,name=bat_soc" json:"bat_soc,omitempty"`
	WirelessSnr            *float32           `protobuf:"fixed32,13,opt,name=wireless_snr" json:"wireless_snr,omitempty"`
	EnvTemp                *float32           `protobuf:"fixed32,14,opt,name=env_temp" json:"env_temp,omitempty"`
	EnvHumid               *float32           `protobuf:"fixed32,15,opt,name=env_humid" json:"env_humid,omitempty"`
	RelayDevice1           *uint32            `protobuf:"varint,16,opt,name=relay_device1" json:"relay_device1,omitempty"`
	RelayDevice2           *uint32            `protobuf:"varint,17,opt,name=relay_device2" json:"relay_device2,omitempty"`
	RelayDevice3           *uint32            `protobuf:"varint,18,opt,name=relay_device3" json:"relay_device3,omitempty"`
	RelayDevice4           *uint32            `protobuf:"varint,19,opt,name=relay_device4" json:"relay_device4,omitempty"`
	RelayDevice5           *uint32            `protobuf:"varint,20,opt,name=relay_device5" json:"relay_device5,omitempty"`
	Lnd_7318U              *uint32            `protobuf:"varint,21,opt,name=lnd_7318u" json:"lnd_7318u,omitempty"`
	Lnd_7128Ec             *uint32            `protobuf:"varint,22,opt,name=lnd_7128ec" json:"lnd_7128ec,omitempty"`
	StatsUptimeMinutes     *uint32            `protobuf:"varint,23,opt,name=stats_uptime_minutes" json:"stats_uptime_minutes,omitempty"`
	StatsAppVersion        *string            `protobuf:"bytes,24,opt,name=stats_app_version" json:"stats_app_version,omitempty"`
	StatsDeviceParams      *string            `protobuf:"bytes,25,opt,name=stats_device_params" json:"stats_device_params,omitempty"`
	StatsTransmittedBytes  *uint32            `protobuf:"varint,26,opt,name=stats_transmitted_bytes" json:"stats_transmitted_bytes,omitempty"`
	StatsReceivedBytes     *uint32            `protobuf:"varint,27,opt,name=stats_received_bytes" json:"stats_received_bytes,omitempty"`
	StatsOneshots          *uint32            `protobuf:"varint,28,opt,name=stats_oneshots" json:"stats_oneshots,omitempty"`
	StatsCommsResets       *uint32            `protobuf:"varint,29,opt,name=stats_comms_resets" json:"stats_comms_resets,omitempty"`
	PmsPm01_0              *uint32            `protobuf:"varint,30,opt,name=pms_pm01_0" json:"pms_pm01_0,omitempty"`
	PmsPm02_5              *uint32            `protobuf:"varint,31,opt,name=pms_pm02_5" json:"pms_pm02_5,omitempty"`
	PmsPm10_0              *uint32            `protobuf:"varint,32,opt,name=pms_pm10_0" json:"pms_pm10_0,omitempty"`
	PmsC00_30              *uint32            `protobuf:"varint,33,opt,name=pms_c00_30" json:"pms_c00_30,omitempty"`
	PmsC00_50              *uint32            `protobuf:"varint,34,opt,name=pms_c00_50" json:"pms_c00_50,omitempty"`
	PmsC01_00              *uint32            `protobuf:"varint,35,opt,name=pms_c01_00" json:"pms_c01_00,omitempty"`
	PmsC02_50              *uint32            `protobuf:"varint,36,opt,name=pms_c02_50" json:"pms_c02_50,omitempty"`
	PmsC05_00              *uint32            `protobuf:"varint,37,opt,name=pms_c05_00" json:"pms_c05_00,omitempty"`
	PmsC10_00              *uint32            `protobuf:"varint,38,opt,name=pms_c10_00" json:"pms_c10_00,omitempty"`
	PmsCsecs               *uint32            `protobuf:"varint,39,opt,name=pms_csecs" json:"pms_csecs,omitempty"`
	OpcPm01_0              *float32           `protobuf:"fixed32,40,opt,name=opc_pm01_0" json:"opc_pm01_0,omitempty"`
	OpcPm02_5              *float32           `protobuf:"fixed32,41,opt,name=opc_pm02_5" json:"opc_pm02_5,omitempty"`
	OpcPm10_0              *float32           `protobuf:"fixed32,42,opt,name=opc_pm10_0" json:"opc_pm10_0,omitempty"`
	OpcC00_38              *uint32            `protobuf:"varint,43,opt,name=opc_c00_38" json:"opc_c00_38,omitempty"`
	OpcC00_54              *uint32            `protobuf:"varint,44,opt,name=opc_c00_54" json:"opc_c00_54,omitempty"`
	OpcC01_00              *uint32            `protobuf:"varint,45,opt,name=opc_c01_00" json:"opc_c01_00,omitempty"`
	OpcC02_10              *uint32            `protobuf:"varint,46,opt,name=opc_c02_10" json:"opc_c02_10,omitempty"`
	OpcC05_00              *uint32            `protobuf:"varint,47,opt,name=opc_c05_00" json:"opc_c05_00,omitempty"`
	OpcC10_00              *uint32            `protobuf:"varint,48,opt,name=opc_c10_00" json:"opc_c10_00,omitempty"`
	OpcCsecs               *uint32            `protobuf:"varint,49,opt,name=opc_csecs" json:"opc_csecs,omitempty"`
	EnvPressure            *float32           `protobuf:"fixed32,50,opt,name=env_pressure" json:"env_pressure,omitempty"`
	StatsCommsPowerFails   *uint32            `protobuf:"varint,51,opt,name=stats_comms_power_fails" json:"stats_comms_power_fails,omitempty"`
	BatCurrent             *float32           `protobuf:"fixed32,52,opt,name=bat_current" json:"bat_current,omitempty"`
	StatsIccid             *string            `protobuf:"bytes,53,opt,name=stats_iccid" json:"stats_iccid,omitempty"`
	StatsMotionEvents      *uint32            `protobuf:"varint,54,opt,name=stats_motion_events" json:"stats_motion_events,omitempty"`
	StatsDfu               *string            `protobuf:"bytes,55,opt,name=stats_dfu" json:"stats_dfu,omitempty"`
	CapturedAtDate         *uint32            `protobuf:"varint,56,opt,name=captured_at_date" json:"captured_at_date,omitempty"`
	CapturedAtTime         *uint32            `protobuf:"varint,57,opt,name=captured_at_time" json:"captured_at_time,omitempty"`
	CapturedAtOffset       *uint32            `protobuf:"varint,58,opt,name=captured_at_offset" json:"captured_at_offset,omitempty"`
	StatsOneshotSeconds    *uint32            `protobuf:"varint,59,opt,name=stats_oneshot_seconds" json:"stats_oneshot_seconds,omitempty"`
	Stamp                  *uint32            `protobuf:"varint,60,opt,name=stamp" json:"stamp,omitempty"`
	StampVersion           *uint32            `protobuf:"varint,61,opt,name=stamp_version" json:"stamp_version,omitempty"`
	StatsCpsi              *string            `protobuf:"bytes,62,opt,name=stats_cpsi" json:"stats_cpsi,omitempty"`
	StatsUptimeDays        *uint32            `protobuf:"varint,63,opt,name=stats_uptime_days" json:"stats_uptime_days,omitempty"`
	StatsDeviceLabel       *string            `protobuf:"bytes,64,opt,name=stats_device_label" json:"stats_device_label,omitempty"`
	StatsGpsParams         *string            `protobuf:"bytes,65,opt,name=stats_gps_params" json:"stats_gps_params,omitempty"`
	StatsServiceParams     *string            `protobuf:"bytes,66,opt,name=stats_service_params" json:"stats_service_params,omitempty"`
	StatsTtnParams         *string            `protobuf:"bytes,67,opt,name=stats_ttn_params" json:"stats_ttn_params,omitempty"`
	StatsSensorParams      *string            `protobuf:"bytes,68,opt,name=stats_sensor_params" json:"stats_sensor_params,omitempty"`
	Lnd_7318C              *uint32            `protobuf:"varint,69,opt,name=lnd_7318c" json:"lnd_7318c,omitempty"`
	StatsBattery           *string            `protobuf:"bytes,70,opt,name=stats_battery" json:"stats_battery,omitempty"`
	StatsModuleFona        *string            `protobuf:"bytes,71,opt,name=stats_module_fona" json:"stats_module_fona,omitempty"`
	StatsModuleLora        *string            `protobuf:"bytes,72,opt,name=stats_module_lora" json:"stats_module_lora,omitempty"`
	Motion                 *bool              `protobuf:"varint,73,opt,name=motion" json:"motion,omitempty"`
	Test                   *bool              `protobuf:"varint,74,opt,name=test" json:"test,omitempty"`
	EncTemp                *float32           `protobuf:"fixed32,75,opt,name=enc_temp" json:"enc_temp,omitempty"`
	EncHumid               *float32           `protobuf:"fixed32,76,opt,name=enc_humid" json:"enc_humid,omitempty"`
	EncPressure            *float32           `protobuf:"fixed32,77,opt,name=enc_pressure" json:"enc_pressure,omitempty"`
	ErrorsOpc              *uint32            `protobuf:"varint,78,opt,name=errors_opc" json:"errors_opc,omitempty"`
	ErrorsPms              *uint32            `protobuf:"varint,79,opt,name=errors_pms" json:"errors_pms,omitempty"`
	ErrorsBme0             *uint32            `protobuf:"varint,80,opt,name=errors_bme0" json:"errors_bme0,omitempty"`
	ErrorsBme1             *uint32            `protobuf:"varint,81,opt,name=errors_bme1" json:"errors_bme1,omitempty"`
	ErrorsLora             *uint32            `protobuf:"varint,82,opt,name=errors_lora" json:"errors_lora,omitempty"`
	ErrorsFona             *uint32            `protobuf:"varint,83,opt,name=errors_fona" json:"errors_fona,omitempty"`
	ErrorsGeiger           *uint32            `protobuf:"varint,84,opt,name=errors_geiger" json:"errors_geiger,omitempty"`
	ErrorsMax01            *uint32            `protobuf:"varint,85,opt,name=errors_max01" json:"errors_max01,omitempty"`
	ErrorsUgps             *uint32            `protobuf:"varint,86,opt,name=errors_ugps" json:"errors_ugps,omitempty"`
	ErrorsTwi              *uint32            `protobuf:"varint,87,opt,name=errors_twi" json:"errors_twi,omitempty"`
	ErrorsTwiInfo          *string            `protobuf:"bytes,88,opt,name=errors_twi_info" json:"errors_twi_info,omitempty"`
	ErrorsLis              *uint32            `protobuf:"varint,89,opt,name=errors_lis" json:"errors_lis,omitempty"`
	ErrorsSpi              *uint32            `protobuf:"varint,90,opt,name=errors_spi" json:"errors_spi,omitempty"`
	ErrorsConnectLora      *uint32            `protobuf:"varint,91,opt,name=errors_connect_lora" json:"errors_connect_lora,omitempty"`
	ErrorsConnectFona      *uint32            `protobuf:"varint,92,opt,name=errors_connect_fona" json:"errors_connect_fona,omitempty"`
	ErrorsConnectWireless  *uint32            `protobuf:"varint,93,opt,name=errors_connect_wireless" json:"errors_connect_wireless,omitempty"`
	ErrorsConnectData      *uint32            `protobuf:"varint,94,opt,name=errors_connect_data" json:"errors_connect_data,omitempty"`
	ErrorsConnectService   *uint32            `protobuf:"varint,95,opt,name=errors_connect_service" json:"errors_connect_service,omitempty"`
	MotionBeganOffset      *uint32            `protobuf:"varint,96,opt,name=motion_began_offset" json:"motion_began_offset,omitempty"`
	ErrorsConnectGateway   *uint32            `protobuf:"varint,97,opt,name=errors_connect_gateway" json:"errors_connect_gateway,omitempty"`
	StatsCommsAntFails     *uint32            `protobuf:"varint,98,opt,name=stats_comms_ant_fails" json:"stats_comms_ant_fails,omitempty"`
	Lnd_712U               *uint32            `protobuf:"varint,99,opt,name=lnd_712u" json:"lnd_712u,omitempty"`
	Lnd_78017W             *uint32            `protobuf:"varint,100,opt,name=lnd_78017w" json:"lnd_78017w,omitempty"`
	XXX_unrecognized       []byte             `json:"-"`
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

func (m *Telecast) GetDeviceId() uint32 {
	if m != nil && m.DeviceId != nil {
		return *m.DeviceId
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

func (m *Telecast) GetReplyType() TelecastReplyType {
	if m != nil && m.ReplyType != nil {
		return *m.ReplyType
	}
	return Telecast_NO_REPLY
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

func (m *Telecast) GetBatVoltage() float32 {
	if m != nil && m.BatVoltage != nil {
		return *m.BatVoltage
	}
	return 0
}

func (m *Telecast) GetBatSoc() float32 {
	if m != nil && m.BatSoc != nil {
		return *m.BatSoc
	}
	return 0
}

func (m *Telecast) GetWirelessSnr() float32 {
	if m != nil && m.WirelessSnr != nil {
		return *m.WirelessSnr
	}
	return 0
}

func (m *Telecast) GetEnvTemp() float32 {
	if m != nil && m.EnvTemp != nil {
		return *m.EnvTemp
	}
	return 0
}

func (m *Telecast) GetEnvHumid() float32 {
	if m != nil && m.EnvHumid != nil {
		return *m.EnvHumid
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

func (m *Telecast) GetLnd_7318U() uint32 {
	if m != nil && m.Lnd_7318U != nil {
		return *m.Lnd_7318U
	}
	return 0
}

func (m *Telecast) GetLnd_7128Ec() uint32 {
	if m != nil && m.Lnd_7128Ec != nil {
		return *m.Lnd_7128Ec
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

func (m *Telecast) GetBatCurrent() float32 {
	if m != nil && m.BatCurrent != nil {
		return *m.BatCurrent
	}
	return 0
}

func (m *Telecast) GetStatsIccid() string {
	if m != nil && m.StatsIccid != nil {
		return *m.StatsIccid
	}
	return ""
}

func (m *Telecast) GetStatsMotionEvents() uint32 {
	if m != nil && m.StatsMotionEvents != nil {
		return *m.StatsMotionEvents
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

func (m *Telecast) GetStatsCpsi() string {
	if m != nil && m.StatsCpsi != nil {
		return *m.StatsCpsi
	}
	return ""
}

func (m *Telecast) GetStatsUptimeDays() uint32 {
	if m != nil && m.StatsUptimeDays != nil {
		return *m.StatsUptimeDays
	}
	return 0
}

func (m *Telecast) GetStatsDeviceLabel() string {
	if m != nil && m.StatsDeviceLabel != nil {
		return *m.StatsDeviceLabel
	}
	return ""
}

func (m *Telecast) GetStatsGpsParams() string {
	if m != nil && m.StatsGpsParams != nil {
		return *m.StatsGpsParams
	}
	return ""
}

func (m *Telecast) GetStatsServiceParams() string {
	if m != nil && m.StatsServiceParams != nil {
		return *m.StatsServiceParams
	}
	return ""
}

func (m *Telecast) GetStatsTtnParams() string {
	if m != nil && m.StatsTtnParams != nil {
		return *m.StatsTtnParams
	}
	return ""
}

func (m *Telecast) GetStatsSensorParams() string {
	if m != nil && m.StatsSensorParams != nil {
		return *m.StatsSensorParams
	}
	return ""
}

func (m *Telecast) GetLnd_7318C() uint32 {
	if m != nil && m.Lnd_7318C != nil {
		return *m.Lnd_7318C
	}
	return 0
}

func (m *Telecast) GetStatsBattery() string {
	if m != nil && m.StatsBattery != nil {
		return *m.StatsBattery
	}
	return ""
}

func (m *Telecast) GetStatsModuleFona() string {
	if m != nil && m.StatsModuleFona != nil {
		return *m.StatsModuleFona
	}
	return ""
}

func (m *Telecast) GetStatsModuleLora() string {
	if m != nil && m.StatsModuleLora != nil {
		return *m.StatsModuleLora
	}
	return ""
}

func (m *Telecast) GetMotion() bool {
	if m != nil && m.Motion != nil {
		return *m.Motion
	}
	return false
}

func (m *Telecast) GetTest() bool {
	if m != nil && m.Test != nil {
		return *m.Test
	}
	return false
}

func (m *Telecast) GetEncTemp() float32 {
	if m != nil && m.EncTemp != nil {
		return *m.EncTemp
	}
	return 0
}

func (m *Telecast) GetEncHumid() float32 {
	if m != nil && m.EncHumid != nil {
		return *m.EncHumid
	}
	return 0
}

func (m *Telecast) GetEncPressure() float32 {
	if m != nil && m.EncPressure != nil {
		return *m.EncPressure
	}
	return 0
}

func (m *Telecast) GetErrorsOpc() uint32 {
	if m != nil && m.ErrorsOpc != nil {
		return *m.ErrorsOpc
	}
	return 0
}

func (m *Telecast) GetErrorsPms() uint32 {
	if m != nil && m.ErrorsPms != nil {
		return *m.ErrorsPms
	}
	return 0
}

func (m *Telecast) GetErrorsBme0() uint32 {
	if m != nil && m.ErrorsBme0 != nil {
		return *m.ErrorsBme0
	}
	return 0
}

func (m *Telecast) GetErrorsBme1() uint32 {
	if m != nil && m.ErrorsBme1 != nil {
		return *m.ErrorsBme1
	}
	return 0
}

func (m *Telecast) GetErrorsLora() uint32 {
	if m != nil && m.ErrorsLora != nil {
		return *m.ErrorsLora
	}
	return 0
}

func (m *Telecast) GetErrorsFona() uint32 {
	if m != nil && m.ErrorsFona != nil {
		return *m.ErrorsFona
	}
	return 0
}

func (m *Telecast) GetErrorsGeiger() uint32 {
	if m != nil && m.ErrorsGeiger != nil {
		return *m.ErrorsGeiger
	}
	return 0
}

func (m *Telecast) GetErrorsMax01() uint32 {
	if m != nil && m.ErrorsMax01 != nil {
		return *m.ErrorsMax01
	}
	return 0
}

func (m *Telecast) GetErrorsUgps() uint32 {
	if m != nil && m.ErrorsUgps != nil {
		return *m.ErrorsUgps
	}
	return 0
}

func (m *Telecast) GetErrorsTwi() uint32 {
	if m != nil && m.ErrorsTwi != nil {
		return *m.ErrorsTwi
	}
	return 0
}

func (m *Telecast) GetErrorsTwiInfo() string {
	if m != nil && m.ErrorsTwiInfo != nil {
		return *m.ErrorsTwiInfo
	}
	return ""
}

func (m *Telecast) GetErrorsLis() uint32 {
	if m != nil && m.ErrorsLis != nil {
		return *m.ErrorsLis
	}
	return 0
}

func (m *Telecast) GetErrorsSpi() uint32 {
	if m != nil && m.ErrorsSpi != nil {
		return *m.ErrorsSpi
	}
	return 0
}

func (m *Telecast) GetErrorsConnectLora() uint32 {
	if m != nil && m.ErrorsConnectLora != nil {
		return *m.ErrorsConnectLora
	}
	return 0
}

func (m *Telecast) GetErrorsConnectFona() uint32 {
	if m != nil && m.ErrorsConnectFona != nil {
		return *m.ErrorsConnectFona
	}
	return 0
}

func (m *Telecast) GetErrorsConnectWireless() uint32 {
	if m != nil && m.ErrorsConnectWireless != nil {
		return *m.ErrorsConnectWireless
	}
	return 0
}

func (m *Telecast) GetErrorsConnectData() uint32 {
	if m != nil && m.ErrorsConnectData != nil {
		return *m.ErrorsConnectData
	}
	return 0
}

func (m *Telecast) GetErrorsConnectService() uint32 {
	if m != nil && m.ErrorsConnectService != nil {
		return *m.ErrorsConnectService
	}
	return 0
}

func (m *Telecast) GetMotionBeganOffset() uint32 {
	if m != nil && m.MotionBeganOffset != nil {
		return *m.MotionBeganOffset
	}
	return 0
}

func (m *Telecast) GetErrorsConnectGateway() uint32 {
	if m != nil && m.ErrorsConnectGateway != nil {
		return *m.ErrorsConnectGateway
	}
	return 0
}

func (m *Telecast) GetStatsCommsAntFails() uint32 {
	if m != nil && m.StatsCommsAntFails != nil {
		return *m.StatsCommsAntFails
	}
	return 0
}

func (m *Telecast) GetLnd_712U() uint32 {
	if m != nil && m.Lnd_712U != nil {
		return *m.Lnd_712U
	}
	return 0
}

func (m *Telecast) GetLnd_78017W() uint32 {
	if m != nil && m.Lnd_78017W != nil {
		return *m.Lnd_78017W
	}
	return 0
}

func init() {
	proto.RegisterType((*Telecast)(nil), "ttproto.Telecast")
	proto.RegisterEnum("ttproto.TelecastDeviceType", TelecastDeviceType_name, TelecastDeviceType_value)
	proto.RegisterEnum("ttproto.TelecastReplyType", TelecastReplyType_name, TelecastReplyType_value)
}

var fileDescriptor0 = []byte{
	// 1164 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x96, 0x7b, 0x57, 0xdb, 0x46,
	0x13, 0xc6, 0x5f, 0x92, 0x00, 0xf6, 0x72, 0x33, 0xe6, 0x36, 0x09, 0x24, 0xe4, 0xa5, 0x6d, 0x92,
	0xde, 0x88, 0x6c, 0xa0, 0xd0, 0x7b, 0x0d, 0x56, 0x28, 0x0d, 0xb5, 0x5d, 0x50, 0xa0, 0xf4, 0xb6,
	0x5d, 0xcb, 0x6b, 0x47, 0xe7, 0xe8, 0x76, 0xb4, 0x2b, 0x53, 0x7f, 0xda, 0x7e, 0x95, 0xae, 0x46,
	0x5a, 0x63, 0x19, 0xf8, 0xcf, 0xfe, 0x3d, 0xb3, 0xab, 0x99, 0xdd, 0x79, 0x46, 0x22, 0x05, 0x29,
	0xb7, 0xc3, 0x28, 0x90, 0x41, 0x79, 0x5a, 0x4a, 0xfc, 0xb1, 0xf5, 0xef, 0x12, 0x29, 0x58, 0xdc,
	0xe5, 0x36, 0x13, 0xb2, 0x5c, 0x21, 0x33, 0x1d, 0xde, 0x77, 0x6c, 0x4e, 0xe5, 0x20, 0xe4, 0x30,
	0xf1, 0x7c, 0xe2, 0xd5, 0x7c, 0x75, 0x63, 0x3b, 0x8b, 0xdd, 0xd6, 0x71, 0xdb, 0x69, 0x90, 0xa5,
	0x62, 0xca, 0x2f, 0xc9, 0x66, 0xdd, 0x6c, 0x9d, 0x99, 0x47, 0x35, 0xcb, 0xac, 0x57, 0x8d, 0xca,
	0xfe, 0x1b, 0xf3, 0xb0, 0x8e, 0xf2, 0x49, 0xfd, 0x5c, 0x46, 0x8e, 0xdf, 0x83, 0x07, 0x6a, 0x9b,
	0x62, 0x79, 0x91, 0x14, 0xb3, 0xbd, 0x9d, 0x0e, 0x3c, 0x54, 0x68, 0xae, 0xbc, 0x40, 0xa6, 0x3d,
	0x2e, 0x04, 0xeb, 0x71, 0x78, 0x84, 0x31, 0x4b, 0x64, 0xc6, 0x66, 0xa1, 0x8c, 0x23, 0xde, 0xa1,
	0x4c, 0xc2, 0x24, 0xc2, 0xd7, 0x84, 0x44, 0x3c, 0x74, 0x07, 0x69, 0x4e, 0x53, 0x98, 0xd3, 0xfa,
	0xed, 0x9c, 0x30, 0x06, 0x53, 0x7a, 0x46, 0x56, 0x6f, 0xa5, 0x74, 0xc1, 0xdc, 0x98, 0xc3, 0x34,
	0x3e, 0xb6, 0x44, 0x0a, 0x2e, 0x93, 0x8e, 0x8c, 0x3b, 0x1c, 0x0a, 0x8a, 0x3c, 0x48, 0x72, 0x73,
	0x03, 0xbf, 0x97, 0xa2, 0x22, 0x22, 0x15, 0xc4, 0xdc, 0x2c, 0x88, 0x28, 0x32, 0x99, 0x24, 0xd7,
	0x66, 0x92, 0xf6, 0x03, 0x57, 0x26, 0x19, 0xcf, 0x60, 0x98, 0x2a, 0x21, 0x81, 0x22, 0xb0, 0x61,
	0x16, 0xc1, 0x32, 0x99, 0xbd, 0x76, 0x22, 0x95, 0x94, 0x10, 0x54, 0xf8, 0x11, 0xcc, 0xe9, 0xdd,
	0xb8, 0xdf, 0xa7, 0x92, 0x7b, 0x21, 0xcc, 0xeb, 0x47, 0x26, 0xe4, 0x7d, 0xec, 0xa9, 0xe3, 0x58,
	0x40, 0xb4, 0x42, 0xe6, 0xd4, 0x42, 0x36, 0xa0, 0xe9, 0x39, 0x55, 0xa0, 0x84, 0xe9, 0x8e, 0xe1,
	0x2a, 0x2c, 0xde, 0x85, 0x77, 0xa0, 0x7c, 0x17, 0xde, 0x85, 0xa5, 0xbb, 0xf0, 0x1e, 0x2c, 0x23,
	0x4e, 0x0a, 0xf7, 0x3b, 0x74, 0x7f, 0xa7, 0x72, 0x10, 0xc3, 0x0a, 0xa2, 0x32, 0x21, 0x88, 0x2a,
	0xd5, 0x03, 0x6e, 0xc3, 0x2a, 0xb2, 0x0d, 0xb2, 0x2c, 0x24, 0x93, 0x82, 0xc6, 0xa1, 0x74, 0x3c,
	0x4e, 0x3d, 0xc7, 0x8f, 0x25, 0x17, 0xb0, 0x86, 0xea, 0x63, 0xb2, 0x98, 0xaa, 0x2c, 0x0c, 0x69,
	0x9f, 0x47, 0xc2, 0x09, 0x7c, 0x00, 0xbc, 0xbb, 0x75, 0xb2, 0x94, 0x4a, 0xd9, 0xd5, 0x87, 0x2c,
	0x62, 0x9e, 0x80, 0xc7, 0x28, 0x6e, 0x92, 0xb5, 0x54, 0x94, 0x11, 0xf3, 0x85, 0xe7, 0x48, 0xa9,
	0xae, 0xbd, 0x3d, 0x48, 0x36, 0x7e, 0x92, 0x7f, 0x6c, 0xc4, 0x6d, 0xee, 0xf4, 0x87, 0xea, 0x3a,
	0xaa, 0xab, 0x64, 0x3e, 0x55, 0x03, 0x9f, 0x8b, 0xf7, 0x81, 0x14, 0xb0, 0x81, 0xfc, 0x09, 0x29,
	0xa7, 0xdc, 0x0e, 0x3c, 0x2f, 0x59, 0x2b, 0xb8, 0xd2, 0x9e, 0xea, 0xe2, 0x42, 0xc5, 0x42, 0xcf,
	0xa8, 0x50, 0x03, 0x9e, 0x8d, 0xb1, 0x2a, 0xdd, 0x83, 0xcd, 0x3c, 0xab, 0x18, 0x2a, 0xee, 0xf9,
	0x28, 0xb3, 0x0d, 0x83, 0xee, 0x18, 0xf0, 0xff, 0x71, 0xb6, 0x67, 0xc0, 0x56, 0x9e, 0xa9, 0x47,
	0x18, 0xf0, 0x41, 0x9e, 0x55, 0x93, 0xb8, 0x0f, 0xf3, 0x6c, 0x2f, 0x89, 0xfb, 0x28, 0xc7, 0x92,
	0xc7, 0x1a, 0xf0, 0x42, 0xdf, 0x11, 0x32, 0xc1, 0x6d, 0x01, 0x2f, 0x75, 0x58, 0x10, 0xda, 0xba,
	0x8c, 0x57, 0xd8, 0x3d, 0x37, 0x2c, 0x29, 0xe3, 0xe3, 0x3c, 0xc3, 0x32, 0x3e, 0x19, 0x65, 0x58,
	0xc6, 0x01, 0x7c, 0x3a, 0xba, 0x1f, 0x96, 0xb1, 0x0b, 0x9f, 0xe5, 0x19, 0x96, 0xf1, 0x79, 0x9e,
	0x55, 0x69, 0xc5, 0x80, 0xed, 0x3c, 0xc3, 0x32, 0x5e, 0xe7, 0x58, 0x5a, 0x86, 0xa1, 0xcb, 0x40,
	0x86, 0x65, 0x54, 0x10, 0x29, 0xaf, 0x24, 0x1e, 0x08, 0xd5, 0x15, 0x09, 0xe5, 0x79, 0xa8, 0x62,
	0x82, 0xc3, 0xb6, 0x48, 0xef, 0x2f, 0x0c, 0xae, 0x79, 0x44, 0xbb, 0xcc, 0x71, 0x05, 0xec, 0xe0,
	0xb2, 0xcc, 0x88, 0x76, 0x1c, 0x45, 0xdc, 0x97, 0xb0, 0x8b, 0xab, 0x14, 0x4c, 0x57, 0x39, 0xb6,
	0xad, 0x1c, 0xb5, 0x97, 0x6f, 0x3f, 0x2f, 0x90, 0xaa, 0x29, 0x29, 0xef, 0xab, 0x05, 0x02, 0xbe,
	0xd0, 0x09, 0x65, 0xbd, 0xd9, 0x8d, 0x61, 0x1f, 0xe3, 0x81, 0x94, 0x46, 0xe6, 0x0f, 0xed, 0x30,
	0xc9, 0xe1, 0x00, 0x83, 0xc7, 0x94, 0xc4, 0x05, 0xf0, 0xa5, 0x6e, 0xb7, 0x51, 0x25, 0xe8, 0x76,
	0x55, 0xbf, 0xc1, 0x57, 0xa8, 0x3d, 0x25, 0x2b, 0xb9, 0x16, 0xa5, 0xaa, 0xf8, 0xc0, 0xef, 0x08,
	0xf8, 0x1a, 0xe5, 0x39, 0x32, 0xa9, 0x64, 0x35, 0x12, 0xbe, 0xd1, 0x1e, 0xc5, 0xbf, 0x43, 0x0f,
	0x7d, 0xab, 0x0f, 0x33, 0x3b, 0x8f, 0x50, 0x38, 0xf0, 0x1d, 0x26, 0x3a, 0xb4, 0x5c, 0x66, 0xc8,
	0x0e, 0x1b, 0x08, 0xf8, 0x3e, 0xdf, 0xfe, 0x99, 0xe5, 0x5c, 0xd6, 0xe6, 0x2e, 0xfc, 0xa0, 0xeb,
	0x4b, 0xb5, 0x5e, 0x28, 0xb4, 0x17, 0x6b, 0xa8, 0x0c, 0xad, 0x26, 0x78, 0x34, 0xea, 0xd4, 0xc3,
	0xfc, 0x3a, 0x29, 0x7d, 0xad, 0x1c, 0xe5, 0x4f, 0x58, 0x70, 0x5f, 0x04, 0x91, 0x16, 0xeb, 0x7a,
	0xe4, 0xeb, 0xe9, 0x62, 0x83, 0x39, 0x52, 0xa3, 0x8a, 0x57, 0x37, 0x28, 0x79, 0x34, 0x80, 0x37,
	0xf9, 0x7a, 0xbc, 0xa0, 0x13, 0xbb, 0x9c, 0x76, 0x03, 0x9f, 0xc1, 0xf1, 0x9d, 0x92, 0x1b, 0x44,
	0x0c, 0x7e, 0x44, 0x69, 0x9e, 0x4c, 0xa5, 0x17, 0x0b, 0x27, 0xea, 0x7f, 0xa1, 0x3c, 0x4b, 0x1e,
	0xa9, 0xf1, 0x20, 0xe1, 0x27, 0xfc, 0x87, 0x33, 0xd7, 0x4e, 0x67, 0xee, 0xdb, 0x9b, 0x99, 0x6b,
	0x67, 0x33, 0xf7, 0x54, 0x8f, 0xeb, 0x04, 0x0d, 0x5b, 0xf0, 0x67, 0xed, 0x11, 0x1e, 0x45, 0x41,
	0xa4, 0x2e, 0x2e, 0xb4, 0xa1, 0xa1, 0xaf, 0x21, 0x63, 0xca, 0x8d, 0xd0, 0xd4, 0x9d, 0x98, 0xb1,
	0xb6, 0xc7, 0x0d, 0x68, 0xdd, 0x86, 0x15, 0xf8, 0x65, 0x0c, 0x62, 0xfe, 0x67, 0x63, 0x10, 0xeb,
	0x3d, 0xd7, 0x27, 0x94, 0xc1, 0x1e, 0x77, 0x7a, 0x3c, 0x02, 0x6b, 0xe8, 0x95, 0x14, 0x7b, 0xec,
	0x1f, 0xa3, 0x02, 0xef, 0xc6, 0x76, 0x88, 0xd5, 0x95, 0xc2, 0xc5, 0x58, 0xa6, 0xf2, 0xda, 0x81,
	0x4b, 0x64, 0x6b, 0x64, 0xe1, 0x86, 0x51, 0xc7, 0xef, 0x06, 0xf0, 0x2b, 0x9e, 0xe1, 0x4d, 0xb0,
	0xeb, 0x08, 0xb8, 0x1a, 0xdb, 0x40, 0x84, 0x0e, 0xfc, 0x86, 0x4c, 0x5d, 0x74, 0xc6, 0x54, 0x07,
	0xfb, 0xdc, 0x96, 0x69, 0x21, 0xbf, 0xdf, 0x23, 0x62, 0x41, 0x7f, 0xa0, 0xa8, 0xfc, 0x3c, 0x26,
	0xea, 0x17, 0x24, 0xfc, 0x79, 0xcf, 0x6a, 0x65, 0x3c, 0x06, 0x7f, 0xa1, 0xa8, 0x5e, 0xe6, 0x63,
	0x62, 0xd6, 0xa1, 0x40, 0xf5, 0xe2, 0xcc, 0xdc, 0x6d, 0xde, 0x63, 0xbe, 0xf6, 0xdf, 0xdf, 0xf7,
	0x2c, 0xee, 0x29, 0x4b, 0x5f, 0xb3, 0x01, 0xb0, 0xbc, 0x3f, 0xd3, 0x51, 0xc3, 0x7c, 0x99, 0x0d,
	0x9a, 0xf6, 0xf0, 0x43, 0x21, 0x7d, 0x15, 0xc6, 0x60, 0xe7, 0x5e, 0x8e, 0x07, 0xea, 0xab, 0xe2,
	0x1a, 0x3a, 0x09, 0xdb, 0xea, 0x13, 0x32, 0xf2, 0x3d, 0xb4, 0x46, 0x96, 0xde, 0x35, 0xde, 0x36,
	0x9a, 0x97, 0x0d, 0x5a, 0x37, 0x2f, 0x4e, 0x8e, 0x4c, 0x6a, 0x5d, 0xb5, 0xcc, 0xd2, 0xff, 0xd4,
	0x66, 0xb3, 0x87, 0xc7, 0xe6, 0xc9, 0xf1, 0x89, 0x49, 0x1b, 0xb5, 0x46, 0xb3, 0x34, 0xa1, 0xec,
	0x5f, 0x3c, 0x6f, 0x9e, 0xd6, 0xce, 0x8e, 0x6a, 0xe7, 0x56, 0xe9, 0x41, 0xb9, 0x48, 0x26, 0x2d,
	0xab, 0xd6, 0x6a, 0x95, 0x1e, 0xaa, 0xa7, 0x4c, 0x59, 0x56, 0xa3, 0x59, 0x37, 0x4b, 0x8f, 0xd2,
	0xdf, 0xc7, 0xea, 0x53, 0xa6, 0x34, 0x59, 0x9e, 0x21, 0xd3, 0x96, 0x75, 0x6e, 0x9e, 0x5d, 0x98,
	0xa5, 0xa9, 0xad, 0x17, 0xa4, 0x78, 0xf3, 0xcd, 0x33, 0x4b, 0x0a, 0x8d, 0x26, 0x3d, 0x33, 0x5b,
	0xa7, 0x57, 0xea, 0x59, 0x2a, 0xae, 0x76, 0x7a, 0xda, 0xbc, 0x34, 0xeb, 0xa5, 0x89, 0xff, 0x02,
	0x00, 0x00, 0xff, 0xff, 0x43, 0xc7, 0xb5, 0xdf, 0xf5, 0x09, 0x00, 0x00,
}
