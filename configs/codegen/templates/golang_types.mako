<%
    import re
    import yacg.model.model as model
    import yacg.templateHelper as templateHelper
    import yacg.model.modelFuncs as modelFuncs
    import yacg.util.stringUtils as stringUtils

    templateFile = 'golang_types.mako'
    templateVersion = '1.1.0'

    packageName = templateParameters.get('modelPackage','<<PLEASE SET modelPackage TEMPLATE PARAM>>')
    optionalPackage = templateParameters.get('optionalPackage','<<PLEASE SET optionalPackage TEMPLATE PARAM>>')
    jsonTypesPackage = templateParameters.get('jsonTypesPackage','<<PLEASE SET jsonTypesPackage TEMPLATE PARAM>>')
    jsonSerialization = templateParameters.get('jsonSerialization',False)

    def printStarForJson(isJson):
        return "*" if isJson else ""

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
            if jsonSerialization:
                ret = printStarForJson(forJson and (not isRequired)) + 'bool'
            else:
                ret = 'bool'
        elif isinstance(typeObj, model.StringType):
            ret = 'string'
        elif isinstance(typeObj, model.BytesType):
            ret = 'byte'
        elif isinstance(typeObj, model.UuidType):
            ret = printStarForJson(forJson and (not isRequired)) + 'uuid.UUID'
        elif isinstance(typeObj, model.EnumType):
            ret = typeObj.name
        elif isinstance(typeObj, model.DateType):
            if jsonSerialization:
                ret = printStarForJson(forJson and (not isRequired)) + 'json_types.JsonDate'
            else:
                ret = 'time.Time'
        elif isinstance(typeObj, model.TimeType):
            if jsonSerialization:
                ret = printStarForJson(forJson and (not isRequired)) + 'json_types.JsonTime'
            else:
                ret = 'time.Time'
        elif isinstance(typeObj, model.DateTimeType):
            if jsonSerialization:
                ret = printStarForJson(forJson and (not isRequired)) + 'json_types.JsonTimestamp'
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

    def printOmitemptyIfNeeded(property):
        if not property.required or property.isArray or isinstance(property.type, model.DictionaryType):
            return ",omitempty"
        else:
            return ""

    def getEnumDefaultValue(type):
        if type.default is not None:
            return secureEnumValues(type.default)
        else:
            return secureEnumValues(type.values[0])

    def secureEnumValues(value):
        pattern = re.compile("^[0-9]")
        return '_' + value if pattern.match(value) else value

    def isEnumDefaultValue(value, type):
        return getEnumDefaultValue(type) == secureEnumValues(value)

%>// Attention, this file is generated. Manual changes get lost with the next
// run of the code generation.
// created by yacg (template: ${templateFile} v${templateVersion})
package ${packageName}


import (
% if modelFuncs.isUuidContained(modelTypes):
    uuid "github.com/google/uuid"
% endif
% if modelFuncs.isAtLeastOneDateRelatedTypeContained(modelTypes):
    % if jsonSerialization:
    json_types "${jsonTypesPackage}"
    % else:
    "time"
    % endif
% endif
% if modelFuncs.containsOptionalProps(modelTypes):
    optional "${optionalPackage}"
% endif
	encJson "encoding/json"
    "reflect"
)

% for type in modelTypes:
    % if modelFuncs.isEnumType(type):
        % if type.description != None:
/* ${templateHelper.addLineBreakToDescription(type.description,4)}
*/
        % endif
type ${type.name} string

const (
    ${type.name}_${getEnumDefaultValue(type)} ${type.name} = "${getEnumDefaultValue(type)}"
        % for value in type.values:
            % if not isEnumDefaultValue(value, type):
    ${type.name}_${value} = "${value}"
            % endif
        % endfor
)

    % endif

    % if hasattr(type, "properties"):
        % if type.description != None:
/* ${templateHelper.addLineBreakToDescription(type.description,4)}
*/
        % endif
type ${type.name} struct {
        % for property in type.properties:

            % if property.description != None:
    // ${property.description}
            % endif
    ${stringUtils.toUpperCamelCase(property.name)} ${printGolangType(property.type, property.isArray, property.required, property.arrayDimensions, False)}
        % endfor
}

// Creates a ${type.name} object
func Make${type.name}() ${type.name} {
    var ret ${type.name}
    // TODO: initialize default values
    return ret
}

func (v ${type.name}) Equals(o ${type.name}) bool {
        % for property in type.properties:
            % if isinstance(property.type, model.DictionaryType):
    if len(v.${stringUtils.toUpperCamelCase(property.name)}) != len(o.${stringUtils.toUpperCamelCase(property.name)}) {
        return false
    }
    for key, vValue := range v.${stringUtils.toUpperCamelCase(property.name)} {
        oValue, exists := o.${stringUtils.toUpperCamelCase(property.name)}[key]
        if (!exists) || (!reflect.DeepEqual(oValue, vValue)) {
            return false
        }
    }
            % elif (not property.isArray) and (not isinstance(property.type, model.ArrayType)) and (modelFuncs.isBaseType(property.type) or isinstance(property.type, model.EnumType)):
    if v.${stringUtils.toUpperCamelCase(property.name)} != o.${stringUtils.toUpperCamelCase(property.name)} {
        return false
    }
            % else:
    if !reflect.DeepEqual(v.${stringUtils.toUpperCamelCase(property.name)}, o.${stringUtils.toUpperCamelCase(property.name)}) {
        return false
    }
            % endif
        % endfor
	return true
}

    % endif
    % if modelFuncs.isEnumType(type) or hasattr(type, "properties"):
        % if modelFuncs.typeIsAsOptionalContained(type.name, modelTypes):
type Optional${type.name} struct {
	Value ${type.name}
	IsSet bool
}

// Creates a Optional${type.name} object
func MakeOptional${type.name}() Optional${type.name} {
    var ret Optional${type.name}
    // TODO: initialize default values
    return ret
}

func (m Optional${type.name}) Set(v ${type.name}) {
	m.Value = v
	m.IsSet = true
}

func (m Optional${type.name}) UnSet() {
	m.IsSet = false
}
        % endif

        % if jsonSerialization:
func (v ${type.name}) MarshalJSON() ([]byte, error) {
            % for property in type.properties:
                % if (not property.required) or property.isArray or isinstance(property.type, model.DictionaryType):
    var _${stringUtils.toUpperCamelCase(property.name)} ${printGolangType(property.type, property.isArray, property.required, property.arrayDimensions, True)}
                    % if (property.isArray) or (isinstance(property.type, model.DictionaryType)):
	if len(v.${stringUtils.toUpperCamelCase(property.name)}) > 0 {
                    % else:
	if v.${stringUtils.toUpperCamelCase(property.name)}.IsSet {
                    % endif
                    % if property.isArray:
		_${stringUtils.toUpperCamelCase(property.name)} = &v.${stringUtils.toUpperCamelCase(property.name)}
                    % elif isinstance(property.type, model.DictionaryType):
		_${stringUtils.toUpperCamelCase(property.name)} = &v.${stringUtils.toUpperCamelCase(property.name)}
                    % elif isinstance(property.type, model.ComplexType):
        _${stringUtils.toUpperCamelCase(property.name)} = &v.${stringUtils.toUpperCamelCase(property.name)}.Value
                    % elif isinstance(property.type, model.EnumType):
		_${stringUtils.toUpperCamelCase(property.name)} = string(v.${stringUtils.toUpperCamelCase(property.name)}.Value)
                    % elif isinstance(property.type, model.DateType):
		_${stringUtils.toUpperCamelCase(property.name)} = &v.${stringUtils.toUpperCamelCase(property.name)}.Value
                    % elif isinstance(property.type, model.DateTimeType):
		_${stringUtils.toUpperCamelCase(property.name)} = &v.${stringUtils.toUpperCamelCase(property.name)}.Value
                    % elif isinstance(property.type, model.TimeType):
		_${stringUtils.toUpperCamelCase(property.name)} = &v.${stringUtils.toUpperCamelCase(property.name)}.Value
                    % elif isinstance(property.type, model.UuidType) or isinstance(property.type, model.BooleanType):
		_${stringUtils.toUpperCamelCase(property.name)} = &v.${stringUtils.toUpperCamelCase(property.name)}.Value
                    % else:
		_${stringUtils.toUpperCamelCase(property.name)} = v.${stringUtils.toUpperCamelCase(property.name)}.Value
                    % endif
	}
                % endif
            % endfor

	return encJson.Marshal(&struct {
            % for property in type.properties:
        ${stringUtils.toUpperCamelCase(property.name)} ${printGolangType(property.type, property.isArray, property.required, property.arrayDimensions, True)} `json:"${property.name}${printOmitemptyIfNeeded(property)}"`
            % endfor
	}{
            % for property in type.properties:
                % if not property.required or property.isArray or isinstance(property.type, model.DictionaryType):
        ${stringUtils.toUpperCamelCase(property.name)}: _${stringUtils.toUpperCamelCase(property.name)},
                % else:
        ${stringUtils.toUpperCamelCase(property.name)}: v.${stringUtils.toUpperCamelCase(property.name)},
                % endif
            % endfor
	})
}

            % if modelFuncs.typeIsAsOptionalContained(type.name, modelTypes):
func (v Optional${type.name}) MarshalJSON() ([]byte, error) {
	if v.IsSet {
		return encJson.Marshal(v.Value)
	} else {
		return []byte("null"), nil
	}
}
            % endif
        % endif
    % endif

% endfor
