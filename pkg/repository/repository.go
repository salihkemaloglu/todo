package repository

import (
	"fmt"

	"github.com/heetch/sqalx"
	"github.com/pkg/errors"
	"github.com/salihkemaloglu/todo/pkg/model"
)

// UpsertObject upsert object according online status
func UpsertObject(o model.Object, node sqalx.Node) error {
	_, err := node.Exec(
		`INSERT INTO "object" (
			object_id,
			online
			)
		VALUES (
			$1,
			$2
		)
		ON CONFLICT ON
			CONSTRAINT unique_object DO UPDATE SET created_at = NOW();
		`,
		o.ObjectID,
		o.Online,
	)

	return errors.Wrap(err, fmt.Sprintf("couldn't insert by object_id: %d", o.ObjectID))
}

// GetObject returns object by object_id
func GetObject(objectID int, node sqalx.Node) (model.Object, error) {
	var o model.Object
	err := node.Get(&o, `
	SELECT
		online,
		created_at
	FROM "object"
		WHERE object_id =$1;
	`, objectID)

	return o, errors.Wrap(err, fmt.Sprintf("couldn't get object by object_id: %d", objectID))
}

// DeleteObject deletes object by object_id
func DeleteObject(objectID int, node sqalx.Node) error {
	_, err := node.Exec(`
	DELETE FROM "object"
    	WHERE object_id=$1;
	`, objectID)
	return errors.Wrap(err, fmt.Sprintf("couldn't delete by object_id: %d", objectID))
}
