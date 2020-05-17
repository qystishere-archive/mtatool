package parser

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/clbanning/mxj"

	gtaItem "github.com/qystishere/mtatool/gta/parser/item"
	"github.com/qystishere/mtatool/mta/parser/item"
)

type Map struct {
	Path  string
	Items []gtaItem.Item

	FileInfo os.FileInfo
}

func LoadMap(path string) (*Map, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	stat, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	m, err := mxj.NewMapXmlReader(file)
	if err != nil {
		return nil, err
	}

	value, err := m.ValueForPath("map")
	if err != nil {
		return nil, err
	}

	items := make([]gtaItem.Item, 0)
	switch value := value.(type) {
	case map[string]interface{}:
		var valuesMaps []map[string]interface{}
		for k, v := range value {
			put := func(valuesMap map[string]interface{}) {
				for k, v := range valuesMap {
					valuesMap[strings.TrimLeft(k, "-")] = v
				}
				valuesMap["type"] = k
				valuesMaps = append(valuesMaps, valuesMap)
			}

			switch v := v.(type) {
			case []interface{}:
				for _, v := range v {
					if v, ok := v.(map[string]interface{}); ok {
						put(v)
					}
				}
			case map[string]interface{}:
				put(v)
			default:
				continue
			}
		}

		for _, valuesMap := range valuesMaps {
			bb, err := json.Marshal(valuesMap)
			if err != nil {
				return nil, err
			}

			var i gtaItem.Item
			switch gtaItem.Type(valuesMap["type"].(string)) {
			case gtaItem.TypeObject:
				var object item.Object
				err = json.Unmarshal(bb, &object)
				i = object
			case gtaItem.TypeVehicle:
				var vehicle item.Vehicle
				err = json.Unmarshal(bb, &vehicle)
				i = vehicle
			case gtaItem.TypePed:
				var ped item.Ped
				err = json.Unmarshal(bb, &ped)
				i = ped
			case gtaItem.TypeMarker:
				var marker item.Marker
				err = json.Unmarshal(bb, &marker)
				i = marker
			case gtaItem.TypePickup:
				var pickup item.Pickup
				err = json.Unmarshal(bb, &pickup)
				i = pickup
			default:
				return nil, fmt.Errorf("unknown item type: %s", valuesMap["type"].(string))
			}

			if err != nil {
				return nil, fmt.Errorf("conversion error: %w", err)
			}

			items = append(items, i)
		}
	default:
		return nil, fmt.Errorf("unknown map type: %T", value)
	}

	return &Map{
		Path:  path,
		Items: items,

		FileInfo: stat,
	}, nil
}
