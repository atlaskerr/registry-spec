package schema

import (
	"github.com/xeipuuv/gojsonschema"
)

// Validator validates.
func Validator() map[MediaType]*gojsonschema.Schema {
	vmap := map[MediaType]*gojsonschema.Schema{
		MediaTypeNamespace:       namespaceValidator(),
		MediaTypeNamespaceCreate: namespaceCreateValidator(),
		MediaTypeNamespaceList:   namespaceListValidator(),

		MediaTypeProject:       projectValidator(),
		MediaTypeProjectCreate: projectCreateValidator(),
		MediaTypeProjectList:   projectListValidator(),

		MediaTypeRepository:       repositoryValidator(),
		MediaTypeRepositoryCreate: repositoryCreateValidator(),
		MediaTypeRepositoryList:   repositoryListValidator(),
	}
	return vmap

}

func namespaceValidator() *gojsonschema.Schema {
	path := "namespace/single.schema.json"
	mediaType := string(MediaTypeNamespace)
	return load(path, mediaType)
}

func namespaceCreateValidator() *gojsonschema.Schema {
	path := "namespace/create.schema.json"
	mediaType := string(MediaTypeNamespaceCreate)
	return load(path, mediaType)
}

func namespaceListValidator() *gojsonschema.Schema {
	path := "namespace/list.schema.json"
	mediaType := string(MediaTypeNamespaceList)
	return load(path, mediaType)
}

func repositoryValidator() *gojsonschema.Schema {
	path := "repository/single.schema.json"
	mediaType := string(MediaTypeRepository)
	return load(path, mediaType)
}

func repositoryCreateValidator() *gojsonschema.Schema {
	path := "repository/create.schema.json"
	mediaType := string(MediaTypeRepositoryCreate)
	return load(path, mediaType)
}

func repositoryListValidator() *gojsonschema.Schema {
	path := "repository/list.schema.json"
	mediaType := string(MediaTypeRepositoryList)
	return load(path, mediaType)
}

func projectValidator() *gojsonschema.Schema {
	path := "project/single.schema.json"
	mediaType := string(MediaTypeProject)
	return load(path, mediaType)
}

func projectCreateValidator() *gojsonschema.Schema {
	path := "project/create.schema.json"
	mediaType := string(MediaTypeProjectCreate)
	return load(path, mediaType)
}

func projectListValidator() *gojsonschema.Schema {
	path := "project/list.schema.json"
	mediaType := string(MediaTypeProjectList)
	return load(path, mediaType)
}