package api

const (
	TimestampFormat = "2006-01-02T15:04:05.999Z"
)

// ============================================================
// CAM MESSAGE STRUCTURE (ETSI ITS PROTOCOL STANDARD)
// ============================================================

// CAMDatagram - Cooperative Awareness Message
// Based on ETSI TS 102 894-2 V1.3.1
type CAMDatagram struct {
	Header CAMHeader `json:"header"`
	Body   CAMBody   `json:"body"`
}

// ============================================================
// CAM HEADER (ETSI Standard)
// ============================================================

type CAMHeader struct {
	ProtocolVersion int    `json:"protocol_version"` // CAM version (typically 2)
	MessageID       int    `json:"message_id"`       // 1 for CAM
	StationID       int    `json:"station_id"`       // Unique identifier of the sending station
	ReferenceTime   string `json:"reference_time"`   // Timestamp in ISO 8601 format
	GenerationTime  int64  `json:"generation_time"`  // Milliseconds since last full hour
}

// ============================================================
// CAM BODY
// ============================================================

type CAMBody struct {
	StationData         StationData         `json:"station_data"`
	HighFrequencyData   *HighFrequencyData  `json:"high_frequency_data,omitempty"`
	LowFrequencyData    *LowFrequencyData   `json:"low_frequency_data,omitempty"`
	SpecialVehicleData  *SpecialVehicleData `json:"special_vehicle_data,omitempty"`
	MotionStateData     *MotionStateData    `json:"motion_state_data,omitempty"`
	RoadHazardWarning   *RoadHazardWarning  `json:"road_hazard_warning,omitempty"`
	PublicTransport     *PublicTransport    `json:"public_transport,omitempty"`
	EmergencyVehicle    *EmergencyVehicle   `json:"emergency_vehicle,omitempty"`
	AdditionalSensors   AdditionalSensors   `json:"additional_sensors"`
}

// ============================================================
// CAM STATION DATA (Required - Common for all vehicles)
// ============================================================

type StationData struct {
	StationType int       `json:"station_type"` // 0=unknown, 1=pedestrian, 2=cyclist, 3=moped, 4=motorcycle, 5=car, 6=light truck, 7=heavy truck
	VehicleID   VehicleID `json:"vehicle_id"`
	Dimensions  Dimensions `json:"dimensions"`
	Position    Position  `json:"position"`
}

type VehicleID struct {
	VIN  string `json:"vin"`
	ID   int    `json:"id"`
	Type string `json:"type"` // Type of vehicle
}

type Dimensions struct {
	Length                   float32 `json:"length"`                   // in meters
	Width                    float32 `json:"width"`                    // in meters
	Height                   float32 `json:"height,omitempty"`         // in meters (optional)
	WheelbaseDistance        float32 `json:"wheelbase_distance,omitempty"` // Front to rear axle (optional)
	TrackWidth               float32 `json:"track_width,omitempty"`    // Left to right wheel distance (optional)
	ApproximateFrontOverhang float32 `json:"front_overhang,omitempty"` // (optional)
	ApproximateRearOverhang  float32 `json:"rear_overhang,omitempty"`  // (optional)
}

type Position struct {
	Latitude                  float64            `json:"latitude"`
	Longitude                 float64            `json:"longitude"`
	PositionConfidenceEllipse PositionConfidence `json:"position_confidence_ellipse"`
	Altitude                  *Altitude          `json:"altitude,omitempty"`
}

type PositionConfidence struct {
	SemiMajorAxisAccuracy    float32 `json:"semi_major_axis_accuracy"`    // in meters
	SemiMinorAxisAccuracy    float32 `json:"semi_minor_axis_accuracy"`    // in meters
	SemiMajorAxisOrientation float32 `json:"semi_major_axis_orientation"` // degrees from north
}

type Altitude struct {
	AltitudeValue      float32 `json:"altitude_value"`      // in meters
	AltitudeConfidence string  `json:"altitude_confidence"` // alt, unavailable, 10cm, 20cm, 50cm, 100cm, 200cm, 500cm, 1000cm, over1000cm
}

// ============================================================
// CAM HIGH FREQUENCY DATA (Position and motion updates ~10-25 Hz)
// ============================================================

type HighFrequencyData struct {
	Heading               Heading              `json:"heading"`
	Speed                 Speed                `json:"speed"`
	DriveDirection        string               `json:"drive_direction"` // forward, backward, unavailable
	LongitudinalControl   LongitudinalControl  `json:"longitudinal_control"`
	LateralControl        *LateralControl      `json:"lateral_control,omitempty"`
	VehicleLength         float32              `json:"vehicle_length,omitempty"` // in meters
	PositionLonConfidence float32              `json:"position_lon_confidence,omitempty"` // confidence in longitude position
	PositionLatConfidence float32              `json:"position_lat_confidence,omitempty"` // confidence in latitude position
	YawRate               *YawRate             `json:"yaw_rate,omitempty"`
	Curvature             *Curvature           `json:"curvature,omitempty"`
}

type Heading struct {
	HeadingValue      float32 `json:"heading_value"`       // 0-359.99 degrees, 0=north, 90=east
	HeadingConfidence float32 `json:"heading_confidence"`  // 0-127 degrees
}

type Speed struct {
	SpeedValue      float32 `json:"speed_value"`       // in m/s
	SpeedConfidence float32 `json:"speed_confidence"`  // confidence value
}

type LongitudinalControl struct {
	AccelerationValue      float32 `json:"acceleration_value"`      // in m/s², -160 to +161
	AccelerationConfidence float32 `json:"acceleration_confidence"` // confidence value
	EngineRPM              *int    `json:"engine_rpm,omitempty"`
	BrakePedalStatus       bool    `json:"brake_pedal_status"`
	BrakeLightsOn          bool    `json:"brake_lights_on"`
	GasPedalStatus         bool    `json:"gas_pedal_status"`
}

type LateralControl struct {
	SteeringWheelAngle      float32 `json:"steering_wheel_angle"`       // in degrees, -127 to +127
	SteeringWheelAngleConf  float32 `json:"steering_wheel_angle_confidence"`
	LateralAcceleration     float32 `json:"lateral_acceleration,omitempty"` // in m/s²
	LateralAccelerationConf float32 `json:"lateral_acceleration_confidence,omitempty"`
}

type YawRate struct {
	YawRateValue      float32 `json:"yaw_rate_value"`      // degrees per second
	YawRateConfidence float32 `json:"yaw_rate_confidence"` // confidence value
}

type Curvature struct {
	CurvatureValue      float32 `json:"curvature_value"`      // 1/meters
	CurvatureConfidence float32 `json:"curvature_confidence"` // confidence value
}

// ============================================================
// CAM LOW FREQUENCY DATA (~1-2 Hz)
// ============================================================

type LowFrequencyData struct {
	VehicleRole                  string                        `json:"vehicle_role,omitempty"` // civilian, publicTransport, specialTransport, dangerous, roadWork
	ExteriorLights               *ExteriorLights               `json:"exterior_lights,omitempty"`
	WiperStatus                  *WiperStatus                  `json:"wiper_status,omitempty"`
	PathHistory                  []PathPoint                   `json:"path_history,omitempty"`
	BasicVehicleContainerLowFreq *BasicVehicleContainerLowFreq `json:"basic_vehicle_container_low_freq,omitempty"`
}

type ExteriorLights struct {
	LowBeamHeadlightsOn    bool `json:"low_beam_headlights_on"`
	HighBeamHeadlightsOn   bool `json:"high_beam_headlights_on"`
	LeftTurnSignalOn       bool `json:"left_turn_signal_on"`
	RightTurnSignalOn      bool `json:"right_turn_signal_on"`
	DaytimeRunningLightsOn bool `json:"daytime_running_lights_on"`
	ReverseLightsOn        bool `json:"reverse_lights_on"`
	FogLightsOn            bool `json:"fog_lights_on"`
	ParkingLightsOn        bool `json:"parking_lights_on"`
	HazardWarningLightsOn  bool `json:"hazard_warning_lights_on"`
}

type WiperStatus struct {
	FrontWiperOn  bool `json:"front_wiper_on"`
	RearWiperOn   bool `json:"rear_wiper_on"`
	WasherFluidOn bool `json:"washer_fluid_on"`
}

type PathPoint struct {
	DeltaLatitude      float32 `json:"delta_latitude"`
	DeltaLongitude     float32 `json:"delta_longitude"`
	DeltaTime          int     `json:"delta_time"` // in 0.1 second units
	PositionConfidence float32 `json:"position_confidence,omitempty"`
}

type BasicVehicleContainerLowFreq struct {
	VehicleLength      float32 `json:"vehicle_length"`
	VehicleWidth       float32 `json:"vehicle_width"`
	VehicleHeight      float32 `json:"vehicle_height,omitempty"`
	GrossDrivingWeight float32 `json:"gross_driving_weight,omitempty"` // in kg
	TrailerPresent     bool    `json:"trailer_present"`
}

// ============================================================
// CAM SPECIAL VEHICLE DATA
// ============================================================

type SpecialVehicleData struct {
	SpecialVehicleType  int     `json:"special_vehicle_type"` // 0=unavailable, 1=fire, 2=ambulance, 3=police, 4=rescue, 5=hazmat, 6=roadwork
	LightBarActive      bool    `json:"light_bar_active"`
	SirenActive         bool    `json:"siren_active"`
	AdditionalEquipment *string `json:"additional_equipment,omitempty"`
}

// ============================================================
// CAM MOTION STATE DATA
// ============================================================

type MotionStateData struct {
	VehicleMoving            bool     `json:"vehicle_moving"`
	LongitudinalAcceleration float32  `json:"longitudinal_acceleration"`
	LateralAcceleration      float32  `json:"lateral_acceleration,omitempty"`
	VerticalAcceleration     float32  `json:"vertical_acceleration,omitempty"`
	PitchRate                *float32 `json:"pitch_rate,omitempty"`
	RollRate                 *float32 `json:"roll_rate,omitempty"`
}

// ============================================================
// CAM ROAD HAZARD WARNING
// ============================================================

type RoadHazardWarning struct {
	EventType       string   `json:"event_type"`       // obstacle, accident, roadwork, debris, pothole, flooding, skidHazard, etc.
	Location        Position `json:"location"`
	Severity        string   `json:"severity"`        // critical, heavy, moderate, low, unknown
	Description     string   `json:"description,omitempty"`
	Relevant        bool     `json:"relevant"`
	DistanceToEvent int      `json:"distance_to_event,omitempty"` // in meters
}

// ============================================================
// CAM PUBLIC TRANSPORT DATA
// ============================================================

type PublicTransport struct {
	PublicTransportType int    `json:"public_transport_type"` // 0=bus, 1=tram, 2=train, 3=metro, 4=taxi, 5=privateHire
	NumberOfSeats       int    `json:"number_of_seats,omitempty"`
	NumberOfOccupants   int    `json:"number_of_occupants,omitempty"`
	DoorOpen            bool   `json:"door_open"`
	NextStopID          string `json:"next_stop_id,omitempty"`
	RouteID             string `json:"route_id,omitempty"`
}

// ============================================================
// CAM EMERGENCY VEHICLE DATA
// ============================================================

type EmergencyVehicle struct {
	VehicleType        int     `json:"vehicle_type"` // 0=fire, 1=ambulance, 2=police, 3=rescue, 4=hazmat
	LightBarActivation int     `json:"light_bar_activation"`
	SirenActivation    int     `json:"siren_activation"`
	LightsPattern      string  `json:"lights_pattern,omitempty"`
	CategoryOfVehicles string  `json:"category_of_vehicles,omitempty"`
	EmergencyPriority  bool    `json:"emergency_priority"`
	HeadingFromVehicle float32 `json:"heading_from_vehicle,omitempty"`
	DirectionToVehicle string  `json:"direction_to_vehicle,omitempty"`
}

// ============================================================
// ADDITIONAL SENSORS (Custom - NOT in standard CAM)
// Keep sensors from your API that are NOT in CAM standard
// ============================================================

type AdditionalSensors struct {
	// Ultrasonic sensors (not in CAM)
	FrontUltrasonic float32 `json:"front_ultrasonic,omitempty"`  // in meters
	RearUltrasonic  float32 `json:"rear_ultrasonic,omitempty"`   // in meters

	// LIDAR sensors (not in CAM)
	FrontLidar float32 `json:"front_lidar,omitempty"` // in meters

	// Individual wheel speeds (not in CAM)
	SpeedFrontLeft  float32 `json:"speed_front_left,omitempty"`  // in m/s
	SpeedFrontRight float32 `json:"speed_front_right,omitempty"` // in m/s
	SpeedRearLeft   float32 `json:"speed_rear_left,omitempty"`   // in m/s
	SpeedRearRight  float32 `json:"speed_rear_right,omitempty"`  // in m/s

	// Voltage sensors (not in CAM)
	Voltage0 float32 `json:"voltage0,omitempty"` // Battery voltage or other
	Voltage1 float32 `json:"voltage1,omitempty"`
	Voltage2 float32 `json:"voltage2,omitempty"`

	// GPS confidence metrics (not in CAM standard)
	GPSSatelliteCount     float32 `json:"gps_satellite_count,omitempty"`
	GPSHorizontalAccuracy float32 `json:"gps_horizontal_accuracy,omitempty"` // in meters
	GPSDirection          float32 `json:"gps_direction,omitempty"`           // degrees

	// User control (not in CAM)
	IsControlledByUser bool `json:"is_controlled_by_user,omitempty"`
}