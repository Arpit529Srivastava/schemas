// Package v1beta1 provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.3.0 DO NOT EDIT.
package v1beta1

import (
	"encoding/json"
	"fmt"

	"github.com/gofrs/uuid"
)

// Defines values for ComponentDefinitionFormat.
const (
	CUE  ComponentDefinitionFormat = "CUE"
	JSON ComponentDefinitionFormat = "JSON"
)

// Defines values for ComponentDefinitionMetadataShape.
const (
	Barrel               ComponentDefinitionMetadataShape = "barrel"
	BottomRoundRectangle ComponentDefinitionMetadataShape = "bottom-round-rectangle"
	ConcaveHexagon       ComponentDefinitionMetadataShape = "concave-hexagon"
	CutRectangle         ComponentDefinitionMetadataShape = "cut-rectangle"
	Diamond              ComponentDefinitionMetadataShape = "diamond"
	Ellipse              ComponentDefinitionMetadataShape = "ellipse"
	Heptagon             ComponentDefinitionMetadataShape = "heptagon"
	Hexagon              ComponentDefinitionMetadataShape = "hexagon"
	Octagon              ComponentDefinitionMetadataShape = "octagon"
	Pentagon             ComponentDefinitionMetadataShape = "pentagon"
	Rectangle            ComponentDefinitionMetadataShape = "rectangle"
	Rhomboid             ComponentDefinitionMetadataShape = "rhomboid"
	RoundDiamond         ComponentDefinitionMetadataShape = "round-diamond"
	RoundHeptagon        ComponentDefinitionMetadataShape = "round-heptagon"
	RoundHexagon         ComponentDefinitionMetadataShape = "round-hexagon"
	RoundOctagon         ComponentDefinitionMetadataShape = "round-octagon"
	RoundPentagon        ComponentDefinitionMetadataShape = "round-pentagon"
	RoundRectangle       ComponentDefinitionMetadataShape = "round-rectangle"
	RoundTag             ComponentDefinitionMetadataShape = "round-tag"
	RoundTriangle        ComponentDefinitionMetadataShape = "round-triangle"
	Star                 ComponentDefinitionMetadataShape = "star"
	Tag                  ComponentDefinitionMetadataShape = "tag"
	Triangle             ComponentDefinitionMetadataShape = "triangle"
	Vee                  ComponentDefinitionMetadataShape = "vee"
)

// Defines values for ComponentDefinitionModelStatus.
const (
	ComponentDefinitionModelStatusDuplicate ComponentDefinitionModelStatus = "duplicate"
	ComponentDefinitionModelStatusEnabled   ComponentDefinitionModelStatus = "enabled"
	ComponentDefinitionModelStatusIgnored   ComponentDefinitionModelStatus = "ignored"
)

// Defines values for ModelDefinitionStatus.
const (
	ModelDefinitionStatusDuplicate ModelDefinitionStatus = "duplicate"
	ModelDefinitionStatusEnabled   ModelDefinitionStatus = "enabled"
	ModelDefinitionStatusIgnored   ModelDefinitionStatus = "ignored"
)

type Component struct {
	// Kind The unique identifier (name) assigned by the registrant to this component. Example: A Kubernetes Pod is of kind 'Pod'.
	Kind string `json:"kind" yaml:"kind"`

	// Schema JSON schema of the object as defined by the registrant.
	Schema string `json:"schema" yaml:"schema"`

	// Version Version of the component produced by the registrant. Example: APIVersion of a Kubernetes Pod.
	Version string `json:"version" yaml:"version"`
}

// ComponentDefinition Components are reusable building blocks for depicting capabilities defined within models. Learn more at https://docs.meshery.io/concepts/components
type ComponentDefinition struct {
	// Component Component and it's properties.
	Component     `json:"component" yaml:"component"`
	Configuration map[string]interface{} `json:"configuration" yaml:"configuration"`

	// Description A written representation of the purpose and characteristics of the component.
	Description string `json:"description" yaml:"description"`

	// DisplayName Name of the component in human-readible format.
	DisplayName string `json:"displayName" yaml:"displayName"`

	// Format Format specifies the format used in the `component.schema` field. JSON is the default.
	Format ComponentDefinitionFormat `json:"format" yaml:"format"`

	// Id Uniquely identifies the entity (i.e. component) as defined in a declaration (i.e. design).
	Id *uuid.UUID `json:"id" yaml:"id"`

	// Metadata Metadata contains additional information associated with the component.
	Metadata *ComponentDefinition_Metadata `json:"metadata,omitempty"`

	// Model Reference to the specific registered model to which the component belongs and from which model version, category, and other properties may be referenced. Learn more at https://docs.meshery.io/concepts/models
	Model ModelDefinition `json:"model"`

	// SchemaVersion Specifies the version of the schema to which the component definition conforms.
	SchemaVersion string `json:"schemaVersion" yaml:"schemaVersion"`

	// Version Version of the component definition.
	Version string `json:"version" yaml:"version"`
}

// ComponentDefinitionFormat Format specifies the format used in the `component.schema` field. JSON is the default.
type ComponentDefinitionFormat string

// ComponentDefinitionMetadataShape Shape of the component used for UI representation.
type ComponentDefinitionMetadataShape string

// ComponentDefinition_Metadata Metadata contains additional information associated with the component.
type ComponentDefinition_Metadata struct {
	// Capabilities Meshery manages components in accordance with their specific capabilities. This field explicitly identifies those capabilities largely by what actions a given component supports; e.g. metric-scrape, sub-interface, and so on. This field is extensible. ComponentDefinitions may define a broad array of capabilities, which are in-turn dynamically interpretted by Meshery for full lifecycle management.
	Capabilities *map[string]interface{} `json:"capabilities" yaml:"capabilities"`

	// Genealogy Genealogy represents the various representational states of the component.
	Genealogy *string `json:"genealogy" yaml:"genealogy"`

	// IsAnnotation Identifies whether the component is semantically meaningful or not; identifies whether the component should be treated as deployable entity or is for purposes of logical representation.
	IsAnnotation *bool `json:"isAnnotation" yaml:"isAnnotation"`

	// PrimaryColor Primary color of the component used for UI representation.
	PrimaryColor string `json:"primaryColor" yaml:"primaryColor"`

	// Published 'published' controls whether the component should be registered in Meshery Registry. When the same 'published' property in Models, is set to 'false', the Model property takes precedence with all Entities in the Model not being registered.
	Published *bool `json:"published" yaml:"published"`

	// SecondaryColor Secondary color of the component used for UI representation.
	SecondaryColor *string `json:"secondaryColor" yaml:"secondaryColor"`

	// Shape Shape of the component used for UI representation.
	Shape ComponentDefinitionMetadataShape `json:"shape" yaml:"shape"`

	// SvgColor Colored SVG of the component used for UI representation on light background.
	SvgColor string `json:"svgColor" yaml:"svgColor"`

	// SvgComplete Complete SVG of the component used for UI representation, often inclusive of background.
	SvgComplete *string `json:"svgComplete" yaml:"svgComplete"`

	// SvgWhite White SVG of the component used for UI representation on dark background.
	SvgWhite             string                 `json:"svgWhite" yaml:"svgWhite"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// ComponentDefinitionModelStatus Status of model, including:
// - duplicate: this component is a duplicate of another. The component that is to be the canonical reference and that is duplicated by other components should not be assigned the 'duplicate' status.
// - maintenance: model is unavailable for a period of time.
// - enabled: model is available for use for all users of this Meshery Server.
// - ignored: model is unavailable for use for all users of this Meshery Server.
type ComponentDefinitionModelStatus string

type Model struct {
	// Version Version of the model as defined by the registrant.
	Version string `json:"version" yaml:"version"`
}

// ModelDefinition Meshery Models serve as a portable unit of packaging to define managed entities, their relationships, and capabilities.
type ModelDefinition struct {
	// Category Category of the model.
	Category string `json:"category" yaml:"category"`

	// Description Description of the model.
	Description *string `json:"description" yaml:"description"`

	// DisplayName Human-readable name for the model.
	DisplayName *string `json:"displayName" yaml:"displayName"`

	// Metadata Metadata containing additional information associated with the model.
	Metadata *ModelDefinition_Metadata `json:"metadata,omitempty"`

	// Model Registrant-defined data associated with the model. Properties pertain to the software being managed (e.g. Kubernetes v1.31)
	Model *Model `json:"model,omitempty" yaml:"model"`

	// Name The unique name for the model within the scope of a registrant.
	Name       string                 `json:"name" yaml:"name"`
	Registrant map[string]interface{} `json:"registrant" yaml:"registrant"`

	// SchemaVersion Specifies the version of the schema used for the definition.
	SchemaVersion *string `json:"schemaVersion" yaml:"schemaVersion"`

	// Status Status of model, including:
	// - duplicate: this component is a duplicate of another. The component that is to be the canonical reference and that is duplicated by other components should not be assigned the 'duplicate' status.
	// - maintenance: model is unavailable for a period of time.
	// - enabled: model is available for use for all users of this Meshery Server.
	// - ignored: model is unavailable for use for all users of this Meshery Server.
	Status *ModelDefinitionStatus `json:"status" yaml:"status"`

	// SubCategory Sub-category of the model.
	SubCategory *string `json:"subCategory" yaml:"subCategory"`

	// Version Version of the model definition.
	Version string `json:"version" yaml:"version"`
}

// ModelDefinition_Metadata Metadata containing additional information associated with the model.
type ModelDefinition_Metadata struct {
	// IsAnnotation Indicates whether the model and its entities should be treated as deployable entities or as logical representations.
	IsAnnotation *bool `json:"isAnnotation" yaml:"isAnnotation"`

	// PrimaryColor Primary color associated with the model.
	PrimaryColor *string `json:"primaryColor" yaml:"primaryColor"`

	// SecondaryColor Secondary color associated with the model.
	SecondaryColor *string `json:"secondaryColor" yaml:"secondaryColor"`

	// SvgColor SVG representation of the model in colored format.
	SvgColor *string `json:"svgColor" yaml:"svgColor"`

	// SvgComplete SVG representation of the complete model.
	SvgComplete *string `json:"svgComplete" yaml:"svgComplete"`

	// SvgWhite SVG representation of the model in white color.
	SvgWhite             *string                `json:"svgWhite" yaml:"svgWhite"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// ModelDefinitionStatus Status of model, including:
// - duplicate: this component is a duplicate of another. The component that is to be the canonical reference and that is duplicated by other components should not be assigned the 'duplicate' status.
// - maintenance: model is unavailable for a period of time.
// - enabled: model is available for use for all users of this Meshery Server.
// - ignored: model is unavailable for use for all users of this Meshery Server.
type ModelDefinitionStatus string

// Getter for additional properties for ComponentDefinition_Metadata. Returns the specified
// element and whether it was found
func (a ComponentDefinition_Metadata) Get(fieldName string) (value interface{}, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for ComponentDefinition_Metadata
func (a *ComponentDefinition_Metadata) Set(fieldName string, value interface{}) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]interface{})
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for ComponentDefinition_Metadata to handle AdditionalProperties
func (a *ComponentDefinition_Metadata) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["capabilities"]; found {
		err = json.Unmarshal(raw, &a.Capabilities)
		if err != nil {
			return fmt.Errorf("error reading 'capabilities': %w", err)
		}
		delete(object, "capabilities")
	}

	if raw, found := object["genealogy"]; found {
		err = json.Unmarshal(raw, &a.Genealogy)
		if err != nil {
			return fmt.Errorf("error reading 'genealogy': %w", err)
		}
		delete(object, "genealogy")
	}

	if raw, found := object["isAnnotation"]; found {
		err = json.Unmarshal(raw, &a.IsAnnotation)
		if err != nil {
			return fmt.Errorf("error reading 'isAnnotation': %w", err)
		}
		delete(object, "isAnnotation")
	}

	if raw, found := object["primaryColor"]; found {
		err = json.Unmarshal(raw, &a.PrimaryColor)
		if err != nil {
			return fmt.Errorf("error reading 'primaryColor': %w", err)
		}
		delete(object, "primaryColor")
	}

	if raw, found := object["published"]; found {
		err = json.Unmarshal(raw, &a.Published)
		if err != nil {
			return fmt.Errorf("error reading 'published': %w", err)
		}
		delete(object, "published")
	}

	if raw, found := object["secondaryColor"]; found {
		err = json.Unmarshal(raw, &a.SecondaryColor)
		if err != nil {
			return fmt.Errorf("error reading 'secondaryColor': %w", err)
		}
		delete(object, "secondaryColor")
	}

	if raw, found := object["shape"]; found {
		err = json.Unmarshal(raw, &a.Shape)
		if err != nil {
			return fmt.Errorf("error reading 'shape': %w", err)
		}
		delete(object, "shape")
	}

	if raw, found := object["svgColor"]; found {
		err = json.Unmarshal(raw, &a.SvgColor)
		if err != nil {
			return fmt.Errorf("error reading 'svgColor': %w", err)
		}
		delete(object, "svgColor")
	}

	if raw, found := object["svgComplete"]; found {
		err = json.Unmarshal(raw, &a.SvgComplete)
		if err != nil {
			return fmt.Errorf("error reading 'svgComplete': %w", err)
		}
		delete(object, "svgComplete")
	}

	if raw, found := object["svgWhite"]; found {
		err = json.Unmarshal(raw, &a.SvgWhite)
		if err != nil {
			return fmt.Errorf("error reading 'svgWhite': %w", err)
		}
		delete(object, "svgWhite")
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]interface{})
		for fieldName, fieldBuf := range object {
			var fieldVal interface{}
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return fmt.Errorf("error unmarshaling field %s: %w", fieldName, err)
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for ComponentDefinition_Metadata to handle AdditionalProperties
func (a ComponentDefinition_Metadata) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	if a.Capabilities != nil {
		object["capabilities"], err = json.Marshal(a.Capabilities)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'capabilities': %w", err)
		}
	}

	if a.Genealogy != nil {
		object["genealogy"], err = json.Marshal(a.Genealogy)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'genealogy': %w", err)
		}
	}

	if a.IsAnnotation != nil {
		object["isAnnotation"], err = json.Marshal(a.IsAnnotation)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'isAnnotation': %w", err)
		}
	}

	object["primaryColor"], err = json.Marshal(a.PrimaryColor)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'primaryColor': %w", err)
	}

	if a.Published != nil {
		object["published"], err = json.Marshal(a.Published)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'published': %w", err)
		}
	}

	if a.SecondaryColor != nil {
		object["secondaryColor"], err = json.Marshal(a.SecondaryColor)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'secondaryColor': %w", err)
		}
	}

	object["shape"], err = json.Marshal(a.Shape)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'shape': %w", err)
	}

	object["svgColor"], err = json.Marshal(a.SvgColor)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'svgColor': %w", err)
	}

	if a.SvgComplete != nil {
		object["svgComplete"], err = json.Marshal(a.SvgComplete)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'svgComplete': %w", err)
		}
	}

	object["svgWhite"], err = json.Marshal(a.SvgWhite)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'svgWhite': %w", err)
	}

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, fmt.Errorf("error marshaling '%s': %w", fieldName, err)
		}
	}
	return json.Marshal(object)
}

// Getter for additional properties for ModelDefinition_Metadata. Returns the specified
// element and whether it was found
func (a ModelDefinition_Metadata) Get(fieldName string) (value interface{}, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for ModelDefinition_Metadata
func (a *ModelDefinition_Metadata) Set(fieldName string, value interface{}) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]interface{})
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for ModelDefinition_Metadata to handle AdditionalProperties
func (a *ModelDefinition_Metadata) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["isAnnotation"]; found {
		err = json.Unmarshal(raw, &a.IsAnnotation)
		if err != nil {
			return fmt.Errorf("error reading 'isAnnotation': %w", err)
		}
		delete(object, "isAnnotation")
	}

	if raw, found := object["primaryColor"]; found {
		err = json.Unmarshal(raw, &a.PrimaryColor)
		if err != nil {
			return fmt.Errorf("error reading 'primaryColor': %w", err)
		}
		delete(object, "primaryColor")
	}

	if raw, found := object["secondaryColor"]; found {
		err = json.Unmarshal(raw, &a.SecondaryColor)
		if err != nil {
			return fmt.Errorf("error reading 'secondaryColor': %w", err)
		}
		delete(object, "secondaryColor")
	}

	if raw, found := object["svgColor"]; found {
		err = json.Unmarshal(raw, &a.SvgColor)
		if err != nil {
			return fmt.Errorf("error reading 'svgColor': %w", err)
		}
		delete(object, "svgColor")
	}

	if raw, found := object["svgComplete"]; found {
		err = json.Unmarshal(raw, &a.SvgComplete)
		if err != nil {
			return fmt.Errorf("error reading 'svgComplete': %w", err)
		}
		delete(object, "svgComplete")
	}

	if raw, found := object["svgWhite"]; found {
		err = json.Unmarshal(raw, &a.SvgWhite)
		if err != nil {
			return fmt.Errorf("error reading 'svgWhite': %w", err)
		}
		delete(object, "svgWhite")
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]interface{})
		for fieldName, fieldBuf := range object {
			var fieldVal interface{}
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return fmt.Errorf("error unmarshaling field %s: %w", fieldName, err)
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for ModelDefinition_Metadata to handle AdditionalProperties
func (a ModelDefinition_Metadata) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	if a.IsAnnotation != nil {
		object["isAnnotation"], err = json.Marshal(a.IsAnnotation)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'isAnnotation': %w", err)
		}
	}

	if a.PrimaryColor != nil {
		object["primaryColor"], err = json.Marshal(a.PrimaryColor)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'primaryColor': %w", err)
		}
	}

	if a.SecondaryColor != nil {
		object["secondaryColor"], err = json.Marshal(a.SecondaryColor)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'secondaryColor': %w", err)
		}
	}

	if a.SvgColor != nil {
		object["svgColor"], err = json.Marshal(a.SvgColor)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'svgColor': %w", err)
		}
	}

	if a.SvgComplete != nil {
		object["svgComplete"], err = json.Marshal(a.SvgComplete)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'svgComplete': %w", err)
		}
	}

	if a.SvgWhite != nil {
		object["svgWhite"], err = json.Marshal(a.SvgWhite)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'svgWhite': %w", err)
		}
	}

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, fmt.Errorf("error marshaling '%s': %w", fieldName, err)
		}
	}
	return json.Marshal(object)
}
