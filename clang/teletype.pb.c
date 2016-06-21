/* Automatically generated nanopb constant definitions */
/* Generated by nanopb-0.3.5 at Tue Jun 21 12:47:11 2016. */

#include "teletype.pb.h"

/* @@protoc_insertion_point(includes) */
#if PB_PROTO_HEADER_VERSION != 30
#error Regenerate this file with the current version of nanopb generator.
#endif



const pb_field_t teletype_Telecast_fields[16] = {
    PB_FIELD(  1, UENUM   , REQUIRED, STATIC  , FIRST, teletype_Telecast, DeviceType, DeviceType, 0),
    PB_FIELD(  2, STRING  , OPTIONAL, STATIC  , OTHER, teletype_Telecast, DeviceIDString, DeviceType, 0),
    PB_FIELD(  3, UINT32  , OPTIONAL, STATIC  , OTHER, teletype_Telecast, DeviceIDNumber, DeviceIDString, 0),
    PB_FIELD(  4, STRING  , OPTIONAL, STATIC  , OTHER, teletype_Telecast, Message, DeviceIDNumber, 0),
    PB_FIELD(  5, STRING  , OPTIONAL, STATIC  , OTHER, teletype_Telecast, CapturedAt, Message, 0),
    PB_FIELD(  6, UENUM   , OPTIONAL, STATIC  , OTHER, teletype_Telecast, Unit, CapturedAt, 0),
    PB_FIELD(  7, UINT32  , OPTIONAL, STATIC  , OTHER, teletype_Telecast, Value, Unit, 0),
    PB_FIELD(  8, FLOAT   , OPTIONAL, STATIC  , OTHER, teletype_Telecast, Latitude, Value, 0),
    PB_FIELD(  9, FLOAT   , OPTIONAL, STATIC  , OTHER, teletype_Telecast, Longitude, Latitude, 0),
    PB_FIELD( 10, UINT32  , OPTIONAL, STATIC  , OTHER, teletype_Telecast, Altitude, Longitude, 0),
    PB_FIELD( 11, FLOAT   , OPTIONAL, STATIC  , OTHER, teletype_Telecast, BatteryVoltage, Altitude, 0),
    PB_FIELD( 12, FLOAT   , OPTIONAL, STATIC  , OTHER, teletype_Telecast, BatterySOC, BatteryVoltage, 0),
    PB_FIELD( 13, FLOAT   , OPTIONAL, STATIC  , OTHER, teletype_Telecast, WirelessSNR, BatterySOC, 0),
    PB_FIELD( 14, FLOAT   , OPTIONAL, STATIC  , OTHER, teletype_Telecast, envTemperature, WirelessSNR, 0),
    PB_FIELD( 15, FLOAT   , OPTIONAL, STATIC  , OTHER, teletype_Telecast, envHumidity, envTemperature, 0),
    PB_LAST_FIELD
};


/* @@protoc_insertion_point(eof) */
