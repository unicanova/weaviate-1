/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 - 2018 Weaviate. All rights reserved.
 * LICENSE: https://github.com/creativesoftwarefdn/weaviate/blob/develop/LICENSE.md
 * AUTHOR: Bob van Luijt (bob@kub.design)
 * See www.creativesoftwarefdn.org for details
 * Contact: @CreativeSofwFdn / bob@kub.design
 */

package models

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ThingHistoryObject thing history object
// swagger:model ThingHistoryObject

type ThingHistoryObject struct {
	ThingCreate

	// Timestamp of creation of this thing history in milliseconds since epoch UTC.
	CreationTimeUnix int64 `json:"creationTimeUnix,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *ThingHistoryObject) UnmarshalJSON(raw []byte) error {

	var aO0 ThingCreate
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.ThingCreate = aO0

	var data struct {
		CreationTimeUnix int64 `json:"creationTimeUnix,omitempty"`
	}
	if err := swag.ReadJSON(raw, &data); err != nil {
		return err
	}

	m.CreationTimeUnix = data.CreationTimeUnix

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m ThingHistoryObject) MarshalJSON() ([]byte, error) {
	var _parts [][]byte

	aO0, err := swag.WriteJSON(m.ThingCreate)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var data struct {
		CreationTimeUnix int64 `json:"creationTimeUnix,omitempty"`
	}

	data.CreationTimeUnix = m.CreationTimeUnix

	jsonData, err := swag.WriteJSON(data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, jsonData)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this thing history object
func (m *ThingHistoryObject) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.ThingCreate.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *ThingHistoryObject) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ThingHistoryObject) UnmarshalBinary(b []byte) error {
	var res ThingHistoryObject
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
