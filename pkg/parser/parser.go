package parser

import (
	"encoding/xml"
	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/sqlite3"
	"github.com/jeotic/nmapper/pkg"
	"io"
	"strconv"
)

func ParseReader(env pkg.ENV, ri io.Reader) (int64, error) {
	dialect := env.Builder

	var decoder *xml.Decoder = xml.NewDecoder(ri)
	var nmap_id int64
	var last_host_id int64
	var last_port_id int64

	for {
		t, err := decoder.RawToken()

		if err == io.EOF {
			return nmap_id, nil
		} else if err != nil {
			return -1, err
		}

		switch t := t.(type) {
		case xml.StartElement:
			switch t.Name.Local {
			case "nmaprun":
				ds := BuildQuery(dialect, "nmap_run", t.Attr)
				id, err := RunQuery(env, ds)

				if err != nil {
					return -1, err
				}

				nmap_id = id
			case "scaninfo":
				t.Attr = AppendIntAttributes(t.Attr, "run_id", nmap_id)

				ds := BuildQuery(dialect, "scan_info", t.Attr)
				_, err := RunQuery(env, ds)

				if err != nil {
					return -1, err
				}
			case "verbose":
				t.Attr = AppendIntAttributes(t.Attr, "run_id", nmap_id)

				ds := BuildQuery(dialect, "verbose", t.Attr)
				_, err := RunQuery(env, ds)

				if err != nil {
					return -1, err
				}
			case "debugging":
				t.Attr = AppendIntAttributes(t.Attr, "run_id", nmap_id)

				ds := BuildQuery(dialect, "debugging", t.Attr)
				_, err := RunQuery(env, ds)

				if err != nil {
					return -1, err
				}
			case "taskbegin":
				t.Attr = AppendIntAttributes(t.Attr, "run_id", nmap_id)
				// Add Type
				t.Attr = AppendAttributes(t.Attr, "type", "BEGIN")

				ds := BuildQuery(dialect, "task", t.Attr)
				_, err := RunQuery(env, ds)

				if err != nil {
					return -1, err
				}
			case "taskend":
				t.Attr = AppendIntAttributes(t.Attr, "run_id", nmap_id)
				// Add Type
				t.Attr = AppendAttributes(t.Attr, "type", "END")

				ds := BuildQuery(dialect, "task", t.Attr)
				_, err := RunQuery(env, ds)

				if err != nil {
					return -1, err
				}
			case "host":
				t.Attr = AppendIntAttributes(t.Attr, "run_id", nmap_id)

				ds := BuildQuery(dialect, "host", t.Attr)
				id, err := RunQuery(env, ds)

				if err != nil {
					return -1, err
				}

				last_host_id = id
			case "status":
				t.Attr = AppendIntAttributes(t.Attr, "run_id", nmap_id)
				t.Attr = AppendIntAttributes(t.Attr, "host_id", last_host_id)

				ds := BuildQuery(dialect, "host_status", t.Attr)
				_, err := RunQuery(env, ds)

				if err != nil {
					return -1, err
				}
			case "address":
				t.Attr = AppendIntAttributes(t.Attr, "run_id", nmap_id)
				t.Attr = AppendIntAttributes(t.Attr, "host_id", last_host_id)

				ds := BuildQuery(dialect, "host_address", t.Attr)
				_, err := RunQuery(env, ds)

				if err != nil {
					return -1, err
				}
			case "times":
				t.Attr = AppendIntAttributes(t.Attr, "run_id", nmap_id)
				t.Attr = AppendIntAttributes(t.Attr, "host_id", last_host_id)

				ds := BuildQuery(dialect, "host_time", t.Attr)
				_, err := RunQuery(env, ds)

				if err != nil {
					return -1, err
				}
			case "hostname":
				t.Attr = AppendIntAttributes(t.Attr, "run_id", nmap_id)
				t.Attr = AppendIntAttributes(t.Attr, "host_id", last_host_id)

				ds := BuildQuery(dialect, "host_name", t.Attr)
				_, err := RunQuery(env, ds)

				if err != nil {
					return -1, err
				}
			case "port":
				t.Attr = AppendIntAttributes(t.Attr, "run_id", nmap_id)
				t.Attr = AppendIntAttributes(t.Attr, "host_id", last_host_id)

				ds := BuildQuery(dialect, "host_port", t.Attr)
				id, err := RunQuery(env, ds)

				if err != nil {
					return -1, err
				}

				last_port_id = id
			case "state":
				t.Attr = AppendIntAttributes(t.Attr, "run_id", nmap_id)
				t.Attr = AppendIntAttributes(t.Attr, "port_id", last_port_id)

				ds := BuildQuery(dialect, "port_state", t.Attr)
				_, err := RunQuery(env, ds)

				if err != nil {
					return -1, err
				}
			case "service":
				t.Attr = AppendIntAttributes(t.Attr, "run_id", nmap_id)
				t.Attr = AppendIntAttributes(t.Attr, "port_id", last_port_id)

				ds := BuildQuery(dialect, "port_service", t.Attr)
				_, err := RunQuery(env, ds)

				if err != nil {
					return -1, err
				}
			case "finished":
				t.Attr = AppendIntAttributes(t.Attr, "run_id", nmap_id)

				ds := BuildQuery(dialect, "run_finished_stat", t.Attr)
				_, err := RunQuery(env, ds)

				if err != nil {
					return -1, err
				}
			case "hosts":
				t.Attr = AppendIntAttributes(t.Attr, "run_id", nmap_id)

				ds := BuildQuery(dialect, "run_host_stat", t.Attr)
				_, err := RunQuery(env, ds)

				if err != nil {
					return -1, err
				}
			}
		}
	}
}

func AppendAttributes(attrs []xml.Attr, key string, value string) []xml.Attr {
	name := xml.Name{
		Space: "",
		Local: key,
	}

	return append(attrs, xml.Attr{Name: name, Value: value})
}

func AppendIntAttributes(attrs []xml.Attr, key string, value int64) []xml.Attr {
	return AppendAttributes(attrs, key, strconv.FormatInt(value, 10))
}

func BuildQuery(builder goqu.DialectWrapper, tableName string, attrs []xml.Attr) *goqu.InsertDataset {
	ds := builder.Insert(tableName).Prepared(true)

	record := goqu.Record{}

	for _, attr := range attrs {
		record[attr.Name.Local] = attr.Value
	}

	return ds.Rows(record)
}

func RunQuery(env pkg.ENV, qs *goqu.InsertDataset) (int64, error) {
	query, args, err := qs.ToSQL()

	if err != nil {
		return -1, err
	}

	res, err := env.DB.DB().Exec(query, args...)

	if err != nil {
		return -1, err
	}

	return res.LastInsertId()
}
