package main

type CellSpec struct {
	CapacityMAh           CapacityMAh           `json:"capacity_mAh"`
	NominalVoltageV       float64               `json:"nominal_voltage_V"`
	StandardCharge        StandardCharge        `json:"standard_charge"`
	MaxChargeVoltageV     float64               `json:"max_charge_voltage_V"`
	MaxChargeCurrentMA    int                   `json:"max_charge_current_mA"`
	StandardDischarge     StandardDischarge     `json:"standard_discharge"`
	MaxDischargeCurrentA  int                   `json:"max_discharge_current_A"`
	MaxTemperatureLimitC  int                   `json:"max_temperature_limit_C"`
	OperatingTemperatureC OperatingTemperatureC `json:"operating_temperature_C"`
	StorageTemperatureC   StorageTemperatureC   `json:"storage_temperature_C"`
}

type CapacityMAh struct {
	Nominal int `json:"nominal"`
	Mininum int `json:"mininum"`
}
type StandardCharge struct {
	ConstantCurrentMA     int     `json:"constant_current_mA"`
	ConstantVoltageV      float64 `json:"constant_voltage_V"`
	EndConditionCurrentMA int     `json:"end_condition_current_mA"`
}
type StandardDischarge struct {
	ConstantCurrentMA    int     `json:"constant_current_mA"`
	EndConditionVoltageV float64 `json:"end_condition_voltage_V"`
}
type Charge struct {
	From int `json:"from"`
	To   int `json:"to"`
}
type Discharge struct {
	From int `json:"from"`
	To   int `json:"to"`
}
type OperatingTemperatureC struct {
	Charge    Charge    `json:"charge"`
	Discharge Discharge `json:"discharge"`
}
type OneMonth struct {
	From int `json:"from"`
	To   int `json:"to"`
}
type ThreeMonths struct {
	From int `json:"from"`
	To   int `json:"to"`
}
type OneYear struct {
	From int `json:"from"`
	To   int `json:"to"`
}
type StorageTemperatureC struct {
	OneMonth    OneMonth    `json:"1_month"`
	ThreeMonths ThreeMonths `json:"3_months"`
	OneYear     OneYear     `json:"1_year"`
}

func (cell CellSpec) AsBattery(spec BatterySpec) CellSpec {
	return CellSpec{
		CapacityMAh: CapacityMAh{
			Nominal: cell.CapacityMAh.Nominal * spec.CellParallel,
			Mininum: cell.CapacityMAh.Mininum * spec.CellParallel,
		},
		NominalVoltageV: cell.NominalVoltageV * float64(spec.CellSeries),
		StandardCharge: StandardCharge{
			ConstantCurrentMA:     cell.StandardCharge.ConstantCurrentMA * spec.CellParallel,
			ConstantVoltageV:      cell.StandardCharge.ConstantVoltageV * float64(spec.CellSeries),
			EndConditionCurrentMA: cell.StandardCharge.EndConditionCurrentMA * spec.CellParallel,
		},
		MaxChargeVoltageV:  cell.MaxChargeVoltageV * float64(spec.CellSeries),
		MaxChargeCurrentMA: cell.MaxChargeCurrentMA * spec.CellParallel,
		StandardDischarge: StandardDischarge{
			ConstantCurrentMA:    cell.StandardDischarge.ConstantCurrentMA * spec.CellParallel,
			EndConditionVoltageV: cell.StandardDischarge.EndConditionVoltageV * float64(spec.CellSeries),
		},
		MaxDischargeCurrentA:  cell.MaxDischargeCurrentA * spec.CellParallel,
		MaxTemperatureLimitC:  cell.MaxTemperatureLimitC,
		OperatingTemperatureC: cell.OperatingTemperatureC,
		StorageTemperatureC:   cell.StorageTemperatureC,
	}
}
