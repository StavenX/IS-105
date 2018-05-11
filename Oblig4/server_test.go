package main

import (
	"testing"
	"time"
)

//Checks if the url holds a json
func TestGetJson(t *testing.T) {
	array := []string{"all_hour.geojson","all_week.geojson","all_day.geojson"}
	for i:=0 ; i < 3; i++ {
		server := GetJson("https://earthquake.usgs.gov/earthquakes/feed/v1.0/summary/"+array[i])
		time.Sleep(time.Second*1)
		if server != nil {
			t.Errorf("There is no Json in " + array[i])
		}
	}
}