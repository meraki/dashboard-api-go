package main

import (
	"fmt"
	"net/url"
	"reflect"
	"strings"

	meraki "github.com/meraki/dashboard-api-go/v4/sdk"
)

var client *meraki.Client

type cmpFunc func(interface{}, interface{}) bool

// Simple comparison function that compares if two values are equal.
func simpleCmp(a, b interface{}) bool {
	return a == b
}
func main() {
	var err error
	fmt.Println("Authenticating")
	client, err = meraki.NewClientWithOptions("https://api.meraki.com/",
		"12f2eb53588c75e28d89e108a05ea0c2487b08cf",
		"true", "AplicationName VendorName")
	if err != nil {
		fmt.Println(err)
		return
	}
	all_response := getAllItems(*client)

	// fmt.Println(len(all_response))

	b := structToMap(all_response)
	c := getDictResult(b, "Ts", "2024-01-24T15:44:04.930727Z", simpleCmp)
	var as meraki.ResponseItemOrganizationsGetOrganizationAPIRequests
	d := mapToStruct(c.(map[string]interface{}), &as)
	if d != nil {
		fmt.Println("Hubo un error")
	}
	fmt.Println(as)
}

// getDictResult function translates the Python function to Go.
func getDictResult(result interface{}, key string, value interface{}, cmpFn cmpFunc) interface{} {
	// Check if the result is a list.
	if reflect.TypeOf(result).Kind() == reflect.Slice {
		resultSlice := reflect.ValueOf(result)

		// If the list has a single element.
		if resultSlice.Len() == 1 {
			// Check if that element is a dictionary.
			if reflect.TypeOf(resultSlice.Index(0).Interface()).Kind() == reflect.Map {
				resultMap := resultSlice.Index(0).Interface().(map[string]interface{})
				// Check if the key and value match the provided ones.
				if val, ok := resultMap[key]; ok && !cmpFn(val, value) {
					return nil
				}
				return resultMap
			}
			return nil
		}

		// Iterate over the elements of the list.
		for i := 0; i < resultSlice.Len(); i++ {
			item := resultSlice.Index(i)
			// Check if the item is a dictionary.
			if reflect.TypeOf(item.Interface()).Kind() == reflect.Map {
				itemMap := item.Interface().(map[string]interface{})
				// Check if the key and value match the provided ones.
				if val, ok := itemMap[key]; !ok || cmpFn(val, value) {
					return itemMap
				}
			}
		}
		return nil
	}

	// If the result is not a list.
	if reflect.TypeOf(result).Kind() != reflect.Map {
		return nil
	}

	// Check if the result is a dictionary.
	resultMap := result.(map[string]interface{})
	// Check if the key and value match the provided ones.
	if val, ok := resultMap[key]; ok && !cmpFn(val, value) {
		return nil
	}

	return resultMap
}

func structToMap(data interface{}) []map[string]interface{} {
	// Verificar si se pasa una slice
	val := reflect.ValueOf(data)
	if val.Kind() != reflect.Slice {
		fmt.Println("El valor proporcionado no es una slice.")
		return nil
	}

	// Obtener el tipo de la estructura dentro de la slice
	elementType := val.Type().Elem()
	if elementType.Kind() == reflect.Ptr {
		elementType = elementType.Elem()
	}
	if elementType.Kind() != reflect.Struct {
		fmt.Println("El tipo dentro de la slice no es una estructura.")
		return nil
	}

	// Convertir cada elemento de la slice a un mapa
	result := make([]map[string]interface{}, val.Len())
	for i := 0; i < val.Len(); i++ {
		element := val.Index(i)
		elementMap := make(map[string]interface{})
		for j := 0; j < element.NumField(); j++ {
			fieldName := elementType.Field(j).Name
			fieldValue := element.Field(j).Interface()
			elementMap[fieldName] = fieldValue
		}
		result[i] = elementMap
	}
	return result
}

func mapToStruct(data map[string]interface{}, target interface{}) error {
	targetType := reflect.TypeOf(target)
	if targetType.Kind() != reflect.Ptr || targetType.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("The target is not a pointer to an struct")
	}
	targetValue := reflect.ValueOf(target).Elem()
	for i := 0; i < targetType.Elem().NumField(); i++ {
		field := targetType.Elem().Field(i)
		fieldName := field.Name
		value, ok := data[fieldName]
		if !ok {
			continue
		}
		fieldValue := reflect.ValueOf(value)
		if fieldValue.Type().ConvertibleTo(field.Type) {
			targetValue.FieldByName(fieldName).Set(fieldValue.Convert(field.Type))
		} else {
			return fmt.Errorf("The valur type %s is not on struct", fieldName)
		}
	}

	return nil
}

func getAllItems(client meraki.Client) meraki.ResponseOrganizationsGetOrganizationAPIRequests {
	var all_response meraki.ResponseOrganizationsGetOrganizationAPIRequests
	response, r2, _ := client.Organizations.GetOrganizationAPIRequests("828099381482762270", &meraki.GetOrganizationAPIRequestsQueryParams{
		PerPage: 1000,
	})
	count := 0
	all_response = append(all_response, *response...)
	for len(*response) >= 1000 {
		count += 1
		fmt.Println(count)
		links := strings.Split(r2.Header().Get("Link"), ",")
		var link string
		if count > 1 {
			link = strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.Split(links[2], ";")[0], ">", ""), "<", ""), client.RestyClient().BaseURL, "")
		} else {
			link = strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.Split(links[1], ";")[0], ">", ""), "<", ""), client.RestyClient().BaseURL, "")
		}
		myUrl, _ := url.Parse(link)
		params, _ := url.ParseQuery(myUrl.RawQuery)
		if params["endingBefore"] != nil {
			response, r2, _ = client.Organizations.GetOrganizationAPIRequests("828099381482762270", &meraki.GetOrganizationAPIRequestsQueryParams{
				PerPage:      1000,
				EndingBefore: params["endingBefore"][0],
			})
			all_response = append(all_response, *response...)
		} else {
			break
		}
	}

	return all_response
}
