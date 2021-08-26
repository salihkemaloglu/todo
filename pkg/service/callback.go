package service

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/pkg/errors"
	"github.com/salihkemaloglu/todo/pkg/model"
	"github.com/salihkemaloglu/todo/pkg/pgu"
	"github.com/salihkemaloglu/todo/pkg/repository"
	"github.com/salihkemaloglu/todo/pkg/util/config"
)

func Callback(o model.Object, config *config.Config) error {
	node, err := pgu.MustOpen(config.Postgres.DatabaseURL)
	if err != nil {
		return errors.Wrap(err, "couldn't open database connection")

	}
	fmt.Println(o.ObjectIDs)
	for _, id := range o.ObjectIDs {
		ro, err := checkOnlineStatus(id, config)
		if err != nil {
			return err
		}

		o, err := repository.GetObject(id, node)
		if err != nil {
			if sql.ErrNoRows.Error() == err.Error()[len(err.Error())-26:] {
				err = repository.UpsertObject(ro, node)
				if err != nil {
					return err
				}
			} else {
				return err
			}
		}

		if !o.Online && time.Now().UTC().Before(o.CreatedAt.Add(time.Second*30)) {
			err = repository.DeleteObject(id, node)
			if err != nil {
				return err
			}
		} else {
			err = repository.UpsertObject(ro, node)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// checkOnlineStatus checks id status whether online or not
func checkOnlineStatus(id int, config *config.Config) (model.Object, error) {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	reqURL := config.OnlineService.BaseURL + strconv.Itoa(id)
	req, err := http.NewRequest("GET", reqURL, nil)
	if err != nil {
		return model.Object{}, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return model.Object{}, err
	}
	defer resp.Body.Close()

	var respObj model.Object
	err = json.NewDecoder(resp.Body).Decode(&respObj)
	return respObj, errors.Wrap(err, "couldn't unmarshal response object")
}
