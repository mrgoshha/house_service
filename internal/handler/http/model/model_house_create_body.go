/*
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package model

type HouseCreateBody struct {
	Address string `json:"address"`

	Year int32 `json:"year"`

	Developer string `json:"developer,omitempty"`
}
