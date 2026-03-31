// Package filter provides scene-based JSON field filtering for Go values.
//
// The quick entrypoints are Select and Omit, which return a value that can be
// passed directly to json.Marshal or framework helpers such as gin.Context.JSON.
//
// If you need the typed helper methods, use SelectFilter or OmitFilter and then
// call JSON, Bytes, Map, Slice, or Interface on the returned Filter.
package filter
