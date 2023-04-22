<%
    import re
    import yacg.model.model as model
    import yacg.templateHelper as templateHelper
    import yacg.model.modelFuncs as modelFuncs
    import yacg.util.stringUtils as stringUtils

    templateFile = 'golang_types_tests.mako'
    templateVersion = '1.1.0'

    packageName = templateParameters.get('modelPackage','<<PLEASE SET modelPackage TEMPLATE PARAM>>')
    exampleData = templateParameters.get('exampleData','<<PLEASE SET exampleData TEMPLATE PARAM>>')
    exampleDataBasePath = templateParameters.get('exampleDataBasePath','<<PLEASE SET exampleDataBasePath TEMPLATE PARAM>>')
    jsonSerialization = templateParameters.get('jsonSerialization',False)

    def printGolangType(typeObj, isArray, isRequired, arrayDimensions, forJson):
        ret = ''
        if typeObj is None:
            return '???'
        elif isinstance(typeObj, model.IntegerType):
            if typeObj.format is None or typeObj.format == model.IntegerTypeFormatEnum.INT32:
                ret = 'int32'
            else:
                ret = 'int'
        elif isinstance(typeObj, model.ObjectType):
            ret = 'interface{}'
        elif isinstance(typeObj, model.NumberType):
            if typeObj.format is None or typeObj.format == model.NumberTypeFormatEnum.DOUBLE:
                ret = 'float64'
            else:
                ret = 'float32'
        elif isinstance(typeObj, model.BooleanType):
            ret = 'bool'
        elif isinstance(typeObj, model.StringType):
            ret = 'string'
        elif isinstance(typeObj, model.BytesType):
            ret = 'byte'
        elif isinstance(typeObj, model.UuidType):
            ret = 'uuid.UUID'
        elif isinstance(typeObj, model.EnumType):
            ret = typeObj.name
        elif isinstance(typeObj, model.DateType):
            if jsonSerialization:
                ret = 'json_types.JsonDate'
            else:
                ret = 'time.Time'
        elif isinstance(typeObj, model.TimeType):
            if jsonSerialization:
                ret = 'json_types.JsonTime'
            else:
                ret = 'time.Time'
        elif isinstance(typeObj, model.DateTimeType):
            if jsonSerialization:
                ret = 'json_types.JsonTimestamp'
            else:
                ret = 'time.Time'
        elif isinstance(typeObj, model.DictionaryType):
            ret = '{}map[string]{}'.format(printStarForJson(forJson),printGolangType(typeObj.valueType, False, True, 0, False))
        elif isinstance(typeObj, model.ComplexType):
            if (not isArray) and (not isRequired):
                ret = printStarForJson(forJson) + typeObj.name
            else:
                ret = typeObj.name
        else:
            ret = '???'

        if (not forJson) and (not isRequired) and (not isArray) and (not isinstance(typeObj, model.DictionaryType)):
            if isinstance(typeObj, model.EnumType) or hasattr(typeObj, "properties"):
                ret = "Optional{}".format(ret)
            else:
                ret = "optional.Optional[{}]".format(ret)
        if isArray:
            ret = printStarForJson(forJson) + ("[]" * arrayDimensions) + ret
        return ret
%>// Attention, this file is generated. Manual changes get lost with the next
// run of the code generation.
// created by yacg (template: ${templateFile} v${templateVersion})
package ${packageName}

import (
	"testing"
% if jsonSerialization:
	encJson "encoding/json"
	json_helper "oth.types/pkg/json_helper"
% endif
)

% for type in modelTypes:
    % if modelFuncs.isEnumType(type):
    % endif
    % if hasattr(type, "properties"):
func TestMake${type.name}(t *testing.T) {
	o1 :=Make${type.name}()
	o2 :=Make${type.name}()
	if ! o1.Equals(o2) {
		t.Error("two fresh created objects are not equal, type: ${type.name}")
		return
	}
}

        % if modelFuncs.typeIsAsOptionalContained(type.name, modelTypes):
func TestMakeOptional${type.name}(t *testing.T) {
	o1 :=MakeOptional${type.name}()
	o2 :=MakeOptional${type.name}()
	if o1.IsSet || o2.IsSet {
		t.Error("Optional type objects have a value after creation, type: ${type.name}")
		return
	}
	if ! o1.Value.Equals(o2.Value) {
		t.Error("two fresh created objects are not equal, type: ${type.name}")
		return
	}
}
        % endif

        % if jsonSerialization:
func TestJson${type.name}(t *testing.T) {
	var x []${type.name}
	err := json_helper.LoadOneObjFromFile(&x, "${exampleDataBasePath}/${exampleData}/${type.name}.json")
	if err != nil {
		t.Errorf("error while loading bundle file: ${exampleData}/${type.name}.json, error: %v", err)
		return
	}
	bytes, err := encJson.Marshal(x)
	if err != nil {
		t.Errorf("error while json serialize type: ${type.name}, error: %v", err)
		return
	}
	var y []${type.name}

	err2 := encJson.Unmarshal(bytes, &y)
	if err2 != nil {
		t.Errorf("error while json unmarshal type: ${type.name}, error: %v", err)
		return
	}


	for i, value := range y {
		value2 := x[i]
		if !value.Equals(value2) {
			b1, _ := encJson.Marshal(value)
			b2, _ := encJson.Marshal(value2)
			t.Errorf("objects are not equal after json marshal/unmarshal, type: ${type.name}\n%s\n%s\n", b1, b2)
			return
		}
	}
}
        % endif
    % endif
% endfor
