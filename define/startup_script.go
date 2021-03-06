package define

import (
	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/output"
	"github.com/sacloud/usacloud/schema"
)

func StartupScriptResource() *schema.Resource {

	commands := map[string]*schema.Command{
		"list": {
			Type:               schema.CommandList,
			Aliases:            []string{"ls", "find", "selector"},
			Params:             startupScriptListParam(),
			TableType:          output.TableSimple,
			TableColumnDefines: startupScriptListColumns(),
			Category:           "basics",
			Order:              10,
		},
		"create": {
			Type:             schema.CommandCreate,
			Params:           startupScriptCreateParam(),
			IncludeFields:    startupScriptDetailIncludes(),
			ExcludeFields:    startupScriptDetailExcludes(),
			UseCustomCommand: true,
			Category:         "basics",
			Order:            20,
		},
		"read": {
			Type:          schema.CommandRead,
			Params:        startupScriptReadParam(),
			IncludeFields: startupScriptDetailIncludes(),
			ExcludeFields: startupScriptDetailExcludes(),
			Category:      "basics",
			Order:         30,
		},
		"update": {
			Type:             schema.CommandUpdate,
			Params:           startupScriptUpdateParam(),
			IncludeFields:    startupScriptDetailIncludes(),
			ExcludeFields:    startupScriptDetailExcludes(),
			UseCustomCommand: true,
			Category:         "basics",
			Order:            40,
		},
		"delete": {
			Type:          schema.CommandDelete,
			Aliases:       []string{"rm"},
			Params:        startupScriptDeleteParam(),
			IncludeFields: startupScriptDetailIncludes(),
			ExcludeFields: startupScriptDetailExcludes(),
			Category:      "basics",
			Order:         50,
		},
	}

	return &schema.Resource{
		Commands:            commands,
		Aliases:             []string{"note"},
		AltResource:         "Note",
		ListResultFieldName: "Notes",
		ResourceCategory:    CategoryCommonItem,
	}
}

func startupScriptListParam() map[string]*schema.Schema {
	return mergeParameterMap(CommonListParam, paramScopeCond, paramTagsCond, paramClassCond)

}

func startupScriptListColumns() []output.ColumnDef {
	return []output.ColumnDef{
		{Name: "ID"},
		{Name: "Class"},
		{Name: "Name"},
		{Name: "Scope"},
	}
}

func startupScriptDetailIncludes() []string {
	return []string{}
}

func startupScriptDetailExcludes() []string {
	return []string{}
}

var allowNoteClasses = []string{
	string(sacloud.NoteClassShell),
	string(sacloud.NoteClassYAMLCloudConfig),
}

func startupScriptCreateParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"script-content": {
			Type:          schema.TypeString,
			HandlerType:   schema.HandlerNoop,
			Aliases:       []string{"note-content"},
			Description:   "set script content",
			ConflictsWith: []string{"script"},
			Category:      "script-input",
			Order:         10,
		},
		"script": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Aliases:      []string{"note"},
			Description:  "set script from file",
			ValidateFunc: validateFileExists(),
			Category:     "script-upload",
			Order:        10,
		},
		"class": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set script class[shell/cloud-config-yaml]",
			ValidateFunc: validateInStrValues(allowNoteClasses...),
			CompleteFunc: completeInStrValues(allowNoteClasses...),
			Required:     true,
			DefaultValue: string(sacloud.NoteClassShell),
			Category:     "basic",
			Order:        20,
		},
		"name":    paramRequiredName,
		"tags":    paramTags,
		"icon-id": paramIconResourceID,
	}
}

func startupScriptReadParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func startupScriptUpdateParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"script-content": {
			Type:          schema.TypeString,
			HandlerType:   schema.HandlerNoop,
			Aliases:       []string{"note-content"},
			Description:   "set script content",
			ConflictsWith: []string{"script"},
			Category:      "script-input",
			Order:         10,
		},
		"script": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Aliases:      []string{"note"},
			Description:  "set script from file",
			ValidateFunc: validateFileExists(),
			Category:     "script-upload",
			Order:        10,
		},
		"class": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set script class[shell/cloud-config-yaml]",
			ValidateFunc: validateInStrValues(allowNoteClasses...),
			CompleteFunc: completeInStrValues(allowNoteClasses...),
			Category:     "basic",
			Order:        20,
		},
		"name":    paramName,
		"tags":    paramTags,
		"icon-id": paramIconResourceID,
	}
}

func startupScriptDeleteParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}
