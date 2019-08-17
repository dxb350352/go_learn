package main
import (
	"math"
)

type Geo struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	ParentId   string `json:"parent_id"`
	ShortName  string `json:"short_name"`
	Level      string `json:"level"`
	CityCode   string `json:"city_code"`
	ZipCode    string `json:"zip_code"`
	MergerName string `json:"merger_name"`
	Longitude  float64 `json:"longitude"`
	Latitude   float64 `json:"latitude"`
	Pinyin     string `json:"pinyin"`
	Point3d    []float64 `json:"point3d"`
}

type GeoXSort []Geo
func (list GeoXSort) Len() int {
	return len(list)
}

func (list GeoXSort) Less(i, j int) bool {
	return list[i].Point3d[0] < list[j].Point3d[0]
}

func (list GeoXSort) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}

type GeoYSort []Geo
func (list GeoYSort) Len() int {
	return len(list)
}

func (list GeoYSort) Less(i, j int) bool {
	return list[i].Point3d[1] < list[j].Point3d[1]
}

func (list GeoYSort) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}

type GeoZSort []Geo
func (list GeoZSort) Len() int {
	return len(list)
}

func (list GeoZSort) Less(i, j int) bool {
	return list[i].Point3d[2] < list[j].Point3d[2]
}

func (list GeoZSort) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}

func (g Geo)SetPoint() {
	g.Point3d[0] = math.Cos(math.Pi / 180 * g.Latitude) * math.Cos(math.Pi / 180 * g.Longitude)
	g.Point3d[1] = math.Cos(math.Pi / 180 * g.Latitude) * math.Sin(math.Pi / 180 * g.Longitude)
	g.Point3d[2] = math.Cos(math.Pi / 180 * g.Latitude)
}

func (g Geo) SquaredDistance(other Geo) float64 {
	x := g.Point3d[0] - other.Point3d[0]
	y := g.Point3d[1] - other.Point3d[1]
	z := g.Point3d[2] - other.Point3d[2]
	return (x * x) + (y * y) + (z * z)
}

func (g Geo) AxisSquaredDistance(other Geo, axis int) {
	return math.Sqrt(g.Compare(other, axis))
}

func (g Geo) Compare(other Geo, axis int) float64 {
	return g.Point3d[axis] - other.Point3d[axis]
}