package nmap

import (
	"github.com/jeotic/nmapper/pkg"
	"log"
)

type Host struct {
	Id        int64
	RunId     int64
	EndTime   int64
	StartTime int64
}

func SelectHosts(env pkg.ENV, query string, args ...interface{}) ([]Host, error) {
	rows, err := env.DB.Query(query, args...)
	var hosts []Host

	if err != nil {
		return hosts, err
	}

	for rows.Next() {
		var host Host

		if err := rows.Scan(&host.Id, &host.RunId, &host.EndTime, &host.StartTime); err != nil {
			log.Println(err.Error())
			return hosts, err
		}

		hosts = append(hosts, host)
	}

	return hosts, nil
}

func GetHosts(env pkg.ENV, run_id int64) ([]Host, error) {
	return SelectHosts(env, "SELECT * FROM host WHERE run_id = ?", run_id)
}

func GetHost(env pkg.ENV, run_id int64, id int64) (*Host, error) {
	runs, err := SelectHosts(env, "SELECT * FROM host WHERE run_id = ? AND id = ?", run_id, id)

	if err != nil {
		return nil, err
	}

	return &runs[0], nil
}
