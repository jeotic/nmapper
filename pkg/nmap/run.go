package nmap

import (
	"github.com/jeotic/nmapper/pkg"
	"log"
)

type Run struct {
	Id               int64
	Start            float64
	Version          float64
	XmlOutputVersion float64
	Args             string
	StartStr         string
	Scanner          string
	ScanInfo         ScanInfo
}

type ScanInfo struct {
	Id          int64
	RunId       int64
	NumServices int64
	Type        JsonNullString
	Protocol    string
	Services    string
}

type Verbose struct {
	Id int64
}

type Debugging struct {
	Id int64
}

func SelectRuns(env pkg.ENV, query string, args ...interface{}) ([]Run, error) {
	rows, err := env.DB.Query(query, args...)
	var runs []Run

	if err != nil {
		return runs, err
	}

	for rows.Next() {
		var run Run
		var scanInfo ScanInfo

		if err := rows.Scan(
			&run.Id,
			&run.Scanner,
			&run.Args,
			&run.Start,
			&run.StartStr,
			&run.Version,
			&run.XmlOutputVersion,
			&scanInfo.Id,
			&scanInfo.Type,
			&scanInfo.Protocol,
			&scanInfo.NumServices,
			&scanInfo.Services,
			&scanInfo.RunId); err != nil {
			log.Println(err.Error())
			return runs, err
		}

		run.ScanInfo = scanInfo

		runs = append(runs, run)
	}

	return runs, nil
}

func GetSelectQuery() string {
	return `SELECT * FROM nmap_run
	JOIN scan_info si on nmap_run.id = si.run_id`
}

func GetRuns(env pkg.ENV) ([]Run, error) {
	return SelectRuns(env, GetSelectQuery())
}

func GetRun(env pkg.ENV, id int64) (*Run, error) {
	runs, err := SelectRuns(env, GetSelectQuery()+" WHERE nmap_run.id = ?", id)

	if err != nil {
		return nil, err
	}

	return &runs[0], nil
}
